// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/common/v1/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of HTTP 1.1 Methods
// RFC 7231: https://tools.ietf.org/html/rfc7231#page-24
type HTTPExtension_Verb int32

const (
	HTTPExtension_NONE    HTTPExtension_Verb = 0
	HTTPExtension_GET     HTTPExtension_Verb = 1
	HTTPExtension_HEAD    HTTPExtension_Verb = 2
	HTTPExtension_POST    HTTPExtension_Verb = 3
	HTTPExtension_PUT     HTTPExtension_Verb = 4
	HTTPExtension_DELETE  HTTPExtension_Verb = 5
	HTTPExtension_CONNECT HTTPExtension_Verb = 6
	HTTPExtension_OPTIONS HTTPExtension_Verb = 7
	HTTPExtension_TRACE   HTTPExtension_Verb = 8
)

var HTTPExtension_Verb_name = map[int32]string{
	0: "NONE",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PUT",
	5: "DELETE",
	6: "CONNECT",
	7: "OPTIONS",
	8: "TRACE",
}

var HTTPExtension_Verb_value = map[string]int32{
	"NONE":    0,
	"GET":     1,
	"HEAD":    2,
	"POST":    3,
	"PUT":     4,
	"DELETE":  5,
	"CONNECT": 6,
	"OPTIONS": 7,
	"TRACE":   8,
}

func (x HTTPExtension_Verb) String() string {
	return proto.EnumName(HTTPExtension_Verb_name, int32(x))
}

func (HTTPExtension_Verb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{0, 0}
}

// Enum describing the supported concurrency for state.
type StateOptions_StateConcurrency int32

const (
	StateOptions_CONCURRENCY_UNSPECIFIED StateOptions_StateConcurrency = 0
	StateOptions_CONCURRENCY_FIRST_WRITE StateOptions_StateConcurrency = 1
	StateOptions_CONCURRENCY_LAST_WRITE  StateOptions_StateConcurrency = 2
)

var StateOptions_StateConcurrency_name = map[int32]string{
	0: "CONCURRENCY_UNSPECIFIED",
	1: "CONCURRENCY_FIRST_WRITE",
	2: "CONCURRENCY_LAST_WRITE",
}

var StateOptions_StateConcurrency_value = map[string]int32{
	"CONCURRENCY_UNSPECIFIED": 0,
	"CONCURRENCY_FIRST_WRITE": 1,
	"CONCURRENCY_LAST_WRITE":  2,
}

func (x StateOptions_StateConcurrency) String() string {
	return proto.EnumName(StateOptions_StateConcurrency_name, int32(x))
}

func (StateOptions_StateConcurrency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{4, 0}
}

// Enum describing the supported consistency for state.
type StateOptions_StateConsistency int32

const (
	StateOptions_CONSISTENCY_UNSPECIFIED StateOptions_StateConsistency = 0
	StateOptions_CONSISTENCY_EVENTUAL    StateOptions_StateConsistency = 1
	StateOptions_CONSISTENCY_STRONG      StateOptions_StateConsistency = 2
)

var StateOptions_StateConsistency_name = map[int32]string{
	0: "CONSISTENCY_UNSPECIFIED",
	1: "CONSISTENCY_EVENTUAL",
	2: "CONSISTENCY_STRONG",
}

var StateOptions_StateConsistency_value = map[string]int32{
	"CONSISTENCY_UNSPECIFIED": 0,
	"CONSISTENCY_EVENTUAL":    1,
	"CONSISTENCY_STRONG":      2,
}

func (x StateOptions_StateConsistency) String() string {
	return proto.EnumName(StateOptions_StateConsistency_name, int32(x))
}

func (StateOptions_StateConsistency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{4, 1}
}

