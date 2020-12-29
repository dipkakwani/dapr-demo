// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/internals/v1/apiversion.proto

package internals

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// APIVersion represents the version of Dapr Runtime API.
type APIVersion int32

const (
	// unspecified apiversion
	APIVersion_APIVERSION_UNSPECIFIED APIVersion = 0
	// Dapr API v1
	APIVersion_V1 APIVersion = 1
)

var APIVersion_name = map[int32]string{
	0: "APIVERSION_UNSPECIFIED",
	1: "V1",
}

var APIVersion_value = map[string]int32{
	"APIVERSION_UNSPECIFIED": 0,
	"V1":                     1,
}

func (x APIVersion) String() string {
	return proto.EnumName(APIVersion_name, int32(x))
}

func (APIVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67e9e4830bd33413, []int{0}
}

func init() {
	proto.RegisterEnum("dapr.proto.internals.v1.APIVersion", APIVersion_name, APIVersion_value)
}

func init() {
	proto.RegisterFile("dapr/proto/internals/v1/apiversion.proto", fileDescriptor_67e9e4830bd33413)
}

var fileDescriptor_67e9e4830bd33413 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x49, 0x2c, 0x28,
	0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x29,
	0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2c, 0xc8, 0x2c, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x03,
	0xcb, 0x0a, 0x89, 0x83, 0x54, 0x42, 0xd8, 0x7a, 0x70, 0x95, 0x7a, 0x65, 0x86, 0x5a, 0x06, 0x5c,
	0x5c, 0x8e, 0x01, 0x9e, 0x61, 0x10, 0xc5, 0x42, 0x52, 0x5c, 0x62, 0x20, 0x9e, 0x6b, 0x50, 0xb0,
	0xa7, 0xbf, 0x5f, 0x7c, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00,
	0x83, 0x10, 0x1b, 0x17, 0x53, 0x98, 0xa1, 0x00, 0xa3, 0x93, 0x79, 0x94, 0x69, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd8, 0x05, 0x10, 0x67, 0x64, 0xa7, 0x63, 0x71,
	0x8a, 0x35, 0x9c, 0x93, 0xc4, 0x06, 0x96, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x2b,
	0x48, 0xe4, 0xb6, 0x00, 0x00, 0x00,
}
