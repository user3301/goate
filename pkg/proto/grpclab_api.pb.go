// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpclab_api.proto

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

type LoginResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdb507d7f20fe0, []int{0}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RegisterResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdb507d7f20fe0, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginResponse)(nil), "grpclab.LoginResponse")
	proto.RegisterType((*RegisterResponse)(nil), "grpclab.RegisterResponse")
}

func init() {
	proto.RegisterFile("grpclab_api.proto", fileDescriptor_01fdb507d7f20fe0)
}

var fileDescriptor_01fdb507d7f20fe0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0xce, 0x49, 0x4c, 0x8a, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87,
	0x0a, 0x49, 0x09, 0xc3, 0xe4, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x21, 0xb2, 0x4a, 0x9a, 0x5c, 0xbc,
	0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c,
	0xec, 0xc5, 0xa5, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30,
	0xae, 0x92, 0x0b, 0x97, 0x40, 0x50, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x11, 0x61, 0xd5, 0x42,
	0x62, 0x5c, 0x6c, 0x45, 0xa9, 0x89, 0xc5, 0xf9, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x50, 0x9e, 0x51, 0x3b, 0x23, 0x97, 0xa0, 0x7b, 0x50, 0x80, 0xb3, 0x4f, 0x62, 0x92, 0x63, 0x80,
	0x67, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x35, 0x17, 0x07, 0xcc, 0x6c, 0x21, 0x11,
	0x3d, 0xa8, 0x43, 0xf5, 0x9c, 0x8b, 0x52, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0x13, 0x73, 0x8a, 0xa5,
	0x24, 0xe1, 0xa2, 0x18, 0x8e, 0x30, 0xe5, 0x62, 0x05, 0xfb, 0x01, 0x87, 0x4e, 0x31, 0xb8, 0x28,
	0x8a, 0x4f, 0x9d, 0x38, 0xa2, 0xd8, 0xf4, 0xc1, 0x81, 0x90, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x1f, 0x4c, 0x89, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GRPCLabAPIServiceClient is the client API for GRPCLabAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCLabAPIServiceClient interface {
	Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResponse, error)
}

type gRPCLabAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCLabAPIServiceClient(cc grpc.ClientConnInterface) GRPCLabAPIServiceClient {
	return &gRPCLabAPIServiceClient{cc}
}

func (c *gRPCLabAPIServiceClient) Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/grpclab.GRPCLabAPIService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCLabAPIServiceClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/grpclab.GRPCLabAPIService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCLabAPIServiceServer is the server API for GRPCLabAPIService service.
type GRPCLabAPIServiceServer interface {
	Register(context.Context, *Credentials) (*RegisterResponse, error)
	Login(context.Context, *Credentials) (*LoginResponse, error)
}

// UnimplementedGRPCLabAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCLabAPIServiceServer struct {
}

func (*UnimplementedGRPCLabAPIServiceServer) Register(ctx context.Context, req *Credentials) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedGRPCLabAPIServiceServer) Login(ctx context.Context, req *Credentials) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterGRPCLabAPIServiceServer(s *grpc.Server, srv GRPCLabAPIServiceServer) {
	s.RegisterService(&_GRPCLabAPIService_serviceDesc, srv)
}

func _GRPCLabAPIService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCLabAPIServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpclab.GRPCLabAPIService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCLabAPIServiceServer).Register(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCLabAPIService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCLabAPIServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpclab.GRPCLabAPIService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCLabAPIServiceServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCLabAPIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpclab.GRPCLabAPIService",
	HandlerType: (*GRPCLabAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GRPCLabAPIService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GRPCLabAPIService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpclab_api.proto",
}
