// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/event.proto

package api

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	event "github.com/mesg-foundation/engine/event"
	types "github.com/mesg-foundation/engine/protobuf/types"
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

// StreamEventRequest defines request to retrieve a stream of events.
type StreamEventRequest struct {
	// Filter used to filter a stream of events.
	Filter               *StreamEventRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StreamEventRequest) Reset()         { *m = StreamEventRequest{} }
func (m *StreamEventRequest) String() string { return proto.CompactTextString(m) }
func (*StreamEventRequest) ProtoMessage()    {}
func (*StreamEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ad091fa003d2ac, []int{0}
}
func (m *StreamEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventRequest.Unmarshal(m, b)
}
func (m *StreamEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventRequest.Marshal(b, m, deterministic)
}
func (m *StreamEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventRequest.Merge(m, src)
}
func (m *StreamEventRequest) XXX_Size() int {
	return xxx_messageInfo_StreamEventRequest.Size(m)
}
func (m *StreamEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventRequest proto.InternalMessageInfo

func (m *StreamEventRequest) GetFilter() *StreamEventRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

// Filter contains filtering criteria.
type StreamEventRequest_Filter struct {
	// hash to filter events.
	Hash github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"hash" validate:"omitempty,hash"`
	// instance's hash to filter events.
	InstanceHash github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=instanceHash,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"instanceHash" validate:"omitempty,hash"`
	// key is the key of the event.
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty" validate:"printascii"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventRequest_Filter) Reset()         { *m = StreamEventRequest_Filter{} }
func (m *StreamEventRequest_Filter) String() string { return proto.CompactTextString(m) }
func (*StreamEventRequest_Filter) ProtoMessage()    {}
func (*StreamEventRequest_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ad091fa003d2ac, []int{0, 0}
}
func (m *StreamEventRequest_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventRequest_Filter.Unmarshal(m, b)
}
func (m *StreamEventRequest_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventRequest_Filter.Marshal(b, m, deterministic)
}
func (m *StreamEventRequest_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventRequest_Filter.Merge(m, src)
}
func (m *StreamEventRequest_Filter) XXX_Size() int {
	return xxx_messageInfo_StreamEventRequest_Filter.Size(m)
}
func (m *StreamEventRequest_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventRequest_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventRequest_Filter proto.InternalMessageInfo

func (m *StreamEventRequest_Filter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// CreateEventRequest defines request for execution update.
type CreateEventRequest struct {
	// instanceHash is hash of instance that can proceed an execution.
	InstanceHash github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=instanceHash,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"instanceHash" validate:"required,hash"`
	// key is the key of the event.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" validate:"required,printascii"`
	// data is the data for the event.
	Data                 *types.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ad091fa003d2ac, []int{1}
}
func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateEventRequest) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

// CreateEventResponse defines response for execution update.
type CreateEventResponse struct {
	// Hash represents event.
	Hash                 github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CreateEventResponse) Reset()         { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()    {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ad091fa003d2ac, []int{2}
}
func (m *CreateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventResponse.Unmarshal(m, b)
}
func (m *CreateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventResponse.Marshal(b, m, deterministic)
}
func (m *CreateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventResponse.Merge(m, src)
}
func (m *CreateEventResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEventResponse.Size(m)
}
func (m *CreateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamEventRequest)(nil), "mesg.api.StreamEventRequest")
	proto.RegisterType((*StreamEventRequest_Filter)(nil), "mesg.api.StreamEventRequest.Filter")
	proto.RegisterType((*CreateEventRequest)(nil), "mesg.api.CreateEventRequest")
	proto.RegisterType((*CreateEventResponse)(nil), "mesg.api.CreateEventResponse")
}

func init() { proto.RegisterFile("protobuf/api/event.proto", fileDescriptor_c0ad091fa003d2ac) }

var fileDescriptor_c0ad091fa003d2ac = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0xd6, 0x82, 0x6d, 0x2f, 0x2c, 0xaa, 0x94, 0x9a, 0x0f, 0x57, 0xcb, 0xa5, 0x08,
	0x6a, 0xa3, 0x00, 0x17, 0x10, 0x87, 0x06, 0xf1, 0x75, 0xc4, 0xdc, 0x38, 0xb1, 0x59, 0x4f, 0x92,
	0x55, 0x6b, 0xef, 0x76, 0x77, 0x1c, 0x29, 0xbf, 0x82, 0x13, 0xff, 0x89, 0xdf, 0xc0, 0xc1, 0x77,
	0x24, 0x4e, 0xf9, 0x05, 0xc8, 0xbb, 0x69, 0x53, 0x27, 0xe5, 0x12, 0x89, 0x53, 0xb2, 0x33, 0x6f,
	0xde, 0x7b, 0x33, 0x7e, 0xa4, 0xaf, 0x8d, 0x42, 0x35, 0xaa, 0xc6, 0x29, 0xd7, 0x32, 0x85, 0x19,
	0x94, 0x98, 0xb8, 0x12, 0xbd, 0x55, 0x80, 0x9d, 0x24, 0x5c, 0xcb, 0x88, 0x4d, 0xd4, 0x44, 0xa5,
	0x57, 0xc0, 0xe6, 0xe5, 0x1e, 0xee, 0x9f, 0x47, 0x47, 0xf7, 0xae, 0xda, 0x38, 0xd7, 0x60, 0x53,
	0x8b, 0xa6, 0x12, 0x4b, 0xaa, 0x28, 0x5a, 0x6b, 0x5e, 0x93, 0x61, 0xbf, 0xbb, 0x84, 0x7e, 0x41,
	0x03, 0xbc, 0x78, 0xd7, 0x54, 0x33, 0xb8, 0xa8, 0xc0, 0x22, 0x7d, 0x4d, 0xc2, 0xb1, 0x3c, 0x47,
	0x30, 0xfd, 0xe0, 0x28, 0x38, 0xde, 0x1b, 0x3c, 0x4a, 0x2e, 0xed, 0x24, 0x9b, 0xe8, 0xe4, 0xbd,
	0x83, 0x66, 0xcb, 0x91, 0xe8, 0x47, 0x97, 0x84, 0xbe, 0x44, 0xc7, 0x64, 0x67, 0xca, 0xed, 0xd4,
	0xb1, 0xec, 0x0f, 0xb3, 0x9f, 0x75, 0xdc, 0xf9, 0x55, 0xc7, 0x27, 0x13, 0x89, 0xd3, 0x6a, 0x94,
	0x08, 0x55, 0xa4, 0x42, 0xd9, 0x42, 0xd9, 0xe5, 0xcf, 0x89, 0xcd, 0xcf, 0xbc, 0xc9, 0xe4, 0x54,
	0x88, 0xd3, 0x3c, 0x37, 0x60, 0xed, 0xa2, 0x8e, 0x0f, 0x67, 0xfc, 0x5c, 0xe6, 0x1c, 0xe1, 0x15,
	0x53, 0x85, 0x44, 0x28, 0x34, 0xce, 0x9f, 0x36, 0xc4, 0x2c, 0x73, 0xfc, 0x74, 0x46, 0xf6, 0x65,
	0x69, 0x91, 0x97, 0x02, 0x3e, 0x36, 0x7a, 0xdd, 0xff, 0xa6, 0xd7, 0xd2, 0xa1, 0x4f, 0x48, 0xef,
	0x0c, 0xe6, 0xfd, 0xde, 0x51, 0x70, 0x7c, 0x7b, 0x78, 0xb8, 0xa8, 0xe3, 0x83, 0xd5, 0xa4, 0x36,
	0xb2, 0x44, 0x6e, 0x85, 0x94, 0x2c, 0x6b, 0x50, 0xec, 0x4f, 0x40, 0xe8, 0x5b, 0x03, 0x1c, 0xa1,
	0x75, 0xeb, 0x6a, 0xcd, 0xbb, 0xbf, 0xd5, 0xe7, 0x6d, 0xbd, 0xf7, 0x57, 0x0e, 0x0c, 0x5c, 0x54,
	0xd2, 0x40, 0x7e, 0xa3, 0xf5, 0x17, 0xde, 0x7a, 0xd7, 0x59, 0x67, 0x8b, 0x3a, 0x7e, 0x78, 0xc3,
	0xe0, 0xfa, 0x0e, 0xf4, 0x31, 0xd9, 0xc9, 0x39, 0x72, 0xb7, 0xf1, 0xde, 0xe0, 0xc0, 0xc7, 0xe2,
	0x32, 0x5f, 0x4d, 0x36, 0x2a, 0x81, 0x99, 0x83, 0xb0, 0x6f, 0xe4, 0x6e, 0x6b, 0x5b, 0xab, 0x55,
	0x69, 0x81, 0x7e, 0x6a, 0x45, 0xe2, 0xe5, 0x56, 0x6b, 0xfa, 0xaf, 0x3e, 0xf8, 0x1e, 0x90, 0x5d,
	0x47, 0x4e, 0x3f, 0x90, 0xd0, 0x6b, 0xd1, 0xfb, 0xab, 0xa4, 0x6e, 0xde, 0x3a, 0x7a, 0xf0, 0x8f,
	0xae, 0xf7, 0xc6, 0x3a, 0xf4, 0x0d, 0x09, 0x7d, 0xc0, 0xaf, 0x13, 0x6d, 0x46, 0x3e, 0xba, 0xe3,
	0xbb, 0xde, 0x98, 0xeb, 0xb0, 0xce, 0xb3, 0x60, 0xb8, 0xfb, 0xb5, 0xc7, 0xb5, 0x1c, 0x85, 0xee,
	0x22, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x93, 0x46, 0xcb, 0xdf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventClient interface {
	// Create creates event with data.
	Create(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	// Stream returns a stream of events that satisfy criteria
	// specified in a request.
	Stream(ctx context.Context, in *StreamEventRequest, opts ...grpc.CallOption) (Event_StreamClient, error)
}

type eventClient struct {
	cc *grpc.ClientConn
}

func NewEventClient(cc *grpc.ClientConn) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) Create(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Event/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Stream(ctx context.Context, in *StreamEventRequest, opts ...grpc.CallOption) (Event_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Event_serviceDesc.Streams[0], "/mesg.api.Event/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Event_StreamClient interface {
	Recv() (*event.Event, error)
	grpc.ClientStream
}

type eventStreamClient struct {
	grpc.ClientStream
}

func (x *eventStreamClient) Recv() (*event.Event, error) {
	m := new(event.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServer is the server API for Event service.
type EventServer interface {
	// Create creates event with data.
	Create(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	// Stream returns a stream of events that satisfy criteria
	// specified in a request.
	Stream(*StreamEventRequest, Event_StreamServer) error
}

// UnimplementedEventServer can be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (*UnimplementedEventServer) Create(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventServer) Stream(req *StreamEventRequest, srv Event_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterEventServer(s *grpc.Server, srv EventServer) {
	s.RegisterService(&_Event_serviceDesc, srv)
}

func _Event_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Event/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Create(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServer).Stream(m, &eventStreamServer{stream})
}

type Event_StreamServer interface {
	Send(*event.Event) error
	grpc.ServerStream
}

type eventStreamServer struct {
	grpc.ServerStream
}

func (x *eventStreamServer) Send(m *event.Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Event_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Event_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Event_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/api/event.proto",
}
