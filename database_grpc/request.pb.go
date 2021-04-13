// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: request.proto

package databaseGrpc

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateItemById  *UpdateItemById  `protobuf:"bytes,1,opt,name=updateItemById,proto3" json:"updateItemById,omitempty"`
	UpdateItemByKey *UpdateItemByKey `protobuf:"bytes,2,opt,name=updateItemByKey,proto3" json:"updateItemByKey,omitempty"`
	AddItem         *AddItem         `protobuf:"bytes,3,opt,name=addItem,proto3" json:"addItem,omitempty"`
	FindItemById    *FindItemById    `protobuf:"bytes,4,opt,name=findItemById,proto3" json:"findItemById,omitempty"`
	FindItemByKey   *FindItemByKey   `protobuf:"bytes,5,opt,name=findItemByKey,proto3" json:"findItemByKey,omitempty"`
	DeleteItemById  *DeleteItemById  `protobuf:"bytes,6,opt,name=deleteItemById,proto3" json:"deleteItemById,omitempty"`
	DeleteItemByKey *DeleteItemByKey `protobuf:"bytes,7,opt,name=deleteItemByKey,proto3" json:"deleteItemByKey,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUpdateItemById() *UpdateItemById {
	if x != nil {
		return x.UpdateItemById
	}
	return nil
}

func (x *Request) GetUpdateItemByKey() *UpdateItemByKey {
	if x != nil {
		return x.UpdateItemByKey
	}
	return nil
}

func (x *Request) GetAddItem() *AddItem {
	if x != nil {
		return x.AddItem
	}
	return nil
}

func (x *Request) GetFindItemById() *FindItemById {
	if x != nil {
		return x.FindItemById
	}
	return nil
}

func (x *Request) GetFindItemByKey() *FindItemByKey {
	if x != nil {
		return x.FindItemByKey
	}
	return nil
}

func (x *Request) GetDeleteItemById() *DeleteItemById {
	if x != nil {
		return x.DeleteItemById
	}
	return nil
}

func (x *Request) GetDeleteItemByKey() *DeleteItemByKey {
	if x != nil {
		return x.DeleteItemByKey
	}
	return nil
}

type MatchItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Match:
	//	*MatchItem_ValString
	//	*MatchItem_ValInt
	//	*MatchItem_ValBool
	//	*MatchItem_ValFloat
	//	*MatchItem_ValDouble
	Match isMatchItem_Match `protobuf_oneof:"match"`
}

func (x *MatchItem) Reset() {
	*x = MatchItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchItem) ProtoMessage() {}

func (x *MatchItem) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchItem.ProtoReflect.Descriptor instead.
func (*MatchItem) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *MatchItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *MatchItem) GetMatch() isMatchItem_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *MatchItem) GetValString() string {
	if x, ok := x.GetMatch().(*MatchItem_ValString); ok {
		return x.ValString
	}
	return ""
}

func (x *MatchItem) GetValInt() int64 {
	if x, ok := x.GetMatch().(*MatchItem_ValInt); ok {
		return x.ValInt
	}
	return 0
}

func (x *MatchItem) GetValBool() bool {
	if x, ok := x.GetMatch().(*MatchItem_ValBool); ok {
		return x.ValBool
	}
	return false
}

func (x *MatchItem) GetValFloat() float32 {
	if x, ok := x.GetMatch().(*MatchItem_ValFloat); ok {
		return x.ValFloat
	}
	return 0
}

func (x *MatchItem) GetValDouble() float64 {
	if x, ok := x.GetMatch().(*MatchItem_ValDouble); ok {
		return x.ValDouble
	}
	return 0
}

type isMatchItem_Match interface {
	isMatchItem_Match()
}

type MatchItem_ValString struct {
	ValString string `protobuf:"bytes,2,opt,name=valString,proto3,oneof"`
}

type MatchItem_ValInt struct {
	ValInt int64 `protobuf:"varint,3,opt,name=valInt,proto3,oneof"`
}

type MatchItem_ValBool struct {
	ValBool bool `protobuf:"varint,4,opt,name=valBool,proto3,oneof"`
}

type MatchItem_ValFloat struct {
	ValFloat float32 `protobuf:"fixed32,5,opt,name=valFloat,proto3,oneof"`
}

type MatchItem_ValDouble struct {
	ValDouble float64 `protobuf:"fixed64,6,opt,name=valDouble,proto3,oneof"`
}

func (*MatchItem_ValString) isMatchItem_Match() {}

func (*MatchItem_ValInt) isMatchItem_Match() {}

func (*MatchItem_ValBool) isMatchItem_Match() {}

func (*MatchItem_ValFloat) isMatchItem_Match() {}

func (*MatchItem_ValDouble) isMatchItem_Match() {}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

type UpdateItemById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId   string                `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	UpdateItem ITEM                  `protobuf:"varint,2,opt,name=updateItem,proto3,enum=databaseGrpc.ITEM" json:"updateItem,omitempty"` // 更新的对象
	Items      map[string]*anypb.Any `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateItemById) Reset() {
	*x = UpdateItemById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemById) ProtoMessage() {}

func (x *UpdateItemById) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemById.ProtoReflect.Descriptor instead.
func (*UpdateItemById) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateItemById) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *UpdateItemById) GetUpdateItem() ITEM {
	if x != nil {
		return x.UpdateItem
	}
	return ITEM_PLAYER
}

func (x *UpdateItemById) GetItems() map[string]*anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateItemByKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateItem ITEM                  `protobuf:"varint,1,opt,name=updateItem,proto3,enum=databaseGrpc.ITEM" json:"updateItem,omitempty"`
	MatchItem  []*MatchItem          `protobuf:"bytes,2,rep,name=matchItem,proto3" json:"matchItem,omitempty"`
	Item       map[string]*anypb.Any `protobuf:"bytes,3,rep,name=item,proto3" json:"item,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateItemByKey) Reset() {
	*x = UpdateItemByKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemByKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemByKey) ProtoMessage() {}

