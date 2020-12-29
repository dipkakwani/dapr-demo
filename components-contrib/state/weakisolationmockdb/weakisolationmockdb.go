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


*/

package weakisolationmockdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/dapr/pkg/logger"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	gohttp "net/http"
)

const (
	keyDelimiter        = "||"
	valueEntityProperty = "Value"
	operationTimeout    = 1000

	hostKey = "host"
	storeNameKey = "storeName"
)

type StateStore struct {
	state.DefaultBulkStore
	baseUrl string
	client gohttp.Client

	logger logger.Logger
}

type stateStoreMetadata struct {
	host string
	storeName string
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

	log.Debugf("WeakIsolationMockDB initialised, host: %s storeName: %s", meta.host, meta.storeName)

	return nil
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
	log.Debugf("fetching %s", req.Key)
	resp := r.DoRequest("GET", r.baseUrl + "/" + req.Key, nil, nil)

	//var JSONData map[string]interface{}
	//json.Unmarshal(resp.RawBody, &JSONData)

	return &state.GetResponse{
		Data: resp.RawBody,
	}, nil
}

func (r *StateStore) Set(req *state.SetRequest) error {
	log.Debugf("saving %s", req.Key)
	b, _ := json.Marshal(req)
	r.DoRequest("POST", r.baseUrl, b, nil)

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

	return &meta, nil
}

func (r *StateStore) DoRequest(method, url string, body []byte, params map[string]string, headers ...string) HTTPResponse {
	if params != nil {
		url += "?"
		for k, v := range params {
			url += k + "=" + v + "&"
		}
		url = url[:len(url)-1]
	}
	req, _ := gohttp.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if len(headers) == 1 {
		req.Header.Set("If-Match", headers[0])
	}
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
