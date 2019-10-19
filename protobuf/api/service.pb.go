// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/service.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	service "github.com/mesg-foundation/engine/service"
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

// The request's data for the `Create` API.
type ServiceServiceCreateRequest struct {
	// Service's sid.
	Sid string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	// Service's name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Service's description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Configurations related to the service
	Configuration service.Service_Configuration `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration"`
	// The list of tasks this service can execute.
	Tasks []*service.Service_Task `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// The list of events this service can emit.
	Events []*service.Service_Event `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	// The container dependencies this service requires.
	Dependencies []*service.Service_Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// Service's repository url.
	Repository string `protobuf:"bytes,8,opt,name=repository,proto3" json:"repository,omitempty"`
	// The hash id of service's source code on IPFS.
	Source               string   `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceServiceCreateRequest) Reset()         { *m = ServiceServiceCreateRequest{} }
func (m *ServiceServiceCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceCreateRequest) ProtoMessage()    {}
func (*ServiceServiceCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{0}
}
func (m *ServiceServiceCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceCreateRequest.Unmarshal(m, b)
}
func (m *ServiceServiceCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceCreateRequest.Marshal(b, m, deterministic)
}
func (m *ServiceServiceCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceCreateRequest.Merge(m, src)
}
func (m *ServiceServiceCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceCreateRequest.Size(m)
}
func (m *ServiceServiceCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceCreateRequest proto.InternalMessageInfo

func (m *ServiceServiceCreateRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *ServiceServiceCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceServiceCreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ServiceServiceCreateRequest) GetConfiguration() service.Service_Configuration {
	if m != nil {
		return m.Configuration
	}
	return service.Service_Configuration{}
}

func (m *ServiceServiceCreateRequest) GetTasks() []*service.Service_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *ServiceServiceCreateRequest) GetEvents() []*service.Service_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ServiceServiceCreateRequest) GetDependencies() []*service.Service_Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ServiceServiceCreateRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ServiceServiceCreateRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// The response's data for the `Create` API.
type ServiceServiceCreateResponse struct {
	// The service's hash created.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ServiceServiceCreateResponse) Reset()         { *m = ServiceServiceCreateResponse{} }
func (m *ServiceServiceCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceCreateResponse) ProtoMessage()    {}
func (*ServiceServiceCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{1}
}
func (m *ServiceServiceCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceCreateResponse.Unmarshal(m, b)
}
func (m *ServiceServiceCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceCreateResponse.Marshal(b, m, deterministic)
}
func (m *ServiceServiceCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceCreateResponse.Merge(m, src)
}
func (m *ServiceServiceCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceCreateResponse.Size(m)
}
func (m *ServiceServiceCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceCreateResponse proto.InternalMessageInfo

// The request's data for the `Get` API.
type ServiceServiceGetRequest struct {
	// The service's hash to fetch.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ServiceServiceGetRequest) Reset()         { *m = ServiceServiceGetRequest{} }
func (m *ServiceServiceGetRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceGetRequest) ProtoMessage()    {}
func (*ServiceServiceGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{2}
}
func (m *ServiceServiceGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceGetRequest.Unmarshal(m, b)
}
func (m *ServiceServiceGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceGetRequest.Marshal(b, m, deterministic)
}
func (m *ServiceServiceGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceGetRequest.Merge(m, src)
}
func (m *ServiceServiceGetRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceGetRequest.Size(m)
}
func (m *ServiceServiceGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceGetRequest proto.InternalMessageInfo

type ServiceServiceGetResponse struct {
	Service              *service.Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ServiceServiceGetResponse) Reset()         { *m = ServiceServiceGetResponse{} }
func (m *ServiceServiceGetResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceGetResponse) ProtoMessage()    {}
func (*ServiceServiceGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{3}
}
func (m *ServiceServiceGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceGetResponse.Unmarshal(m, b)
}
func (m *ServiceServiceGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceGetResponse.Marshal(b, m, deterministic)
}
func (m *ServiceServiceGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceGetResponse.Merge(m, src)
}
func (m *ServiceServiceGetResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceGetResponse.Size(m)
}
func (m *ServiceServiceGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceGetResponse proto.InternalMessageInfo

func (m *ServiceServiceGetResponse) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

// The request's data for the `List` API.
type ServiceServiceListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceServiceListRequest) Reset()         { *m = ServiceServiceListRequest{} }
func (m *ServiceServiceListRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceListRequest) ProtoMessage()    {}
func (*ServiceServiceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{4}
}
func (m *ServiceServiceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceListRequest.Unmarshal(m, b)
}
func (m *ServiceServiceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceListRequest.Marshal(b, m, deterministic)
}
func (m *ServiceServiceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceListRequest.Merge(m, src)
}
func (m *ServiceServiceListRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceListRequest.Size(m)
}
func (m *ServiceServiceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceListRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type ServiceServiceListResponse struct {
	// List of services that match the request's filters.
	Services             []*service.Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServiceServiceListResponse) Reset()         { *m = ServiceServiceListResponse{} }
func (m *ServiceServiceListResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceListResponse) ProtoMessage()    {}
func (*ServiceServiceListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{5}
}
func (m *ServiceServiceListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceListResponse.Unmarshal(m, b)
}
func (m *ServiceServiceListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceListResponse.Marshal(b, m, deterministic)
}
func (m *ServiceServiceListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceListResponse.Merge(m, src)
}
func (m *ServiceServiceListResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceListResponse.Size(m)
}
func (m *ServiceServiceListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceListResponse proto.InternalMessageInfo

func (m *ServiceServiceListResponse) GetServices() []*service.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// The request's data for the `List` API.
type ServiceServiceExistsRequest struct {
	// The service's hash of the existing service. This hash is nil if exists is fals.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ServiceServiceExistsRequest) Reset()         { *m = ServiceServiceExistsRequest{} }
func (m *ServiceServiceExistsRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceExistsRequest) ProtoMessage()    {}
func (*ServiceServiceExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{6}
}
func (m *ServiceServiceExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceExistsRequest.Unmarshal(m, b)
}
func (m *ServiceServiceExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceExistsRequest.Marshal(b, m, deterministic)
}
func (m *ServiceServiceExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceExistsRequest.Merge(m, src)
}
func (m *ServiceServiceExistsRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceExistsRequest.Size(m)
}
func (m *ServiceServiceExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceExistsRequest proto.InternalMessageInfo

// The response's data for the `Exists` API.
type ServiceServiceExistsResponse struct {
	// True if a service already exists, false otherwise.
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceServiceExistsResponse) Reset()         { *m = ServiceServiceExistsResponse{} }
func (m *ServiceServiceExistsResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceExistsResponse) ProtoMessage()    {}
func (*ServiceServiceExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{7}
}
func (m *ServiceServiceExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceExistsResponse.Unmarshal(m, b)
}
func (m *ServiceServiceExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceExistsResponse.Marshal(b, m, deterministic)
}
func (m *ServiceServiceExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceExistsResponse.Merge(m, src)
}
func (m *ServiceServiceExistsResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceExistsResponse.Size(m)
}
func (m *ServiceServiceExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceExistsResponse proto.InternalMessageInfo

func (m *ServiceServiceExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type ServiceServiceHashRequest struct {
	// Service's sid.
	Sid string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	// Service's name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Service's description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Configurations related to the service
	Configuration service.Service_Configuration `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration"`
	// The list of tasks this service can execute.
	Tasks []*service.Service_Task `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// The list of events this service can emit.
	Events []*service.Service_Event `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	// The container dependencies this service requires.
	Dependencies []*service.Service_Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// Service's repository url.
	Repository string `protobuf:"bytes,8,opt,name=repository,proto3" json:"repository,omitempty"`
	// The hash id of service's source code on IPFS.
	Source               string   `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceServiceHashRequest) Reset()         { *m = ServiceServiceHashRequest{} }
func (m *ServiceServiceHashRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceHashRequest) ProtoMessage()    {}
func (*ServiceServiceHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{8}
}
func (m *ServiceServiceHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceHashRequest.Unmarshal(m, b)
}
func (m *ServiceServiceHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceHashRequest.Marshal(b, m, deterministic)
}
func (m *ServiceServiceHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceHashRequest.Merge(m, src)
}
func (m *ServiceServiceHashRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceHashRequest.Size(m)
}
func (m *ServiceServiceHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceHashRequest proto.InternalMessageInfo

func (m *ServiceServiceHashRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *ServiceServiceHashRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceServiceHashRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ServiceServiceHashRequest) GetConfiguration() service.Service_Configuration {
	if m != nil {
		return m.Configuration
	}
	return service.Service_Configuration{}
}

func (m *ServiceServiceHashRequest) GetTasks() []*service.Service_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *ServiceServiceHashRequest) GetEvents() []*service.Service_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ServiceServiceHashRequest) GetDependencies() []*service.Service_Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ServiceServiceHashRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ServiceServiceHashRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// The request's data for the `Hash` API.
type ServiceServiceHashResponse struct {
	// Hash of the service calculated.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *ServiceServiceHashResponse) Reset()         { *m = ServiceServiceHashResponse{} }
func (m *ServiceServiceHashResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceServiceHashResponse) ProtoMessage()    {}
func (*ServiceServiceHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0615fe53b372bcb1, []int{9}
}
func (m *ServiceServiceHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceServiceHashResponse.Unmarshal(m, b)
}
func (m *ServiceServiceHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceServiceHashResponse.Marshal(b, m, deterministic)
}
func (m *ServiceServiceHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceServiceHashResponse.Merge(m, src)
}
func (m *ServiceServiceHashResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceServiceHashResponse.Size(m)
}
func (m *ServiceServiceHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceServiceHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceServiceHashResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ServiceServiceCreateRequest)(nil), "mesg.api.v1.ServiceServiceCreateRequest")
	proto.RegisterType((*ServiceServiceCreateResponse)(nil), "mesg.api.v1.ServiceServiceCreateResponse")
	proto.RegisterType((*ServiceServiceGetRequest)(nil), "mesg.api.v1.ServiceServiceGetRequest")
	proto.RegisterType((*ServiceServiceGetResponse)(nil), "mesg.api.v1.ServiceServiceGetResponse")
	proto.RegisterType((*ServiceServiceListRequest)(nil), "mesg.api.v1.ServiceServiceListRequest")
	proto.RegisterType((*ServiceServiceListResponse)(nil), "mesg.api.v1.ServiceServiceListResponse")
	proto.RegisterType((*ServiceServiceExistsRequest)(nil), "mesg.api.v1.ServiceServiceExistsRequest")
	proto.RegisterType((*ServiceServiceExistsResponse)(nil), "mesg.api.v1.ServiceServiceExistsResponse")
	proto.RegisterType((*ServiceServiceHashRequest)(nil), "mesg.api.v1.ServiceServiceHashRequest")
	proto.RegisterType((*ServiceServiceHashResponse)(nil), "mesg.api.v1.ServiceServiceHashResponse")
}

func init() { proto.RegisterFile("protobuf/api/service.proto", fileDescriptor_0615fe53b372bcb1) }

var fileDescriptor_0615fe53b372bcb1 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0x06, 0x0b, 0x5d, 0xf6, 0xb0, 0x1a, 0x33, 0x17, 0xa4, 0x16, 0xe2, 0x62, 0xe3, 0x0f, 0xc6,
	0x38, 0x15, 0xd6, 0xf8, 0x00, 0xac, 0x04, 0x2f, 0x34, 0xd9, 0x54, 0x63, 0xa2, 0x37, 0x9b, 0x52,
	0x0e, 0x65, 0xb2, 0xa1, 0x53, 0x3b, 0x53, 0x22, 0x0f, 0x63, 0x7c, 0x1d, 0x9f, 0xc1, 0x8b, 0x8d,
	0x8f, 0x62, 0x3a, 0x2d, 0xd8, 0x12, 0xe8, 0x7a, 0xc1, 0x85, 0x17, 0x5e, 0x31, 0x73, 0xe6, 0x3b,
	0xe7, 0xfb, 0x38, 0x67, 0xbe, 0x29, 0x98, 0x61, 0xc4, 0x25, 0x9f, 0xc4, 0x33, 0xdb, 0x0d, 0x99,
	0x2d, 0x30, 0x5a, 0x32, 0x0f, 0xa9, 0x0a, 0x92, 0xe6, 0x02, 0x85, 0x4f, 0xdd, 0x90, 0xd1, 0x65,
	0xdf, 0xb4, 0x7c, 0xee, 0x73, 0x7b, 0x83, 0x4e, 0x76, 0x6a, 0xa3, 0x56, 0x69, 0x82, 0xd9, 0xd9,
	0x1c, 0xcb, 0x55, 0x88, 0xa2, 0x58, 0xce, 0xfa, 0xae, 0x41, 0xfb, 0x7d, 0x1a, 0xc9, 0x7e, 0xce,
	0x23, 0x74, 0x25, 0x3a, 0xf8, 0x25, 0x46, 0x21, 0xc9, 0x5d, 0xd0, 0x04, 0x9b, 0x1a, 0xd5, 0x6e,
	0xb5, 0x77, 0xec, 0x24, 0x4b, 0x42, 0xa0, 0x16, 0xb8, 0x0b, 0x34, 0x6e, 0xa9, 0x90, 0x5a, 0x93,
	0x2e, 0x34, 0xa7, 0x28, 0xbc, 0x88, 0x85, 0x92, 0xf1, 0xc0, 0xd0, 0xd4, 0x51, 0x3e, 0x44, 0x2e,
	0xe0, 0xb6, 0xc7, 0x83, 0x19, 0xf3, 0xe3, 0xc8, 0x55, 0x98, 0x5a, 0xb7, 0xda, 0x6b, 0x0e, 0x1e,
	0x52, 0xf5, 0x77, 0x94, 0x32, 0xba, 0xec, 0xd3, 0x4c, 0x03, 0x3d, 0xcf, 0x63, 0x87, 0xb5, 0x1f,
	0xd7, 0xa7, 0x15, 0xa7, 0x58, 0x80, 0xf4, 0xa1, 0x2e, 0x5d, 0x71, 0x25, 0x8c, 0x7a, 0x57, 0xeb,
	0x35, 0x07, 0xed, 0x3d, 0x95, 0x3e, 0xb8, 0xe2, 0xca, 0x49, 0x91, 0xe4, 0x25, 0xe8, 0xb8, 0xc4,
	0x40, 0x0a, 0x43, 0x57, 0x39, 0x9d, 0x3d, 0x39, 0xa3, 0x04, 0xe4, 0x64, 0x58, 0x32, 0x82, 0x93,
	0x29, 0x86, 0x18, 0x4c, 0x31, 0xf0, 0x18, 0x0a, 0xe3, 0x48, 0xe5, 0x3e, 0xd8, 0x93, 0xfb, 0x7a,
	0x0d, 0x5d, 0x39, 0x85, 0x34, 0x72, 0x1f, 0x20, 0xc2, 0x90, 0x0b, 0x26, 0x79, 0xb4, 0x32, 0x1a,
	0xaa, 0x45, 0xb9, 0x08, 0x69, 0x81, 0x2e, 0x78, 0x1c, 0x79, 0x68, 0x1c, 0xab, 0xb3, 0x6c, 0x67,
	0xf9, 0xd0, 0xd9, 0x3d, 0x20, 0x11, 0xf2, 0x40, 0x20, 0x19, 0x43, 0x6d, 0xee, 0x8a, 0xb9, 0x1a,
	0xd1, 0xc9, 0xf0, 0x2c, 0x69, 0xd5, 0xcf, 0xeb, 0xd3, 0x67, 0x3e, 0x93, 0xf3, 0x78, 0x42, 0x3d,
	0xbe, 0xb0, 0x13, 0xa1, 0xcf, 0x67, 0x3c, 0x0e, 0xa6, 0xaa, 0x73, 0x36, 0x06, 0x3e, 0x0b, 0xd0,
	0x4e, 0xb2, 0xe8, 0x1b, 0x57, 0xcc, 0x1d, 0x55, 0xc0, 0xf2, 0xc0, 0x28, 0x12, 0x8d, 0x51, 0xae,
	0xaf, 0xc1, 0xc1, 0x48, 0xde, 0xc1, 0xbd, 0x1d, 0x24, 0xd9, 0x5f, 0x79, 0x01, 0x47, 0xd9, 0xed,
	0x54, 0x44, 0xcd, 0x41, 0x6b, 0x77, 0x93, 0x9d, 0x35, 0xcc, 0x6a, 0x6f, 0x97, 0x7b, 0xcb, 0xc4,
	0x5a, 0xb4, 0x75, 0x01, 0xe6, 0xae, 0xc3, 0x8c, 0x6c, 0x00, 0x8d, 0xac, 0x8a, 0x30, 0xaa, 0x6a,
	0xa4, 0xfb, 0xd8, 0x36, 0x38, 0x6b, 0xb6, 0x6d, 0x96, 0xd1, 0x57, 0x26, 0xa4, 0x38, 0x78, 0x97,
	0x5e, 0x6d, 0xcf, 0x7c, 0xcd, 0x93, 0x69, 0x6f, 0x81, 0x8e, 0x2a, 0xa2, 0xa8, 0x1a, 0x4e, 0xb6,
	0xb3, 0xbe, 0x69, 0xdb, 0xfd, 0x50, 0x45, 0xff, 0x7b, 0xf9, 0x1f, 0xf1, 0x32, 0x6e, 0xdf, 0xc8,
	0x74, 0x3c, 0x07, 0x76, 0xf2, 0xe0, 0x97, 0x06, 0x77, 0x8a, 0x3c, 0xe4, 0x12, 0xf4, 0xf4, 0xdd,
	0x20, 0x3d, 0x9a, 0xfb, 0x82, 0xd0, 0x92, 0xb7, 0xdf, 0x7c, 0xfa, 0x17, 0xc8, 0x54, 0xba, 0x55,
	0x21, 0x1f, 0x41, 0x1b, 0xa3, 0x24, 0x8f, 0x4a, 0x72, 0xfe, 0xbc, 0x27, 0xe6, 0xe3, 0x9b, 0x60,
	0x9b, 0xba, 0x9f, 0xa0, 0x96, 0xd8, 0x96, 0x94, 0x65, 0xe4, 0x4c, 0x6f, 0x3e, 0xb9, 0x11, 0xb7,
	0x29, 0x7d, 0x09, 0x7a, 0xea, 0xab, 0xd2, 0x9e, 0x14, 0x2c, 0x5e, 0xda, 0x93, 0xa2, 0x49, 0x53,
	0xed, 0xc9, 0x54, 0x4a, 0xb5, 0xe7, 0x0c, 0x5a, 0xaa, 0x3d, 0x7f, 0x53, 0xac, 0xca, 0xb0, 0xfe,
	0x59, 0x73, 0x43, 0x36, 0xd1, 0xd5, 0x57, 0xfc, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x0b, 0x6a, 0x0f, 0x32, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceServiceClient interface {
	// Create a Service from a Service Definition.
	// It will return an unique identifier which is used to interact with the Service.
	Create(ctx context.Context, in *ServiceServiceCreateRequest, opts ...grpc.CallOption) (*ServiceServiceCreateResponse, error)
	// Get returns a Service matching the criteria of the request.
	Get(ctx context.Context, in *ServiceServiceGetRequest, opts ...grpc.CallOption) (*ServiceServiceGetResponse, error)
	// List returns services specified in a request.
	List(ctx context.Context, in *ServiceServiceListRequest, opts ...grpc.CallOption) (*ServiceServiceListResponse, error)
	// Exists return if a service already exists.
	Exists(ctx context.Context, in *ServiceServiceExistsRequest, opts ...grpc.CallOption) (*ServiceServiceExistsResponse, error)
	// Hash return the hash of a service
	Hash(ctx context.Context, in *ServiceServiceHashRequest, opts ...grpc.CallOption) (*ServiceServiceHashResponse, error)
}

type serviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewServiceServiceClient(cc *grpc.ClientConn) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) Create(ctx context.Context, in *ServiceServiceCreateRequest, opts ...grpc.CallOption) (*ServiceServiceCreateResponse, error) {
	out := new(ServiceServiceCreateResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.ServiceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Get(ctx context.Context, in *ServiceServiceGetRequest, opts ...grpc.CallOption) (*ServiceServiceGetResponse, error) {
	out := new(ServiceServiceGetResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.ServiceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) List(ctx context.Context, in *ServiceServiceListRequest, opts ...grpc.CallOption) (*ServiceServiceListResponse, error) {
	out := new(ServiceServiceListResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.ServiceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Exists(ctx context.Context, in *ServiceServiceExistsRequest, opts ...grpc.CallOption) (*ServiceServiceExistsResponse, error) {
	out := new(ServiceServiceExistsResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.ServiceService/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Hash(ctx context.Context, in *ServiceServiceHashRequest, opts ...grpc.CallOption) (*ServiceServiceHashResponse, error) {
	out := new(ServiceServiceHashResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.ServiceService/Hash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
type ServiceServiceServer interface {
	// Create a Service from a Service Definition.
	// It will return an unique identifier which is used to interact with the Service.
	Create(context.Context, *ServiceServiceCreateRequest) (*ServiceServiceCreateResponse, error)
	// Get returns a Service matching the criteria of the request.
	Get(context.Context, *ServiceServiceGetRequest) (*ServiceServiceGetResponse, error)
	// List returns services specified in a request.
	List(context.Context, *ServiceServiceListRequest) (*ServiceServiceListResponse, error)
	// Exists return if a service already exists.
	Exists(context.Context, *ServiceServiceExistsRequest) (*ServiceServiceExistsResponse, error)
	// Hash return the hash of a service
	Hash(context.Context, *ServiceServiceHashRequest) (*ServiceServiceHashResponse, error)
}

// UnimplementedServiceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (*UnimplementedServiceServiceServer) Create(ctx context.Context, req *ServiceServiceCreateRequest) (*ServiceServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedServiceServiceServer) Get(ctx context.Context, req *ServiceServiceGetRequest) (*ServiceServiceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedServiceServiceServer) List(ctx context.Context, req *ServiceServiceListRequest) (*ServiceServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedServiceServiceServer) Exists(ctx context.Context, req *ServiceServiceExistsRequest) (*ServiceServiceExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (*UnimplementedServiceServiceServer) Hash(ctx context.Context, req *ServiceServiceHashRequest) (*ServiceServiceHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash not implemented")
}

func RegisterServiceServiceServer(s *grpc.Server, srv ServiceServiceServer) {
	s.RegisterService(&_ServiceService_serviceDesc, srv)
}

func _ServiceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.ServiceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Create(ctx, req.(*ServiceServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceServiceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.ServiceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Get(ctx, req.(*ServiceServiceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.ServiceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).List(ctx, req.(*ServiceServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceServiceExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.ServiceService/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Exists(ctx, req.(*ServiceServiceExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceServiceHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.ServiceService/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Hash(ctx, req.(*ServiceServiceHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.v1.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServiceService_List_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _ServiceService_Exists_Handler,
		},
		{
			MethodName: "Hash",
			Handler:    _ServiceService_Hash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/service.proto",
}