func (x *UpdateItemByKey) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemByKey.ProtoReflect.Descriptor instead.
func (*UpdateItemByKey) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItemByKey) GetUpdateItem() ITEM {
	if x != nil {
		return x.UpdateItem
	}
	return ITEM_PLAYER
}

func (x *UpdateItemByKey) GetMatchItem() []*MatchItem {
	if x != nil {
		return x.MatchItem
	}
	return nil
}

func (x *UpdateItemByKey) GetItem() map[string]*anypb.Any {
	if x != nil {
		return x.Item
	}
	return nil
}

type AddItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddItem ITEM       `protobuf:"varint,1,opt,name=addItem,proto3,enum=databaseGrpc.ITEM" json:"addItem,omitempty"`
	Item    *anypb.Any `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *AddItem) Reset() {
	*x = AddItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItem) ProtoMessage() {}

func (x *AddItem) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItem.ProtoReflect.Descriptor instead.
func (*AddItem) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{5}
}

func (x *AddItem) GetAddItem() ITEM {
	if x != nil {
		return x.AddItem
	}
	return ITEM_PLAYER
}

func (x *AddItem) GetItem() *anypb.Any {
	if x != nil {
		return x.Item
	}
	return nil
}

type FindItemById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FindItem ITEM   `protobuf:"varint,1,opt,name=findItem,proto3,enum=databaseGrpc.ITEM" json:"findItem,omitempty"`
	ItemId   string `protobuf:"bytes,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *FindItemById) Reset() {
	*x = FindItemById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindItemById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindItemById) ProtoMessage() {}

func (x *FindItemById) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindItemById.ProtoReflect.Descriptor instead.
func (*FindItemById) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{6}
}

func (x *FindItemById) GetFindItem() ITEM {
	if x != nil {
		return x.FindItem
	}
	return ITEM_PLAYER
}

func (x *FindItemById) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type FindItemByKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FindItem  ITEM         `protobuf:"varint,1,opt,name=FindItem,proto3,enum=databaseGrpc.ITEM" json:"FindItem,omitempty"`
	MatchItem []*MatchItem `protobuf:"bytes,2,rep,name=matchItem,proto3" json:"matchItem,omitempty"`
}

func (x *FindItemByKey) Reset() {
	*x = FindItemByKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindItemByKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindItemByKey) ProtoMessage() {}

func (x *FindItemByKey) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindItemByKey.ProtoReflect.Descriptor instead.
func (*FindItemByKey) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{7}
}

func (x *FindItemByKey) GetFindItem() ITEM {
	if x != nil {
		return x.FindItem
	}
	return ITEM_PLAYER
}

func (x *FindItemByKey) GetMatchItem() []*MatchItem {
	if x != nil {
		return x.MatchItem
	}
	return nil
}

