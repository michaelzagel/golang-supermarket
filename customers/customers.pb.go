// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: customers/customers.proto

package customers

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	DeliveryAddress string `protobuf:"bytes,2,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterInput) Reset() {
	*x = RegisterInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInput) ProtoMessage() {}

func (x *RegisterInput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInput.ProtoReflect.Descriptor instead.
func (*RegisterInput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterInput) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *RegisterInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *RegisterOutput) Reset() {
	*x = RegisterOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOutput) ProtoMessage() {}

func (x *RegisterOutput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOutput.ProtoReflect.Descriptor instead.
func (*RegisterOutput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterOutput) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetInfoInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId *int64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3,oneof" json:"customer_id,omitempty"`
	Email      *string `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
}

func (x *GetInfoInput) Reset() {
	*x = GetInfoInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoInput) ProtoMessage() {}

func (x *GetInfoInput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoInput.ProtoReflect.Descriptor instead.
func (*GetInfoInput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoInput) GetCustomerId() int64 {
	if x != nil && x.CustomerId != nil {
		return *x.CustomerId
	}
	return 0
}

func (x *GetInfoInput) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type GetInfoOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	DeliveryAddress string `protobuf:"bytes,2,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	CustomerId      int64  `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetInfoOutput) Reset() {
	*x = GetInfoOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoOutput) ProtoMessage() {}

func (x *GetInfoOutput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoOutput.ProtoReflect.Descriptor instead.
func (*GetInfoOutput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoOutput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetInfoOutput) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *GetInfoOutput) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type VerifyPasswordInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *VerifyPasswordInput) Reset() {
	*x = VerifyPasswordInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordInput) ProtoMessage() {}

func (x *VerifyPasswordInput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordInput.ProtoReflect.Descriptor instead.
func (*VerifyPasswordInput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyPasswordInput) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *VerifyPasswordInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type VerifyPasswordOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VerifyPasswordOutput) Reset() {
	*x = VerifyPasswordOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customers_customers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordOutput) ProtoMessage() {}

func (x *VerifyPasswordOutput) ProtoReflect() protoreflect.Message {
	mi := &file_customers_customers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordOutput.ProtoReflect.Descriptor instead.
func (*VerifyPasswordOutput) Descriptor() ([]byte, []int) {
	return file_customers_customers_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyPasswordOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_customers_customers_proto protoreflect.FileDescriptor

var file_customers_customers_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x31, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x71, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe3, 0x01, 0x0a, 0x09,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x19, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customers_customers_proto_rawDescOnce sync.Once
	file_customers_customers_proto_rawDescData = file_customers_customers_proto_rawDesc
)

func file_customers_customers_proto_rawDescGZIP() []byte {
	file_customers_customers_proto_rawDescOnce.Do(func() {
		file_customers_customers_proto_rawDescData = protoimpl.X.CompressGZIP(file_customers_customers_proto_rawDescData)
	})
	return file_customers_customers_proto_rawDescData
}

var file_customers_customers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_customers_customers_proto_goTypes = []interface{}{
	(*RegisterInput)(nil),        // 0: customers.RegisterInput
	(*RegisterOutput)(nil),       // 1: customers.RegisterOutput
	(*GetInfoInput)(nil),         // 2: customers.GetInfoInput
	(*GetInfoOutput)(nil),        // 3: customers.GetInfoOutput
	(*VerifyPasswordInput)(nil),  // 4: customers.VerifyPasswordInput
	(*VerifyPasswordOutput)(nil), // 5: customers.VerifyPasswordOutput
}
var file_customers_customers_proto_depIdxs = []int32{
	0, // 0: customers.Customers.Register:input_type -> customers.RegisterInput
	2, // 1: customers.Customers.GetInfo:input_type -> customers.GetInfoInput
	4, // 2: customers.Customers.VerifyPassword:input_type -> customers.VerifyPasswordInput
	1, // 3: customers.Customers.Register:output_type -> customers.RegisterOutput
	3, // 4: customers.Customers.GetInfo:output_type -> customers.GetInfoOutput
	5, // 5: customers.Customers.VerifyPassword:output_type -> customers.VerifyPasswordOutput
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customers_customers_proto_init() }
func file_customers_customers_proto_init() {
	if File_customers_customers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customers_customers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInput); i {
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
		file_customers_customers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOutput); i {
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
		file_customers_customers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoInput); i {
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
		file_customers_customers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoOutput); i {
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
		file_customers_customers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordInput); i {
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
		file_customers_customers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordOutput); i {
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
	file_customers_customers_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customers_customers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customers_customers_proto_goTypes,
		DependencyIndexes: file_customers_customers_proto_depIdxs,
		MessageInfos:      file_customers_customers_proto_msgTypes,
	}.Build()
	File_customers_customers_proto = out.File
	file_customers_customers_proto_rawDesc = nil
	file_customers_customers_proto_goTypes = nil
	file_customers_customers_proto_depIdxs = nil
}
