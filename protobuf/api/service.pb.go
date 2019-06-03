// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/service.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import definition "github.com/mesg-foundation/core/protobuf/definition"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateServiceRequest struct {
	Definition           *definition.Service `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_cdb2d5a31910dd99, []int{0}
}
func (m *CreateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceRequest.Unmarshal(m, b)
}
func (m *CreateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceRequest.Marshal(b, m, deterministic)
}
func (dst *CreateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceRequest.Merge(dst, src)
}
func (m *CreateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceRequest.Size(m)
}
func (m *CreateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceRequest proto.InternalMessageInfo

func (m *CreateServiceRequest) GetDefinition() *definition.Service {
	if m != nil {
		return m.Definition
	}
	return nil
}

type CreateServiceResponse struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceResponse) Reset()         { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()    {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_cdb2d5a31910dd99, []int{1}
}
func (m *CreateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceResponse.Unmarshal(m, b)
}
func (m *CreateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceResponse.Marshal(b, m, deterministic)
}
func (dst *CreateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceResponse.Merge(dst, src)
}
func (m *CreateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceResponse.Size(m)
}
func (m *CreateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceResponse proto.InternalMessageInfo

func (m *CreateServiceResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *CreateServiceResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateServiceRequest)(nil), "api.CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "api.CreateServiceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceXClient is the client API for ServiceX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceXClient interface {
	Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
}

type serviceXClient struct {
	cc *grpc.ClientConn
}

func NewServiceXClient(cc *grpc.ClientConn) ServiceXClient {
	return &serviceXClient{cc}
}

func (c *serviceXClient) Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/api.ServiceX/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceXServer is the server API for ServiceX service.
type ServiceXServer interface {
	Create(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
}

func RegisterServiceXServer(s *grpc.Server, srv ServiceXServer) {
	s.RegisterService(&_ServiceX_serviceDesc, srv)
}

func _ServiceX_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceXServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceX/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceXServer).Create(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceX_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServiceX",
	HandlerType: (*ServiceXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceX_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/service.proto",
}

func init() { proto.RegisterFile("protobuf/api/service.proto", fileDescriptor_service_cdb2d5a31910dd99) }

var fileDescriptor_service_cdb2d5a31910dd99 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x03, 0x0b, 0x0a, 0x31, 0x27, 0x16, 0x64, 0x4a, 0x29, 0xc2, 0x15, 0xa4, 0xa4, 0xa6, 0x65,
	0xe6, 0x65, 0x96, 0x64, 0xe6, 0xe7, 0xa1, 0xaa, 0x53, 0xf2, 0xe6, 0x12, 0x71, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x0d, 0x86, 0x08, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x73, 0x71,
	0x21, 0xf4, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xeb, 0x21, 0x84, 0xf4, 0x60, 0xea,
	0x91, 0x94, 0x29, 0xd9, 0x72, 0x89, 0xa2, 0x19, 0x56, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24,
	0xc0, 0xc5, 0x5c, 0x9c, 0x99, 0x02, 0x36, 0x86, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe2, 0x62, 0xc9,
	0x48, 0x2c, 0xce, 0x90, 0x60, 0x02, 0x0b, 0x81, 0xd9, 0x46, 0xbe, 0x5c, 0x1c, 0x50, 0x8d, 0x11,
	0x42, 0x8e, 0x5c, 0x6c, 0x10, 0xa3, 0x84, 0x24, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xb0, 0x39, 0x52,
	0x4a, 0x0a, 0x9b, 0x14, 0xc4, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x0f, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0xb6, 0x9f, 0xac, 0x27, 0x01, 0x00, 0x00,
}