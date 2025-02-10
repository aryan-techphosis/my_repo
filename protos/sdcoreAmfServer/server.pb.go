// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: server.proto

package sdcoreAmfServer

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

type MsgType int32

const (
	MsgType_UNKNOWN      MsgType = 0
	MsgType_INIT_MSG     MsgType = 1
	MsgType_GNB_MSG      MsgType = 2
	MsgType_AMF_MSG      MsgType = 3
	MsgType_REDIRECT_MSG MsgType = 4
	MsgType_GNB_DISC     MsgType = 5
	MsgType_GNB_CONN     MsgType = 6
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "UNKNOWN",
		1: "INIT_MSG",
		2: "GNB_MSG",
		3: "AMF_MSG",
		4: "REDIRECT_MSG",
		5: "GNB_DISC",
		6: "GNB_CONN",
	}
	MsgType_value = map[string]int32{
		"UNKNOWN":      0,
		"INIT_MSG":     1,
		"GNB_MSG":      2,
		"AMF_MSG":      3,
		"REDIRECT_MSG": 4,
		"GNB_DISC":     5,
		"GNB_CONN":     6,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type SctplbMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SctplbId   string  `protobuf:"bytes,1,opt,name=SctplbId,proto3" json:"SctplbId,omitempty"`
	Msgtype    MsgType `protobuf:"varint,2,opt,name=Msgtype,proto3,enum=sdcoreAmfServer.MsgType" json:"Msgtype,omitempty"`
	GnbIpAddr  string  `protobuf:"bytes,3,opt,name=GnbIpAddr,proto3" json:"GnbIpAddr,omitempty"`
	VerboseMsg string  `protobuf:"bytes,4,opt,name=VerboseMsg,proto3" json:"VerboseMsg,omitempty"`
	Msg        []byte  `protobuf:"bytes,5,opt,name=Msg,proto3" json:"Msg,omitempty"`
	GnbId      string  `protobuf:"bytes,6,opt,name=GnbId,proto3" json:"GnbId,omitempty"`
}

func (x *SctplbMessage) Reset() {
	*x = SctplbMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SctplbMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SctplbMessage) ProtoMessage() {}

func (x *SctplbMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SctplbMessage.ProtoReflect.Descriptor instead.
func (*SctplbMessage) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *SctplbMessage) GetSctplbId() string {
	if x != nil {
		return x.SctplbId
	}
	return ""
}

func (x *SctplbMessage) GetMsgtype() MsgType {
	if x != nil {
		return x.Msgtype
	}
	return MsgType_UNKNOWN
}

func (x *SctplbMessage) GetGnbIpAddr() string {
	if x != nil {
		return x.GnbIpAddr
	}
	return ""
}

func (x *SctplbMessage) GetVerboseMsg() string {
	if x != nil {
		return x.VerboseMsg
	}
	return ""
}

func (x *SctplbMessage) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SctplbMessage) GetGnbId() string {
	if x != nil {
		return x.GnbId
	}
	return ""
}

type AmfMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmfId      string  `protobuf:"bytes,1,opt,name=AmfId,proto3" json:"AmfId,omitempty"`
	RedirectId string  `protobuf:"bytes,2,opt,name=RedirectId,proto3" json:"RedirectId,omitempty"`
	Msgtype    MsgType `protobuf:"varint,3,opt,name=Msgtype,proto3,enum=sdcoreAmfServer.MsgType" json:"Msgtype,omitempty"`
	GnbIpAddr  string  `protobuf:"bytes,4,opt,name=GnbIpAddr,proto3" json:"GnbIpAddr,omitempty"`
	GnbId      string  `protobuf:"bytes,5,opt,name=GnbId,proto3" json:"GnbId,omitempty"`
	VerboseMsg string  `protobuf:"bytes,6,opt,name=VerboseMsg,proto3" json:"VerboseMsg,omitempty"`
	Msg        []byte  `protobuf:"bytes,7,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *AmfMessage) Reset() {
	*x = AmfMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmfMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmfMessage) ProtoMessage() {}

func (x *AmfMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmfMessage.ProtoReflect.Descriptor instead.
func (*AmfMessage) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *AmfMessage) GetAmfId() string {
	if x != nil {
		return x.AmfId
	}
	return ""
}

func (x *AmfMessage) GetRedirectId() string {
	if x != nil {
		return x.RedirectId
	}
	return ""
}

func (x *AmfMessage) GetMsgtype() MsgType {
	if x != nil {
		return x.Msgtype
	}
	return MsgType_UNKNOWN
}

func (x *AmfMessage) GetGnbIpAddr() string {
	if x != nil {
		return x.GnbIpAddr
	}
	return ""
}

func (x *AmfMessage) GetGnbId() string {
	if x != nil {
		return x.GnbId
	}
	return ""
}

func (x *AmfMessage) GetVerboseMsg() string {
	if x != nil {
		return x.VerboseMsg
	}
	return ""
}

func (x *AmfMessage) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0xc5, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x74, 0x70, 0x6c, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63, 0x74, 0x70, 0x6c, 0x62, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x63, 0x74, 0x70, 0x6c, 0x62, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x07, 0x4d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6e, 0x62, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6e, 0x62, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x6e, 0x62, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x47, 0x6e, 0x62, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x0a, 0x41, 0x6d, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x6d, 0x66, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x6d, 0x66, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6e, 0x62, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6e, 0x62, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x47, 0x6e, 0x62, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47,
	0x6e, 0x62, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x4d,
	0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x56, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x65, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x2a, 0x6c, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x47, 0x4e, 0x42, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4d, 0x46,
	0x5f, 0x4d, 0x53, 0x47, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x4e, 0x42, 0x5f,
	0x44, 0x49, 0x53, 0x43, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x4e, 0x42, 0x5f, 0x43, 0x4f,
	0x4e, 0x4e, 0x10, 0x06, 0x32, 0x61, 0x0a, 0x0b, 0x4e, 0x67, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x74, 0x70, 0x6c, 0x62, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6d, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x64, 0x63,
	0x6f, 0x72, 0x65, 0x41, 0x6d, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_proto_goTypes = []interface{}{
	(MsgType)(0),          // 0: sdcoreAmfServer.msgType
	(*SctplbMessage)(nil), // 1: sdcoreAmfServer.SctplbMessage
	(*AmfMessage)(nil),    // 2: sdcoreAmfServer.AmfMessage
}
var file_server_proto_depIdxs = []int32{
	0, // 0: sdcoreAmfServer.SctplbMessage.Msgtype:type_name -> sdcoreAmfServer.msgType
	0, // 1: sdcoreAmfServer.AmfMessage.Msgtype:type_name -> sdcoreAmfServer.msgType
	1, // 2: sdcoreAmfServer.NgapService.HandleMessage:input_type -> sdcoreAmfServer.SctplbMessage
	2, // 3: sdcoreAmfServer.NgapService.HandleMessage:output_type -> sdcoreAmfServer.AmfMessage
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SctplbMessage); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmfMessage); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
