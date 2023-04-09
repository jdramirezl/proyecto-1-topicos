// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mom/proto/message.proto

package message

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Type int32

const (
	Type_QUEUE Type = 0
	Type_TOPIC Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "QUEUE",
		1: "TOPIC",
	}
	Type_value = map[string]int32{
		"QUEUE": 0,
		"TOPIC": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_mom_proto_message_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_mom_proto_message_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_mom_proto_message_proto_rawDescGZIP(), []int{0}
}

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    Type   `protobuf:"varint,3,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Payload string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_mom_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MessageRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_QUEUE
}

func (x *MessageRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ConsumeMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type Type   `protobuf:"varint,2,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Ip   string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ConsumeMessageRequest) Reset() {
	*x = ConsumeMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeMessageRequest) ProtoMessage() {}

func (x *ConsumeMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mom_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeMessageRequest.ProtoReflect.Descriptor instead.
func (*ConsumeMessageRequest) Descriptor() ([]byte, []int) {
	return file_mom_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumeMessageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConsumeMessageRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_QUEUE
}

func (x *ConsumeMessageRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ConsumeMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ConsumeMessageResponse) Reset() {
	*x = ConsumeMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mom_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeMessageResponse) ProtoMessage() {}

func (x *ConsumeMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mom_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeMessageResponse.ProtoReflect.Descriptor instead.
func (*ConsumeMessageResponse) Descriptor() ([]byte, []int) {
	return file_mom_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumeMessageResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_mom_proto_message_proto protoreflect.FileDescriptor

var file_mom_proto_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x56, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x1c, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x55, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x01, 0x32, 0xce, 0x01, 0x0a, 0x0e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x18, 0x5a, 0x16,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x3b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mom_proto_message_proto_rawDescOnce sync.Once
	file_mom_proto_message_proto_rawDescData = file_mom_proto_message_proto_rawDesc
)

func file_mom_proto_message_proto_rawDescGZIP() []byte {
	file_mom_proto_message_proto_rawDescOnce.Do(func() {
		file_mom_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_mom_proto_message_proto_rawDescData)
	})
	return file_mom_proto_message_proto_rawDescData
}

var file_mom_proto_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mom_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mom_proto_message_proto_goTypes = []interface{}{
	(Type)(0),                      // 0: Type
	(*MessageRequest)(nil),         // 1: MessageRequest
	(*ConsumeMessageRequest)(nil),  // 2: ConsumeMessageRequest
	(*ConsumeMessageResponse)(nil), // 3: ConsumeMessageResponse
	(*empty.Empty)(nil),            // 4: google.protobuf.Empty
}
var file_mom_proto_message_proto_depIdxs = []int32{
	0, // 0: MessageRequest.type:type_name -> Type
	0, // 1: ConsumeMessageRequest.type:type_name -> Type
	1, // 2: MessageService.AddMessage:input_type -> MessageRequest
	1, // 3: MessageService.RemoveMessage:input_type -> MessageRequest
	2, // 4: MessageService.ConsumeMessage:input_type -> ConsumeMessageRequest
	4, // 5: MessageService.AddMessage:output_type -> google.protobuf.Empty
	4, // 6: MessageService.RemoveMessage:output_type -> google.protobuf.Empty
	3, // 7: MessageService.ConsumeMessage:output_type -> ConsumeMessageResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mom_proto_message_proto_init() }
func file_mom_proto_message_proto_init() {
	if File_mom_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mom_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_mom_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeMessageRequest); i {
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
		file_mom_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeMessageResponse); i {
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
			RawDescriptor: file_mom_proto_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mom_proto_message_proto_goTypes,
		DependencyIndexes: file_mom_proto_message_proto_depIdxs,
		EnumInfos:         file_mom_proto_message_proto_enumTypes,
		MessageInfos:      file_mom_proto_message_proto_msgTypes,
	}.Build()
	File_mom_proto_message_proto = out.File
	file_mom_proto_message_proto_rawDesc = nil
	file_mom_proto_message_proto_goTypes = nil
	file_mom_proto_message_proto_depIdxs = nil
}