type DeleteItemById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteItem ITEM   `protobuf:"varint,1,opt,name=deleteItem,proto3,enum=databaseGrpc.ITEM" json:"deleteItem,omitempty"`
	ItemId     string `protobuf:"bytes,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *DeleteItemById) Reset() {
	*x = DeleteItemById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemById) ProtoMessage() {}

func (x *DeleteItemById) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemById.ProtoReflect.Descriptor instead.
func (*DeleteItemById) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemById) GetDeleteItem() ITEM {
	if x != nil {
		return x.DeleteItem
	}
	return ITEM_PLAYER
}

func (x *DeleteItemById) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type DeleteItemByKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteItem ITEM         `protobuf:"varint,1,opt,name=DeleteItem,proto3,enum=databaseGrpc.ITEM" json:"DeleteItem,omitempty"`
	MatchItem  []*MatchItem `protobuf:"bytes,2,rep,name=matchItem,proto3" json:"matchItem,omitempty"`
}

func (x *DeleteItemByKey) Reset() {
	*x = DeleteItemByKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemByKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemByKey) ProtoMessage() {}

func (x *DeleteItemByKey) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemByKey.ProtoReflect.Descriptor instead.
func (*DeleteItemByKey) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemByKey) GetDeleteItem() ITEM {
	if x != nil {
		return x.DeleteItem
	}
	return ITEM_PLAYER
}

func (x *DeleteItemByKey) GetMatchItem() []*MatchItem {
	if x != nil {
		return x.MatchItem
	}
	return nil
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x1a, 0x0f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79,
	0x4b, 0x65, 0x79, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x42, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x0e,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x47,
	0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x49, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x49, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x76, 0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x0a,
	0x08, 0x76, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x00, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x09, 0x76,
	0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x1b, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f,
	0x70, 0x22, 0xef, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x4e, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x88, 0x02, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a,
	0x4d, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x61,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x56, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x0d, 0x46, 0x69, 0x6e,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d,
	0x52, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x09, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65,
	0x6d, 0x22, 0x5c, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22,
	0x7c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x32, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x54, 0x45, 0x4d, 0x52, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_request_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: databaseGrpc.Request
	(*MatchItem)(nil),       // 1: databaseGrpc.MatchItem
	(*Operation)(nil),       // 2: databaseGrpc.Operation
	(*UpdateItemById)(nil),  // 3: databaseGrpc.UpdateItemById
	(*UpdateItemByKey)(nil), // 4: databaseGrpc.UpdateItemByKey
	(*AddItem)(nil),         // 5: databaseGrpc.AddItem
	(*FindItemById)(nil),    // 6: databaseGrpc.FindItemById
	(*FindItemByKey)(nil),   // 7: databaseGrpc.FindItemByKey
	(*DeleteItemById)(nil),  // 8: databaseGrpc.DeleteItemById
	(*DeleteItemByKey)(nil), // 9: databaseGrpc.DeleteItemByKey
	nil,                     // 10: databaseGrpc.UpdateItemById.ItemsEntry
	nil,                     // 11: databaseGrpc.UpdateItemByKey.ItemEntry
	(ITEM)(0),               // 12: databaseGrpc.ITEM
	(*anypb.Any)(nil),       // 13: google.protobuf.Any
}
var file_request_proto_depIdxs = []int32{
	3,  // 0: databaseGrpc.Request.updateItemById:type_name -> databaseGrpc.UpdateItemById
	4,  // 1: databaseGrpc.Request.updateItemByKey:type_name -> databaseGrpc.UpdateItemByKey
	5,  // 2: databaseGrpc.Request.addItem:type_name -> databaseGrpc.AddItem
	6,  // 3: databaseGrpc.Request.findItemById:type_name -> databaseGrpc.FindItemById
	7,  // 4: databaseGrpc.Request.findItemByKey:type_name -> databaseGrpc.FindItemByKey
	8,  // 5: databaseGrpc.Request.deleteItemById:type_name -> databaseGrpc.DeleteItemById
	9,  // 6: databaseGrpc.Request.deleteItemByKey:type_name -> databaseGrpc.DeleteItemByKey
	12, // 7: databaseGrpc.UpdateItemById.updateItem:type_name -> databaseGrpc.ITEM
	10, // 8: databaseGrpc.UpdateItemById.items:type_name -> databaseGrpc.UpdateItemById.ItemsEntry
	12, // 9: databaseGrpc.UpdateItemByKey.updateItem:type_name -> databaseGrpc.ITEM
	1,  // 10: databaseGrpc.UpdateItemByKey.matchItem:type_name -> databaseGrpc.MatchItem
	11, // 11: databaseGrpc.UpdateItemByKey.item:type_name -> databaseGrpc.UpdateItemByKey.ItemEntry
	12, // 12: databaseGrpc.AddItem.addItem:type_name -> databaseGrpc.ITEM
	13, // 13: databaseGrpc.AddItem.item:type_name -> google.protobuf.Any
	12, // 14: databaseGrpc.FindItemById.findItem:type_name -> databaseGrpc.ITEM
	12, // 15: databaseGrpc.FindItemByKey.FindItem:type_name -> databaseGrpc.ITEM
	1,  // 16: databaseGrpc.FindItemByKey.matchItem:type_name -> databaseGrpc.MatchItem
	12, // 17: databaseGrpc.DeleteItemById.deleteItem:type_name -> databaseGrpc.ITEM
	12, // 18: databaseGrpc.DeleteItemByKey.DeleteItem:type_name -> databaseGrpc.ITEM
	1,  // 19: databaseGrpc.DeleteItemByKey.matchItem:type_name -> databaseGrpc.MatchItem
	13, // 20: databaseGrpc.UpdateItemById.ItemsEntry.value:type_name -> google.protobuf.Any
	13, // 21: databaseGrpc.UpdateItemByKey.ItemEntry.value:type_name -> google.protobuf.Any
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	file_type_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchItem); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemById); i {
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
		file_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemByKey); i {
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
		file_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItem); i {
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
		file_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindItemById); i {
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
		file_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindItemByKey); i {
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
		file_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemById); i {
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
		file_request_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemByKey); i {
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
	file_request_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MatchItem_ValString)(nil),
		(*MatchItem_ValInt)(nil),
		(*MatchItem_ValBool)(nil),
		(*MatchItem_ValFloat)(nil),
		(*MatchItem_ValDouble)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
