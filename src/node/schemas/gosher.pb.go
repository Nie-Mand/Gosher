// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/gosher.proto

package schemas

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

type FileSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileBytes   string `protobuf:"bytes,1,opt,name=fileBytes,proto3" json:"fileBytes,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *FileSendRequest) Reset() {
	*x = FileSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSendRequest) ProtoMessage() {}

func (x *FileSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSendRequest.ProtoReflect.Descriptor instead.
func (*FileSendRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{0}
}

func (x *FileSendRequest) GetFileBytes() string {
	if x != nil {
		return x.FileBytes
	}
	return ""
}

func (x *FileSendRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	File        []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[1]
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
	return file_proto_gosher_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Request) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	File []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{3}
}

type PingForFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PingForFileRequest) Reset() {
	*x = PingForFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingForFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingForFileRequest) ProtoMessage() {}

func (x *PingForFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingForFileRequest.ProtoReflect.Descriptor instead.
func (*PingForFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{4}
}

func (x *PingForFileRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PingForFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Who      string `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *PingForFileResponse) Reset() {
	*x = PingForFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingForFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingForFileResponse) ProtoMessage() {}

func (x *PingForFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingForFileResponse.ProtoReflect.Descriptor instead.
func (*PingForFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{5}
}

func (x *PingForFileResponse) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *PingForFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type ListenForFilePingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName    string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListenForFilePingsRequest) Reset() {
	*x = ListenForFilePingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenForFilePingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenForFilePingsRequest) ProtoMessage() {}

