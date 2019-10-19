// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
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

// Service represents the service's type.
type Service struct {
	// Service's hash.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,10,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash" hash:"-" validate:"required"`
	// Service's sid.
	Sid string `protobuf:"bytes,12,opt,name=sid,proto3" json:"sid,omitempty" hash:"name:12" validate:"required,printascii,max=63,domain"`
	// Service's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"required,printascii"`
	// Service's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// Configurations related to the service
	Configuration Service_Configuration `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration" hash:"name:8" validate:"required"`
	// The list of tasks this service can execute.
	Tasks []*Service_Task `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty" hash:"name:5" validate:"dive,required"`
	// The list of events this service can emit.
	Events []*Service_Event `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty" hash:"name:6" validate:"dive,required"`
	// The container dependencies this service requires.
	Dependencies []*Service_Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty" hash:"name:7" validate:"dive,required"`
	// Service's repository url.
	Repository string `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty" hash:"name:9" validate:"omitempty,uri"`
	// The hash id of service's source code on IPFS.
	Source               string   `protobuf:"bytes,13,opt,name=source,proto3" json:"source,omitempty" hash:"name:13" validate:"required,printascii"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

// Events are emitted by the service whenever the service wants.
type Service_Event struct {
	// Event's key.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" hash:"name:4" validate:"printascii"`
	// Event's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Event's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// List of data of this event.
	Data                 []*Service_Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" hash:"name:3" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Event) Reset()         { *m = Service_Event{} }
func (m *Service_Event) String() string { return proto.CompactTextString(m) }
func (*Service_Event) ProtoMessage()    {}
func (*Service_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 0}
}
func (m *Service_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Event.Unmarshal(m, b)
}
func (m *Service_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Event.Marshal(b, m, deterministic)
}
func (m *Service_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Event.Merge(m, src)
}
func (m *Service_Event) XXX_Size() int {
	return xxx_messageInfo_Service_Event.Size(m)
}
func (m *Service_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Event proto.InternalMessageInfo

// Task is a function that requires inputs and returns output.
type Service_Task struct {
	// Task's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Task's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Task's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// List inputs of this task.
	Inputs []*Service_Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty" hash:"name:6" validate:"dive,required"`
	// List of tasks outputs.
	Outputs              []*Service_Parameter `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty" hash:"name:7" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Task) Reset()         { *m = Service_Task{} }
func (m *Service_Task) String() string { return proto.CompactTextString(m) }
func (*Service_Task) ProtoMessage()    {}
func (*Service_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 1}
}
func (m *Service_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Task.Unmarshal(m, b)
}
func (m *Service_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Task.Marshal(b, m, deterministic)
}
func (m *Service_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Task.Merge(m, src)
}
func (m *Service_Task) XXX_Size() int {
	return xxx_messageInfo_Service_Task.Size(m)
}
func (m *Service_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Task proto.InternalMessageInfo

// Parameter describes the task's inputs, the task's outputs, and the event's data.
type Service_Parameter struct {
	// Parameter's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Parameter's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Parameter's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// Parameter's type: `String`, `Number`, `Boolean`, `Object` or `Any`.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" hash:"name:3" validate:"required,printascii,oneof=String Number Boolean Object Any"`
	// Set the parameter as optional.
	Optional bool `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty" hash:"name:4"`
	// Mark a parameter as an array of the defined type.
	Repeated bool `protobuf:"varint,9,opt,name=repeated,proto3" json:"repeated,omitempty" hash:"name:9"`
	// Optional object structure type when type is set to `Object`.
	Object               []*Service_Parameter `protobuf:"bytes,10,rep,name=object,proto3" json:"object,omitempty" hash:"name:10" validate:"unique,dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Parameter) Reset()         { *m = Service_Parameter{} }
func (m *Service_Parameter) String() string { return proto.CompactTextString(m) }
func (*Service_Parameter) ProtoMessage()    {}
func (*Service_Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 2}
}
func (m *Service_Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Parameter.Unmarshal(m, b)
}
func (m *Service_Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Parameter.Marshal(b, m, deterministic)
}
func (m *Service_Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Parameter.Merge(m, src)
}
func (m *Service_Parameter) XXX_Size() int {
	return xxx_messageInfo_Service_Parameter.Size(m)
}
func (m *Service_Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Parameter proto.InternalMessageInfo

// A configuration is the configuration of the main container of the service's instance.
type Service_Configuration struct {
	// List of volumes.
	Volumes []string `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty" hash:"name:1" validate:"unique,dive,printascii"`
	// List of volumes mounted from other dependencies.
	VolumesFrom []string `protobuf:"bytes,2,rep,name=volumes_from,json=volumesFrom,proto3" json:"volumes_from,omitempty" hash:"name:2" validate:"unique,dive,printascii"`
	// List of ports the container exposes.
	Ports []string `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty" hash:"name:3" validate:"unique,dive,portmap"`
	// Args to pass to the container.
	Args []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty" hash:"name:5" validate:"dive,printascii"`
	// Command to run the container.
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" hash:"name:4" validate:"printascii"`
	// Default env vars to apply to service's instance on runtime.
	Env                  []string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty" hash:"name:6" validate:"unique,dive,env"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Configuration) Reset()         { *m = Service_Configuration{} }
func (m *Service_Configuration) String() string { return proto.CompactTextString(m) }
func (*Service_Configuration) ProtoMessage()    {}
func (*Service_Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 3}
}
func (m *Service_Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Configuration.Unmarshal(m, b)
}
func (m *Service_Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Configuration.Marshal(b, m, deterministic)
}
func (m *Service_Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Configuration.Merge(m, src)
}
func (m *Service_Configuration) XXX_Size() int {
	return xxx_messageInfo_Service_Configuration.Size(m)
}
func (m *Service_Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Configuration proto.InternalMessageInfo

// A dependency is a configuration of an other container that runs separately from the service.
type Service_Dependency struct {
	// Dependency's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Image's name of the container.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty" hash:"name:1" validate:"printascii"`
	// List of volumes.
	Volumes []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty" hash:"name:2" validate:"unique,dive,printascii"`
	// List of volumes mounted from other dependencies.
	VolumesFrom []string `protobuf:"bytes,3,rep,name=volumes_from,json=volumesFrom,proto3" json:"volumes_from,omitempty" hash:"name:3" validate:"unique,dive,printascii"`
	// List of ports the container exposes.
	Ports []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty" hash:"name:4" validate:"unique,dive,portmap"`
	// Args to pass to the container.
	Args []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty" hash:"name:6" validate:"dive,printascii"`
	// Command to run the container.
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" hash:"name:5" validate:"printascii"`
	// Default env vars to apply to service's instance on runtime.
	Env                  []string `protobuf:"bytes,9,rep,name=env,proto3" json:"env,omitempty" hash:"name:9" validate:"unique,dive,env"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Dependency) Reset()         { *m = Service_Dependency{} }
func (m *Service_Dependency) String() string { return proto.CompactTextString(m) }
func (*Service_Dependency) ProtoMessage()    {}
func (*Service_Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 4}
}
func (m *Service_Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Dependency.Unmarshal(m, b)
}
func (m *Service_Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Dependency.Marshal(b, m, deterministic)
}
func (m *Service_Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Dependency.Merge(m, src)
}
func (m *Service_Dependency) XXX_Size() int {
	return xxx_messageInfo_Service_Dependency.Size(m)
}
func (m *Service_Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Dependency proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Service)(nil), "mesg.types.v1.Service")
	proto.RegisterType((*Service_Event)(nil), "mesg.types.v1.Service.Event")
	proto.RegisterType((*Service_Task)(nil), "mesg.types.v1.Service.Task")
	proto.RegisterType((*Service_Parameter)(nil), "mesg.types.v1.Service.Parameter")
	proto.RegisterType((*Service_Configuration)(nil), "mesg.types.v1.Service.Configuration")
	proto.RegisterType((*Service_Dependency)(nil), "mesg.types.v1.Service.Dependency")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x97, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xb7, 0xb5, 0x93, 0x34, 0xd3, 0xe6, 0x82, 0xb9, 0x1a, 0x85, 0x15, 0xf6, 0x1a, 0x04,
	0x59, 0x48, 0x93, 0x6d, 0xb2, 0x5b, 0xb6, 0x5d, 0x15, 0xd8, 0xb0, 0xac, 0x58, 0x24, 0xfe, 0xb9,
	0x08, 0x24, 0x24, 0xc4, 0x4e, 0xec, 0x89, 0x3b, 0x34, 0x9e, 0x71, 0x67, 0xc6, 0x11, 0x79, 0x02,
	0x1e, 0x81, 0x97, 0xe0, 0x09, 0xb8, 0xe5, 0x66, 0x9f, 0x81, 0x0b, 0x4b, 0xbc, 0x00, 0x17, 0x7e,
	0x02, 0xe4, 0x49, 0xd2, 0x75, 0x1a, 0x3b, 0xfd, 0xb3, 0xdc, 0xc0, 0x5d, 0xac, 0xce, 0xf7, 0xfb,
	0xc6, 0xe7, 0x7c, 0x67, 0xa6, 0x06, 0x0d, 0x49, 0xc4, 0x84, 0x7a, 0xa4, 0x13, 0x09, 0xae, 0x38,
	0x6c, 0x84, 0x44, 0x06, 0x1d, 0x35, 0x8d, 0x88, 0xec, 0x4c, 0xf6, 0x9a, 0x4e, 0xc0, 0x03, 0xde,
	0xd5, 0x7f, 0x1a, 0xc6, 0xa3, 0x6e, 0xf6, 0xa4, 0x1f, 0xf4, 0xaf, 0x99, 0xc4, 0xf9, 0x1b, 0x81,
	0xda, 0xf1, 0x0c, 0x02, 0x03, 0x60, 0x9e, 0x60, 0x79, 0x82, 0x80, 0xbd, 0xd1, 0xda, 0x19, 0x1c,
	0xbf, 0x48, 0xac, 0x5b, 0x7f, 0x26, 0xd6, 0x7b, 0x01, 0x55, 0x27, 0xf1, 0xb0, 0xe3, 0xf1, 0xb0,
	0x9b, 0xf1, 0x77, 0x47, 0x3c, 0x66, 0x3e, 0x56, 0x94, 0xb3, 0x2e, 0x61, 0x01, 0x65, 0xa4, 0x9b,
	0xa9, 0x3a, 0x9f, 0x62, 0x79, 0x92, 0x26, 0xd6, 0xed, 0xec, 0xe1, 0xd0, 0xd9, 0x75, 0xec, 0x09,
	0x1e, 0x53, 0x1f, 0x2b, 0x72, 0xe8, 0x08, 0x72, 0x16, 0x53, 0x41, 0x7c, 0xc7, 0xd5, 0x06, 0xf0,
	0x6b, 0x60, 0x48, 0xea, 0xa3, 0x1d, 0x7b, 0xa3, 0x55, 0x1f, 0x7c, 0x98, 0x26, 0xd6, 0xa3, 0x99,
	0x88, 0xe1, 0x90, 0x1c, 0xee, 0xf5, 0x8a, 0xa4, 0xed, 0x48, 0x50, 0xa6, 0xb0, 0xf4, 0x28, 0x6d,
	0x87, 0xf8, 0xe7, 0xa3, 0xfd, 0x7e, 0xdb, 0xe7, 0x21, 0xa6, 0xcc, 0x71, 0x33, 0x16, 0x7c, 0x02,
	0xcc, 0x4c, 0x8d, 0x36, 0x34, 0xf3, 0x5e, 0x9a, 0x58, 0xed, 0x3c, 0xf3, 0x12, 0xa4, 0xe3, 0x6a,
	0x35, 0x7c, 0x06, 0xb6, 0x7d, 0x22, 0x3d, 0x41, 0xa3, 0xec, 0xf5, 0xd0, 0xa6, 0x86, 0xbd, 0x93,
	0x26, 0xd6, 0x9b, 0x39, 0xd8, 0xd2, 0xfe, 0xf2, 0x8c, 0xbc, 0x16, 0xc6, 0xa0, 0xe1, 0x71, 0x36,
	0xa2, 0x41, 0x2c, 0x74, 0xad, 0xd0, 0x96, 0xbd, 0xd1, 0xda, 0xee, 0xbd, 0xd5, 0x59, 0xea, 0x51,
	0x67, 0x5e, 0xfb, 0xce, 0xc7, 0xf9, 0xb5, 0x83, 0xbb, 0x59, 0xed, 0xd3, 0xc4, 0xba, 0x93, 0xb3,
	0x7d, 0x58, 0x5c, 0xd1, 0x65, 0x17, 0xf8, 0x03, 0xa8, 0x28, 0x2c, 0x4f, 0x25, 0xaa, 0xd8, 0x46,
	0x6b, 0xbb, 0xf7, 0x7a, 0x89, 0xdd, 0x37, 0x58, 0x9e, 0x0e, 0xde, 0x4d, 0x13, 0xeb, 0xed, 0x9c,
	0xc3, 0x83, 0xbc, 0x83, 0x4f, 0x27, 0xa4, 0xfd, 0xd2, 0x66, 0x46, 0x85, 0xcf, 0x41, 0x95, 0x4c,
	0x08, 0x53, 0x12, 0x55, 0x35, 0xff, 0x76, 0x09, 0xff, 0x93, 0x6c, 0xd1, 0x8a, 0xc1, 0xfe, 0x1a,
	0x83, 0x39, 0x17, 0x9e, 0x81, 0x1d, 0x9f, 0x44, 0x84, 0xf9, 0x84, 0x79, 0x94, 0x48, 0x54, 0xd3,
	0x3e, 0x77, 0x4a, 0x7c, 0x9e, 0x2c, 0x96, 0x4e, 0x57, 0xcc, 0xde, 0x5f, 0x63, 0xb6, 0x64, 0x01,
	0x3f, 0x03, 0x40, 0x90, 0x88, 0x4b, 0xaa, 0xb8, 0x98, 0xa2, 0xba, 0x6e, 0xfa, 0x45, 0xda, 0x41,
	0x9e, 0xc6, 0x43, 0xaa, 0x48, 0x18, 0xa9, 0x69, 0x3b, 0x16, 0xd4, 0x71, 0x73, 0x6a, 0xf8, 0x0c,
	0x54, 0x25, 0x8f, 0x85, 0x47, 0x50, 0x43, 0x73, 0xf6, 0xd2, 0xc4, 0xda, 0xcd, 0x27, 0xb1, 0x7f,
	0x69, 0x14, 0xe7, 0x80, 0xe6, 0x6f, 0x9b, 0xa0, 0xa2, 0xeb, 0x08, 0x0f, 0x80, 0x71, 0x4a, 0xa6,
	0xc8, 0x2c, 0x8c, 0xe3, 0xfd, 0xb2, 0x38, 0x66, 0x1a, 0xf8, 0x68, 0x69, 0x2e, 0x2e, 0x6a, 0xf7,
	0xca, 0xb4, 0xff, 0xfa, 0x38, 0x3c, 0x07, 0xa6, 0x8f, 0x15, 0x46, 0x86, 0x6e, 0xa7, 0x5d, 0xd2,
	0xce, 0xaf, 0xb0, 0xc0, 0x21, 0x51, 0x44, 0xac, 0xd4, 0xbf, 0xbf, 0xa6, 0x9b, 0x9a, 0xdc, 0xfc,
	0xd5, 0x00, 0x66, 0x16, 0xeb, 0x45, 0xb5, 0xb6, 0x0a, 0x77, 0xfb, 0xf0, 0x3f, 0x51, 0x2d, 0x1f,
	0x54, 0x29, 0x8b, 0xe2, 0xf3, 0x31, 0xbb, 0x7e, 0xbd, 0xd6, 0x8e, 0xda, 0x8c, 0x0d, 0x47, 0xa0,
	0xc6, 0x63, 0xa5, 0x6d, 0x6a, 0x37, 0xb4, 0x59, 0x37, 0x64, 0x0b, 0x78, 0xf3, 0x17, 0x13, 0xd4,
	0xcf, 0x11, 0xff, 0x87, 0xf6, 0x9c, 0x02, 0x33, 0xab, 0x11, 0x32, 0x34, 0xe3, 0xbb, 0x34, 0xb1,
	0x8e, 0xcb, 0xa2, 0x5a, 0x74, 0x7f, 0x71, 0x46, 0xf8, 0xe8, 0xe8, 0x58, 0x09, 0xca, 0x02, 0xfb,
	0x8b, 0x38, 0x1c, 0x12, 0x61, 0x0f, 0x38, 0x1f, 0x13, 0xcc, 0xec, 0x2f, 0x87, 0x3f, 0x11, 0x4f,
	0xd9, 0x8f, 0xd9, 0xd4, 0x71, 0xb5, 0x09, 0xdc, 0x05, 0x5b, 0x5c, 0xdb, 0xe2, 0xb1, 0x3e, 0x01,
	0xb6, 0x06, 0xaf, 0xa5, 0x89, 0xd5, 0x58, 0x3a, 0x01, 0xdc, 0xf3, 0x25, 0xd9, 0x72, 0x41, 0x22,
	0x82, 0x15, 0xf1, 0xf5, 0x51, 0xb6, 0xba, 0xfc, 0xc0, 0x71, 0xcf, 0x97, 0xc0, 0x31, 0xa8, 0x72,
	0x6d, 0x89, 0xc0, 0x15, 0x23, 0xd0, 0x4b, 0x13, 0xab, 0x93, 0x2f, 0xfb, 0xbd, 0xfc, 0xfb, 0xc6,
	0x8c, 0x9e, 0xc5, 0xa4, 0x7d, 0x31, 0x71, 0x33, 0x8f, 0xe6, 0x1f, 0x06, 0x68, 0x2c, 0xdd, 0x74,
	0xf0, 0x73, 0x50, 0x9b, 0xf0, 0x71, 0x1c, 0x12, 0x89, 0x36, 0x6c, 0xa3, 0x55, 0x1f, 0xf4, 0xd3,
	0xc4, 0xea, 0x96, 0x75, 0x35, 0x4f, 0xcf, 0x77, 0x67, 0xc1, 0x80, 0xdf, 0x82, 0x9d, 0xf9, 0xcf,
	0x1f, 0x47, 0x82, 0x87, 0x68, 0xb3, 0x90, 0xd9, 0xbb, 0x0a, 0x73, 0x7b, 0x0e, 0x7a, 0x2a, 0x78,
	0x08, 0x9f, 0x82, 0x4a, 0xc4, 0x85, 0x92, 0xfa, 0xfc, 0x5a, 0xfd, 0xff, 0xa2, 0x5f, 0x0a, 0xe4,
	0x42, 0x85, 0x38, 0x72, 0xdc, 0x99, 0x1c, 0x7e, 0x04, 0x4c, 0x2c, 0x02, 0x89, 0x4c, 0x8d, 0x69,
	0xa7, 0x89, 0xd5, 0x5a, 0x7b, 0x01, 0x2f, 0xc5, 0x38, 0x53, 0xc2, 0xc7, 0xa0, 0xe6, 0xf1, 0x30,
	0xc4, 0xcc, 0x47, 0x95, 0xeb, 0xdd, 0x07, 0x0b, 0x1d, 0xfc, 0x00, 0x18, 0x84, 0x4d, 0xf4, 0xd1,
	0xb2, 0xba, 0x87, 0xfd, 0xb2, 0x57, 0x21, 0x6c, 0xe2, 0xb8, 0x99, 0xb0, 0xf9, 0xbb, 0x09, 0xc0,
	0xcb, 0x8b, 0xf7, 0x55, 0x06, 0xfa, 0x08, 0x54, 0x68, 0x88, 0x83, 0x6b, 0x4f, 0xf4, 0x4c, 0x95,
	0x0f, 0xcf, 0x2b, 0x34, 0xba, 0x34, 0x3c, 0x46, 0x21, 0xb3, 0x7f, 0xf3, 0xf0, 0x98, 0x85, 0xe1,
	0xb9, 0x7f, 0xdd, 0xf0, 0x5c, 0xa1, 0x71, 0x37, 0x0d, 0xcf, 0x83, 0xab, 0x86, 0xa7, 0x5e, 0xb8,
	0x87, 0x83, 0x4b, 0xc3, 0x33, 0xe8, 0xbe, 0xf8, 0xeb, 0x8d, 0x5b, 0xdf, 0xdf, 0xbd, 0xfc, 0xa3,
	0x62, 0xfe, 0x69, 0x33, 0xac, 0xea, 0x0f, 0x95, 0xfe, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x9c, 0x92, 0x40, 0xec, 0x0c, 0x00, 0x00,
}
