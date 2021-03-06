// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/v1/core.proto

package corev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStudentRequest) Reset() {
	*x = ListStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentRequest) ProtoMessage() {}

func (x *ListStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentRequest.ProtoReflect.Descriptor instead.
func (*ListStudentRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{0}
}

type ListStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Point int32   `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *ListStudentResponse) Reset() {
	*x = ListStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentResponse) ProtoMessage() {}

func (x *ListStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentResponse.ProtoReflect.Descriptor instead.
func (*ListStudentResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *ListStudentResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ListStudentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListStudentResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListStudentResponse) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type GetStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Point int32  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetStudentRequest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type GetStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id      int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Point   int32  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
	Success string `protobuf:"bytes,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetStudentResponse) Reset() {
	*x = GetStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentResponse) ProtoMessage() {}

func (x *GetStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentResponse.ProtoReflect.Descriptor instead.
func (*GetStudentResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *GetStudentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStudentResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetStudentResponse) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *GetStudentResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Point int32  `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{4}
}

func (x *CreateStudentRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStudentRequest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type CreateStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Point   int32  `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *CreateStudentResponse) Reset() {
	*x = CreateStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentResponse) ProtoMessage() {}

func (x *CreateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentResponse.ProtoReflect.Descriptor instead.
func (*CreateStudentResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{5}
}

func (x *CreateStudentResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

func (x *CreateStudentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStudentResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStudentResponse) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type UpdateStudentResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Id    int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Point int32 `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"` //google.protobuf.FieldMask update_mask = 2;
}

func (x *UpdateStudentResquest) Reset() {
	*x = UpdateStudentResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentResquest) ProtoMessage() {}

func (x *UpdateStudentResquest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentResquest.ProtoReflect.Descriptor instead.
func (*UpdateStudentResquest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStudentResquest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateStudentResquest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentResquest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type UpdateStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updatesuccess string `protobuf:"bytes,1,opt,name=updatesuccess,proto3" json:"updatesuccess,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id            int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Point         int32  `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
	User          *User  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateStudentResponse) Reset() {
	*x = UpdateStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentResponse) ProtoMessage() {}

func (x *UpdateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentResponse.ProtoReflect.Descriptor instead.
func (*UpdateStudentResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStudentResponse) GetUpdatesuccess() string {
	if x != nil {
		return x.Updatesuccess
	}
	return ""
}

func (x *UpdateStudentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStudentResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentResponse) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *UpdateStudentResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteStudentResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Point int32  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *DeleteStudentResquest) Reset() {
	*x = DeleteStudentResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentResquest) ProtoMessage() {}

func (x *DeleteStudentResquest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentResquest.ProtoReflect.Descriptor instead.
func (*DeleteStudentResquest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteStudentResquest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteStudentResquest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteStudentResquest) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

type DeleteStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deletesuccess string `protobuf:"bytes,1,opt,name=Deletesuccess,proto3" json:"Deletesuccess,omitempty"`
}

func (x *DeleteStudentResponse) Reset() {
	*x = DeleteStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentResponse) ProtoMessage() {}

func (x *DeleteStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentResponse.ProtoReflect.Descriptor instead.
func (*DeleteStudentResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStudentResponse) GetDeletesuccess() string {
	if x != nil {
		return x.Deletesuccess
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentName string `protobuf:"bytes,1,opt,name=Student_name,json=StudentName,proto3" json:"Student_name,omitempty"`
	ID          int32  `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	POINT       int32  `protobuf:"varint,3,opt,name=POINT,proto3" json:"POINT,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{10}
}

func (x *User) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetPOINT() int32 {
	if x != nil {
		return x.POINT
	}
	return 0
}

type Messages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Messages) Reset() {
	*x = Messages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Messages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Messages) ProtoMessage() {}

func (x *Messages) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Messages.ProtoReflect.Descriptor instead.
func (*Messages) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{11}
}

func (x *Messages) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{12}
}

func (x *Ids) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_core_v1_core_proto protoreflect.FileDescriptor

var file_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x73, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x22, 0x1e, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x19, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x32, 0x81, 0x04, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x69,
	0x64, 0x12, 0x5b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x67,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x32, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x62, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x2a, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x8d, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x61, 0x6e, 0x6f, 0x69, 0x73, 0x61, 0x62, 0x6f, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43, 0x6f, 0x72,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_core_proto_rawDescOnce sync.Once
	file_core_v1_core_proto_rawDescData = file_core_v1_core_proto_rawDesc
)

