// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/book/v1/book.proto

package v1

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SaleInfo *SaleInfo `protobuf:"bytes,3,opt,name=saleInfo,proto3" json:"saleInfo,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetSaleInfo() *SaleInfo {
	if x != nil {
		return x.SaleInfo
	}
	return nil
}

type SaleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaledAt      string `protobuf:"bytes,1,opt,name=saledAt,proto3" json:"saledAt,omitempty"`
	CustomerId   int64  `protobuf:"varint,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	CustomerName string `protobuf:"bytes,3,opt,name=customerName,proto3" json:"customerName,omitempty"`
}

func (x *SaleInfo) Reset() {
	*x = SaleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleInfo) ProtoMessage() {}

func (x *SaleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleInfo.ProtoReflect.Descriptor instead.
func (*SaleInfo) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{1}
}

func (x *SaleInfo) GetSaledAt() string {
	if x != nil {
		return x.SaledAt
	}
	return ""
}

func (x *SaleInfo) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SaleInfo) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

type BookReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *Book  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BookReply) Reset() {
	*x = BookReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookReply) ProtoMessage() {}

func (x *BookReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookReply.ProtoReflect.Descriptor instead.
func (*BookReply) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookReply) GetData() *Book {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BookReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FindBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindBookRequest) Reset() {
	*x = FindBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookRequest) ProtoMessage() {}

func (x *FindBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookRequest.ProtoReflect.Descriptor instead.
func (*FindBookRequest) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{3}
}

func (x *FindBookRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SaleBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId int64 `protobuf:"varint,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *SaleBookRequest) Reset() {
	*x = SaleBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleBookRequest) ProtoMessage() {}

func (x *SaleBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleBookRequest.ProtoReflect.Descriptor instead.
func (*SaleBookRequest) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{4}
}

func (x *SaleBookRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaleBookRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type NewBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewBookRequest) Reset() {
	*x = NewBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBookRequest) ProtoMessage() {}

func (x *NewBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBookRequest.ProtoReflect.Descriptor instead.
func (*NewBookRequest) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{5}
}

func (x *NewBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBookReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBookReply) Reset() {
	*x = DeleteBookReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_book_v1_book_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookReply) ProtoMessage() {}

func (x *DeleteBookReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_book_v1_book_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookReply.ProtoReflect.Descriptor instead.
func (*DeleteBookReply) Descriptor() ([]byte, []int) {
	return file_api_book_v1_book_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_book_v1_book_proto protoreflect.FileDescriptor

var file_api_book_v1_book_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x22, 0x59, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x68, 0x0a, 0x08,
	0x53, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x21, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0f, 0x53, 0x61, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x85,
	0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x61,
	0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_book_v1_book_proto_rawDescOnce sync.Once
	file_api_book_v1_book_proto_rawDescData = file_api_book_v1_book_proto_rawDesc
)

func file_api_book_v1_book_proto_rawDescGZIP() []byte {
	file_api_book_v1_book_proto_rawDescOnce.Do(func() {
		file_api_book_v1_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_book_v1_book_proto_rawDescData)
	})
	return file_api_book_v1_book_proto_rawDescData
}

var file_api_book_v1_book_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_book_v1_book_proto_goTypes = []interface{}{
	(*Book)(nil),              // 0: book.v1.Book
	(*SaleInfo)(nil),          // 1: book.v1.SaleInfo
	(*BookReply)(nil),         // 2: book.v1.BookReply
	(*FindBookRequest)(nil),   // 3: book.v1.FindBookRequest
	(*SaleBookRequest)(nil),   // 4: book.v1.SaleBookRequest
	(*NewBookRequest)(nil),    // 5: book.v1.NewBookRequest
	(*DeleteBookRequest)(nil), // 6: book.v1.DeleteBookRequest
	(*DeleteBookReply)(nil),   // 7: book.v1.DeleteBookReply
}
var file_api_book_v1_book_proto_depIdxs = []int32{
	1, // 0: book.v1.Book.saleInfo:type_name -> book.v1.SaleInfo
	0, // 1: book.v1.BookReply.data:type_name -> book.v1.Book
	3, // 2: book.v1.BookService.FindBook:input_type -> book.v1.FindBookRequest
	4, // 3: book.v1.BookService.SaleBook:input_type -> book.v1.SaleBookRequest
	5, // 4: book.v1.BookService.NewBook:input_type -> book.v1.NewBookRequest
	6, // 5: book.v1.BookService.DeleteBook:input_type -> book.v1.DeleteBookRequest
	2, // 6: book.v1.BookService.FindBook:output_type -> book.v1.BookReply
	2, // 7: book.v1.BookService.SaleBook:output_type -> book.v1.BookReply
	2, // 8: book.v1.BookService.NewBook:output_type -> book.v1.BookReply
	7, // 9: book.v1.BookService.DeleteBook:output_type -> book.v1.DeleteBookReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_book_v1_book_proto_init() }
func file_api_book_v1_book_proto_init() {
	if File_api_book_v1_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_book_v1_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_api_book_v1_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleInfo); i {
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
		file_api_book_v1_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookReply); i {
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
		file_api_book_v1_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBookRequest); i {
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
		file_api_book_v1_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleBookRequest); i {
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
		file_api_book_v1_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBookRequest); i {
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
		file_api_book_v1_book_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
		file_api_book_v1_book_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookReply); i {
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
			RawDescriptor: file_api_book_v1_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_book_v1_book_proto_goTypes,
		DependencyIndexes: file_api_book_v1_book_proto_depIdxs,
		MessageInfos:      file_api_book_v1_book_proto_msgTypes,
	}.Build()
	File_api_book_v1_book_proto = out.File
	file_api_book_v1_book_proto_rawDesc = nil
	file_api_book_v1_book_proto_goTypes = nil
	file_api_book_v1_book_proto_depIdxs = nil
}