// Code generated by protoc-gen-go. DO NOT EDIT.
// source: math.proto

package proto

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

type Request struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Response struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f139a3799a86a974, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("math.proto", fileDescriptor_f139a3799a86a974) }

var fileDescriptor_f139a3799a86a974 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c, 0xec, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0x79, 0xa5, 0xb9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x20, 0xa6, 0x92, 0x12, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x54, 0x01, 0x94, 0x67, 0x64,
	0xc2, 0xc5, 0xe2, 0x9b, 0x58, 0x92, 0x21, 0xa4, 0xc3, 0xc5, 0xec, 0x9b, 0x58, 0x21, 0xc4, 0x07,
	0x31, 0x5e, 0x0f, 0x6a, 0xa8, 0x14, 0x3f, 0x9c, 0x0f, 0x31, 0x47, 0x89, 0x41, 0x83, 0xd1, 0x80,
	0x31, 0x89, 0x0d, 0x2c, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xb7, 0x4c, 0x3f, 0x92,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MathClient is the client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathClient interface {
	Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error)
}

type mathClient struct {
	cc *grpc.ClientConn
}

func NewMathClient(cc *grpc.ClientConn) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Math_serviceDesc.Streams[0], "/proto.Math/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathMaxClient{stream}
	return x, nil
}

type Math_MaxClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type mathMaxClient struct {
	grpc.ClientStream
}

func (x *mathMaxClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mathMaxClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MathServer is the server API for Math service.
type MathServer interface {
	Max(Math_MaxServer) error
}

// UnimplementedMathServer can be embedded to have forward compatible implementations.
type UnimplementedMathServer struct {
}

func (*UnimplementedMathServer) Max(srv Math_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MathServer).Max(&mathMaxServer{stream})
}

type Math_MaxServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type mathMaxServer struct {
	grpc.ServerStream
}

func (x *mathMaxServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mathMaxServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Math",
	HandlerType: (*MathServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Max",
			Handler:       _Math_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "math.proto",
}
