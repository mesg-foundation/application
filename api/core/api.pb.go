// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/api/core/api.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	github.com/mesg-foundation/core/api/core/api.proto

It has these top-level messages:
	ListenEventRequest
	ExecuteTaskRequest
	ListenResultRequest
	StartServiceRequest
	StopServiceRequest
	EventData
	ExecuteTaskReply
	ResultData
	StartServiceReply
	StopServiceReply
	DeployServiceRequest
	DeployServiceReply
	DeleteServiceRequest
	DeleteServiceReply
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/mesg-foundation/core/service"

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

type ListenEventRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
	EventKey  string `protobuf:"bytes,2,opt,name=eventKey" json:"eventKey,omitempty"`
}

func (m *ListenEventRequest) Reset()                    { *m = ListenEventRequest{} }
func (m *ListenEventRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenEventRequest) ProtoMessage()               {}
func (*ListenEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenEventRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *ListenEventRequest) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

type ExecuteTaskRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
	TaskKey   string `protobuf:"bytes,2,opt,name=taskKey" json:"taskKey,omitempty"`
	TaskData  string `protobuf:"bytes,3,opt,name=taskData" json:"taskData,omitempty"`
}

func (m *ExecuteTaskRequest) Reset()                    { *m = ExecuteTaskRequest{} }
func (m *ExecuteTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskRequest) ProtoMessage()               {}
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteTaskRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *ExecuteTaskRequest) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ExecuteTaskRequest) GetTaskData() string {
	if m != nil {
		return m.TaskData
	}
	return ""
}

type ListenResultRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
	TaskKey   string `protobuf:"bytes,2,opt,name=taskKey" json:"taskKey,omitempty"`
	OutputKey string `protobuf:"bytes,3,opt,name=outputKey" json:"outputKey,omitempty"`
}

func (m *ListenResultRequest) Reset()                    { *m = ListenResultRequest{} }
func (m *ListenResultRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenResultRequest) ProtoMessage()               {}
func (*ListenResultRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListenResultRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *ListenResultRequest) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ListenResultRequest) GetOutputKey() string {
	if m != nil {
		return m.OutputKey
	}
	return ""
}

type StartServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *StartServiceRequest) Reset()                    { *m = StartServiceRequest{} }
func (m *StartServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StartServiceRequest) ProtoMessage()               {}
func (*StartServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StartServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type StopServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *StopServiceRequest) Reset()                    { *m = StopServiceRequest{} }
func (m *StopServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StopServiceRequest) ProtoMessage()               {}
func (*StopServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StopServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type EventData struct {
	Error     string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	EventKey  string `protobuf:"bytes,2,opt,name=eventKey" json:"eventKey,omitempty"`
	EventData string `protobuf:"bytes,3,opt,name=eventData" json:"eventData,omitempty"`
}

func (m *EventData) Reset()                    { *m = EventData{} }
func (m *EventData) String() string            { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()               {}
func (*EventData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EventData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *EventData) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

func (m *EventData) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

type ExecuteTaskReply struct {
	ExecutionID string `protobuf:"bytes,1,opt,name=executionID" json:"executionID,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ExecuteTaskReply) Reset()                    { *m = ExecuteTaskReply{} }
func (m *ExecuteTaskReply) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskReply) ProtoMessage()               {}
func (*ExecuteTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ExecuteTaskReply) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ExecuteTaskReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResultData struct {
	Error       string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	ExecutionID string `protobuf:"bytes,2,opt,name=executionID" json:"executionID,omitempty"`
	TaskKey     string `protobuf:"bytes,3,opt,name=taskKey" json:"taskKey,omitempty"`
	OutputKey   string `protobuf:"bytes,4,opt,name=outputKey" json:"outputKey,omitempty"`
	OutputData  string `protobuf:"bytes,5,opt,name=outputData" json:"outputData,omitempty"`
}

func (m *ResultData) Reset()                    { *m = ResultData{} }
func (m *ResultData) String() string            { return proto.CompactTextString(m) }
func (*ResultData) ProtoMessage()               {}
func (*ResultData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResultData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResultData) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ResultData) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ResultData) GetOutputKey() string {
	if m != nil {
		return m.OutputKey
	}
	return ""
}

func (m *ResultData) GetOutputData() string {
	if m != nil {
		return m.OutputData
	}
	return ""
}

type StartServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StartServiceReply) Reset()                    { *m = StartServiceReply{} }
func (m *StartServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StartServiceReply) ProtoMessage()               {}
func (*StartServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StartServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StopServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StopServiceReply) Reset()                    { *m = StopServiceReply{} }
func (m *StopServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StopServiceReply) ProtoMessage()               {}
func (*StopServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StopServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeployServiceRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *DeployServiceRequest) Reset()                    { *m = DeployServiceRequest{} }
func (m *DeployServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployServiceRequest) ProtoMessage()               {}
func (*DeployServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeployServiceRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeployServiceReply struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *DeployServiceReply) Reset()                    { *m = DeployServiceReply{} }
func (m *DeployServiceReply) String() string            { return proto.CompactTextString(m) }
func (*DeployServiceReply) ProtoMessage()               {}
func (*DeployServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeployServiceReply) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type DeleteServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *DeleteServiceRequest) Reset()                    { *m = DeleteServiceRequest{} }
func (m *DeleteServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceRequest) ProtoMessage()               {}
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type DeleteServiceReply struct {
}

func (m *DeleteServiceReply) Reset()                    { *m = DeleteServiceReply{} }
func (m *DeleteServiceReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceReply) ProtoMessage()               {}
func (*DeleteServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*ListenEventRequest)(nil), "api.ListenEventRequest")
	proto.RegisterType((*ExecuteTaskRequest)(nil), "api.ExecuteTaskRequest")
	proto.RegisterType((*ListenResultRequest)(nil), "api.ListenResultRequest")
	proto.RegisterType((*StartServiceRequest)(nil), "api.StartServiceRequest")
	proto.RegisterType((*StopServiceRequest)(nil), "api.StopServiceRequest")
	proto.RegisterType((*EventData)(nil), "api.EventData")
	proto.RegisterType((*ExecuteTaskReply)(nil), "api.ExecuteTaskReply")
	proto.RegisterType((*ResultData)(nil), "api.ResultData")
	proto.RegisterType((*StartServiceReply)(nil), "api.StartServiceReply")
	proto.RegisterType((*StopServiceReply)(nil), "api.StopServiceReply")
	proto.RegisterType((*DeployServiceRequest)(nil), "api.DeployServiceRequest")
	proto.RegisterType((*DeployServiceReply)(nil), "api.DeployServiceReply")
	proto.RegisterType((*DeleteServiceRequest)(nil), "api.DeleteServiceRequest")
	proto.RegisterType((*DeleteServiceReply)(nil), "api.DeleteServiceReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Core service

type CoreClient interface {
	ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Core_ListenEventClient, error)
	ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error)
	ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Core_ListenResultClient, error)
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error)
	DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (*DeployServiceReply, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Core_ListenEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Core_serviceDesc.Streams[0], c.cc, "/api.Core/ListenEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreListenEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Core_ListenEventClient interface {
	Recv() (*EventData, error)
	grpc.ClientStream
}

type coreListenEventClient struct {
	grpc.ClientStream
}

func (x *coreListenEventClient) Recv() (*EventData, error) {
	m := new(EventData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreClient) ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error) {
	out := new(ExecuteTaskReply)
	err := grpc.Invoke(ctx, "/api.Core/ExecuteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Core_ListenResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Core_serviceDesc.Streams[1], c.cc, "/api.Core/ListenResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreListenResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Core_ListenResultClient interface {
	Recv() (*ResultData, error)
	grpc.ClientStream
}

type coreListenResultClient struct {
	grpc.ClientStream
}

func (x *coreListenResultClient) Recv() (*ResultData, error) {
	m := new(ResultData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error) {
	out := new(StartServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/StartService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error) {
	out := new(StopServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/StopService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (*DeployServiceReply, error) {
	out := new(DeployServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/DeployService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error) {
	out := new(DeleteServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/DeleteService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Core service

type CoreServer interface {
	ListenEvent(*ListenEventRequest, Core_ListenEventServer) error
	ExecuteTask(context.Context, *ExecuteTaskRequest) (*ExecuteTaskReply, error)
	ListenResult(*ListenResultRequest, Core_ListenResultServer) error
	StartService(context.Context, *StartServiceRequest) (*StartServiceReply, error)
	StopService(context.Context, *StopServiceRequest) (*StopServiceReply, error)
	DeployService(context.Context, *DeployServiceRequest) (*DeployServiceReply, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceReply, error)
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_ListenEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServer).ListenEvent(m, &coreListenEventServer{stream})
}

type Core_ListenEventServer interface {
	Send(*EventData) error
	grpc.ServerStream
}

type coreListenEventServer struct {
	grpc.ServerStream
}

func (x *coreListenEventServer) Send(m *EventData) error {
	return x.ServerStream.SendMsg(m)
}

func _Core_ExecuteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ExecuteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/ExecuteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ExecuteTask(ctx, req.(*ExecuteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListenResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServer).ListenResult(m, &coreListenResultServer{stream})
}

type Core_ListenResultServer interface {
	Send(*ResultData) error
	grpc.ServerStream
}

type coreListenResultServer struct {
	grpc.ServerStream
}

func (x *coreListenResultServer) Send(m *ResultData) error {
	return x.ServerStream.SendMsg(m)
}

func _Core_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeployService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeployService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/DeployService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeployService(ctx, req.(*DeployServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTask",
			Handler:    _Core_ExecuteTask_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _Core_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _Core_StopService_Handler,
		},
		{
			MethodName: "DeployService",
			Handler:    _Core_DeployService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Core_DeleteService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenEvent",
			Handler:       _Core_ListenEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenResult",
			Handler:       _Core_ListenResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/mesg-foundation/core/api/core/api.proto",
}

func init() { proto.RegisterFile("github.com/mesg-foundation/core/api/core/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xce, 0xab, 0x2d, 0x19, 0x17, 0x08, 0xdb, 0x40, 0x8d, 0x55, 0xa1, 0x6a, 0x4f, 0x05, 0x89,
	0x04, 0xa5, 0x70, 0x41, 0x42, 0x95, 0x42, 0x72, 0xe0, 0x21, 0x0e, 0x09, 0x27, 0x38, 0xb9, 0x61,
	0x68, 0xad, 0xb8, 0xde, 0xc5, 0x5e, 0x57, 0xf8, 0xbf, 0xf0, 0xb7, 0xf8, 0x3f, 0x68, 0x77, 0xfd,
	0x58, 0x3f, 0x30, 0xad, 0x38, 0x25, 0xf3, 0xfa, 0xbe, 0x19, 0x7f, 0x33, 0x0b, 0xb3, 0x0b, 0x4f,
	0x5c, 0xc6, 0xe7, 0x93, 0x0d, 0xbb, 0x9a, 0x5e, 0x61, 0x74, 0xf1, 0xfc, 0x3b, 0x8b, 0x83, 0x6f,
	0xae, 0xf0, 0x58, 0x30, 0xdd, 0xb0, 0x10, 0xa7, 0x2e, 0xf7, 0xf2, 0x3f, 0x13, 0x1e, 0x32, 0xc1,
	0x48, 0xdf, 0xe5, 0x9e, 0xf3, 0xea, 0x5f, 0x85, 0x11, 0x86, 0xd7, 0xde, 0x26, 0xff, 0xd5, 0xb5,
	0xf4, 0x13, 0x90, 0x8f, 0x5e, 0x24, 0x30, 0x58, 0x5e, 0x63, 0x20, 0x56, 0xf8, 0x23, 0xc6, 0x48,
	0x90, 0x23, 0x18, 0xa6, 0x69, 0xef, 0x16, 0x76, 0xf7, 0xb8, 0x7b, 0x32, 0x5c, 0x15, 0x0e, 0xe2,
	0xc0, 0x1d, 0x94, 0xd9, 0x1f, 0x30, 0xb1, 0x7b, 0x2a, 0x98, 0xdb, 0xf4, 0x12, 0xc8, 0xf2, 0x27,
	0x6e, 0x62, 0x81, 0x9f, 0xdd, 0x68, 0x7b, 0x33, 0x3c, 0x1b, 0xf6, 0x84, 0x1b, 0x6d, 0x0b, 0xb8,
	0xcc, 0x94, 0x4c, 0xf2, 0xef, 0xc2, 0x15, 0xae, 0xdd, 0xd7, 0x4c, 0x99, 0x4d, 0xb7, 0x70, 0xa0,
	0x3b, 0x5f, 0x61, 0x14, 0xfb, 0xe2, 0x7f, 0xa9, 0x8e, 0x60, 0xc8, 0x62, 0xc1, 0x63, 0x35, 0x95,
	0xe6, 0x2a, 0x1c, 0xf4, 0x14, 0x0e, 0xd6, 0xc2, 0x0d, 0xc5, 0x5a, 0x23, 0xdd, 0x88, 0x8c, 0xce,
	0x80, 0xac, 0x05, 0xe3, 0xb7, 0xaa, 0xf9, 0x0a, 0x43, 0xa5, 0x84, 0x1c, 0x91, 0x8c, 0x61, 0x07,
	0xc3, 0x90, 0x85, 0x69, 0x9a, 0x36, 0xda, 0x3e, 0xbf, 0x04, 0xc7, 0xac, 0x3c, 0x9b, 0x22, 0x77,
	0xd0, 0xf7, 0x30, 0x2a, 0x89, 0xc3, 0xfd, 0x84, 0x1c, 0x83, 0x85, 0xca, 0xe7, 0xb1, 0x20, 0x6f,
	0xc8, 0x74, 0x15, 0x5d, 0xf4, 0x8c, 0x2e, 0xe8, 0xaf, 0x2e, 0x80, 0xfe, 0xf2, 0x2d, 0xad, 0x56,
	0xc0, 0x7b, 0x75, 0x70, 0x43, 0x90, 0x7e, 0x8b, 0x20, 0x83, 0x8a, 0x20, 0xe4, 0x09, 0x80, 0x36,
	0xd4, 0xa4, 0x3b, 0x2a, 0x6c, 0x78, 0xe8, 0x53, 0x78, 0x50, 0x16, 0x4c, 0xce, 0xda, 0xd8, 0x24,
	0x3d, 0x81, 0x51, 0x49, 0xa6, 0xbf, 0x67, 0xce, 0x61, 0xbc, 0x40, 0xee, 0xb3, 0xa4, 0x22, 0xe9,
	0x33, 0xd8, 0x4b, 0x15, 0x54, 0xf9, 0xd6, 0x6c, 0x34, 0xc9, 0xae, 0x2c, 0xcb, 0xcc, 0x12, 0xe4,
	0x52, 0x54, 0x30, 0x24, 0x5f, 0xfb, 0x52, 0xbc, 0x94, 0xbc, 0x3e, 0x0a, 0xbc, 0xd5, 0x2a, 0x8d,
	0x25, 0x53, 0xa9, 0x8a, 0xfb, 0xc9, 0xec, 0x77, 0x1f, 0x06, 0x6f, 0x59, 0x88, 0xe4, 0x35, 0x58,
	0xc6, 0xe5, 0x93, 0xc3, 0x89, 0x7c, 0x50, 0xea, 0x6f, 0x81, 0x73, 0x4f, 0x05, 0xf2, 0xa5, 0xa4,
	0x9d, 0x17, 0x5d, 0x72, 0x06, 0x96, 0xb1, 0x48, 0x69, 0x6d, 0xfd, 0xee, 0x9d, 0x87, 0xf5, 0x00,
	0xf7, 0x13, 0xda, 0x21, 0x6f, 0x60, 0xdf, 0x3c, 0x5e, 0x62, 0x1b, 0xec, 0xa5, 0x7b, 0x76, 0xee,
	0xab, 0x48, 0xb1, 0x69, 0x8a, 0x7f, 0x0e, 0xfb, 0xa6, 0xba, 0x69, 0x79, 0xc3, 0x85, 0x3a, 0x8f,
	0x1a, 0x22, 0xba, 0x85, 0x33, 0xb0, 0x0c, 0xd9, 0xd3, 0x19, 0xea, 0xf7, 0x9a, 0xce, 0x50, 0xdd,
	0x10, 0xda, 0x21, 0x4b, 0xb8, 0x5b, 0x52, 0x92, 0x3c, 0x56, 0x99, 0x4d, 0x1b, 0xe2, 0x1c, 0x36,
	0x85, 0x0c, 0x18, 0x43, 0xa6, 0x1c, 0xa6, 0x2e, 0x78, 0x0e, 0x53, 0x55, 0x95, 0x76, 0xe6, 0xbb,
	0x5f, 0x06, 0xf2, 0x99, 0x3f, 0xdf, 0x55, 0xef, 0xfa, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0xc6, 0xe2, 0xbc, 0x49, 0x06, 0x00, 0x00,
}
