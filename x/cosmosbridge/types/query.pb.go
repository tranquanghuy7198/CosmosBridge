// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmosbridge/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cosmosbridge/query.proto", fileDescriptor_68be86a5fc051fd3) }

var fileDescriptor_68be86a5fc051fd3 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x2a, 0x29, 0x4a, 0xcc, 0x2b, 0x2c, 0x4d, 0xcc, 0x4b, 0xcf,
	0x28, 0xad, 0x34, 0x37, 0xb4, 0xb4, 0xd0, 0x43, 0x56, 0x8a, 0xc2, 0x91, 0x92, 0x49, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0x86, 0x98, 0x24, 0xa5, 0x05, 0x51, 0xab, 0x9f, 0x94, 0x58, 0x0c, 0xb5, 0x42,
	0xbf, 0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0xbf, 0x20, 0x31, 0x3d, 0x33, 0x0f, 0xac, 0x18,
	0xa2, 0xd6, 0x88, 0x9d, 0x8b, 0x35, 0x10, 0xa4, 0xc2, 0x29, 0xfc, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x6c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0xd1, 0xdd, 0xa8, 0xef, 0x0c, 0xb6, 0xca, 0x09, 0xe2, 0x9d, 0x0a, 0x7d, 0x14, 0xdf, 0x95,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x2d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xca,
	0x3c, 0xae, 0x41, 0xfa, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tranquanghuy7198.cosmosbridge.cosmosbridge.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cosmosbridge/query.proto",
}