// HTTPExtension includes HTTP verb and querystring
// when Dapr runtime delivers HTTP content.
//
// For example, when callers calls http invoke api
// POST http://localhost:3500/v1.0/invoke/<app_id>/method/<method>?query1=value1&query2=value2
//
// Dapr runtime will parse POST as a verb and extract querystring to quersytring map.
type HTTPExtension struct {
	// Required. HTTP verb.
	Verb HTTPExtension_Verb `protobuf:"varint,1,opt,name=verb,proto3,enum=dapr.proto.common.v1.HTTPExtension_Verb" json:"verb,omitempty"`
	// querystring includes HTTP querystring.
	Querystring          map[string]string `protobuf:"bytes,2,rep,name=querystring,proto3" json:"querystring,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTPExtension) Reset()         { *m = HTTPExtension{} }
func (m *HTTPExtension) String() string { return proto.CompactTextString(m) }
func (*HTTPExtension) ProtoMessage()    {}
func (*HTTPExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{0}
}

func (m *HTTPExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPExtension.Unmarshal(m, b)
}
func (m *HTTPExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPExtension.Marshal(b, m, deterministic)
}
func (m *HTTPExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPExtension.Merge(m, src)
}
func (m *HTTPExtension) XXX_Size() int {
	return xxx_messageInfo_HTTPExtension.Size(m)
}
func (m *HTTPExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPExtension.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPExtension proto.InternalMessageInfo

func (m *HTTPExtension) GetVerb() HTTPExtension_Verb {
	if m != nil {
		return m.Verb
	}
	return HTTPExtension_NONE
}

func (m *HTTPExtension) GetQuerystring() map[string]string {
	if m != nil {
		return m.Querystring
	}
	return nil
}

// InvokeRequest is the message to invoke a method with the data.
// This message is used in InvokeService of Dapr gRPC Service and OnInvoke
// of AppCallback gRPC service.
type InvokeRequest struct {
	// Required. method is a method name which will be invoked by caller.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Required. Bytes value or Protobuf message which caller sent.
	// Dapr treats Any.value as bytes type if Any.type_url is unset.
	Data *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// The type of data content.
	//
	// This field is required if data delivers http request body
	// Otherwise, this is optional.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// HTTP specific fields if request conveys http-compatible request.
	//
	// This field is required for http-compatible request. Otherwise,
	// this field is optional.
	HttpExtension        *HTTPExtension `protobuf:"bytes,4,opt,name=http_extension,json=httpExtension,proto3" json:"http_extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InvokeRequest) Reset()         { *m = InvokeRequest{} }
func (m *InvokeRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeRequest) ProtoMessage()    {}
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{1}
}

