// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: notification/notification.proto

package notification

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

type NotifyWindowInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WindowTo   string `protobuf:"bytes,1,opt,name=window_to,json=windowTo,proto3" json:"window_to,omitempty"`
	OrderId    int64  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	WindowFrom string `protobuf:"bytes,3,opt,name=window_from,json=windowFrom,proto3" json:"window_from,omitempty"`
}

func (x *NotifyWindowInput) Reset() {
	*x = NotifyWindowInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyWindowInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyWindowInput) ProtoMessage() {}

func (x *NotifyWindowInput) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyWindowInput.ProtoReflect.Descriptor instead.
func (*NotifyWindowInput) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyWindowInput) GetWindowTo() string {
	if x != nil {
		return x.WindowTo
	}
	return ""
}

func (x *NotifyWindowInput) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *NotifyWindowInput) GetWindowFrom() string {
	if x != nil {
		return x.WindowFrom
	}
	return ""
}

type NotifyWindowOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NotifyWindowOutput) Reset() {
	*x = NotifyWindowOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyWindowOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyWindowOutput) ProtoMessage() {}

func (x *NotifyWindowOutput) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyWindowOutput.ProtoReflect.Descriptor instead.
func (*NotifyWindowOutput) Descriptor() ([]byte, []int) {
	return file_notification_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyWindowOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_notification_notification_proto protoreflect.FileDescriptor

var file_notification_notification_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x6c, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x54,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x22, 0x2e, 0x0a,
	0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x63, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x1f, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x20,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x00, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notification_proto_rawDescOnce sync.Once
	file_notification_notification_proto_rawDescData = file_notification_notification_proto_rawDesc
)

func file_notification_notification_proto_rawDescGZIP() []byte {
	file_notification_notification_proto_rawDescOnce.Do(func() {
		file_notification_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notification_proto_rawDescData)
	})
	return file_notification_notification_proto_rawDescData
}

var file_notification_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notification_notification_proto_goTypes = []interface{}{
	(*NotifyWindowInput)(nil),  // 0: notification.NotifyWindowInput
	(*NotifyWindowOutput)(nil), // 1: notification.NotifyWindowOutput
}
var file_notification_notification_proto_depIdxs = []int32{
	0, // 0: notification.Notification.NotifyWindow:input_type -> notification.NotifyWindowInput
	1, // 1: notification.Notification.NotifyWindow:output_type -> notification.NotifyWindowOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notification_notification_proto_init() }
func file_notification_notification_proto_init() {
	if File_notification_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyWindowInput); i {
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
		file_notification_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyWindowOutput); i {
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
			RawDescriptor: file_notification_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_notification_proto_goTypes,
		DependencyIndexes: file_notification_notification_proto_depIdxs,
		MessageInfos:      file_notification_notification_proto_msgTypes,
	}.Build()
	File_notification_notification_proto = out.File
	file_notification_notification_proto_rawDesc = nil
	file_notification_notification_proto_goTypes = nil
	file_notification_notification_proto_depIdxs = nil
}
