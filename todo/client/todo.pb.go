// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Todo struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// @inject_tag: sql:",notnull,default:false"
	Completed bool `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateTodoRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{1}
}
func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(dst, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{2}
}
func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(dst, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTodosRequest struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodosRequest) Reset()         { *m = CreateTodosRequest{} }
func (m *CreateTodosRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodosRequest) ProtoMessage()    {}
func (*CreateTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{3}
}
func (m *CreateTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodosRequest.Unmarshal(m, b)
}
func (m *CreateTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodosRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodosRequest.Merge(dst, src)
}
func (m *CreateTodosRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodosRequest.Size(m)
}
func (m *CreateTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodosRequest proto.InternalMessageInfo

func (m *CreateTodosRequest) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateTodosResponse struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodosResponse) Reset()         { *m = CreateTodosResponse{} }
func (m *CreateTodosResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodosResponse) ProtoMessage()    {}
func (*CreateTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{4}
}
func (m *CreateTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodosResponse.Unmarshal(m, b)
}
func (m *CreateTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodosResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodosResponse.Merge(dst, src)
}
func (m *CreateTodosResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodosResponse.Size(m)
}
func (m *CreateTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodosResponse proto.InternalMessageInfo

func (m *CreateTodosResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{5}
}
func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (dst *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(dst, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoResponse struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{6}
}
func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (dst *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(dst, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type ListTodoRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	NotCompleted         bool     `protobuf:"varint,2,opt,name=not_completed,json=notCompleted,proto3" json:"not_completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRequest) Reset()         { *m = ListTodoRequest{} }
func (m *ListTodoRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoRequest) ProtoMessage()    {}
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{7}
}
func (m *ListTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRequest.Unmarshal(m, b)
}
func (m *ListTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRequest.Marshal(b, m, deterministic)
}
func (dst *ListTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRequest.Merge(dst, src)
}
func (m *ListTodoRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoRequest.Size(m)
}
func (m *ListTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRequest proto.InternalMessageInfo

func (m *ListTodoRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListTodoRequest) GetNotCompleted() bool {
	if m != nil {
		return m.NotCompleted
	}
	return false
}

type ListTodoResponse struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{8}
}
func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(dst, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{9}
}
func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(dst, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{10}
}
func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(dst, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

type UpdateTodoRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRequest) Reset()         { *m = UpdateTodoRequest{} }
func (m *UpdateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRequest) ProtoMessage()    {}
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{11}
}
func (m *UpdateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRequest.Unmarshal(m, b)
}
func (m *UpdateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRequest.Merge(dst, src)
}
func (m *UpdateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRequest.Size(m)
}
func (m *UpdateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRequest proto.InternalMessageInfo

func (m *UpdateTodoRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoResponse) Reset()         { *m = UpdateTodoResponse{} }
func (m *UpdateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoResponse) ProtoMessage()    {}
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{12}
}
func (m *UpdateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoResponse.Unmarshal(m, b)
}
func (m *UpdateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoResponse.Merge(dst, src)
}
func (m *UpdateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoResponse.Size(m)
}
func (m *UpdateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoResponse proto.InternalMessageInfo

type UpdateTodosRequest struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodosRequest) Reset()         { *m = UpdateTodosRequest{} }
func (m *UpdateTodosRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodosRequest) ProtoMessage()    {}
func (*UpdateTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{13}
}
func (m *UpdateTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodosRequest.Unmarshal(m, b)
}
func (m *UpdateTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodosRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodosRequest.Merge(dst, src)
}
func (m *UpdateTodosRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodosRequest.Size(m)
}
func (m *UpdateTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodosRequest proto.InternalMessageInfo

func (m *UpdateTodosRequest) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateTodosResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodosResponse) Reset()         { *m = UpdateTodosResponse{} }
func (m *UpdateTodosResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodosResponse) ProtoMessage()    {}
func (*UpdateTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_7c24c30d5b413c68, []int{14}
}
func (m *UpdateTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodosResponse.Unmarshal(m, b)
}
func (m *UpdateTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodosResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodosResponse.Merge(dst, src)
}
func (m *UpdateTodosResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodosResponse.Size(m)
}
func (m *UpdateTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodosResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Todo)(nil), "todo.client.v1.Todo")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.client.v1.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.client.v1.CreateTodoResponse")
	proto.RegisterType((*CreateTodosRequest)(nil), "todo.client.v1.CreateTodosRequest")
	proto.RegisterType((*CreateTodosResponse)(nil), "todo.client.v1.CreateTodosResponse")
	proto.RegisterType((*GetTodoRequest)(nil), "todo.client.v1.GetTodoRequest")
	proto.RegisterType((*GetTodoResponse)(nil), "todo.client.v1.GetTodoResponse")
	proto.RegisterType((*ListTodoRequest)(nil), "todo.client.v1.ListTodoRequest")
	proto.RegisterType((*ListTodoResponse)(nil), "todo.client.v1.ListTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "todo.client.v1.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.client.v1.DeleteTodoResponse")
	proto.RegisterType((*UpdateTodoRequest)(nil), "todo.client.v1.UpdateTodoRequest")
	proto.RegisterType((*UpdateTodoResponse)(nil), "todo.client.v1.UpdateTodoResponse")
	proto.RegisterType((*UpdateTodosRequest)(nil), "todo.client.v1.UpdateTodosRequest")
	proto.RegisterType((*UpdateTodosResponse)(nil), "todo.client.v1.UpdateTodosResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	// Bulk version of CreateTodo
	CreateTodos(ctx context.Context, in *CreateTodosRequest, opts ...grpc.CallOption) (*CreateTodosResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
	ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
	UpdateTodos(ctx context.Context, in *UpdateTodosRequest, opts ...grpc.CallOption) (*UpdateTodosResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateTodos(ctx context.Context, in *CreateTodosRequest, opts ...grpc.CallOption) (*CreateTodosResponse, error) {
	out := new(CreateTodosResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/CreateTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error) {
	out := new(ListTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodos(ctx context.Context, in *UpdateTodosRequest, opts ...grpc.CallOption) (*UpdateTodosResponse, error) {
	out := new(UpdateTodosResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v1.TodoService/UpdateTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	// Bulk version of CreateTodo
	CreateTodos(context.Context, *CreateTodosRequest) (*CreateTodosResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
	ListTodo(context.Context, *ListTodoRequest) (*ListTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error)
	UpdateTodos(context.Context, *UpdateTodosRequest) (*UpdateTodosResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/CreateTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodos(ctx, req.(*CreateTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v1.TodoService/UpdateTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodos(ctx, req.(*UpdateTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.client.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "CreateTodos",
			Handler:    _TodoService_CreateTodos_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "UpdateTodos",
			Handler:    _TodoService_UpdateTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_7c24c30d5b413c68) }

var fileDescriptor_todo_7c24c30d5b413c68 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe5, 0x24, 0xce, 0x3f, 0x99, 0xfc, 0x9b, 0xa4, 0xd3, 0x14, 0x59, 0x56, 0x45, 0x8d,
	0x83, 0x44, 0x94, 0x83, 0xad, 0x86, 0x13, 0x45, 0xa0, 0x96, 0x22, 0x71, 0xe9, 0xc9, 0x94, 0x0b,
	0x97, 0xc8, 0x89, 0x97, 0x6a, 0x89, 0xe3, 0x35, 0xf1, 0x26, 0x17, 0xc4, 0x85, 0x57, 0xe0, 0xd1,
	0x78, 0x03, 0xc4, 0x5b, 0x70, 0x41, 0x5e, 0xaf, 0x6b, 0xc7, 0x4e, 0x0c, 0xf4, 0x16, 0xef, 0x7e,
	0xf3, 0x7d, 0x33, 0x3b, 0x3f, 0x05, 0x80, 0x33, 0x8f, 0x59, 0xe1, 0x8a, 0x71, 0x86, 0x5d, 0xf1,
	0x7b, 0xee, 0x53, 0x12, 0x70, 0x6b, 0x73, 0xa6, 0x9f, 0xdc, 0x32, 0x76, 0xeb, 0x13, 0xdb, 0x0d,
	0xa9, 0xed, 0x06, 0x01, 0xe3, 0x2e, 0xa7, 0x2c, 0x88, 0x12, 0xb5, 0x7e, 0x2a, 0x6f, 0xc5, 0xd7,
	0x6c, 0xfd, 0xc1, 0xe6, 0x74, 0x49, 0x22, 0xee, 0x2e, 0xc3, 0x44, 0x60, 0xfe, 0x50, 0xa0, 0x71,
	0xc3, 0x3c, 0x86, 0x5d, 0xa8, 0x51, 0x4f, 0x53, 0x0c, 0x65, 0xd4, 0x76, 0x6a, 0xd4, 0xc3, 0x01,
	0xa8, 0x9c, 0x72, 0x9f, 0x68, 0x35, 0x71, 0x94, 0x7c, 0xa0, 0x01, 0x1d, 0x8f, 0x44, 0xf3, 0x15,
	0x0d, 0xe3, 0x14, 0xad, 0x2e, 0xee, 0xf2, 0x47, 0x78, 0x02, 0xed, 0x39, 0x5b, 0x86, 0x3e, 0xe1,
	0xc4, 0xd3, 0x1a, 0x86, 0x32, 0x6a, 0x39, 0xd9, 0x01, 0x3e, 0x03, 0x98, 0xaf, 0x88, 0xcb, 0x89,
	0x37, 0x75, 0xb9, 0xa6, 0x1a, 0xca, 0xa8, 0x33, 0xd1, 0xad, 0xa4, 0x49, 0x2b, 0x6d, 0xd2, 0xba,
	0x49, 0x9b, 0x74, 0xda, 0x52, 0x7d, 0xc9, 0xe3, 0xd2, 0x75, 0xe8, 0xa5, 0xa5, 0xcd, 0x3f, 0x97,
	0x4a, 0xf5, 0x25, 0x37, 0x5f, 0xc0, 0xe1, 0x95, 0xf0, 0x89, 0x27, 0x75, 0xc8, 0xa7, 0x35, 0x89,
	0x38, 0x8e, 0xa0, 0x41, 0x39, 0x59, 0x8a, 0x91, 0x3b, 0x93, 0x81, 0xb5, 0xfd, 0xae, 0x96, 0x90,
	0x0a, 0x85, 0xf9, 0x18, 0x30, 0x5f, 0x1e, 0x85, 0x2c, 0x88, 0x48, 0xf1, 0xc1, 0xcc, 0x8b, 0xbc,
	0x2a, 0x4a, 0x53, 0xc6, 0xa0, 0xc6, 0x1e, 0x91, 0xa6, 0x18, 0xf5, 0xbd, 0x31, 0x89, 0xc4, 0x7c,
	0x02, 0x47, 0x5b, 0x0e, 0x32, 0xa8, 0x0f, 0x75, 0xea, 0x25, 0x06, 0x6d, 0x27, 0xfe, 0x69, 0x1a,
	0xd0, 0x7d, 0x43, 0x78, 0x7e, 0x98, 0x62, 0x33, 0xcf, 0xa1, 0x77, 0xa7, 0x90, 0x36, 0x7f, 0x3f,
	0xef, 0x35, 0xf4, 0xae, 0x69, 0xb4, 0xe5, 0x3f, 0x00, 0xd5, 0xa7, 0x4b, 0xca, 0x45, 0xb5, 0xea,
	0x24, 0x1f, 0x38, 0x84, 0x83, 0x80, 0xf1, 0x69, 0xb6, 0xef, 0x9a, 0xd8, 0xf7, 0xff, 0x01, 0xe3,
	0x57, 0xe9, 0x99, 0xf9, 0x12, 0xfa, 0x99, 0x9b, 0xec, 0xe5, 0x5f, 0x5e, 0x65, 0x08, 0x87, 0xaf,
	0x49, 0x6c, 0x55, 0x35, 0xef, 0x00, 0x30, 0x2f, 0x4a, 0x62, 0xe2, 0xbd, 0xbf, 0x13, 0x10, 0xdc,
	0x6f, 0xef, 0x03, 0xc0, 0x7c, 0xb9, 0x34, 0xbd, 0xc8, 0x9f, 0xde, 0x6b, 0xcf, 0xc7, 0x70, 0xb4,
	0xe5, 0x90, 0x18, 0x4f, 0x7e, 0xa9, 0xd0, 0x89, 0x4f, 0xde, 0x92, 0xd5, 0x86, 0xce, 0x09, 0x2e,
	0x00, 0x32, 0x1c, 0xf0, 0x51, 0xd1, 0xb1, 0x44, 0xb4, 0x6e, 0x56, 0x49, 0x64, 0xf7, 0x0f, 0xbe,
	0x7e, 0xff, 0xf9, 0xad, 0xd6, 0x37, 0x5b, 0xf6, 0xe6, 0xcc, 0x8e, 0xe5, 0xe7, 0x62, 0x56, 0x0c,
	0xa1, 0x93, 0x63, 0x0f, 0x2b, 0xac, 0xd2, 0x91, 0xf5, 0x61, 0xa5, 0x46, 0xe6, 0x69, 0x22, 0x0f,
	0xcd, 0x83, 0x34, 0xcf, 0x9e, 0xad, 0xfd, 0xc5, 0xb9, 0x32, 0x46, 0x17, 0xfe, 0x93, 0x88, 0xe2,
	0xc3, 0xa2, 0xd3, 0x36, 0xdd, 0xfa, 0xe9, 0xde, 0x7b, 0x99, 0x72, 0x2c, 0x52, 0x7a, 0x98, 0xa5,
	0x7c, 0xa6, 0xde, 0x17, 0x9c, 0x42, 0x2b, 0x45, 0x0f, 0x4b, 0x1e, 0x05, 0xc4, 0x75, 0x63, 0xbf,
	0x40, 0xa6, 0xf4, 0x45, 0x0a, 0xe0, 0xdd, 0xdb, 0xe1, 0x47, 0x80, 0x0c, 0xbb, 0xf2, 0x8a, 0x4a,
	0xdc, 0x96, 0x57, 0xb4, 0x83, 0x5a, 0x39, 0xcc, 0xb8, 0x30, 0xcc, 0x02, 0x20, 0xa3, 0xa6, 0x9c,
	0x55, 0x02, 0xbd, 0x9c, 0xb5, 0x03, 0x66, 0x89, 0x83, 0xbe, 0x03, 0x87, 0x1c, 0xa2, 0x58, 0x61,
	0xb5, 0x1f, 0x87, 0x1d, 0x8c, 0xa7, 0x38, 0xe8, 0x25, 0x1c, 0x5e, 0xb5, 0xde, 0x37, 0x93, 0xd2,
	0x59, 0x53, 0xfc, 0x99, 0x3f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x21, 0xce, 0x68, 0x1d, 0xf6,
	0x06, 0x00, 0x00,
}