func (m *InvokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeRequest.Unmarshal(m, b)
}
func (m *InvokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeRequest.Marshal(b, m, deterministic)
}
func (m *InvokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeRequest.Merge(m, src)
}
func (m *InvokeRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeRequest.Size(m)
}
func (m *InvokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeRequest proto.InternalMessageInfo

func (m *InvokeRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *InvokeRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *InvokeRequest) GetHttpExtension() *HTTPExtension {
	if m != nil {
		return m.HttpExtension
	}
	return nil
}

// InvokeResponse is the response message inclduing data and its content type
// from app callback.
// This message is used in InvokeService of Dapr gRPC Service and OnInvoke
// of AppCallback gRPC service.
type InvokeResponse struct {
	// Required. The content body of InvokeService response.
	Data *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Required. The type of data content.
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeResponse) Reset()         { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string { return proto.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()    {}
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{2}
}

func (m *InvokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeResponse.Unmarshal(m, b)
}
func (m *InvokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeResponse.Marshal(b, m, deterministic)
}
func (m *InvokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeResponse.Merge(m, src)
}
func (m *InvokeResponse) XXX_Size() int {
	return xxx_messageInfo_InvokeResponse.Size(m)
}
func (m *InvokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeResponse proto.InternalMessageInfo

func (m *InvokeResponse) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeResponse) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

// StateItem represents state key, value, and additional options to save state.
type StateItem struct {
	// Required. The state key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Required. The state data for key
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// The entity tag which represents the specific version of data.
	// The exact ETag format is defined by the corresponding data store.
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// The metadata which will be passed to state store component.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Options for concurrency and consistency to save the state.
	Options              *StateOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StateItem) Reset()         { *m = StateItem{} }
func (m *StateItem) String() string { return proto.CompactTextString(m) }
func (*StateItem) ProtoMessage()    {}
func (*StateItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{3}
}

func (m *StateItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateItem.Unmarshal(m, b)
}
func (m *StateItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateItem.Marshal(b, m, deterministic)
}
func (m *StateItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateItem.Merge(m, src)
}
func (m *StateItem) XXX_Size() int {
	return xxx_messageInfo_StateItem.Size(m)
}
func (m *StateItem) XXX_DiscardUnknown() {
	xxx_messageInfo_StateItem.DiscardUnknown(m)
}

var xxx_messageInfo_StateItem proto.InternalMessageInfo

func (m *StateItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StateItem) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StateItem) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *StateItem) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StateItem) GetOptions() *StateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// StateOptions configures concurrency and consistency for state operations
type StateOptions struct {
	Concurrency          StateOptions_StateConcurrency `protobuf:"varint,1,opt,name=concurrency,proto3,enum=dapr.proto.common.v1.StateOptions_StateConcurrency" json:"concurrency,omitempty"`
	Consistency          StateOptions_StateConsistency `protobuf:"varint,2,opt,name=consistency,proto3,enum=dapr.proto.common.v1.StateOptions_StateConsistency" json:"consistency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *StateOptions) Reset()         { *m = StateOptions{} }
func (m *StateOptions) String() string { return proto.CompactTextString(m) }
func (*StateOptions) ProtoMessage()    {}
func (*StateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c448f683ad359d5, []int{4}
}

func (m *StateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateOptions.Unmarshal(m, b)
}
func (m *StateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateOptions.Marshal(b, m, deterministic)
}
func (m *StateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateOptions.Merge(m, src)
}
func (m *StateOptions) XXX_Size() int {
	return xxx_messageInfo_StateOptions.Size(m)
}
func (m *StateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateOptions proto.InternalMessageInfo

func (m *StateOptions) GetConcurrency() StateOptions_StateConcurrency {
	if m != nil {
		return m.Concurrency
	}
	return StateOptions_CONCURRENCY_UNSPECIFIED
}

func (m *StateOptions) GetConsistency() StateOptions_StateConsistency {
	if m != nil {
		return m.Consistency
	}
	return StateOptions_CONSISTENCY_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("dapr.proto.common.v1.HTTPExtension_Verb", HTTPExtension_Verb_name, HTTPExtension_Verb_value)
	proto.RegisterEnum("dapr.proto.common.v1.StateOptions_StateConcurrency", StateOptions_StateConcurrency_name, StateOptions_StateConcurrency_value)
	proto.RegisterEnum("dapr.proto.common.v1.StateOptions_StateConsistency", StateOptions_StateConsistency_name, StateOptions_StateConsistency_value)
	proto.RegisterType((*HTTPExtension)(nil), "dapr.proto.common.v1.HTTPExtension")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.common.v1.HTTPExtension.QuerystringEntry")
	proto.RegisterType((*InvokeRequest)(nil), "dapr.proto.common.v1.InvokeRequest")
	proto.RegisterType((*InvokeResponse)(nil), "dapr.proto.common.v1.InvokeResponse")
	proto.RegisterType((*StateItem)(nil), "dapr.proto.common.v1.StateItem")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.common.v1.StateItem.MetadataEntry")
	proto.RegisterType((*StateOptions)(nil), "dapr.proto.common.v1.StateOptions")
}

func init() { proto.RegisterFile("dapr/proto/common/v1/common.proto", fileDescriptor_0c448f683ad359d5) }

var fileDescriptor_0c448f683ad359d5 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc6, 0x8e, 0x93, 0xc0, 0x24, 0x41, 0xab, 0x55, 0xc4, 0x2f, 0xbf, 0x70, 0x01, 0xf7, 0x92,
	0x4b, 0x1d, 0x05, 0x7a, 0xa8, 0x0a, 0xaa, 0x14, 0x9c, 0x05, 0x5c, 0x51, 0x3b, 0xb5, 0x1d, 0xaa,
	0x56, 0xaa, 0x22, 0x27, 0x6c, 0x1d, 0x0b, 0x62, 0x1b, 0x7b, 0x13, 0xd5, 0x8f, 0xd0, 0x43, 0x5f,
	0xa4, 0x4f, 0xd1, 0x27, 0xe8, 0x33, 0x55, 0x5e, 0x3b, 0x21, 0xa4, 0xa8, 0xc0, 0x6d, 0xfe, 0x7d,
	0xdf, 0x7c, 0x33, 0x63, 0x2f, 0xec, 0x5f, 0x39, 0x61, 0xd4, 0x0e, 0xa3, 0x80, 0x05, 0xed, 0x71,
	0x30, 0x9d, 0x06, 0x7e, 0x7b, 0xde, 0xc9, 0x2d, 0x85, 0x87, 0x71, 0x3d, 0x2d, 0xc9, 0x6c, 0x25,
	0x4f, 0xcc, 0x3b, 0xcd, 0xff, 0xdd, 0x20, 0x70, 0x6f, 0x68, 0x06, 0x1d, 0xcd, 0xbe, 0xb6, 0x1d,
	0x3f, 0xc9, 0x8a, 0xe4, 0xdf, 0x22, 0xd4, 0xce, 0x6d, 0xbb, 0x4f, 0xbe, 0x31, 0xea, 0xc7, 0x5e,
	0xe0, 0xe3, 0x63, 0x90, 0xe6, 0x34, 0x1a, 0x35, 0x84, 0x3d, 0xa1, 0xb5, 0x7d, 0xd0, 0x52, 0x1e,
	0x62, 0x54, 0xee, 0x41, 0x94, 0x4b, 0x1a, 0x8d, 0x4c, 0x8e, 0xc2, 0x97, 0x50, 0xb9, 0x9d, 0xd1,
	0x28, 0x89, 0x59, 0xe4, 0xf9, 0x6e, 0x43, 0xdc, 0x2b, 0xb4, 0x2a, 0x07, 0xaf, 0x9e, 0x42, 0xf2,
	0xe1, 0x0e, 0x46, 0x7c, 0x16, 0x25, 0xe6, 0x2a, 0x51, 0xf3, 0x2d, 0xa0, 0xf5, 0x02, 0x8c, 0xa0,
	0x70, 0x4d, 0x13, 0x2e, 0x74, 0xcb, 0x4c, 0x4d, 0x5c, 0x87, 0xe2, 0xdc, 0xb9, 0x99, 0xd1, 0x86,
	0xc8, 0x63, 0x99, 0xf3, 0x46, 0x7c, 0x2d, 0xc8, 0x2e, 0x48, 0xa9, 0x4a, 0xbc, 0x09, 0x92, 0x6e,
	0xe8, 0x04, 0x6d, 0xe0, 0x32, 0x14, 0xce, 0x88, 0x8d, 0x84, 0x34, 0x74, 0x4e, 0xba, 0x3d, 0x24,
	0xa6, 0x56, 0xdf, 0xb0, 0x6c, 0x54, 0x48, 0x93, 0xfd, 0x81, 0x8d, 0x24, 0x0c, 0x50, 0xea, 0x91,
	0x0b, 0x62, 0x13, 0x54, 0xc4, 0x15, 0x28, 0xab, 0x86, 0xae, 0x13, 0xd5, 0x46, 0xa5, 0xd4, 0x31,
	0xfa, 0xb6, 0x66, 0xe8, 0x16, 0x2a, 0xe3, 0x2d, 0x28, 0xda, 0x66, 0x57, 0x25, 0x68, 0x53, 0xfe,
	0x25, 0x40, 0x4d, 0xf3, 0xe7, 0xc1, 0x35, 0x35, 0xe9, 0xed, 0x8c, 0xc6, 0x0c, 0xef, 0x40, 0x69,
	0x4a, 0xd9, 0x24, 0xb8, 0xca, 0x95, 0xe6, 0x1e, 0x6e, 0x81, 0x74, 0xe5, 0x30, 0x87, 0x6b, 0xad,
	0x1c, 0xd4, 0x95, 0xec, 0x48, 0xca, 0xe2, 0x48, 0x4a, 0xd7, 0x4f, 0x4c, 0x5e, 0x81, 0xf7, 0xa1,
	0x3a, 0x0e, 0x7c, 0x46, 0x7d, 0x36, 0x64, 0x49, 0x48, 0x1b, 0x05, 0xce, 0x53, 0xc9, 0x63, 0x76,
	0x12, 0x52, 0xfc, 0x0e, 0xb6, 0x27, 0x8c, 0x85, 0x43, 0xba, 0xd8, 0x67, 0x43, 0xe2, 0xb4, 0x2f,
	0x9e, 0xb0, 0x7a, 0xb3, 0x96, 0x42, 0x97, 0xae, 0xfc, 0x05, 0xb6, 0x17, 0x13, 0xc4, 0x61, 0xe0,
	0xc7, 0x74, 0x29, 0x55, 0x78, 0xb6, 0x54, 0xf1, 0x2f, 0xa9, 0xf2, 0x0f, 0x11, 0xb6, 0x2c, 0xe6,
	0x30, 0xaa, 0x31, 0x3a, 0x7d, 0xec, 0x88, 0xd5, 0xfc, 0x88, 0x18, 0x83, 0x44, 0x99, 0xe3, 0xe6,
	0xb3, 0x73, 0x1b, 0x6b, 0xb0, 0x39, 0xa5, 0xcc, 0xe1, 0xd2, 0x24, 0xfe, 0xa5, 0xbd, 0x7c, 0x78,
	0xdc, 0x65, 0x3b, 0xe5, 0x7d, 0x5e, 0x9f, 0x7d, 0x62, 0x4b, 0x38, 0x3e, 0x86, 0x72, 0x10, 0x32,
	0x2f, 0xf0, 0xe3, 0x46, 0x91, 0x0f, 0x29, 0xff, 0x83, 0xc9, 0xc8, 0x2a, 0xcd, 0x05, 0xa4, 0x79,
	0x04, 0xb5, 0x7b, 0xc4, 0xcf, 0xfa, 0x34, 0xbf, 0x17, 0xa0, 0xba, 0x4a, 0x8b, 0x07, 0x90, 0xee,
	0x6b, 0x3c, 0x8b, 0x22, 0xea, 0x8f, 0x93, 0xfc, 0x47, 0x3c, 0x7c, 0x5c, 0x4f, 0xe6, 0xa8, 0x77,
	0x50, 0x73, 0x95, 0x27, 0xa7, 0x8d, 0xbd, 0x98, 0x71, 0x5a, 0xf1, 0xd9, 0xb4, 0x0b, 0xa8, 0xb9,
	0xca, 0x23, 0x4f, 0x00, 0xad, 0xf7, 0xc5, 0xbb, 0xf0, 0x9f, 0x6a, 0xe8, 0xea, 0xc0, 0x34, 0x89,
	0xae, 0x7e, 0x1a, 0x0e, 0x74, 0xab, 0x4f, 0x54, 0xed, 0x54, 0x23, 0x3d, 0xb4, 0xb1, 0x9e, 0x3c,
	0xd5, 0x4c, 0xcb, 0x1e, 0x7e, 0x34, 0x35, 0x9b, 0x20, 0x01, 0x37, 0x61, 0x67, 0x35, 0x79, 0xd1,
	0x5d, 0xe6, 0x44, 0xd9, 0xb9, 0xeb, 0xb4, 0xe8, 0x9e, 0x93, 0x59, 0x9a, 0x65, 0x3f, 0xd0, 0xa9,
	0x01, 0xf5, 0xd5, 0x24, 0xb9, 0x24, 0xba, 0x3d, 0xe8, 0x5e, 0x20, 0x01, 0xef, 0x00, 0x5e, 0xcd,
	0x58, 0xb6, 0x69, 0xe8, 0x67, 0x48, 0x3c, 0xf1, 0x00, 0xbc, 0x20, 0x5b, 0xc9, 0xbc, 0x73, 0x52,
	0x55, 0xf9, 0x42, 0xfa, 0xe9, 0x72, 0xe2, 0xcf, 0x6d, 0xd7, 0x63, 0x93, 0xd9, 0x28, 0xdd, 0x52,
	0x9b, 0xbf, 0xc4, 0xd9, 0x73, 0x7c, 0xed, 0xae, 0x3f, 0xc9, 0x47, 0x99, 0xf5, 0x53, 0xdc, 0xed,
	0xa5, 0x44, 0xea, 0x8d, 0x47, 0x7d, 0xa6, 0x74, 0x67, 0x2c, 0x70, 0xa9, 0xaf, 0x9c, 0x45, 0xe1,
	0x58, 0x99, 0x77, 0x46, 0x25, 0x8e, 0x3a, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x3f, 0x00,
	0x68, 0xd6, 0x05, 0x00, 0x00,
}
