// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/operator/v1/operator.proto

package operator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ComponentUpdateEvent includes the updated component event.
type ComponentUpdateEvent struct {
	Component            []byte   `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentUpdateEvent) Reset()         { *m = ComponentUpdateEvent{} }
func (m *ComponentUpdateEvent) String() string { return proto.CompactTextString(m) }
func (*ComponentUpdateEvent) ProtoMessage()    {}
func (*ComponentUpdateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6e6e3126ef3d27, []int{0}
}

func (m *ComponentUpdateEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentUpdateEvent.Unmarshal(m, b)
}
func (m *ComponentUpdateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentUpdateEvent.Marshal(b, m, deterministic)
}
func (m *ComponentUpdateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentUpdateEvent.Merge(m, src)
}
func (m *ComponentUpdateEvent) XXX_Size() int {
	return xxx_messageInfo_ComponentUpdateEvent.Size(m)
}
func (m *ComponentUpdateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentUpdateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentUpdateEvent proto.InternalMessageInfo

func (m *ComponentUpdateEvent) GetComponent() []byte {
	if m != nil {
		return m.Component
	}
	return nil
}

// ListComponentResponse includes the list of available components.
type ListComponentResponse struct {
	Components           [][]byte `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListComponentResponse) Reset()         { *m = ListComponentResponse{} }
func (m *ListComponentResponse) String() string { return proto.CompactTextString(m) }
func (*ListComponentResponse) ProtoMessage()    {}
func (*ListComponentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6e6e3126ef3d27, []int{1}
}

func (m *ListComponentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListComponentResponse.Unmarshal(m, b)
}
func (m *ListComponentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListComponentResponse.Marshal(b, m, deterministic)
}
func (m *ListComponentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListComponentResponse.Merge(m, src)
}
func (m *ListComponentResponse) XXX_Size() int {
	return xxx_messageInfo_ListComponentResponse.Size(m)
}
func (m *ListComponentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListComponentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListComponentResponse proto.InternalMessageInfo

func (m *ListComponentResponse) GetComponents() [][]byte {
	if m != nil {
		return m.Components
	}
	return nil
}

// GetConfigurationRequest is the request message to get the configuration.
type GetConfigurationRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigurationRequest) Reset()         { *m = GetConfigurationRequest{} }
func (m *GetConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigurationRequest) ProtoMessage()    {}
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6e6e3126ef3d27, []int{2}
}

