// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

/*
WeakIsolationMockDB connector for DAPR.

Sample configuration in yaml:

	apiVersion: dapr.io/v1alpha1
	kind: Component
	metadata:
	  name: statestore
	spec:
	  type: state.weakisolationmockdb
	  metadata:
	  - name: host
		value: <host:ip>
	  - name: storeName
		value: <storeName>
	  - name: replicas
		value: <numberOfReplicas>


*/

package weakisolationmockdb

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/dapr/pkg/logger"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	gohttp "net/http"
	"strconv"
	"time"
)

const (
	keyDelimiter        = "||"
	valueEntityProperty = "Value"
	operationTimeout    = 1000

	hostKey = "host"
	storeNameKey = "storeName"
	replicasKey = "replicas"
)

type StateStore struct {
	state.DefaultBulkStore
	baseUrl string
	client gohttp.Client
	replicas int
	clientId uint32


	logger logger.Logger
}

type stateStoreMetadata struct {
	host string
	storeName string
	replicas int
}

type HTTPResponse struct {
	StatusCode  int
	ContentType string
	RawHeader   gohttp.Header
	RawBody     []byte
	JSONBody    interface{}
	ErrorBody   map[string]string
}

func NewWeakIsolationMockDBStateStore(logger logger.Logger) *StateStore {
	s := &StateStore{
		logger: logger,
	}
	s.DefaultBulkStore = state.NewDefaultBulkStore(s)

	return s
}

// Initialises DB
func (r *StateStore) Init(metadata state.Metadata) error {
	meta, err := getMetadata(metadata.Properties)
	if err != nil {
		return err
	}

	r.client = gohttp.Client{}
	r.baseUrl = meta.host + "/" + meta.storeName
	r.replicas = meta.replicas
	r.clientId = genId()

	log.Infof("WeakIsolationMockDB initialised client id: %s, host: %s storeName: %s",
					r.clientId, meta.host, meta.storeName)

	return nil
}

func genId() uint32 {
	u1 := uuid.New()
	id := binary.BigEndian.Uint32(u1[:4])
	return id
}

func (r *StateStore) Delete(req *state.DeleteRequest) error {
	log.Debugf("delete %s", req.Key)
	// Unsupported
	return nil
}

func (r *StateStore) BulkDelete(req []state.DeleteRequest) error {
	log.Debugf("bulk delete %v key(s)", len(req))
	for _, s := range req {
		err := r.Delete(&s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *StateStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	log.Debugf("GET request %s", req.Key)
	var resp HTTPResponse
	var body []byte
	etag := -1

	if req.Options.Consistency == state.Strong {
		// Send request to all replicas and pick the value with the largest ETag
		for i := 0; i < r.replicas; i++ {
			resp = r.DoRequest("GET", r.baseUrl + "/" + req.Key, nil, fmt.Sprint(i))
			respEtag, _ := strconv.Atoi(resp.RawHeader.Get("ETag"))

			if respEtag > etag {
				body = resp.RawBody
				etag = respEtag
			}
		}

	} else {
		// Randomly choose one replica
		randomSource := rand.NewSource(time.Now().UnixNano())
		randGen := rand.New(randomSource)
		sessionId := fmt.Sprint(randGen.Intn(r.replicas))
		resp = r.DoRequest("GET", r.baseUrl + "/" + req.Key, nil, sessionId)
		body = resp.RawBody
		etag, _ = strconv.Atoi(resp.RawHeader.Get("ETag"))
	}

	return &state.GetResponse{
		Data: body,
		ETag: strconv.Itoa(etag),
	}, nil
}

func (r *StateStore) Set(req *state.SetRequest) error {
	log.Debugf("SET request %s", req.Key)
	b, _ := json.Marshal(req)

	if req.Options.Consistency == state.Strong {
		// Send request to all replicas in sequential order, making sure the first replica
		// has the same ETag
		resp := r.DoRequest("POST", r.baseUrl, b, "0")

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			// Failed in the first replica
			return errors.New(resp.ErrorBody["description"])
		}

		// Skip ETag for the rest of replicas
		req.ETag = ""
		b, _ := json.Marshal(req)

		for i := 1; i < r.replicas; i++ {
			r.DoRequest("POST", r.baseUrl, b, fmt.Sprint(i))
		}

	} else {
		// Randomly choose one replica
		randomSource := rand.NewSource(time.Now().UnixNano())
		randGen := rand.New(randomSource)
		sessionId := fmt.Sprint(randGen.Intn(r.replicas))
		resp := r.DoRequest("POST", r.baseUrl, b, sessionId)
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return errors.New(resp.ErrorBody["description"])
		}
	}
	return nil
}

func (r *StateStore) BulkSet(req []state.SetRequest) error {
	log.Debugf("bulk set %v key(s)", len(req))

	for _, s := range req {
		err := r.Set(&s)
		if err != nil {
			return err
		}
	}

	return nil
}


func getMetadata(metadata map[string]string) (*stateStoreMetadata, error) {
	meta := stateStoreMetadata{}

	if val, ok := metadata[hostKey]; ok && val != "" {
		meta.host = val
	} else {
		return nil, errors.New(fmt.Sprintf("missing or empty %s field from metadata", hostKey))
	}

	if val, ok := metadata[storeNameKey]; ok && val != "" {
		meta.storeName = val
	} else {
		return nil, errors.New(fmt.Sprintf("missing or empty %s field from metadata", storeNameKey))
	}

	if val, ok := metadata[replicasKey]; ok && val != "" {
		meta.replicas, _ =  strconv.Atoi(val)
	} else {
		return nil, errors.New(fmt.Sprintf("missing or empty %s field from metadata", replicasKey))
	}

	return &meta, nil
}

func (r *StateStore) DoRequest(method, url string, body []byte, sessionId string, headers ...string) HTTPResponse {
	req, _ := gohttp.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if len(headers) == 1 {
		req.Header.Set("If-Match", headers[0])
	}

	req.Header.Set("session-id", sessionId)

	res, err := r.client.Do(req)
	if err != nil {
		panic(fmt.Errorf("failed to request: %v", err))
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := HTTPResponse{
		StatusCode:  res.StatusCode,
		ContentType: res.Header.Get("Content-Type"),
		RawHeader:   res.Header,
		RawBody:     bodyBytes,
	}

	if response.ContentType == "application/json" {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			json.Unmarshal(bodyBytes, &response.JSONBody)
		} else {
			json.Unmarshal(bodyBytes, &response.ErrorBody)
		}
	}

	return response
}