func (x *ListenForFilePingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenForFilePingsRequest.ProtoReflect.Descriptor instead.
func (*ListenForFilePingsRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{6}
}

func (x *ListenForFilePingsRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ListenForFilePingsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ListenForFilePingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListenForFilePingsResponse) Reset() {
	*x = ListenForFilePingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenForFilePingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenForFilePingsResponse) ProtoMessage() {}

func (x *ListenForFilePingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenForFilePingsResponse.ProtoReflect.Descriptor instead.
func (*ListenForFilePingsResponse) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{7}
}

func (x *ListenForFilePingsResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RequestFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Who      string `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *RequestFileRequest) Reset() {
	*x = RequestFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFileRequest) ProtoMessage() {}

func (x *RequestFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFileRequest.ProtoReflect.Descriptor instead.
func (*RequestFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{8}
}

func (x *RequestFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *RequestFileRequest) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

type RequestFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	File []byte `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *RequestFileResponse) Reset() {
	*x = RequestFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFileResponse) ProtoMessage() {}

func (x *RequestFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFileResponse.ProtoReflect.Descriptor instead.
func (*RequestFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{9}
}

func (x *RequestFileResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RequestFileResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type SeedFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	File []byte `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *SeedFileRequest) Reset() {
	*x = SeedFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedFileRequest) ProtoMessage() {}

func (x *SeedFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedFileRequest.ProtoReflect.Descriptor instead.
func (*SeedFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{10}
}

func (x *SeedFileRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SeedFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type SeedFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Who      string `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *SeedFileResponse) Reset() {
	*x = SeedFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gosher_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedFileResponse) ProtoMessage() {}

func (x *SeedFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gosher_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedFileResponse.ProtoReflect.Descriptor instead.
func (*SeedFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_gosher_proto_rawDescGZIP(), []int{11}
}

func (x *SeedFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SeedFileResponse) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

var File_proto_gosher_proto protoreflect.FileDescriptor

var file_proto_gosher_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x22, 0x51, 0x0a,
	0x0f, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x65, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x0a, 0x0a,
	0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x12, 0x50, 0x69, 0x6e,
	0x67, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x43, 0x0a, 0x13, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3e, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x42, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x77, 0x68, 0x6f, 0x22, 0x3b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x53, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x53,
	0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77,
	0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x32, 0xb5, 0x03,
	0x0a, 0x06, 0x47, 0x6f, 0x73, 0x68, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48,
	0x69, 0x12, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x48, 0x69, 0x12, 0x11, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x63, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x4c, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45,
	0x0a, 0x08, 0x53, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53,
	0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gosher_proto_rawDescOnce sync.Once
	file_proto_gosher_proto_rawDescData = file_proto_gosher_proto_rawDesc
)

func file_proto_gosher_proto_rawDescGZIP() []byte {
	file_proto_gosher_proto_rawDescOnce.Do(func() {
		file_proto_gosher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gosher_proto_rawDescData)
	})
	return file_proto_gosher_proto_rawDescData
}

var file_proto_gosher_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_gosher_proto_goTypes = []interface{}{
	(*FileSendRequest)(nil),            // 0: schemas.FileSendRequest
	(*Request)(nil),                    // 1: schemas.Request
	(*Response)(nil),                   // 2: schemas.Response
	(*Identity)(nil),                   // 3: schemas.Identity
	(*PingForFileRequest)(nil),         // 4: schemas.PingForFileRequest
	(*PingForFileResponse)(nil),        // 5: schemas.PingForFileResponse
	(*ListenForFilePingsRequest)(nil),  // 6: schemas.ListenForFilePingsRequest
	(*ListenForFilePingsResponse)(nil), // 7: schemas.ListenForFilePingsResponse
	(*RequestFileRequest)(nil),         // 8: schemas.RequestFileRequest
	(*RequestFileResponse)(nil),        // 9: schemas.RequestFileResponse
	(*SeedFileRequest)(nil),            // 10: schemas.SeedFileRequest
	(*SeedFileResponse)(nil),           // 11: schemas.SeedFileResponse
}
var file_proto_gosher_proto_depIdxs = []int32{
	1,  // 0: schemas.Gosher.SayHi:input_type -> schemas.Request
	3,  // 1: schemas.Gosher.ReceiveHi:input_type -> schemas.Identity
	4,  // 2: schemas.Gosher.PingForFile:input_type -> schemas.PingForFileRequest
	6,  // 3: schemas.Gosher.ListenForFilePings:input_type -> schemas.ListenForFilePingsRequest
	8,  // 4: schemas.Gosher.RequestFile:input_type -> schemas.RequestFileRequest
	10, // 5: schemas.Gosher.SeedFile:input_type -> schemas.SeedFileRequest
	2,  // 6: schemas.Gosher.SayHi:output_type -> schemas.Response
	2,  // 7: schemas.Gosher.ReceiveHi:output_type -> schemas.Response
	5,  // 8: schemas.Gosher.PingForFile:output_type -> schemas.PingForFileResponse
	7,  // 9: schemas.Gosher.ListenForFilePings:output_type -> schemas.ListenForFilePingsResponse
	9,  // 10: schemas.Gosher.RequestFile:output_type -> schemas.RequestFileResponse
	11, // 11: schemas.Gosher.SeedFile:output_type -> schemas.SeedFileResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gosher_proto_init() }
func file_proto_gosher_proto_init() {
	if File_proto_gosher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gosher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSendRequest); i {
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
		file_proto_gosher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_gosher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_gosher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_proto_gosher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingForFileRequest); i {
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
		file_proto_gosher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingForFileResponse); i {
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
		file_proto_gosher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenForFilePingsRequest); i {
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
		file_proto_gosher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenForFilePingsResponse); i {
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
		file_proto_gosher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFileRequest); i {
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
		file_proto_gosher_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFileResponse); i {
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
		file_proto_gosher_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedFileRequest); i {
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
		file_proto_gosher_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedFileResponse); i {
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
			RawDescriptor: file_proto_gosher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gosher_proto_goTypes,
		DependencyIndexes: file_proto_gosher_proto_depIdxs,
		MessageInfos:      file_proto_gosher_proto_msgTypes,
	}.Build()
	File_proto_gosher_proto = out.File
	file_proto_gosher_proto_rawDesc = nil
	file_proto_gosher_proto_goTypes = nil
	file_proto_gosher_proto_depIdxs = nil
}