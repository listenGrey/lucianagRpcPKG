// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0--rc1
// source: chat/chat.proto

package chat

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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *QA) Reset() {
	*x = QA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QA) ProtoMessage() {}

func (x *QA) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QA.ProtoReflect.Descriptor instead.
func (*QA) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *QA) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *QA) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Qas  []*QA  `protobuf:"bytes,3,rep,name=qas,proto3" json:"qas,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Chat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chat) GetQas() []*QA {
	if x != nil {
		return x.Qas
	}
	return nil
}

type ChatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatList []*Chat `protobuf:"bytes,1,rep,name=chatList,proto3" json:"chatList,omitempty"`
}

func (x *ChatList) Reset() {
	*x = ChatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatList) ProtoMessage() {}

func (x *ChatList) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatList.ProtoReflect.Descriptor instead.
func (*ChatList) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatList) GetChatList() []*Chat {
	if x != nil {
		return x.ChatList
	}
	return nil
}

type UserChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Cid  int64  `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserChat) Reset() {
	*x = UserChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChat) ProtoMessage() {}

func (x *UserChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChat.ProtoReflect.Descriptor instead.
func (*UserChat) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *UserChat) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserChat) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *UserChat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b,
	0x47, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x02, 0x51, 0x41, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x03, 0x71, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c,
	0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x51, 0x41,
	0x52, 0x03, 0x71, 0x61, 0x73, 0x22, 0x3c, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70,
	0x63, 0x50, 0x4b, 0x47, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x4c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x75, 0x63,
	0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x49, 0x44, 0x1a, 0x18,
	0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x32, 0x40, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x12, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b,
	0x47, 0x2e, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52,
	0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x22, 0x00, 0x32, 0x49, 0x0a, 0x07,
	0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x18, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63,
	0x50, 0x4b, 0x47, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x17, 0x2e, 0x6c,
	0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x32, 0x4b, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70,
	0x63, 0x50, 0x4b, 0x47, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x75, 0x63, 0x69,
	0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x00, 0x32, 0x49, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x12, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b,
	0x47, 0x2e, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52,
	0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData = file_chat_chat_proto_rawDesc
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_chat_proto_rawDescData)
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chat_chat_proto_goTypes = []interface{}{
	(*ID)(nil),       // 0: lucianagRpcPKG.ID
	(*QA)(nil),       // 1: lucianagRpcPKG.QA
	(*Chat)(nil),     // 2: lucianagRpcPKG.Chat
	(*ChatList)(nil), // 3: lucianagRpcPKG.ChatList
	(*UserChat)(nil), // 4: lucianagRpcPKG.UserChat
	(*Success)(nil),  // 5: lucianagRpcPKG.Success
}
var file_chat_chat_proto_depIdxs = []int32{
	1, // 0: lucianagRpcPKG.Chat.qas:type_name -> lucianagRpcPKG.QA
	2, // 1: lucianagRpcPKG.ChatList.chatList:type_name -> lucianagRpcPKG.Chat
	0, // 2: lucianagRpcPKG.GetChatList.GetChatList:input_type -> lucianagRpcPKG.ID
	0, // 3: lucianagRpcPKG.GetChat.GetChat:input_type -> lucianagRpcPKG.ID
	4, // 4: lucianagRpcPKG.NewChat.NewChat:input_type -> lucianagRpcPKG.UserChat
	2, // 5: lucianagRpcPKG.RenameChat.RenameChat:input_type -> lucianagRpcPKG.Chat
	0, // 6: lucianagRpcPKG.DeleteChat.DeleteChat:input_type -> lucianagRpcPKG.ID
	3, // 7: lucianagRpcPKG.GetChatList.GetChatList:output_type -> lucianagRpcPKG.ChatList
	2, // 8: lucianagRpcPKG.GetChat.GetChat:output_type -> lucianagRpcPKG.Chat
	5, // 9: lucianagRpcPKG.NewChat.NewChat:output_type -> lucianagRpcPKG.Success
	5, // 10: lucianagRpcPKG.RenameChat.RenameChat:output_type -> lucianagRpcPKG.Success
	5, // 11: lucianagRpcPKG.DeleteChat.DeleteChat:output_type -> lucianagRpcPKG.Success
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QA); i {
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
		file_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatList); i {
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
		file_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChat); i {
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
		file_chat_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}
