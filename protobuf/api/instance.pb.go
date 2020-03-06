// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/instance.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_cosmos_address "github.com/mesg-foundation/engine/cosmos/address"
	instance "github.com/mesg-foundation/engine/instance"
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

// The request's data for the `Get` API.
type GetInstanceRequest struct {
	Hash                 github_com_mesg_foundation_engine_cosmos_address.InstAddress `protobuf:"bytes,1,opt,name=hash,proto3,casttype=github.com/mesg-foundation/engine/cosmos/address.InstAddress" json:"hash,omitempty" validate:"required,instaddress"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *GetInstanceRequest) Reset()         { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()    {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{0}
}
func (m *GetInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstanceRequest.Unmarshal(m, b)
}
func (m *GetInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstanceRequest.Marshal(b, m, deterministic)
}
func (m *GetInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstanceRequest.Merge(m, src)
}
func (m *GetInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstanceRequest.Size(m)
}
func (m *GetInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstanceRequest proto.InternalMessageInfo

func (m *GetInstanceRequest) GetHash() github_com_mesg_foundation_engine_cosmos_address.InstAddress {
	if m != nil {
		return m.Hash
	}
	return nil
}

// The request's data for the `List` API.
type ListInstanceRequest struct {
	// Filter used to filter a list of instance.
	Filter               *ListInstanceRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListInstanceRequest) Reset()         { *m = ListInstanceRequest{} }
func (m *ListInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstanceRequest) ProtoMessage()    {}
func (*ListInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{1}
}
func (m *ListInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceRequest.Unmarshal(m, b)
}
func (m *ListInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceRequest.Marshal(b, m, deterministic)
}
func (m *ListInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceRequest.Merge(m, src)
}
func (m *ListInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_ListInstanceRequest.Size(m)
}
func (m *ListInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceRequest proto.InternalMessageInfo

func (m *ListInstanceRequest) GetFilter() *ListInstanceRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

// Filter contains filtering criteria.
type ListInstanceRequest_Filter struct {
	// Service hash to filter executions.
	ServiceHash          github_com_mesg_foundation_engine_cosmos_address.ServAddress `protobuf:"bytes,1,opt,name=serviceHash,proto3,casttype=github.com/mesg-foundation/engine/cosmos/address.ServAddress" json:"serviceHash,omitempty" validate:"omitempty,servaddress"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *ListInstanceRequest_Filter) Reset()         { *m = ListInstanceRequest_Filter{} }
