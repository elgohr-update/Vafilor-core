// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secret.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateSecretRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Secret               *Secret  `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSecretRequest) Reset()         { *m = CreateSecretRequest{} }
func (m *CreateSecretRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSecretRequest) ProtoMessage()    {}
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{0}
}

func (m *CreateSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretRequest.Unmarshal(m, b)
}
func (m *CreateSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretRequest.Marshal(b, m, deterministic)
}
func (m *CreateSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretRequest.Merge(m, src)
}
func (m *CreateSecretRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSecretRequest.Size(m)
}
func (m *CreateSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretRequest proto.InternalMessageInfo

func (m *CreateSecretRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateSecretRequest) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

type Secret struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{1}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSecretRequest)(nil), "api.CreateSecretRequest")
	proto.RegisterType((*Secret)(nil), "api.Secret")
	proto.RegisterMapType((map[string]string)(nil), "api.Secret.DataEntry")
}

func init() { proto.RegisterFile("secret.proto", fileDescriptor_6acf428160d7a216) }

var fileDescriptor_6acf428160d7a216 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0xc9, 0xb6, 0x16, 0x76, 0x5a, 0x41, 0xe2, 0x1f, 0x96, 0xb5, 0x87, 0xba, 0x5e, 0x2a,
	0x42, 0x42, 0xeb, 0x41, 0xe9, 0x55, 0xfb, 0x02, 0xdb, 0x8b, 0xd7, 0x69, 0x1d, 0xcb, 0x62, 0xbb,
	0x89, 0x9b, 0xb4, 0x50, 0x44, 0x10, 0xf1, 0x0d, 0x7c, 0x34, 0x5f, 0xc1, 0x07, 0x91, 0x4d, 0x62,
	0xed, 0xc1, 0xdb, 0xe4, 0x37, 0x93, 0x2f, 0x93, 0x0f, 0x3a, 0x86, 0x66, 0x15, 0x59, 0xa1, 0x2b,
	0x65, 0x15, 0x6f, 0xa0, 0x2e, 0xd2, 0xd3, 0xb9, 0x52, 0xf3, 0x05, 0x49, 0x17, 0x4d, 0x57, 0x8f,
	0x92, 0x96, 0xda, 0x6e, 0xfc, 0x44, 0xda, 0x0d, 0x4d, 0xd4, 0x85, 0xc4, 0xb2, 0x54, 0x16, 0x6d,
	0xa1, 0x4a, 0xe3, 0xbb, 0xd9, 0x3d, 0x1c, 0xde, 0x56, 0x84, 0x96, 0x26, 0x8e, 0x9a, 0xd3, 0xf3,
	0x8a, 0x8c, 0xe5, 0x5d, 0x88, 0x4b, 0x5c, 0x92, 0xd1, 0x38, 0xa3, 0x84, 0xf5, 0x58, 0x3f, 0xce,
	0xff, 0x02, 0x7e, 0x0e, 0x2d, 0xbf, 0x44, 0x12, 0xf5, 0x58, 0xbf, 0x3d, 0x6c, 0x0b, 0xd4, 0x85,
	0x08, 0x84, 0xd0, 0xca, 0xde, 0x18, 0xb4, 0x7c, 0xc4, 0x39, 0x34, 0xeb, 0xcb, 0x01, 0xe4, 0x6a,
	0x7e, 0x01, 0xcd, 0x07, 0xb4, 0x98, 0x44, 0xbd, 0x46, 0xbf, 0x3d, 0x3c, 0xde, 0x21, 0x88, 0x3b,
	0xb4, 0x38, 0x2e, 0x6d, 0xb5, 0xc9, 0xdd, 0x48, 0x7a, 0x0d, 0xf1, 0x36, 0xe2, 0x07, 0xd0, 0x78,
	0xa2, 0x4d, 0x40, 0xd5, 0x25, 0x3f, 0x82, 0xbd, 0x35, 0x2e, 0x56, 0xe4, 0x96, 0x89, 0x73, 0x7f,
	0x18, 0x45, 0x37, 0x6c, 0xf8, 0xc1, 0x60, 0xdf, 0x33, 0x27, 0x54, 0xad, 0x8b, 0x19, 0x71, 0x03,
	0x9d, 0xdd, 0xef, 0xf2, 0xc4, 0xbd, 0xfb, 0x8f, 0x81, 0xf4, 0x44, 0x78, 0x6f, 0xe2, 0x57, 0xaa,
	0x18, 0xd7, 0x52, 0xb3, 0xc1, 0xfb, 0xd7, 0xf7, 0x67, 0x74, 0x99, 0x9d, 0xd5, 0x42, 0x8d, 0x5c,
	0x0f, 0xa6, 0x64, 0x71, 0x20, 0x5f, 0xb6, 0x76, 0x5e, 0xa5, 0x17, 0x60, 0x46, 0xc1, 0xc4, 0xb4,
	0xe5, 0x10, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x92, 0x3d, 0x4a, 0xba, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretServiceClient interface {
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type secretServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretServiceClient(cc *grpc.ClientConn) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.SecretService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
type SecretServiceServer interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*empty.Empty, error)
}

// UnimplementedSecretServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (*UnimplementedSecretServiceServer) CreateSecret(ctx context.Context, req *CreateSecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}

func RegisterSecretServiceServer(s *grpc.Server, srv SecretServiceServer) {
	s.RegisterService(&_SecretService_serviceDesc, srv)
}

func _SecretService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SecretService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _SecretService_CreateSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret.proto",
}
