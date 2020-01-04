// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MultiplyRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyRequest) Reset()         { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()    {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *MultiplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyRequest.Unmarshal(m, b)
}
func (m *MultiplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyRequest.Marshal(b, m, deterministic)
}
func (m *MultiplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyRequest.Merge(m, src)
}
func (m *MultiplyRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplyRequest.Size(m)
}
func (m *MultiplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyRequest proto.InternalMessageInfo

func (m *MultiplyRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *MultiplyRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type MultiplyReply struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyReply) Reset()         { *m = MultiplyReply{} }
func (m *MultiplyReply) String() string { return proto.CompactTextString(m) }
func (*MultiplyReply) ProtoMessage()    {}
func (*MultiplyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *MultiplyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyReply.Unmarshal(m, b)
}
func (m *MultiplyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyReply.Marshal(b, m, deterministic)
}
func (m *MultiplyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyReply.Merge(m, src)
}
func (m *MultiplyReply) XXX_Size() int {
	return xxx_messageInfo_MultiplyReply.Size(m)
}
func (m *MultiplyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyReply.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyReply proto.InternalMessageInfo

func (m *MultiplyReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*MultiplyRequest)(nil), "main.MultiplyRequest")
	proto.RegisterType((*MultiplyReply)(nil), "main.MultiplyReply")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53,
	0xd2, 0xe5, 0xe2, 0xf7, 0x2d, 0xcd, 0x29, 0xc9, 0x2c, 0xc8, 0xa9, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0xac, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xac,
	0x00, 0xf1, 0x2a, 0x25, 0x98, 0x20, 0xbc, 0x4a, 0x25, 0x75, 0x2e, 0x5e, 0x84, 0xf2, 0x82, 0x9c,
	0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x0e, 0x28, 0xcf, 0xc8,
	0x8d, 0x8b, 0x2b, 0x24, 0xb5, 0xb8, 0x24, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0xc8, 0x82, 0x8b,
	0x03, 0xa6, 0x4d, 0x48, 0x54, 0x0f, 0x64, 0xb1, 0x1e, 0x9a, 0xad, 0x52, 0xc2, 0xe8, 0xc2, 0x05,
	0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xc7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0x2e, 0xbe, 0x71, 0xbd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServerClient is the client API for TestServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServerClient interface {
	// Mutiplies two values
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error)
}

type testServerClient struct {
	cc *grpc.ClientConn
}

func NewTestServerClient(cc *grpc.ClientConn) TestServerClient {
	return &testServerClient{cc}
}

func (c *testServerClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error) {
	out := new(MultiplyReply)
	err := c.cc.Invoke(ctx, "/main.TestServer/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServerServer is the server API for TestServer service.
type TestServerServer interface {
	// Mutiplies two values
	Multiply(context.Context, *MultiplyRequest) (*MultiplyReply, error)
}

// UnimplementedTestServerServer can be embedded to have forward compatible implementations.
type UnimplementedTestServerServer struct {
}

func (*UnimplementedTestServerServer) Multiply(ctx context.Context, req *MultiplyRequest) (*MultiplyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterTestServerServer(s *grpc.Server, srv TestServerServer) {
	s.RegisterService(&_TestServer_serviceDesc, srv)
}

func _TestServer_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServerServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestServer/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServerServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.TestServer",
	HandlerType: (*TestServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _TestServer_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}