func file_core_v1_core_proto_rawDescGZIP() []byte {
	file_core_v1_core_proto_rawDescOnce.Do(func() {
		file_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_core_proto_rawDescData)
	})
	return file_core_v1_core_proto_rawDescData
}

var file_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_core_v1_core_proto_goTypes = []interface{}{
	(*ListStudentRequest)(nil),    // 0: core.v1.ListStudentRequest
	(*ListStudentResponse)(nil),   // 1: core.v1.ListStudentResponse
	(*GetStudentRequest)(nil),     // 2: core.v1.GetStudentRequest
	(*GetStudentResponse)(nil),    // 3: core.v1.GetStudentResponse
	(*CreateStudentRequest)(nil),  // 4: core.v1.CreateStudentRequest
	(*CreateStudentResponse)(nil), // 5: core.v1.CreateStudentResponse
	(*UpdateStudentResquest)(nil), // 6: core.v1.UpdateStudentResquest
	(*UpdateStudentResponse)(nil), // 7: core.v1.UpdateStudentResponse
	(*DeleteStudentResquest)(nil), // 8: core.v1.DeleteStudentResquest
	(*DeleteStudentResponse)(nil), // 9: core.v1.DeleteStudentResponse
	(*User)(nil),                  // 10: core.v1.User
	(*Messages)(nil),              // 11: core.v1.messages
	(*Ids)(nil),                   // 12: core.v1.ids
}
var file_core_v1_core_proto_depIdxs = []int32{
	10, // 0: core.v1.ListStudentResponse.user:type_name -> core.v1.User
	10, // 1: core.v1.CreateStudentRequest.user:type_name -> core.v1.User
	10, // 2: core.v1.UpdateStudentResquest.user:type_name -> core.v1.User
	10, // 3: core.v1.UpdateStudentResponse.user:type_name -> core.v1.User
	2,  // 4: core.v1.Studentservice.GetStudent:input_type -> core.v1.GetStudentRequest
	0,  // 5: core.v1.Studentservice.ListStudent:input_type -> core.v1.ListStudentRequest
	4,  // 6: core.v1.Studentservice.CreateStudent:input_type -> core.v1.CreateStudentRequest
	6,  // 7: core.v1.Studentservice.UpdateStudent:input_type -> core.v1.UpdateStudentResquest
	8,  // 8: core.v1.Studentservice.DeleteStudent:input_type -> core.v1.DeleteStudentResquest
	3,  // 9: core.v1.Studentservice.GetStudent:output_type -> core.v1.GetStudentResponse
	1,  // 10: core.v1.Studentservice.ListStudent:output_type -> core.v1.ListStudentResponse
	5,  // 11: core.v1.Studentservice.CreateStudent:output_type -> core.v1.CreateStudentResponse
	7,  // 12: core.v1.Studentservice.UpdateStudent:output_type -> core.v1.UpdateStudentResponse
	9,  // 13: core.v1.Studentservice.DeleteStudent:output_type -> core.v1.DeleteStudentResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_core_v1_core_proto_init() }
func file_core_v1_core_proto_init() {
	if File_core_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentRequest); i {
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
		file_core_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentResponse); i {
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
		file_core_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentRequest); i {
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
		file_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentResponse); i {
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
		file_core_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentRequest); i {
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
		file_core_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentResponse); i {
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
		file_core_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentResquest); i {
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
		file_core_v1_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentResponse); i {
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
		file_core_v1_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentResquest); i {
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
		file_core_v1_core_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentResponse); i {
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
		file_core_v1_core_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_core_v1_core_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Messages); i {
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
		file_core_v1_core_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
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
			RawDescriptor: file_core_v1_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_v1_core_proto_goTypes,
		DependencyIndexes: file_core_v1_core_proto_depIdxs,
		MessageInfos:      file_core_v1_core_proto_msgTypes,
	}.Build()
	File_core_v1_core_proto = out.File
	file_core_v1_core_proto_rawDesc = nil
	file_core_v1_core_proto_goTypes = nil
	file_core_v1_core_proto_depIdxs = nil
}
