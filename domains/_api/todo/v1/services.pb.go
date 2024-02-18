// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: todo/v1/services.proto

package todoV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateTodoRequest_UpdateStrategy int32

const (
	UpdateTodoRequest_PATCH UpdateTodoRequest_UpdateStrategy = 0
	UpdateTodoRequest_ALL   UpdateTodoRequest_UpdateStrategy = 1
)

// Enum value maps for UpdateTodoRequest_UpdateStrategy.
var (
	UpdateTodoRequest_UpdateStrategy_name = map[int32]string{
		0: "PATCH",
		1: "ALL",
	}
	UpdateTodoRequest_UpdateStrategy_value = map[string]int32{
		"PATCH": 0,
		"ALL":   1,
	}
)

func (x UpdateTodoRequest_UpdateStrategy) Enum() *UpdateTodoRequest_UpdateStrategy {
	p := new(UpdateTodoRequest_UpdateStrategy)
	*p = x
	return p
}

func (x UpdateTodoRequest_UpdateStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateTodoRequest_UpdateStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_v1_services_proto_enumTypes[0].Descriptor()
}

func (UpdateTodoRequest_UpdateStrategy) Type() protoreflect.EnumType {
	return &file_todo_v1_services_proto_enumTypes[0]
}

func (x UpdateTodoRequest_UpdateStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateTodoRequest_UpdateStrategy.Descriptor instead.
func (UpdateTodoRequest_UpdateStrategy) EnumDescriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{4, 0}
}

type DeleteTodoRequest_DeleteStrategy int32

const (
	DeleteTodoRequest_SOFT DeleteTodoRequest_DeleteStrategy = 0
	DeleteTodoRequest_HARD DeleteTodoRequest_DeleteStrategy = 1
)

// Enum value maps for DeleteTodoRequest_DeleteStrategy.
var (
	DeleteTodoRequest_DeleteStrategy_name = map[int32]string{
		0: "SOFT",
		1: "HARD",
	}
	DeleteTodoRequest_DeleteStrategy_value = map[string]int32{
		"SOFT": 0,
		"HARD": 1,
	}
)

func (x DeleteTodoRequest_DeleteStrategy) Enum() *DeleteTodoRequest_DeleteStrategy {
	p := new(DeleteTodoRequest_DeleteStrategy)
	*p = x
	return p
}

func (x DeleteTodoRequest_DeleteStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteTodoRequest_DeleteStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_v1_services_proto_enumTypes[1].Descriptor()
}

func (DeleteTodoRequest_DeleteStrategy) Type() protoreflect.EnumType {
	return &file_todo_v1_services_proto_enumTypes[1]
}

func (x DeleteTodoRequest_DeleteStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteTodoRequest_DeleteStrategy.Descriptor instead.
func (DeleteTodoRequest_DeleteStrategy) EnumDescriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{5, 0}
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateManyTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*CreateTodoRequest `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *CreateManyTodosRequest) Reset() {
	*x = CreateManyTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManyTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManyTodosRequest) ProtoMessage() {}

func (x *CreateManyTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManyTodosRequest.ProtoReflect.Descriptor instead.
func (*CreateManyTodosRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{1}
}

func (x *CreateManyTodosRequest) GetTodos() []*CreateTodoRequest {
	if x != nil {
		return x.Todos
	}
	return nil
}

type ReadTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadTodoRequest) Reset() {
	*x = ReadTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTodoRequest) ProtoMessage() {}

func (x *ReadTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTodoRequest.ProtoReflect.Descriptor instead.
func (*ReadTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{2}
}

func (x *ReadTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadManyTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageOffset int32 `protobuf:"varint,2,opt,name=page_offset,json=pageOffset,proto3" json:"page_offset,omitempty"`
}

func (x *ReadManyTodosRequest) Reset() {
	*x = ReadManyTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadManyTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadManyTodosRequest) ProtoMessage() {}

func (x *ReadManyTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadManyTodosRequest.ProtoReflect.Descriptor instead.
func (*ReadManyTodosRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{3}
}

func (x *ReadManyTodosRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ReadManyTodosRequest) GetPageOffset() int32 {
	if x != nil {
		return x.PageOffset
	}
	return 0
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateStrategy UpdateTodoRequest_UpdateStrategy `protobuf:"varint,1,opt,name=update_strategy,json=updateStrategy,proto3,enum=UpdateTodoRequest_UpdateStrategy" json:"update_strategy,omitempty"`
	Todo           *Todo                            `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTodoRequest) GetUpdateStrategy() UpdateTodoRequest_UpdateStrategy {
	if x != nil {
		return x.UpdateStrategy
	}
	return UpdateTodoRequest_PATCH
}

func (x *UpdateTodoRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteStrategy DeleteTodoRequest_DeleteStrategy `protobuf:"varint,1,opt,name=delete_strategy,json=deleteStrategy,proto3,enum=DeleteTodoRequest_DeleteStrategy" json:"delete_strategy,omitempty"`
	Id             string                           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTodoRequest) GetDeleteStrategy() DeleteTodoRequest_DeleteStrategy {
	if x != nil {
		return x.DeleteStrategy
	}
	return DeleteTodoRequest_SOFT
}

func (x *DeleteTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{6}
}

func (x *TodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type CreateManyTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created int32 `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Failed  int32 `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CreateManyTodosResponse) Reset() {
	*x = CreateManyTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManyTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManyTodosResponse) ProtoMessage() {}

