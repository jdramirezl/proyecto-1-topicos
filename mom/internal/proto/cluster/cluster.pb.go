// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/cluster.proto

package cluster

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
	return file_proto_cluster_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_proto_cluster_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{0}
}

type SystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    Type   `protobuf:"varint,3,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *SystemRequest) Reset() {
	*x = SystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemRequest) ProtoMessage() {}

func (x *SystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemRequest.ProtoReflect.Descriptor instead.
func (*SystemRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *SystemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_QUEUE
}

func (x *SystemRequest) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type SubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type Type   `protobuf:"varint,2,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Ip   string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *SubscriberRequest) Reset() {
	*x = SubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberRequest) ProtoMessage() {}

func (x *SubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberRequest.ProtoReflect.Descriptor instead.
func (*SubscriberRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriberRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubscriberRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_QUEUE
}

func (x *SubscriberRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type PeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *PeerRequest) Reset() {
	*x = PeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRequest) ProtoMessage() {}

func (x *PeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRequest.ProtoReflect.Descriptor instead.
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *PeerRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type MasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *MasterRequest) Reset() {
	*x = MasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterRequest) ProtoMessage() {}

func (x *MasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterRequest.ProtoReflect.Descriptor instead.
func (*MasterRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *MasterRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type PeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip []string `protobuf:"bytes,1,rep,name=ip,proto3" json:"ip,omitempty"`
}

func (x *PeerResponse) Reset() {
	*x = PeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerResponse) ProtoMessage() {}

func (x *PeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerResponse.ProtoReflect.Descriptor instead.
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *PeerResponse) GetIp() []string {
	if x != nil {
		return x.Ip
	}
	return nil
}

type ElectLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime int64 `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
}

func (x *ElectLeaderResponse) Reset() {
	*x = ElectLeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectLeaderResponse) ProtoMessage() {}

func (x *ElectLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectLeaderResponse.ProtoReflect.Descriptor instead.
func (*ElectLeaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *ElectLeaderResponse) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

var File_proto_cluster_proto protoreflect.FileDescriptor

var file_proto_cluster_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x52, 0x0a, 0x11,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x22, 0x1d, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22,
	0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x1e, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x2d, 0x0a, 0x13, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2a, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x50, 0x49, 0x43,
	0x10, 0x01, 0x32, 0xaa, 0x05, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x2e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x0c, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x4e, 0x65,
	0x77, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x18, 0x5a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2f, 0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_cluster_proto_rawDescOnce sync.Once
	file_proto_cluster_proto_rawDescData = file_proto_cluster_proto_rawDesc
)

func file_proto_cluster_proto_rawDescGZIP() []byte {
	file_proto_cluster_proto_rawDescOnce.Do(func() {
		file_proto_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cluster_proto_rawDescData)
	})
	return file_proto_cluster_proto_rawDescData
}

var file_proto_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cluster_proto_goTypes = []interface{}{
	(Type)(0),                   // 0: Type
	(*SystemRequest)(nil),       // 1: SystemRequest
	(*SubscriberRequest)(nil),   // 2: SubscriberRequest
	(*PeerRequest)(nil),         // 3: PeerRequest
	(*ConnectionRequest)(nil),   // 4: ConnectionRequest
	(*MasterRequest)(nil),       // 5: MasterRequest
	(*PeerResponse)(nil),        // 6: PeerResponse
	(*ElectLeaderResponse)(nil), // 7: electLeaderResponse
	(*empty.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_proto_cluster_proto_depIdxs = []int32{
	0,  // 0: SystemRequest.type:type_name -> Type
	0,  // 1: SubscriberRequest.type:type_name -> Type
	1,  // 2: ClusterService.AddMessagingSystem:input_type -> SystemRequest
	1,  // 3: ClusterService.RemoveMessagingSystem:input_type -> SystemRequest
	2,  // 4: ClusterService.AddSubscriber:input_type -> SubscriberRequest
	2,  // 5: ClusterService.RemoveSubscriber:input_type -> SubscriberRequest
	4,  // 6: ClusterService.AddConnection:input_type -> ConnectionRequest
	4,  // 7: ClusterService.RemoveConnection:input_type -> ConnectionRequest
	3,  // 8: ClusterService.AddPeer:input_type -> PeerRequest
	3,  // 9: ClusterService.RemovePeer:input_type -> PeerRequest
	5,  // 10: ClusterService.NewMaster:input_type -> MasterRequest
	8,  // 11: ClusterService.Heartbeat:input_type -> google.protobuf.Empty
	8,  // 12: ClusterService.ElectLeader:input_type -> google.protobuf.Empty
	8,  // 13: ClusterService.AddMessagingSystem:output_type -> google.protobuf.Empty
	8,  // 14: ClusterService.RemoveMessagingSystem:output_type -> google.protobuf.Empty
	8,  // 15: ClusterService.AddSubscriber:output_type -> google.protobuf.Empty
	8,  // 16: ClusterService.RemoveSubscriber:output_type -> google.protobuf.Empty
	8,  // 17: ClusterService.AddConnection:output_type -> google.protobuf.Empty
	8,  // 18: ClusterService.RemoveConnection:output_type -> google.protobuf.Empty
	6,  // 19: ClusterService.AddPeer:output_type -> PeerResponse
	8,  // 20: ClusterService.RemovePeer:output_type -> google.protobuf.Empty
	8,  // 21: ClusterService.NewMaster:output_type -> google.protobuf.Empty
	8,  // 22: ClusterService.Heartbeat:output_type -> google.protobuf.Empty
	7,  // 23: ClusterService.ElectLeader:output_type -> electLeaderResponse
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cluster_proto_init() }
func file_proto_cluster_proto_init() {
	if File_proto_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemRequest); i {
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
		file_proto_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriberRequest); i {
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
		file_proto_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRequest); i {
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
		file_proto_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_proto_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterRequest); i {
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
		file_proto_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerResponse); i {
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
		file_proto_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectLeaderResponse); i {
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
			RawDescriptor: file_proto_cluster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cluster_proto_goTypes,
		DependencyIndexes: file_proto_cluster_proto_depIdxs,
		EnumInfos:         file_proto_cluster_proto_enumTypes,
		MessageInfos:      file_proto_cluster_proto_msgTypes,
	}.Build()
	File_proto_cluster_proto = out.File
	file_proto_cluster_proto_rawDesc = nil
	file_proto_cluster_proto_goTypes = nil
	file_proto_cluster_proto_depIdxs = nil
}