func (m *GetConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigurationRequest.Unmarshal(m, b)
}
func (m *GetConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigurationRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigurationRequest.Merge(m, src)
}
func (m *GetConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigurationRequest.Size(m)
}
func (m *GetConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigurationRequest proto.InternalMessageInfo

func (m *GetConfigurationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetConfigurationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// GetConfigurationResponse includes the requested configuration.
type GetConfigurationResponse struct {
	Configuration        []byte   `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigurationResponse) Reset()         { *m = GetConfigurationResponse{} }
func (m *GetConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*GetConfigurationResponse) ProtoMessage()    {}
func (*GetConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6e6e3126ef3d27, []int{3}
}

func (m *GetConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigurationResponse.Unmarshal(m, b)
}
func (m *GetConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigurationResponse.Marshal(b, m, deterministic)
}
func (m *GetConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigurationResponse.Merge(m, src)
}
func (m *GetConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_GetConfigurationResponse.Size(m)
}
func (m *GetConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigurationResponse proto.InternalMessageInfo

func (m *GetConfigurationResponse) GetConfiguration() []byte {
	if m != nil {
		return m.Configuration
	}
	return nil
}

// ListSubscriptionsResponse includes pub/sub subscriptions.
type ListSubscriptionsResponse struct {
	Subscriptions        [][]byte `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSubscriptionsResponse) Reset()         { *m = ListSubscriptionsResponse{} }
func (m *ListSubscriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSubscriptionsResponse) ProtoMessage()    {}
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e6e6e3126ef3d27, []int{4}
}

func (m *ListSubscriptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSubscriptionsResponse.Unmarshal(m, b)
}
func (m *ListSubscriptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSubscriptionsResponse.Marshal(b, m, deterministic)
}
func (m *ListSubscriptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSubscriptionsResponse.Merge(m, src)
}
func (m *ListSubscriptionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSubscriptionsResponse.Size(m)
}
func (m *ListSubscriptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSubscriptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSubscriptionsResponse proto.InternalMessageInfo

func (m *ListSubscriptionsResponse) GetSubscriptions() [][]byte {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func init() {
	proto.RegisterType((*ComponentUpdateEvent)(nil), "dapr.proto.operator.v1.ComponentUpdateEvent")
	proto.RegisterType((*ListComponentResponse)(nil), "dapr.proto.operator.v1.ListComponentResponse")
	proto.RegisterType((*GetConfigurationRequest)(nil), "dapr.proto.operator.v1.GetConfigurationRequest")
	proto.RegisterType((*GetConfigurationResponse)(nil), "dapr.proto.operator.v1.GetConfigurationResponse")
	proto.RegisterType((*ListSubscriptionsResponse)(nil), "dapr.proto.operator.v1.ListSubscriptionsResponse")
}

func init() {
	proto.RegisterFile("dapr/proto/operator/v1/operator.proto", fileDescriptor_4e6e6e3126ef3d27)
}

var fileDescriptor_4e6e6e3126ef3d27 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x6d, 0x5e, 0x1f, 0x8f, 0xd7, 0x4b, 0xfd, 0x1a, 0xb4, 0xc6, 0x2a, 0x52, 0x06, 0x85, 0x2e,
	0x74, 0xd2, 0x5a, 0xc5, 0x85, 0x1b, 0xb5, 0x14, 0x17, 0x0a, 0x42, 0xc4, 0x85, 0xba, 0x31, 0x49,
	0xa7, 0x31, 0x68, 0x66, 0xc6, 0xcc, 0xa4, 0xe2, 0x8f, 0xf1, 0xbf, 0x4a, 0x3e, 0x9a, 0xa6, 0x6d,
	0x22, 0xae, 0xe6, 0x72, 0xce, 0x9d, 0x33, 0x73, 0xcf, 0xe1, 0xc2, 0xfe, 0xd0, 0x12, 0x81, 0x21,
	0x02, 0xae, 0xb8, 0xc1, 0x05, 0x0d, 0x2c, 0xc5, 0x03, 0x63, 0xdc, 0xcd, 0x6a, 0x12, 0x53, 0xa8,
	0x11, 0xb5, 0x25, 0x35, 0xc9, 0xa8, 0x71, 0xb7, 0xb9, 0xed, 0x72, 0xee, 0xbe, 0xd1, 0x44, 0xc0,
	0x0e, 0x47, 0x06, 0xf5, 0x85, 0xfa, 0x4c, 0x1a, 0xf1, 0x31, 0xac, 0xf7, 0xb9, 0x2f, 0x38, 0xa3,
	0x4c, 0xdd, 0x8b, 0xa1, 0xa5, 0xe8, 0x60, 0x4c, 0x99, 0x42, 0x3b, 0x50, 0x73, 0x26, 0xb8, 0xae,
	0xb5, 0xb4, 0x76, 0xdd, 0x9c, 0x02, 0xf8, 0x14, 0x36, 0x6e, 0x3c, 0xa9, 0xb2, 0x9b, 0x26, 0x95,
	0x82, 0x33, 0x49, 0xd1, 0x2e, 0x40, 0xd6, 0x25, 0x75, 0xad, 0x55, 0x6d, 0xd7, 0xcd, 0x1c, 0x82,
	0xaf, 0x61, 0xf3, 0x8a, 0xaa, 0x3e, 0x67, 0x23, 0xcf, 0x0d, 0x03, 0x4b, 0x79, 0x9c, 0x99, 0xf4,
	0x3d, 0xa4, 0x52, 0x21, 0x04, 0x7f, 0x99, 0xe5, 0xd3, 0xf8, 0xb1, 0x9a, 0x19, 0xd7, 0xd1, 0x2f,
	0xa2, 0x53, 0x0a, 0xcb, 0xa1, 0xfa, 0x9f, 0x98, 0x98, 0x02, 0xf8, 0x1c, 0xf4, 0x45, 0xb1, 0xf4,
	0x23, 0x7b, 0xb0, 0xe4, 0xe4, 0x89, 0x74, 0x86, 0x59, 0x10, 0x5f, 0xc0, 0x56, 0x34, 0xc7, 0x5d,
	0x68, 0x4b, 0x27, 0xf0, 0x44, 0x84, 0xc9, 0xbc, 0x84, 0xcc, 0x13, 0xe9, 0x38, 0xb3, 0xe0, 0xd1,
	0x57, 0x15, 0xfe, 0xdf, 0xa6, 0x6e, 0xa3, 0x27, 0x58, 0x99, 0x73, 0x13, 0x35, 0x48, 0x62, 0x3f,
	0x99, 0xd8, 0x4f, 0x06, 0x91, 0xfd, 0xcd, 0x03, 0x52, 0x1c, 0x17, 0x29, 0x8a, 0x03, 0x57, 0x3a,
	0x1a, 0x7a, 0x80, 0xe5, 0x19, 0xd3, 0x65, 0xa9, 0xf6, 0x61, 0x99, 0x76, 0x61, 0x68, 0xb8, 0x82,
	0x3e, 0x60, 0x75, 0xde, 0x49, 0x64, 0x94, 0x89, 0x94, 0x04, 0xd8, 0xec, 0xfc, 0xfe, 0x42, 0xf6,
	0xf0, 0x33, 0xac, 0x2d, 0x04, 0x50, 0x3a, 0x56, 0xf7, 0xa7, 0xb1, 0x0a, 0x33, 0xc4, 0x95, 0xcb,
	0x93, 0xc7, 0x9e, 0xeb, 0xa9, 0x97, 0xd0, 0x26, 0x0e, 0xf7, 0x8d, 0x78, 0x93, 0x92, 0x75, 0x7a,
	0x75, 0x17, 0x57, 0xea, 0x6c, 0x52, 0xdb, 0xff, 0x62, 0xae, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0xc7, 0xdc, 0xeb, 0x7c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatorClient interface {
	// Sends events to Dapr sidecars upon component changes.
	ComponentUpdate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Operator_ComponentUpdateClient, error)
	// Returns a list of available components
	ListComponents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListComponentResponse, error)
	// Returns a given configuration by name
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	// Returns a list of pub/sub subscriptions
	ListSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error)
}

