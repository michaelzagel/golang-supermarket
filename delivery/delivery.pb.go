// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: delivery/delivery.proto

package delivery

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

type SelectWindowInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryAddress string `protobuf:"bytes,1,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
}

func (x *SelectWindowInput) Reset() {
	*x = SelectWindowInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectWindowInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectWindowInput) ProtoMessage() {}

func (x *SelectWindowInput) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectWindowInput.ProtoReflect.Descriptor instead.
func (*SelectWindowInput) Descriptor() ([]byte, []int) {
	return file_delivery_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *SelectWindowInput) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

type SelectWindowOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WindowFrom string `protobuf:"bytes,1,opt,name=window_from,json=windowFrom,proto3" json:"window_from,omitempty"`
	WindowTo   string `protobuf:"bytes,2,opt,name=window_to,json=windowTo,proto3" json:"window_to,omitempty"`
}

func (x *SelectWindowOutput) Reset() {
	*x = SelectWindowOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectWindowOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectWindowOutput) ProtoMessage() {}

func (x *SelectWindowOutput) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectWindowOutput.ProtoReflect.Descriptor instead.
func (*SelectWindowOutput) Descriptor() ([]byte, []int) {
	return file_delivery_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *SelectWindowOutput) GetWindowFrom() string {
	if x != nil {
		return x.WindowFrom
	}
	return ""
}

func (x *SelectWindowOutput) GetWindowTo() string {
	if x != nil {
		return x.WindowTo
	}
	return ""
}

var File_delivery_delivery_proto protoreflect.FileDescriptor

var file_delivery_delivery_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x22, 0x3e, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x54, 0x6f, 0x32, 0x57, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00,
	0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x3b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delivery_delivery_proto_rawDescOnce sync.Once
	file_delivery_delivery_proto_rawDescData = file_delivery_delivery_proto_rawDesc
)

func file_delivery_delivery_proto_rawDescGZIP() []byte {
	file_delivery_delivery_proto_rawDescOnce.Do(func() {
		file_delivery_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_delivery_proto_rawDescData)
	})
	return file_delivery_delivery_proto_rawDescData
}

var file_delivery_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_delivery_delivery_proto_goTypes = []interface{}{
	(*SelectWindowInput)(nil),  // 0: delivery.SelectWindowInput
	(*SelectWindowOutput)(nil), // 1: delivery.SelectWindowOutput
}
var file_delivery_delivery_proto_depIdxs = []int32{
	0, // 0: delivery.Delivery.SelectWindow:input_type -> delivery.SelectWindowInput
	1, // 1: delivery.Delivery.SelectWindow:output_type -> delivery.SelectWindowOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_delivery_delivery_proto_init() }
func file_delivery_delivery_proto_init() {
	if File_delivery_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectWindowInput); i {
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
		file_delivery_delivery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectWindowOutput); i {
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
			RawDescriptor: file_delivery_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_delivery_proto_goTypes,
		DependencyIndexes: file_delivery_delivery_proto_depIdxs,
		MessageInfos:      file_delivery_delivery_proto_msgTypes,
	}.Build()
	File_delivery_delivery_proto = out.File
	file_delivery_delivery_proto_rawDesc = nil
	file_delivery_delivery_proto_goTypes = nil
	file_delivery_delivery_proto_depIdxs = nil
}