func (x *CreateManyTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManyTodosResponse.ProtoReflect.Descriptor instead.
func (*CreateManyTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{7}
}

func (x *CreateManyTodosResponse) GetCreated() int32 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *CreateManyTodosResponse) GetFailed() int32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

func (x *CreateManyTodosResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ReadManyTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos          []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	NextPageOffset int32   `protobuf:"varint,2,opt,name=next_page_offset,json=nextPageOffset,proto3" json:"next_page_offset,omitempty"`
}

func (x *ReadManyTodosResponse) Reset() {
	*x = ReadManyTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_services_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadManyTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadManyTodosResponse) ProtoMessage() {}

func (x *ReadManyTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_services_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadManyTodosResponse.ProtoReflect.Descriptor instead.
func (*ReadManyTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_v1_services_proto_rawDescGZIP(), []int{8}
}

func (x *ReadManyTodosResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

func (x *ReadManyTodosResponse) GetNextPageOffset() int32 {
	if x != nil {
		return x.NextPageOffset
	}
	return 0
}

var File_todo_v1_services_proto protoreflect.FileDescriptor

var file_todo_v1_services_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x54, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x4f, 0x46, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x52, 0x44, 0x10,
	0x01, 0x22, 0x29, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x61, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x5e, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x05,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32,
	0xbb, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x27, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64,
	0x4d, 0x61, 0x6e, 0x79, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x71, 0x75,
	0x65, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_todo_v1_services_proto_rawDescOnce sync.Once
	file_todo_v1_services_proto_rawDescData = file_todo_v1_services_proto_rawDesc
)

func file_todo_v1_services_proto_rawDescGZIP() []byte {
	file_todo_v1_services_proto_rawDescOnce.Do(func() {
		file_todo_v1_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_v1_services_proto_rawDescData)
	})
	return file_todo_v1_services_proto_rawDescData
}

var file_todo_v1_services_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_todo_v1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_todo_v1_services_proto_goTypes = []interface{}{
	(UpdateTodoRequest_UpdateStrategy)(0), // 0: UpdateTodoRequest.UpdateStrategy
	(DeleteTodoRequest_DeleteStrategy)(0), // 1: DeleteTodoRequest.DeleteStrategy
	(*CreateTodoRequest)(nil),             // 2: CreateTodoRequest
	(*CreateManyTodosRequest)(nil),        // 3: CreateManyTodosRequest
	(*ReadTodoRequest)(nil),               // 4: ReadTodoRequest
	(*ReadManyTodosRequest)(nil),          // 5: ReadManyTodosRequest
	(*UpdateTodoRequest)(nil),             // 6: UpdateTodoRequest
	(*DeleteTodoRequest)(nil),             // 7: DeleteTodoRequest
	(*TodoResponse)(nil),                  // 8: TodoResponse
	(*CreateManyTodosResponse)(nil),       // 9: CreateManyTodosResponse
	(*ReadManyTodosResponse)(nil),         // 10: ReadManyTodosResponse
	(*Todo)(nil),                          // 11: Todo
}
var file_todo_v1_services_proto_depIdxs = []int32{
	2,  // 0: CreateManyTodosRequest.todos:type_name -> CreateTodoRequest
	0,  // 1: UpdateTodoRequest.update_strategy:type_name -> UpdateTodoRequest.UpdateStrategy
	11, // 2: UpdateTodoRequest.todo:type_name -> Todo
	1,  // 3: DeleteTodoRequest.delete_strategy:type_name -> DeleteTodoRequest.DeleteStrategy
	11, // 4: TodoResponse.todo:type_name -> Todo
	11, // 5: ReadManyTodosResponse.todos:type_name -> Todo
	2,  // 6: TodoService.Create:input_type -> CreateTodoRequest
	3,  // 7: TodoService.CreateMany:input_type -> CreateManyTodosRequest
	4,  // 8: TodoService.Read:input_type -> ReadTodoRequest
	5,  // 9: TodoService.ReadMany:input_type -> ReadManyTodosRequest
	6,  // 10: TodoService.Update:input_type -> UpdateTodoRequest
	7,  // 11: TodoService.Delete:input_type -> DeleteTodoRequest
	8,  // 12: TodoService.Create:output_type -> TodoResponse
	9,  // 13: TodoService.CreateMany:output_type -> CreateManyTodosResponse
	8,  // 14: TodoService.Read:output_type -> TodoResponse
	10, // 15: TodoService.ReadMany:output_type -> ReadManyTodosResponse
	8,  // 16: TodoService.Update:output_type -> TodoResponse
	8,  // 17: TodoService.Delete:output_type -> TodoResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_todo_v1_services_proto_init() }
func file_todo_v1_services_proto_init() {
	if File_todo_v1_services_proto != nil {
		return
	}
	file_todo_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_todo_v1_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManyTodosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadManyTodosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManyTodosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_v1_services_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadManyTodosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_v1_services_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_v1_services_proto_goTypes,
		DependencyIndexes: file_todo_v1_services_proto_depIdxs,
		EnumInfos:         file_todo_v1_services_proto_enumTypes,
		MessageInfos:      file_todo_v1_services_proto_msgTypes,
	}.Build()
	File_todo_v1_services_proto = out.File
	file_todo_v1_services_proto_rawDesc = nil
	file_todo_v1_services_proto_goTypes = nil
	file_todo_v1_services_proto_depIdxs = nil
}
