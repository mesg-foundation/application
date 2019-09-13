// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: process.proto

package process

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	types "github.com/mesg-foundation/engine/protobuf/types"
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

// Type of condition available to compare the values.
type Process_Node_Filter_Condition_Predicate int32

const (
	// Predicate not defined.
	Process_Node_Filter_Condition_Unknown Process_Node_Filter_Condition_Predicate = 0
	// Equal
	Process_Node_Filter_Condition_EQ Process_Node_Filter_Condition_Predicate = 1
)

var Process_Node_Filter_Condition_Predicate_name = map[int32]string{
	0: "Unknown",
	1: "EQ",
}

var Process_Node_Filter_Condition_Predicate_value = map[string]int32{
	"Unknown": 0,
	"EQ":      1,
}

func (x Process_Node_Filter_Condition_Predicate) String() string {
	return proto.EnumName(Process_Node_Filter_Condition_Predicate_name, int32(x))
}

func (Process_Node_Filter_Condition_Predicate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 4, 0, 0}
}

// A process is a configuration to trigger a specific task when certains conditions of a trigger are valid.
type Process struct {
	// Process's hash
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty" hash:"-" validate:"required"`
	// Process's key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" hash:"name:2" validate:"required"`
	// Nodes with information related to the execution to trigger.
	Nodes []*Process_Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty" hash:"name:4" validate:"dive,required"`
	// Edges to create the link between the nodes.
	Edges                []*Process_Edge `protobuf:"bytes,5,rep,name=edges,proto3" json:"edges,omitempty" hash:"name:5" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0}
}
func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

// Node of the process
type Process_Node struct {
	// Types that are valid to be assigned to Type:
	//	*Process_Node_Result_
	//	*Process_Node_Event_
	//	*Process_Node_Task_
	//	*Process_Node_Map_
	//	*Process_Node_Filter_
	Type                 isProcess_Node_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Process_Node) Reset()         { *m = Process_Node{} }
func (m *Process_Node) String() string { return proto.CompactTextString(m) }
func (*Process_Node) ProtoMessage()    {}
func (*Process_Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0}
}
func (m *Process_Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node.Unmarshal(m, b)
}
func (m *Process_Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node.Marshal(b, m, deterministic)
}
func (m *Process_Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node.Merge(m, src)
}
func (m *Process_Node) XXX_Size() int {
	return xxx_messageInfo_Process_Node.Size(m)
}
func (m *Process_Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node proto.InternalMessageInfo

type isProcess_Node_Type interface {
	isProcess_Node_Type()
	Equal(interface{}) bool
}

type Process_Node_Result_ struct {
	Result *Process_Node_Result `protobuf:"bytes,1,opt,name=result,proto3,oneof" json:"result,omitempty" hash:"name:1"`
}
type Process_Node_Event_ struct {
	Event *Process_Node_Event `protobuf:"bytes,2,opt,name=event,proto3,oneof" json:"event,omitempty" hash:"name:2"`
}
type Process_Node_Task_ struct {
	Task *Process_Node_Task `protobuf:"bytes,3,opt,name=task,proto3,oneof" json:"task,omitempty" hash:"name:3"`
}
type Process_Node_Map_ struct {
	Map *Process_Node_Map `protobuf:"bytes,4,opt,name=map,proto3,oneof" json:"map,omitempty" hash:"name:4"`
}
type Process_Node_Filter_ struct {
	Filter *Process_Node_Filter `protobuf:"bytes,5,opt,name=filter,proto3,oneof" json:"filter,omitempty" hash:"name:5"`
}

func (*Process_Node_Result_) isProcess_Node_Type() {}
func (*Process_Node_Event_) isProcess_Node_Type()  {}
func (*Process_Node_Task_) isProcess_Node_Type()   {}
func (*Process_Node_Map_) isProcess_Node_Type()    {}
func (*Process_Node_Filter_) isProcess_Node_Type() {}

func (m *Process_Node) GetType() isProcess_Node_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Process_Node) GetResult() *Process_Node_Result {
	if x, ok := m.GetType().(*Process_Node_Result_); ok {
		return x.Result
	}
	return nil
}

func (m *Process_Node) GetEvent() *Process_Node_Event {
	if x, ok := m.GetType().(*Process_Node_Event_); ok {
		return x.Event
	}
	return nil
}

func (m *Process_Node) GetTask() *Process_Node_Task {
	if x, ok := m.GetType().(*Process_Node_Task_); ok {
		return x.Task
	}
	return nil
}

func (m *Process_Node) GetMap() *Process_Node_Map {
	if x, ok := m.GetType().(*Process_Node_Map_); ok {
		return x.Map
	}
	return nil
}

func (m *Process_Node) GetFilter() *Process_Node_Filter {
	if x, ok := m.GetType().(*Process_Node_Filter_); ok {
		return x.Filter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Process_Node) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Process_Node_Result_)(nil),
		(*Process_Node_Event_)(nil),
		(*Process_Node_Task_)(nil),
		(*Process_Node_Map_)(nil),
		(*Process_Node_Filter_)(nil),
	}
}

type Process_Node_Result struct {
	// Key that identifies the node.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// Hash of the instance that triggers the process.
	InstanceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=instanceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash" hash:"name:2" validate:"required"`
	// Key of the task that triggers the process.
	TaskKey              string   `protobuf:"bytes,3,opt,name=taskKey,proto3" json:"taskKey,omitempty" hash:"name:3" validate:"printascii,required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Node_Result) Reset()         { *m = Process_Node_Result{} }
func (m *Process_Node_Result) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Result) ProtoMessage()    {}
func (*Process_Node_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 0}
}
func (m *Process_Node_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Result.Unmarshal(m, b)
}
func (m *Process_Node_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Result.Marshal(b, m, deterministic)
}
func (m *Process_Node_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Result.Merge(m, src)
}
func (m *Process_Node_Result) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Result.Size(m)
}
func (m *Process_Node_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Result proto.InternalMessageInfo

type Process_Node_Event struct {
	// Key that identifies the node.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// Hash of the instance that triggers the process.
	InstanceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=instanceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash" hash:"name:2" validate:"required"`
	// Key of the event that triggers the process.
	EventKey             string   `protobuf:"bytes,3,opt,name=eventKey,proto3" json:"eventKey,omitempty" hash:"name:3" validate:"printascii,required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Node_Event) Reset()         { *m = Process_Node_Event{} }
func (m *Process_Node_Event) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Event) ProtoMessage()    {}
func (*Process_Node_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 1}
}
func (m *Process_Node_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Event.Unmarshal(m, b)
}
func (m *Process_Node_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Event.Marshal(b, m, deterministic)
}
func (m *Process_Node_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Event.Merge(m, src)
}
func (m *Process_Node_Event) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Event.Size(m)
}
func (m *Process_Node_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Event proto.InternalMessageInfo

type Process_Node_Task struct {
	// Key that identifies the node.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// Hash of the instance to execute.
	InstanceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=instanceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash" hash:"name:2" validate:"required"`
	// Task of the instance to execute.
	TaskKey              string   `protobuf:"bytes,3,opt,name=taskKey,proto3" json:"taskKey,omitempty" hash:"name:3" validate:"printascii,required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Node_Task) Reset()         { *m = Process_Node_Task{} }
func (m *Process_Node_Task) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Task) ProtoMessage()    {}
func (*Process_Node_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 2}
}
func (m *Process_Node_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Task.Unmarshal(m, b)
}
func (m *Process_Node_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Task.Marshal(b, m, deterministic)
}
func (m *Process_Node_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Task.Merge(m, src)
}
func (m *Process_Node_Task) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Task.Size(m)
}
func (m *Process_Node_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Task proto.InternalMessageInfo

type Process_Node_Map struct {
	// Key of the mapping.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// Outputs of the mapping.
	Outputs              []*Process_Node_Map_Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty" hash:"name:2" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Process_Node_Map) Reset()         { *m = Process_Node_Map{} }
func (m *Process_Node_Map) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Map) ProtoMessage()    {}
func (*Process_Node_Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 3}
}
func (m *Process_Node_Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Map.Unmarshal(m, b)
}
func (m *Process_Node_Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Map.Marshal(b, m, deterministic)
}
func (m *Process_Node_Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Map.Merge(m, src)
}
func (m *Process_Node_Map) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Map.Size(m)
}
func (m *Process_Node_Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Map proto.InternalMessageInfo

type Process_Node_Map_Output struct {
	// Key of the output.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// Types that are valid to be assigned to Value:
	//	*Process_Node_Map_Output_Ref
	//	*Process_Node_Map_Output_Constant
	Value                isProcess_Node_Map_Output_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Process_Node_Map_Output) Reset()         { *m = Process_Node_Map_Output{} }
func (m *Process_Node_Map_Output) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Map_Output) ProtoMessage()    {}
func (*Process_Node_Map_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 3, 0}
}
func (m *Process_Node_Map_Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Map_Output.Unmarshal(m, b)
}
func (m *Process_Node_Map_Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Map_Output.Marshal(b, m, deterministic)
}
func (m *Process_Node_Map_Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Map_Output.Merge(m, src)
}
func (m *Process_Node_Map_Output) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Map_Output.Size(m)
}
func (m *Process_Node_Map_Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Map_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Map_Output proto.InternalMessageInfo

type isProcess_Node_Map_Output_Value interface {
	isProcess_Node_Map_Output_Value()
	Equal(interface{}) bool
}

type Process_Node_Map_Output_Ref struct {
	Ref *Process_Node_Map_Output_Reference `protobuf:"bytes,2,opt,name=ref,proto3,oneof" json:"ref,omitempty" hash:"name:2" validate:"required_without=Constant"`
}
type Process_Node_Map_Output_Constant struct {
	Constant *types.Value `protobuf:"bytes,3,opt,name=constant,proto3,oneof" json:"constant,omitempty" hash:"name:3" validate:"required_without=Ref"`
}

func (*Process_Node_Map_Output_Ref) isProcess_Node_Map_Output_Value()      {}
func (*Process_Node_Map_Output_Constant) isProcess_Node_Map_Output_Value() {}

func (m *Process_Node_Map_Output) GetValue() isProcess_Node_Map_Output_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Process_Node_Map_Output) GetRef() *Process_Node_Map_Output_Reference {
	if x, ok := m.GetValue().(*Process_Node_Map_Output_Ref); ok {
		return x.Ref
	}
	return nil
}

func (m *Process_Node_Map_Output) GetConstant() *types.Value {
	if x, ok := m.GetValue().(*Process_Node_Map_Output_Constant); ok {
		return x.Constant
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Process_Node_Map_Output) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Process_Node_Map_Output_Ref)(nil),
		(*Process_Node_Map_Output_Constant)(nil),
	}
}

type Process_Node_Map_Output_Reference struct {
	// Key of the node in the graph. If empty, will be using the src of the edge.
	NodeKey string `protobuf:"bytes,1,opt,name=nodeKey,proto3" json:"nodeKey,omitempty" hash:"name:1" validate:"required"`
	// Key of a specific parameter of the referenced node's output data.
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" hash:"name:2" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Node_Map_Output_Reference) Reset()         { *m = Process_Node_Map_Output_Reference{} }
func (m *Process_Node_Map_Output_Reference) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Map_Output_Reference) ProtoMessage()    {}
func (*Process_Node_Map_Output_Reference) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 3, 0, 0}
}
func (m *Process_Node_Map_Output_Reference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Map_Output_Reference.Unmarshal(m, b)
}
func (m *Process_Node_Map_Output_Reference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Map_Output_Reference.Marshal(b, m, deterministic)
}
func (m *Process_Node_Map_Output_Reference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Map_Output_Reference.Merge(m, src)
}
func (m *Process_Node_Map_Output_Reference) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Map_Output_Reference.Size(m)
}
func (m *Process_Node_Map_Output_Reference) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Map_Output_Reference.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Map_Output_Reference proto.InternalMessageInfo

type Process_Node_Filter struct {
	// Key for the filter
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required"`
	// List of condition to apply for this filter
	Conditions           []Process_Node_Filter_Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions" hash:"name:2"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Process_Node_Filter) Reset()         { *m = Process_Node_Filter{} }
func (m *Process_Node_Filter) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Filter) ProtoMessage()    {}
func (*Process_Node_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 4}
}
func (m *Process_Node_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Filter.Unmarshal(m, b)
}
func (m *Process_Node_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Filter.Marshal(b, m, deterministic)
}
func (m *Process_Node_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Filter.Merge(m, src)
}
func (m *Process_Node_Filter) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Filter.Size(m)
}
func (m *Process_Node_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Filter proto.InternalMessageInfo

type Process_Node_Filter_Condition struct {
	// Key to check.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" hash:"name:1" validate:"required,printascii"`
	// Type of condition to apply.
	Predicate Process_Node_Filter_Condition_Predicate `protobuf:"varint,2,opt,name=predicate,proto3,enum=types.Process_Node_Filter_Condition_Predicate" json:"predicate,omitempty" hash:"name:2" validate:"required"`
	// Value of the filter.
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" hash:"name:3"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Node_Filter_Condition) Reset()         { *m = Process_Node_Filter_Condition{} }
func (m *Process_Node_Filter_Condition) String() string { return proto.CompactTextString(m) }
func (*Process_Node_Filter_Condition) ProtoMessage()    {}
func (*Process_Node_Filter_Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0, 4, 0}
}
func (m *Process_Node_Filter_Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Node_Filter_Condition.Unmarshal(m, b)
}
func (m *Process_Node_Filter_Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Node_Filter_Condition.Marshal(b, m, deterministic)
}
func (m *Process_Node_Filter_Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Node_Filter_Condition.Merge(m, src)
}
func (m *Process_Node_Filter_Condition) XXX_Size() int {
	return xxx_messageInfo_Process_Node_Filter_Condition.Size(m)
}
func (m *Process_Node_Filter_Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Node_Filter_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Node_Filter_Condition proto.InternalMessageInfo

type Process_Edge struct {
	// Source of the edge.
	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty" hash:"name:1" validate:"required"`
	// Destination of the edge.
	Dst                  string   `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty" hash:"name:2" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Edge) Reset()         { *m = Process_Edge{} }
func (m *Process_Edge) String() string { return proto.CompactTextString(m) }
func (*Process_Edge) ProtoMessage()    {}
func (*Process_Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 1}
}
func (m *Process_Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Edge.Unmarshal(m, b)
}
func (m *Process_Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Edge.Marshal(b, m, deterministic)
}
func (m *Process_Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Edge.Merge(m, src)
}
func (m *Process_Edge) XXX_Size() int {
	return xxx_messageInfo_Process_Edge.Size(m)
}
func (m *Process_Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Edge proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("types.Process_Node_Filter_Condition_Predicate", Process_Node_Filter_Condition_Predicate_name, Process_Node_Filter_Condition_Predicate_value)
	proto.RegisterType((*Process)(nil), "types.Process")
	proto.RegisterType((*Process_Node)(nil), "types.Process.Node")
	proto.RegisterType((*Process_Node_Result)(nil), "types.Process.Node.Result")
	proto.RegisterType((*Process_Node_Event)(nil), "types.Process.Node.Event")
	proto.RegisterType((*Process_Node_Task)(nil), "types.Process.Node.Task")
	proto.RegisterType((*Process_Node_Map)(nil), "types.Process.Node.Map")
	proto.RegisterType((*Process_Node_Map_Output)(nil), "types.Process.Node.Map.Output")
	proto.RegisterType((*Process_Node_Map_Output_Reference)(nil), "types.Process.Node.Map.Output.Reference")
	proto.RegisterType((*Process_Node_Filter)(nil), "types.Process.Node.Filter")
	proto.RegisterType((*Process_Node_Filter_Condition)(nil), "types.Process.Node.Filter.Condition")
	proto.RegisterType((*Process_Edge)(nil), "types.Process.Edge")
}

func init() { proto.RegisterFile("process.proto", fileDescriptor_54c4d0e8c0aaf5c3) }

var fileDescriptor_54c4d0e8c0aaf5c3 = []byte{
	// 861 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xcf, 0x6e, 0xe3, 0x54,
	0x14, 0xc6, 0x13, 0xdb, 0x49, 0x26, 0x27, 0x33, 0x68, 0xc6, 0x0c, 0xc2, 0x18, 0x54, 0x07, 0x8b,
	0x3f, 0x01, 0x5a, 0x87, 0xba, 0x1d, 0x90, 0x2a, 0x01, 0xc2, 0x43, 0x61, 0xc4, 0x30, 0x50, 0x22,
	0x8a, 0x10, 0x2c, 0x90, 0x6b, 0x9f, 0x24, 0x56, 0x12, 0xdb, 0xf8, 0x5e, 0xa7, 0x44, 0x62, 0xcb,
	0x02, 0x81, 0xc4, 0x1b, 0xb0, 0x66, 0x07, 0xaf, 0xc0, 0xae, 0x4f, 0xc0, 0x82, 0x85, 0x25, 0x78,
	0x84, 0x3c, 0x01, 0xba, 0xd7, 0x76, 0x9a, 0x62, 0x27, 0x6d, 0xc3, 0x0a, 0x89, 0x5d, 0x1d, 0x9f,
	0xef, 0x77, 0x4e, 0xbf, 0x73, 0xee, 0xb9, 0x86, 0x5b, 0x61, 0x14, 0x38, 0x48, 0x88, 0x11, 0x46,
	0x01, 0x0d, 0xe4, 0x1a, 0x9d, 0x85, 0x48, 0x54, 0x7d, 0x10, 0x0c, 0x82, 0x2e, 0xff, 0xe9, 0x24,
	0xee, 0x77, 0xd9, 0x13, 0x7f, 0xe0, 0x7f, 0xa5, 0xa1, 0xea, 0xd3, 0x8b, 0xd7, 0x5c, 0xd3, 0x25,
	0x34, 0x8a, 0x1d, 0x9a, 0xbe, 0xd4, 0x7f, 0xbb, 0x0b, 0x8d, 0xa3, 0x94, 0x2c, 0xef, 0x83, 0x34,
	0xb4, 0xc9, 0x50, 0xa9, 0xb6, 0xab, 0x9d, 0x9b, 0x56, 0x7b, 0x9e, 0x68, 0xcf, 0xb0, 0xe7, 0x03,
	0x7d, 0x47, 0x6f, 0x4f, 0xed, 0xb1, 0xe7, 0xda, 0x14, 0x0f, 0xf4, 0x08, 0xbf, 0x8a, 0xbd, 0x08,
	0x5d, 0xbd, 0xc7, 0xa3, 0xe5, 0xd7, 0x41, 0x1c, 0xe1, 0x4c, 0x11, 0xda, 0xd5, 0x4e, 0xd3, 0x7a,
	0x7e, 0x9e, 0x68, 0xcf, 0xa6, 0x22, 0xdf, 0x9e, 0xe0, 0x81, 0x59, 0xae, 0x64, 0x0a, 0xf9, 0x18,
	0x6a, 0x7e, 0xe0, 0x22, 0x51, 0xa4, 0xb6, 0xd8, 0x69, 0x99, 0x8f, 0x1b, 0xbc, 0x3c, 0x23, 0xab,
	0xc6, 0xf8, 0x30, 0x70, 0xd1, 0x7a, 0x79, 0x9e, 0x68, 0x2f, 0x2c, 0xf1, 0xf6, 0x97, 0x79, 0xae,
	0x37, 0xc5, 0xed, 0x73, 0x68, 0x4a, 0x63, 0x58, 0x74, 0x07, 0x48, 0x94, 0x5a, 0x29, 0xf6, 0xd0,
	0x1d, 0x14, 0xb1, 0xf7, 0xd6, 0x61, 0x39, 0x4d, 0xfd, 0xe9, 0x0e, 0x48, 0xac, 0x24, 0xf9, 0x3d,
	0xa8, 0x47, 0x48, 0xe2, 0x31, 0xe5, 0x3e, 0xb5, 0x4c, 0xb5, 0xa4, 0x6e, 0xa3, 0xc7, 0x23, 0xac,
	0x3b, 0xf3, 0x44, 0xbb, 0xb5, 0x94, 0x67, 0x57, 0x7f, 0x50, 0xe9, 0x65, 0x72, 0xf9, 0x1d, 0xa8,
	0xe1, 0x14, 0x7d, 0xca, 0xad, 0x6b, 0x99, 0x4f, 0x95, 0x71, 0x0e, 0x59, 0x40, 0x01, 0x63, 0x32,
	0x4c, 0x2a, 0x96, 0xdf, 0x06, 0x89, 0xda, 0x64, 0xa4, 0x88, 0x1c, 0xa2, 0x94, 0x41, 0x3e, 0xb1,
	0xc9, 0xa8, 0xc0, 0xd8, 0x63, 0x0c, 0x2e, 0x95, 0xdf, 0x04, 0x71, 0x62, 0x87, 0x8a, 0xc4, 0x09,
	0x4f, 0x96, 0x11, 0x1e, 0xd9, 0x61, 0x01, 0xb0, 0xcf, 0x00, 0x4c, 0xc8, 0x1c, 0xe9, 0x7b, 0x63,
	0x8a, 0x91, 0x52, 0x5b, 0xed, 0xc8, 0xbb, 0x3c, 0xa2, 0x40, 0xb9, 0xc7, 0x1d, 0x49, 0xe5, 0xea,
	0xf7, 0x02, 0xd4, 0x53, 0xe7, 0xf2, 0xa9, 0xaa, 0x96, 0x4e, 0xd5, 0xee, 0x9a, 0xa9, 0xfa, 0x06,
	0x6e, 0x7a, 0x3e, 0xa1, 0xb6, 0xef, 0xe0, 0x03, 0x36, 0xcc, 0x02, 0x1f, 0xe6, 0xcf, 0xce, 0x12,
	0xad, 0xf2, 0x47, 0xa2, 0xbd, 0x32, 0xf0, 0xe8, 0x30, 0x3e, 0x31, 0x9c, 0x60, 0xd2, 0x9d, 0x20,
	0x19, 0xec, 0xf4, 0x83, 0xd8, 0x77, 0x6d, 0xea, 0x05, 0x7e, 0x17, 0xfd, 0x81, 0xe7, 0x63, 0x97,
	0xe5, 0x31, 0x98, 0xf4, 0x6a, 0xa3, 0x7c, 0x21, 0x9b, 0xfc, 0x3e, 0x34, 0x98, 0xa5, 0x0f, 0x71,
	0xc6, 0x1b, 0xd2, 0xb4, 0x5e, 0x9d, 0x27, 0xda, 0xf6, 0x05, 0xdb, 0x97, 0x28, 0x61, 0xe4, 0xf9,
	0xd4, 0x26, 0x8e, 0xe7, 0x2d, 0xcd, 0x5b, 0x0e, 0x50, 0x7f, 0x10, 0xa0, 0xc6, 0xfb, 0xff, 0x5f,
	0x35, 0xe3, 0x03, 0xb8, 0xc1, 0x67, 0xf4, 0xdf, 0xb8, 0xb1, 0x20, 0xa8, 0xdf, 0x09, 0x20, 0xb1,
	0x49, 0xfe, 0x7f, 0x34, 0x66, 0xea, 0xef, 0x12, 0x88, 0x8f, 0xec, 0x70, 0x73, 0x2b, 0x5c, 0x68,
	0x04, 0x31, 0x0d, 0x63, 0x4a, 0x14, 0x81, 0xaf, 0xc9, 0xad, 0x15, 0xc7, 0xde, 0xf8, 0x88, 0x87,
	0x15, 0x36, 0xa6, 0xb9, 0x66, 0x63, 0xe6, 0x68, 0xf5, 0x17, 0x11, 0xea, 0xa9, 0x7e, 0xf3, 0x4a,
	0x29, 0x88, 0x11, 0xf6, 0xb3, 0x1d, 0xd9, 0x59, 0x5f, 0xa5, 0xd1, 0xc3, 0x3e, 0x46, 0xe8, 0x3b,
	0x68, 0xbd, 0x36, 0x4f, 0x34, 0xf3, 0xb2, 0x16, 0x7d, 0x79, 0xea, 0xd1, 0x61, 0x10, 0xd3, 0x37,
	0xee, 0x07, 0xbc, 0x53, 0x94, 0xaf, 0xb4, 0x08, 0xfb, 0x32, 0xc2, 0x0d, 0x27, 0xfb, 0x2d, 0xdb,
	0xac, 0x77, 0x0d, 0x36, 0x14, 0x46, 0x7e, 0x97, 0x1a, 0x9f, 0xda, 0xe3, 0x18, 0xad, 0xdd, 0x79,
	0xa2, 0xed, 0xac, 0xea, 0x61, 0x21, 0x4d, 0x0f, 0xfb, 0x2c, 0xc3, 0x02, 0xad, 0x7e, 0x5b, 0x85,
	0xe6, 0xa2, 0x66, 0xf9, 0x2d, 0x68, 0xb0, 0x2b, 0xec, 0xe1, 0x75, 0x7d, 0xca, 0x55, 0x1b, 0x5f,
	0xc5, 0x56, 0x03, 0x6a, 0x53, 0xf6, 0xff, 0xa8, 0xbf, 0x8a, 0x50, 0x4f, 0x37, 0xf5, 0xe6, 0x1d,
	0xfb, 0x02, 0xc0, 0x09, 0x7c, 0xd7, 0x63, 0x27, 0x27, 0x1f, 0xaf, 0xe7, 0x56, 0x5f, 0x09, 0xc6,
	0xfd, 0x3c, 0xd8, 0x7a, 0x82, 0x1d, 0xc5, 0xc2, 0x5d, 0xd7, 0x5b, 0xc2, 0xa9, 0x3f, 0x0a, 0xd0,
	0x5c, 0x08, 0x64, 0x6b, 0xb9, 0xc6, 0x7f, 0x9e, 0xa7, 0xd2, 0x1a, 0xb7, 0xcf, 0x0f, 0x56, 0x56,
	0x2e, 0x81, 0x66, 0x18, 0xa1, 0xeb, 0x39, 0x36, 0x45, 0x6e, 0xdd, 0x63, 0xa6, 0x71, 0x95, 0x6a,
	0x8d, 0xa3, 0x5c, 0x75, 0x55, 0xab, 0xcf, 0xf3, 0xc8, 0x2f, 0x66, 0x86, 0x67, 0xab, 0xa0, 0x78,
	0x39, 0xf7, 0xd2, 0xf7, 0x7a, 0x1b, 0x9a, 0x8b, 0x3c, 0x72, 0x0b, 0x1a, 0xc7, 0xfe, 0xc8, 0x0f,
	0x4e, 0xfd, 0xdb, 0x15, 0xb9, 0x0e, 0xc2, 0xe1, 0xc7, 0xb7, 0xab, 0x56, 0x1d, 0x24, 0x56, 0xad,
	0xfa, 0x35, 0x48, 0xec, 0xdb, 0x86, 0xf5, 0x8d, 0x44, 0xce, 0x35, 0xfb, 0x46, 0x22, 0x87, 0x09,
	0x5d, 0x42, 0xaf, 0x39, 0x3d, 0x2e, 0xa1, 0xd6, 0xde, 0xd9, 0x9f, 0x5b, 0x95, 0x9f, 0xff, 0xda,
	0xaa, 0x7e, 0xfe, 0xd2, 0xe5, 0x3b, 0x34, 0xfb, 0x8c, 0x3d, 0xa9, 0xf3, 0x93, 0xb4, 0xf7, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xab, 0x58, 0xa5, 0xd8, 0x0a, 0x00, 0x00,
}

func (this *Process) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process)
	if !ok {
		that2, ok := that.(Process)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if len(this.Nodes) != len(that1.Nodes) {
		return false
	}
	for i := range this.Nodes {
		if !this.Nodes[i].Equal(that1.Nodes[i]) {
			return false
		}
	}
	if len(this.Edges) != len(that1.Edges) {
		return false
	}
	for i := range this.Edges {
		if !this.Edges[i].Equal(that1.Edges[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node)
	if !ok {
		that2, ok := that.(Process_Node)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Result_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Result_)
	if !ok {
		that2, ok := that.(Process_Node_Result_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Result.Equal(that1.Result) {
		return false
	}
	return true
}
func (this *Process_Node_Event_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Event_)
	if !ok {
		that2, ok := that.(Process_Node_Event_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Event.Equal(that1.Event) {
		return false
	}
	return true
}
func (this *Process_Node_Task_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Task_)
	if !ok {
		that2, ok := that.(Process_Node_Task_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Task.Equal(that1.Task) {
		return false
	}
	return true
}
func (this *Process_Node_Map_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map_)
	if !ok {
		that2, ok := that.(Process_Node_Map_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Map.Equal(that1.Map) {
		return false
	}
	return true
}
func (this *Process_Node_Filter_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Filter_)
	if !ok {
		that2, ok := that.(Process_Node_Filter_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Filter.Equal(that1.Filter) {
		return false
	}
	return true
}
func (this *Process_Node_Result) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Result)
	if !ok {
		that2, ok := that.(Process_Node_Result)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.InstanceHash.Equal(that1.InstanceHash) {
		return false
	}
	if this.TaskKey != that1.TaskKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Event)
	if !ok {
		that2, ok := that.(Process_Node_Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.InstanceHash.Equal(that1.InstanceHash) {
		return false
	}
	if this.EventKey != that1.EventKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Task) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Task)
	if !ok {
		that2, ok := that.(Process_Node_Task)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.InstanceHash.Equal(that1.InstanceHash) {
		return false
	}
	if this.TaskKey != that1.TaskKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Map) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map)
	if !ok {
		that2, ok := that.(Process_Node_Map)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if len(this.Outputs) != len(that1.Outputs) {
		return false
	}
	for i := range this.Outputs {
		if !this.Outputs[i].Equal(that1.Outputs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Map_Output) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map_Output)
	if !ok {
		that2, ok := that.(Process_Node_Map_Output)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Map_Output_Ref) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map_Output_Ref)
	if !ok {
		that2, ok := that.(Process_Node_Map_Output_Ref)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	return true
}
func (this *Process_Node_Map_Output_Constant) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map_Output_Constant)
	if !ok {
		that2, ok := that.(Process_Node_Map_Output_Constant)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Constant.Equal(that1.Constant) {
		return false
	}
	return true
}
func (this *Process_Node_Map_Output_Reference) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Map_Output_Reference)
	if !ok {
		that2, ok := that.(Process_Node_Map_Output_Reference)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NodeKey != that1.NodeKey {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Filter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Filter)
	if !ok {
		that2, ok := that.(Process_Node_Filter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if len(this.Conditions) != len(that1.Conditions) {
		return false
	}
	for i := range this.Conditions {
		if !this.Conditions[i].Equal(&that1.Conditions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Node_Filter_Condition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Node_Filter_Condition)
	if !ok {
		that2, ok := that.(Process_Node_Filter_Condition)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Predicate != that1.Predicate {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Process_Edge) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Process_Edge)
	if !ok {
		that2, ok := that.(Process_Edge)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Src != that1.Src {
		return false
	}
	if this.Dst != that1.Dst {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