func (m *ListInstanceRequest_Filter) String() string { return proto.CompactTextString(m) }
func (*ListInstanceRequest_Filter) ProtoMessage()    {}
func (*ListInstanceRequest_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{1, 0}
}
func (m *ListInstanceRequest_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceRequest_Filter.Unmarshal(m, b)
}
func (m *ListInstanceRequest_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceRequest_Filter.Marshal(b, m, deterministic)
}
func (m *ListInstanceRequest_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceRequest_Filter.Merge(m, src)
}
func (m *ListInstanceRequest_Filter) XXX_Size() int {
	return xxx_messageInfo_ListInstanceRequest_Filter.Size(m)
}
func (m *ListInstanceRequest_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceRequest_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceRequest_Filter proto.InternalMessageInfo

func (m *ListInstanceRequest_Filter) GetServiceHash() github_com_mesg_foundation_engine_cosmos_address.ServAddress {
	if m != nil {
		return m.ServiceHash
	}
	return nil
}

// The response's data for the `List` API.
type ListInstanceResponse struct {
	// List of instances that match the request's filters.
	Instances            []*instance.Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListInstanceResponse) Reset()         { *m = ListInstanceResponse{} }
func (m *ListInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstanceResponse) ProtoMessage()    {}
func (*ListInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{2}
}
func (m *ListInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstanceResponse.Unmarshal(m, b)
}
func (m *ListInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstanceResponse.Marshal(b, m, deterministic)
}
func (m *ListInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstanceResponse.Merge(m, src)
}
func (m *ListInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_ListInstanceResponse.Size(m)
}
func (m *ListInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstanceResponse proto.InternalMessageInfo

func (m *ListInstanceResponse) GetInstances() []*instance.Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
	proto.RegisterType((*GetInstanceRequest)(nil), "mesg.api.GetInstanceRequest")
	proto.RegisterType((*ListInstanceRequest)(nil), "mesg.api.ListInstanceRequest")
	proto.RegisterType((*ListInstanceRequest_Filter)(nil), "mesg.api.ListInstanceRequest.Filter")
	proto.RegisterType((*ListInstanceResponse)(nil), "mesg.api.ListInstanceResponse")
}

func init() { proto.RegisterFile("protobuf/api/instance.proto", fileDescriptor_71d44b8f4a870f63) }

var fileDescriptor_71d44b8f4a870f63 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0x65, 0x05, 0x87, 0xb8, 0xe5, 0xaa, 0x3d, 0x0a, 0xe4, 0x3b, 0xce, 0x68, 0x75, 0x05, 0x05,
	0x59, 0x4b, 0x4e, 0x97, 0xd0, 0x84, 0x22, 0x24, 0x51, 0x2a, 0xd2, 0xa5, 0x5b, 0xec, 0xc1, 0xac,
	0x84, 0xbd, 0xc6, 0xbb, 0x26, 0xe2, 0x07, 0xd2, 0x45, 0x4a, 0x7e, 0x90, 0x36, 0x3d, 0x65, 0xaa,
	0xc8, 0x6b, 0x8c, 0x93, 0x40, 0xd2, 0xa4, 0xdb, 0x79, 0xf3, 0x66, 0xe6, 0xbd, 0x99, 0xc5, 0x7f,
	0xe2, 0x44, 0x6a, 0x39, 0x49, 0xa7, 0x0e, 0x8f, 0x85, 0x23, 0x22, 0xa5, 0x79, 0xe4, 0x01, 0x33,
	0x28, 0x69, 0x84, 0xa0, 0x02, 0xc6, 0x63, 0x61, 0xd1, 0x40, 0x06, 0xd2, 0xd9, 0x71, 0xb3, 0xc8,
	0x04, 0xe6, 0x95, 0xb3, 0xad, 0xce, 0x2e, 0xad, 0x57, 0x31, 0xa8, 0x0f, 0xcd, 0xe8, 0x03, 0xc2,
	0x64, 0x04, 0xfa, 0x72, 0x8b, 0x8e, 0x61, 0x91, 0x82, 0xd2, 0xe4, 0x0e, 0xd7, 0x66, 0x5c, 0xcd,
	0xda, 0xa8, 0x8b, 0x7a, 0xbf, 0x86, 0xde, 0x66, 0x6d, 0xdb, 0x4b, 0x3e, 0x17, 0x3e, 0xd7, 0x70,
	0x42, 0x13, 0x58, 0xa4, 0x22, 0x01, 0xbf, 0x6f, 0x7a, 0xf9, 0x7e, 0x02, 0x4a, 0xd1, 0x97, 0xb5,
	0x3d, 0x08, 0x84, 0x9e, 0xa5, 0x13, 0xe6, 0xc9, 0xd0, 0xc9, 0x34, 0x1e, 0x4d, 0x65, 0x1a, 0xf9,
	0x5c, 0x0b, 0x19, 0x39, 0x10, 0x05, 0x22, 0x02, 0xc7, 0x93, 0x2a, 0x94, 0xca, 0xd9, 0x16, 0xb1,
	0x6c, 0xec, 0x59, 0xfe, 0x1e, 0x9b, 0x81, 0xf4, 0x19, 0xe1, 0xdf, 0xd7, 0x42, 0xed, 0x09, 0x1a,
	0xe0, 0xfa, 0x54, 0xcc, 0x35, 0x24, 0x46, 0x52, 0xd3, 0xfd, 0xcf, 0x8a, 0x2d, 0xb0, 0x03, 0x74,
	0x76, 0x6e, 0xb8, 0xe3, 0x6d, 0x8d, 0xf5, 0x84, 0x70, 0x3d, 0x87, 0xc8, 0x3d, 0xc2, 0x4d, 0x05,
	0xc9, 0x52, 0x78, 0x70, 0x51, 0x3a, 0xf4, 0x37, 0x6b, 0xbb, 0x5b, 0x3a, 0x94, 0xa1, 0xd0, 0x10,
	0xc6, 0x7a, 0xd5, 0xcf, 0xa8, 0xdf, 0xb2, 0x78, 0x03, 0xc9, 0xb2, 0xb0, 0xf8, 0x76, 0x30, 0xbd,
	0xc2, 0xad, 0xf7, 0xca, 0x55, 0x2c, 0x23, 0x05, 0xc4, 0xc5, 0x3f, 0x8b, 0x1b, 0xa9, 0x36, 0xea,
	0x56, 0x7b, 0x4d, 0xb7, 0x95, 0x9b, 0x35, 0x07, 0x64, 0xbb, 0x82, 0x92, 0xe6, 0x3e, 0x22, 0xdc,
	0x28, 0x70, 0x72, 0x8a, 0xab, 0x23, 0xd0, 0xe4, 0x6f, 0xb9, 0xa1, 0xfd, 0x03, 0x5b, 0x07, 0x5b,
	0xd2, 0x0a, 0x19, 0xe1, 0x5a, 0xa6, 0x8a, 0x74, 0xbe, 0xdc, 0xaf, 0xf5, 0xef, 0xb3, 0x74, 0x6e,
	0x82, 0x56, 0x86, 0x3f, 0x6e, 0xab, 0x3c, 0x16, 0x93, 0xba, 0xf9, 0x66, 0xc7, 0xaf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x14, 0x50, 0xe8, 0x4c, 0xd2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstanceClient interface {
	// Get returns an Instance matching the criteria of the request.
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*instance.Instance, error)
	// List returns all Instances matching the criteria of the request.
	List(ctx context.Context, in *ListInstanceRequest, opts ...grpc.CallOption) (*ListInstanceResponse, error)
}

type instanceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceClient(cc *grpc.ClientConn) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*instance.Instance, error) {
	out := new(instance.Instance)
	err := c.cc.Invoke(ctx, "/mesg.api.Instance/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) List(ctx context.Context, in *ListInstanceRequest, opts ...grpc.CallOption) (*ListInstanceResponse, error) {
	out := new(ListInstanceResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Instance/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
type InstanceServer interface {
	// Get returns an Instance matching the criteria of the request.
	Get(context.Context, *GetInstanceRequest) (*instance.Instance, error)
	// List returns all Instances matching the criteria of the request.
	List(context.Context, *ListInstanceRequest) (*ListInstanceResponse, error)
}

// UnimplementedInstanceServer can be embedded to have forward compatible implementations.
type UnimplementedInstanceServer struct {
}

func (*UnimplementedInstanceServer) Get(ctx context.Context, req *GetInstanceRequest) (*instance.Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInstanceServer) List(ctx context.Context, req *ListInstanceRequest) (*ListInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterInstanceServer(s *grpc.Server, srv InstanceServer) {
	s.RegisterService(&_Instance_serviceDesc, srv)
}

func _Instance_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Instance/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Instance/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).List(ctx, req.(*ListInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Instance_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Instance_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/instance.proto",
}
