// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: wes.proto

package proto_wes

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MQTTRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic  string               `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Time   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Record []byte               `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *MQTTRecord) Reset() {
	*x = MQTTRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQTTRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQTTRecord) ProtoMessage() {}

func (x *MQTTRecord) ProtoReflect() protoreflect.Message {
	mi := &file_wes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQTTRecord.ProtoReflect.Descriptor instead.
func (*MQTTRecord) Descriptor() ([]byte, []int) {
	return file_wes_proto_rawDescGZIP(), []int{0}
}

func (x *MQTTRecord) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *MQTTRecord) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *MQTTRecord) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

type CLIRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd     string               `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	AmrID   int32                `protobuf:"varint,2,opt,name=amrID,proto3" json:"amrID,omitempty"`
	RouteID int32                `protobuf:"varint,3,opt,name=routeID,proto3" json:"routeID,omitempty"`
	Time    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Record  []byte               `protobuf:"bytes,5,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *CLIRecord) Reset() {
	*x = CLIRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLIRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLIRecord) ProtoMessage() {}

func (x *CLIRecord) ProtoReflect() protoreflect.Message {
	mi := &file_wes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLIRecord.ProtoReflect.Descriptor instead.
func (*CLIRecord) Descriptor() ([]byte, []int) {
	return file_wes_proto_rawDescGZIP(), []int{1}
}

func (x *CLIRecord) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *CLIRecord) GetAmrID() int32 {
	if x != nil {
		return x.AmrID
	}
	return 0
}

func (x *CLIRecord) GetRouteID() int32 {
	if x != nil {
		return x.RouteID
	}
	return 0
}

func (x *CLIRecord) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *CLIRecord) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

type WMSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int32                `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Time    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Record  []byte               `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *WMSRecord) Reset() {
	*x = WMSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WMSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WMSRecord) ProtoMessage() {}

func (x *WMSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_wes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WMSRecord.ProtoReflect.Descriptor instead.
func (*WMSRecord) Descriptor() ([]byte, []int) {
	return file_wes_proto_rawDescGZIP(), []int{2}
}

func (x *WMSRecord) GetOrderID() int32 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *WMSRecord) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *WMSRecord) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

type AMRRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic  string               `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Time   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Record []byte               `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *AMRRecord) Reset() {
	*x = AMRRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AMRRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AMRRecord) ProtoMessage() {}

func (x *AMRRecord) ProtoReflect() protoreflect.Message {
	mi := &file_wes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AMRRecord.ProtoReflect.Descriptor instead.
func (*AMRRecord) Descriptor() ([]byte, []int) {
	return file_wes_proto_rawDescGZIP(), []int{3}
}

func (x *AMRRecord) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *AMRRecord) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *AMRRecord) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

var File_wes_proto protoreflect.FileDescriptor

var file_wes_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x77, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0a, 0x4d, 0x51, 0x54, 0x54, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x09, 0x43, 0x4c, 0x49, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6d, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x61, 0x6d, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x6d, 0x0a, 0x09, 0x57,
	0x4d, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x69, 0x0a, 0x09, 0x41, 0x4d,
	0x52, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6b, 0x75, 0x72, 0x69, 0x6e, 0x30, 0x30, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wes_proto_rawDescOnce sync.Once
	file_wes_proto_rawDescData = file_wes_proto_rawDesc
)

func file_wes_proto_rawDescGZIP() []byte {
	file_wes_proto_rawDescOnce.Do(func() {
		file_wes_proto_rawDescData = protoimpl.X.CompressGZIP(file_wes_proto_rawDescData)
	})
	return file_wes_proto_rawDescData
}

var file_wes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wes_proto_goTypes = []interface{}{
	(*MQTTRecord)(nil),          // 0: proto.wes.MQTTRecord
	(*CLIRecord)(nil),           // 1: proto.wes.CLIRecord
	(*WMSRecord)(nil),           // 2: proto.wes.WMSRecord
	(*AMRRecord)(nil),           // 3: proto.wes.AMRRecord
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_wes_proto_depIdxs = []int32{
	4, // 0: proto.wes.MQTTRecord.time:type_name -> google.protobuf.Timestamp
	4, // 1: proto.wes.CLIRecord.time:type_name -> google.protobuf.Timestamp
	4, // 2: proto.wes.WMSRecord.time:type_name -> google.protobuf.Timestamp
	4, // 3: proto.wes.AMRRecord.time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wes_proto_init() }
func file_wes_proto_init() {
	if File_wes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQTTRecord); i {
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
		file_wes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLIRecord); i {
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
		file_wes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WMSRecord); i {
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
		file_wes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AMRRecord); i {
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
			RawDescriptor: file_wes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wes_proto_goTypes,
		DependencyIndexes: file_wes_proto_depIdxs,
		MessageInfos:      file_wes_proto_msgTypes,
	}.Build()
	File_wes_proto = out.File
	file_wes_proto_rawDesc = nil
	file_wes_proto_goTypes = nil
	file_wes_proto_depIdxs = nil
}