type operatorClient struct {
	cc *grpc.ClientConn
}

func NewOperatorClient(cc *grpc.ClientConn) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) ComponentUpdate(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Operator_ComponentUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Operator_serviceDesc.Streams[0], "/dapr.proto.operator.v1.Operator/ComponentUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &operatorComponentUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Operator_ComponentUpdateClient interface {
	Recv() (*ComponentUpdateEvent, error)
	grpc.ClientStream
}

type operatorComponentUpdateClient struct {
	grpc.ClientStream
}

func (x *operatorComponentUpdateClient) Recv() (*ComponentUpdateEvent, error) {
	m := new(ComponentUpdateEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *operatorClient) ListComponents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListComponentResponse, error) {
	out := new(ListComponentResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/ListComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) ListSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error) {
	out := new(ListSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.operator.v1.Operator/ListSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
type OperatorServer interface {
	// Sends events to Dapr sidecars upon component changes.
	ComponentUpdate(*empty.Empty, Operator_ComponentUpdateServer) error
	// Returns a list of available components
	ListComponents(context.Context, *empty.Empty) (*ListComponentResponse, error)
	// Returns a given configuration by name
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	// Returns a list of pub/sub subscriptions
	ListSubscriptions(context.Context, *empty.Empty) (*ListSubscriptionsResponse, error)
}

// UnimplementedOperatorServer can be embedded to have forward compatible implementations.
type UnimplementedOperatorServer struct {
}

func (*UnimplementedOperatorServer) ComponentUpdate(req *empty.Empty, srv Operator_ComponentUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method ComponentUpdate not implemented")
}
func (*UnimplementedOperatorServer) ListComponents(ctx context.Context, req *empty.Empty) (*ListComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComponents not implemented")
}
func (*UnimplementedOperatorServer) GetConfiguration(ctx context.Context, req *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (*UnimplementedOperatorServer) ListSubscriptions(ctx context.Context, req *empty.Empty) (*ListSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}

func RegisterOperatorServer(s *grpc.Server, srv OperatorServer) {
	s.RegisterService(&_Operator_serviceDesc, srv)
}

func _Operator_ComponentUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperatorServer).ComponentUpdate(m, &operatorComponentUpdateServer{stream})
}

type Operator_ComponentUpdateServer interface {
	Send(*ComponentUpdateEvent) error
	grpc.ServerStream
}

type operatorComponentUpdateServer struct {
	grpc.ServerStream
}

func (x *operatorComponentUpdateServer) Send(m *ComponentUpdateEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Operator_ListComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/ListComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListComponents(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_ListSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ListSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.operator.v1.Operator/ListSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ListSubscriptions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.operator.v1.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComponents",
			Handler:    _Operator_ListComponents_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Operator_GetConfiguration_Handler,
		},
		{
			MethodName: "ListSubscriptions",
			Handler:    _Operator_ListSubscriptions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComponentUpdate",
			Handler:       _Operator_ComponentUpdate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dapr/proto/operator/v1/operator.proto",
}