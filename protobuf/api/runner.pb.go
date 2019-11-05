// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/runner.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	runner "github.com/mesg-foundation/engine/runner"
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
type GetRunnerRequest struct {
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *GetRunnerRequest) Reset()         { *m = GetRunnerRequest{} }
func (m *GetRunnerRequest) String() string { return proto.CompactTextString(m) }
func (*GetRunnerRequest) ProtoMessage()    {}
func (*GetRunnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{0}
}
func (m *GetRunnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunnerRequest.Unmarshal(m, b)
}
func (m *GetRunnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunnerRequest.Marshal(b, m, deterministic)
}
func (m *GetRunnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunnerRequest.Merge(m, src)
}
func (m *GetRunnerRequest) XXX_Size() int {
	return xxx_messageInfo_GetRunnerRequest.Size(m)
}
func (m *GetRunnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRunnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRunnerRequest proto.InternalMessageInfo

// The request's data for the `List` API.
type ListRunnerRequest struct {
	// Filter by Instance hash.
	InstanceHash         github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=instanceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ListRunnerRequest) Reset()         { *m = ListRunnerRequest{} }
func (m *ListRunnerRequest) String() string { return proto.CompactTextString(m) }
func (*ListRunnerRequest) ProtoMessage()    {}
func (*ListRunnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{1}
}
func (m *ListRunnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunnerRequest.Unmarshal(m, b)
}
func (m *ListRunnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunnerRequest.Marshal(b, m, deterministic)
}
func (m *ListRunnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunnerRequest.Merge(m, src)
}
func (m *ListRunnerRequest) XXX_Size() int {
	return xxx_messageInfo_ListRunnerRequest.Size(m)
}
func (m *ListRunnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunnerRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type ListRunnerResponse struct {
	// List of runners that match the request's filters.
	Runners              []*runner.Runner `protobuf:"bytes,1,rep,name=runners,proto3" json:"runners,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListRunnerResponse) Reset()         { *m = ListRunnerResponse{} }
func (m *ListRunnerResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunnerResponse) ProtoMessage()    {}
func (*ListRunnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{2}
}
func (m *ListRunnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunnerResponse.Unmarshal(m, b)
}
func (m *ListRunnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunnerResponse.Marshal(b, m, deterministic)
}
func (m *ListRunnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunnerResponse.Merge(m, src)
}
func (m *ListRunnerResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunnerResponse.Size(m)
}
func (m *ListRunnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunnerResponse proto.InternalMessageInfo

func (m *ListRunnerResponse) GetRunners() []*runner.Runner {
	if m != nil {
		return m.Runners
	}
	return nil
}

// The request's data for the `Create` API.
type CreateRunnerRequest struct {
	// Service's hash to start the runner with.
	ServiceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=serviceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"serviceHash"`
	// Environmental variables to start the runner with.
	Env                  []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRunnerRequest) Reset()         { *m = CreateRunnerRequest{} }
func (m *CreateRunnerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRunnerRequest) ProtoMessage()    {}
func (*CreateRunnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{3}
}
func (m *CreateRunnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRunnerRequest.Unmarshal(m, b)
}
func (m *CreateRunnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRunnerRequest.Marshal(b, m, deterministic)
}
func (m *CreateRunnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRunnerRequest.Merge(m, src)
}
func (m *CreateRunnerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRunnerRequest.Size(m)
}
func (m *CreateRunnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRunnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRunnerRequest proto.InternalMessageInfo

func (m *CreateRunnerRequest) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// The response's data for the `Create` API.
type CreateRunnerResponse struct {
	// The runner's hash created.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CreateRunnerResponse) Reset()         { *m = CreateRunnerResponse{} }
func (m *CreateRunnerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRunnerResponse) ProtoMessage()    {}
func (*CreateRunnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{4}
}
func (m *CreateRunnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRunnerResponse.Unmarshal(m, b)
}
func (m *CreateRunnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRunnerResponse.Marshal(b, m, deterministic)
}
func (m *CreateRunnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRunnerResponse.Merge(m, src)
}
func (m *CreateRunnerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRunnerResponse.Size(m)
}
func (m *CreateRunnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRunnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRunnerResponse proto.InternalMessageInfo

// The request's data for the `Delete` API.
type DeleteRunnerRequest struct {
	// Runner's hash
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	// If true, any persistent data (volumes) that belongs to the runner's container and its dependencies will also be deleted.
	DeleteData           bool     `protobuf:"varint,2,opt,name=deleteData,proto3" json:"deleteData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRunnerRequest) Reset()         { *m = DeleteRunnerRequest{} }
func (m *DeleteRunnerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRunnerRequest) ProtoMessage()    {}
func (*DeleteRunnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{5}
}
func (m *DeleteRunnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRunnerRequest.Unmarshal(m, b)
}
func (m *DeleteRunnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRunnerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRunnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRunnerRequest.Merge(m, src)
}
func (m *DeleteRunnerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRunnerRequest.Size(m)
}
func (m *DeleteRunnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRunnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRunnerRequest proto.InternalMessageInfo

func (m *DeleteRunnerRequest) GetDeleteData() bool {
	if m != nil {
		return m.DeleteData
	}
	return false
}

// The response's data for the `Delete` API.
type DeleteRunnerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRunnerResponse) Reset()         { *m = DeleteRunnerResponse{} }
func (m *DeleteRunnerResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRunnerResponse) ProtoMessage()    {}
func (*DeleteRunnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66b76d1ada2874b2, []int{6}
}
func (m *DeleteRunnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRunnerResponse.Unmarshal(m, b)
}
func (m *DeleteRunnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRunnerResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRunnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRunnerResponse.Merge(m, src)
}
func (m *DeleteRunnerResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRunnerResponse.Size(m)
}
func (m *DeleteRunnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRunnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRunnerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetRunnerRequest)(nil), "mesg.api.GetRunnerRequest")
	proto.RegisterType((*ListRunnerRequest)(nil), "mesg.api.ListRunnerRequest")
	proto.RegisterType((*ListRunnerResponse)(nil), "mesg.api.ListRunnerResponse")
	proto.RegisterType((*CreateRunnerRequest)(nil), "mesg.api.CreateRunnerRequest")
	proto.RegisterType((*CreateRunnerResponse)(nil), "mesg.api.CreateRunnerResponse")
	proto.RegisterType((*DeleteRunnerRequest)(nil), "mesg.api.DeleteRunnerRequest")
	proto.RegisterType((*DeleteRunnerResponse)(nil), "mesg.api.DeleteRunnerResponse")
}

func init() { proto.RegisterFile("protobuf/api/runner.proto", fileDescriptor_66b76d1ada2874b2) }

var fileDescriptor_66b76d1ada2874b2 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xed, 0xb8, 0x84, 0x32, 0xed, 0xa1, 0x6c, 0x2b, 0x64, 0x5c, 0x68, 0xad, 0x3d, 0x59,
	0x02, 0xd6, 0x52, 0x7b, 0xe0, 0x9e, 0x56, 0x4a, 0x91, 0x38, 0x59, 0x42, 0x48, 0x70, 0x40, 0x9b,
	0x66, 0x62, 0xaf, 0x94, 0xec, 0x1a, 0xef, 0x3a, 0x12, 0x97, 0x3c, 0x09, 0x0f, 0xc4, 0x33, 0x70,
	0xc8, 0xb3, 0x20, 0xaf, 0xf3, 0x61, 0x07, 0xf7, 0x94, 0xdc, 0x76, 0x67, 0x66, 0x7f, 0xf3, 0xf5,
	0x5f, 0x78, 0x9d, 0x17, 0xca, 0xa8, 0x51, 0x39, 0x89, 0x79, 0x2e, 0xe2, 0xa2, 0x94, 0x12, 0x0b,
	0x66, 0x6d, 0xe4, 0x78, 0x86, 0x3a, 0x65, 0x3c, 0x17, 0x01, 0x4d, 0x55, 0xaa, 0xe2, 0x4d, 0x64,
	0x75, 0xb3, 0x17, 0x7b, 0xaa, 0xa3, 0x83, 0xcb, 0x8d, 0xdb, 0xfc, 0xca, 0x51, 0xb7, 0x50, 0xf4,
	0x3b, 0x9c, 0x0d, 0xd1, 0x24, 0xd6, 0x94, 0xe0, 0xcf, 0x12, 0xb5, 0x21, 0x43, 0x38, 0xca, 0xb8,
	0xce, 0x7c, 0x37, 0x74, 0xa3, 0xd3, 0xc1, 0xed, 0x9f, 0xe5, 0xb5, 0xf3, 0x77, 0x79, 0xfd, 0x2e,
	0x15, 0x26, 0x2b, 0x47, 0xec, 0x51, 0xcd, 0xe2, 0x2a, 0xff, 0x87, 0x89, 0x2a, 0xe5, 0x98, 0x1b,
	0xa1, 0x64, 0x8c, 0x32, 0x15, 0x12, 0xe3, 0xea, 0x15, 0x7b, 0xe0, 0x3a, 0x4b, 0x2c, 0x80, 0x4e,
	0xe1, 0xe5, 0x67, 0xa1, 0x77, 0xe8, 0x5f, 0xe1, 0x54, 0x48, 0x6d, 0xb8, 0x7c, 0xc4, 0x87, 0x3d,
	0xb3, 0xb4, 0x40, 0x74, 0x00, 0xa4, 0x99, 0x4d, 0xe7, 0x4a, 0x6a, 0x24, 0xef, 0xe1, 0x79, 0xdd,
	0xb0, 0xf6, 0xdd, 0xd0, 0x8b, 0x4e, 0x6e, 0x08, 0xb3, 0xd3, 0xb3, 0xb3, 0x60, 0xab, 0xe0, 0x75,
	0x08, 0x5d, 0xc0, 0xf9, 0x5d, 0x81, 0xdc, 0x60, 0xbb, 0xe6, 0x2f, 0x70, 0xa2, 0xb1, 0x98, 0x8b,
	0xfd, 0x4b, 0x6e, 0x72, 0xc8, 0x19, 0x78, 0x28, 0xe7, 0x7e, 0x2f, 0xf4, 0xa2, 0x17, 0x49, 0x75,
	0xa4, 0x3f, 0xe0, 0xa2, 0x9d, 0x7f, 0xd5, 0xc5, 0xc1, 0x56, 0xb2, 0x80, 0xf3, 0x7b, 0x9c, 0xe2,
	0x6e, 0x83, 0x87, 0xe2, 0x93, 0x2b, 0x80, 0xb1, 0xe5, 0xdf, 0x73, 0xc3, 0xfd, 0x5e, 0xe8, 0x46,
	0xc7, 0x49, 0xc3, 0x42, 0x5f, 0xc1, 0x45, 0x3b, 0x7f, 0xdd, 0xe0, 0xcd, 0xef, 0x1e, 0xf4, 0x6b,
	0x13, 0xf9, 0x08, 0xde, 0x10, 0x0d, 0x09, 0xd8, 0x5a, 0xe5, 0x6c, 0x57, 0xa1, 0x41, 0xc7, 0x0e,
	0xa9, 0x43, 0xee, 0xe0, 0xa8, 0x12, 0x00, 0xb9, 0xdc, 0xbe, 0xfc, 0x4f, 0x7e, 0xc1, 0x9b, 0x6e,
	0x67, 0x5d, 0x06, 0x75, 0xc8, 0x27, 0xe8, 0xd7, 0x1b, 0x20, 0x6f, 0xb7, 0x91, 0x1d, 0x9a, 0x08,
	0xae, 0x9e, 0x72, 0x37, 0x51, 0x75, 0xaf, 0x4d, 0x54, 0xc7, 0xf4, 0x9b, 0xa8, 0xae, 0xe1, 0x50,
	0x67, 0xf0, 0xec, 0x9b, 0xc7, 0x73, 0x31, 0xea, 0xdb, 0x4f, 0x7b, 0xfb, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x8a, 0x86, 0x3a, 0x1c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunnerClient interface {
	// Get returns an Runner matching the criteria of the request.
	Get(ctx context.Context, in *GetRunnerRequest, opts ...grpc.CallOption) (*runner.Runner, error)
	// List returns all Runners matching the criteria of the request.
	List(ctx context.Context, in *ListRunnerRequest, opts ...grpc.CallOption) (*ListRunnerResponse, error)
	// Create an Runner from a Service's hash and custom environmental variables.
	// It will return an unique identifier to identify the runner.
	Create(ctx context.Context, in *CreateRunnerRequest, opts ...grpc.CallOption) (*CreateRunnerResponse, error)
	// Delete an Runner.
	Delete(ctx context.Context, in *DeleteRunnerRequest, opts ...grpc.CallOption) (*DeleteRunnerResponse, error)
}

type runnerClient struct {
	cc *grpc.ClientConn
}

func NewRunnerClient(cc *grpc.ClientConn) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) Get(ctx context.Context, in *GetRunnerRequest, opts ...grpc.CallOption) (*runner.Runner, error) {
	out := new(runner.Runner)
	err := c.cc.Invoke(ctx, "/mesg.api.Runner/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) List(ctx context.Context, in *ListRunnerRequest, opts ...grpc.CallOption) (*ListRunnerResponse, error) {
	out := new(ListRunnerResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Runner/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Create(ctx context.Context, in *CreateRunnerRequest, opts ...grpc.CallOption) (*CreateRunnerResponse, error) {
	out := new(CreateRunnerResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Runner/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Delete(ctx context.Context, in *DeleteRunnerRequest, opts ...grpc.CallOption) (*DeleteRunnerResponse, error) {
	out := new(DeleteRunnerResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Runner/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServer is the server API for Runner service.
type RunnerServer interface {
	// Get returns an Runner matching the criteria of the request.
	Get(context.Context, *GetRunnerRequest) (*runner.Runner, error)
	// List returns all Runners matching the criteria of the request.
	List(context.Context, *ListRunnerRequest) (*ListRunnerResponse, error)
	// Create an Runner from a Service's hash and custom environmental variables.
	// It will return an unique identifier to identify the runner.
	Create(context.Context, *CreateRunnerRequest) (*CreateRunnerResponse, error)
	// Delete an Runner.
	Delete(context.Context, *DeleteRunnerRequest) (*DeleteRunnerResponse, error)
}

// UnimplementedRunnerServer can be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (*UnimplementedRunnerServer) Get(ctx context.Context, req *GetRunnerRequest) (*runner.Runner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRunnerServer) List(ctx context.Context, req *ListRunnerRequest) (*ListRunnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedRunnerServer) Create(ctx context.Context, req *CreateRunnerRequest) (*CreateRunnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRunnerServer) Delete(ctx context.Context, req *DeleteRunnerRequest) (*DeleteRunnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRunnerServer(s *grpc.Server, srv RunnerServer) {
	s.RegisterService(&_Runner_serviceDesc, srv)
}

func _Runner_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Runner/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Get(ctx, req.(*GetRunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Runner/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).List(ctx, req.(*ListRunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Runner/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Create(ctx, req.(*CreateRunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRunnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Runner/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Delete(ctx, req.(*DeleteRunnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Runner_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Runner_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Runner_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Runner_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/runner.proto",
}
