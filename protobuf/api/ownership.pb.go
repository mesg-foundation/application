// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/ownership.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	ownership "github.com/mesg-foundation/engine/ownership"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The request's data for the `List` API.
type OwnershipServiceListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnershipServiceListRequest) Reset()         { *m = OwnershipServiceListRequest{} }
func (m *OwnershipServiceListRequest) String() string { return proto.CompactTextString(m) }
func (*OwnershipServiceListRequest) ProtoMessage()    {}
func (*OwnershipServiceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3de3887ff6b0398, []int{0}
}
func (m *OwnershipServiceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnershipServiceListRequest.Unmarshal(m, b)
}
func (m *OwnershipServiceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnershipServiceListRequest.Marshal(b, m, deterministic)
}
func (m *OwnershipServiceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipServiceListRequest.Merge(m, src)
}
func (m *OwnershipServiceListRequest) XXX_Size() int {
	return xxx_messageInfo_OwnershipServiceListRequest.Size(m)
}
func (m *OwnershipServiceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipServiceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipServiceListRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type OwnershipServiceListResponse struct {
	// List of ownerships that match the request's filters.
	Ownerships           []*ownership.Ownership `protobuf:"bytes,1,rep,name=ownerships,proto3" json:"ownerships,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OwnershipServiceListResponse) Reset()         { *m = OwnershipServiceListResponse{} }
func (m *OwnershipServiceListResponse) String() string { return proto.CompactTextString(m) }
func (*OwnershipServiceListResponse) ProtoMessage()    {}
func (*OwnershipServiceListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3de3887ff6b0398, []int{1}
}
func (m *OwnershipServiceListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnershipServiceListResponse.Unmarshal(m, b)
}
func (m *OwnershipServiceListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnershipServiceListResponse.Marshal(b, m, deterministic)
}
func (m *OwnershipServiceListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipServiceListResponse.Merge(m, src)
}
func (m *OwnershipServiceListResponse) XXX_Size() int {
	return xxx_messageInfo_OwnershipServiceListResponse.Size(m)
}
func (m *OwnershipServiceListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipServiceListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipServiceListResponse proto.InternalMessageInfo

func (m *OwnershipServiceListResponse) GetOwnerships() []*ownership.Ownership {
	if m != nil {
		return m.Ownerships
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnershipServiceListRequest)(nil), "mesg.api.v1.OwnershipServiceListRequest")
	proto.RegisterType((*OwnershipServiceListResponse)(nil), "mesg.api.v1.OwnershipServiceListResponse")
}

func init() { proto.RegisterFile("protobuf/api/ownership.proto", fileDescriptor_e3de3887ff6b0398) }

var fileDescriptor_e3de3887ff6b0398 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x2f, 0xcf, 0x4b, 0x2d, 0x2a, 0xce,
	0xc8, 0x2c, 0xd0, 0x03, 0x0b, 0x0b, 0x71, 0xe7, 0xa6, 0x16, 0xa7, 0xeb, 0x25, 0x16, 0x64, 0xea,
	0x95, 0x19, 0x4a, 0x29, 0xa5, 0xe7, 0xa7, 0xe7, 0xeb, 0xc3, 0xd5, 0x83, 0x78, 0x60, 0x0e, 0x98,
	0x05, 0xd1, 0x20, 0x25, 0x07, 0x97, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x46, 0x37, 0x50, 0x49, 0x96,
	0x4b, 0xda, 0x1f, 0x26, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0xea, 0x93, 0x59, 0x5c, 0x12,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x14, 0xc1, 0x25, 0x83, 0x5d, 0xba, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0xc8, 0x82, 0x8b, 0x0b, 0x6e, 0x62, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7,
	0x91, 0x84, 0x1e, 0xd8, 0x91, 0x60, 0xfb, 0xf4, 0xca, 0x0c, 0xf5, 0xe0, 0x06, 0x04, 0x21, 0xa9,
	0x35, 0x2a, 0xe4, 0x12, 0x40, 0x37, 0x59, 0x28, 0x96, 0x8b, 0x05, 0x64, 0xba, 0x90, 0x86, 0x1e,
	0x92, 0x37, 0xf5, 0xf0, 0xb8, 0x4f, 0x4a, 0x93, 0x08, 0x95, 0x10, 0xa7, 0x2a, 0x31, 0x38, 0xb1,
	0x46, 0x31, 0x27, 0x16, 0x64, 0x26, 0xb1, 0x81, 0x7d, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x15, 0x2a, 0x15, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OwnershipServiceClient is the client API for OwnershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OwnershipServiceClient interface {
	// List returns ownerships specified in a request.
	List(ctx context.Context, in *OwnershipServiceListRequest, opts ...grpc.CallOption) (*OwnershipServiceListResponse, error)
}

type ownershipServiceClient struct {
	cc *grpc.ClientConn
}

func NewOwnershipServiceClient(cc *grpc.ClientConn) OwnershipServiceClient {
	return &ownershipServiceClient{cc}
}

func (c *ownershipServiceClient) List(ctx context.Context, in *OwnershipServiceListRequest, opts ...grpc.CallOption) (*OwnershipServiceListResponse, error) {
	out := new(OwnershipServiceListResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.OwnershipService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnershipServiceServer is the server API for OwnershipService service.
type OwnershipServiceServer interface {
	// List returns ownerships specified in a request.
	List(context.Context, *OwnershipServiceListRequest) (*OwnershipServiceListResponse, error)
}

// UnimplementedOwnershipServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOwnershipServiceServer struct {
}

func (*UnimplementedOwnershipServiceServer) List(ctx context.Context, req *OwnershipServiceListRequest) (*OwnershipServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterOwnershipServiceServer(s *grpc.Server, srv OwnershipServiceServer) {
	s.RegisterService(&_OwnershipService_serviceDesc, srv)
}

func _OwnershipService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnershipServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnershipServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.OwnershipService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnershipServiceServer).List(ctx, req.(*OwnershipServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OwnershipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.v1.OwnershipService",
	HandlerType: (*OwnershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _OwnershipService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/ownership.proto",
}
