// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payments/payments.proto

package payments

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

type TakePaymentInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCard string  `protobuf:"bytes,1,opt,name=credit_card,json=creditCard,proto3" json:"credit_card,omitempty"`
	Amount     float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TakePaymentInput) Reset() {
	*x = TakePaymentInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_payments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakePaymentInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakePaymentInput) ProtoMessage() {}

func (x *TakePaymentInput) ProtoReflect() protoreflect.Message {
	mi := &file_payments_payments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakePaymentInput.ProtoReflect.Descriptor instead.
func (*TakePaymentInput) Descriptor() ([]byte, []int) {
	return file_payments_payments_proto_rawDescGZIP(), []int{0}
}

func (x *TakePaymentInput) GetCreditCard() string {
	if x != nil {
		return x.CreditCard
	}
	return ""
}

func (x *TakePaymentInput) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TakePaymentOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TakePaymentOutput) Reset() {
	*x = TakePaymentOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payments_payments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakePaymentOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakePaymentOutput) ProtoMessage() {}

func (x *TakePaymentOutput) ProtoReflect() protoreflect.Message {
	mi := &file_payments_payments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakePaymentOutput.ProtoReflect.Descriptor instead.
func (*TakePaymentOutput) Descriptor() ([]byte, []int) {
	return file_payments_payments_proto_rawDescGZIP(), []int{1}
}

func (x *TakePaymentOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_payments_payments_proto protoreflect.FileDescriptor

var file_payments_payments_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x10, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2d, 0x0a, 0x11, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0x54, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0b, 0x54,
	0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payments_payments_proto_rawDescOnce sync.Once
	file_payments_payments_proto_rawDescData = file_payments_payments_proto_rawDesc
)

func file_payments_payments_proto_rawDescGZIP() []byte {
	file_payments_payments_proto_rawDescOnce.Do(func() {
		file_payments_payments_proto_rawDescData = protoimpl.X.CompressGZIP(file_payments_payments_proto_rawDescData)
	})
	return file_payments_payments_proto_rawDescData
}

var file_payments_payments_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payments_payments_proto_goTypes = []interface{}{
	(*TakePaymentInput)(nil),  // 0: payments.TakePaymentInput
	(*TakePaymentOutput)(nil), // 1: payments.TakePaymentOutput
}
var file_payments_payments_proto_depIdxs = []int32{
	0, // 0: payments.Payments.TakePayment:input_type -> payments.TakePaymentInput
	1, // 1: payments.Payments.TakePayment:output_type -> payments.TakePaymentOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payments_payments_proto_init() }
func file_payments_payments_proto_init() {
	if File_payments_payments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payments_payments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakePaymentInput); i {
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
		file_payments_payments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakePaymentOutput); i {
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
			RawDescriptor: file_payments_payments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payments_payments_proto_goTypes,
		DependencyIndexes: file_payments_payments_proto_depIdxs,
		MessageInfos:      file_payments_payments_proto_msgTypes,
	}.Build()
	File_payments_payments_proto = out.File
	file_payments_payments_proto_rawDesc = nil
	file_payments_payments_proto_goTypes = nil
	file_payments_payments_proto_depIdxs = nil
}
