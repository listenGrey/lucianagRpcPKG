// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0--rc1
// source: user/user.proto

package user

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

// 用于注册时校验用户是否存在
type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Exist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *Exist) Reset() {
	*x = Exist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exist) ProtoMessage() {}

func (x *Exist) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exist.ProtoReflect.Descriptor instead.
func (*Exist) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Exist) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

// 发送注册信息
type RegisterFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterFrom) Reset() {
	*x = RegisterFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterFrom) ProtoMessage() {}

func (x *RegisterFrom) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterFrom.ProtoReflect.Descriptor instead.
func (*RegisterFrom) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterFrom) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterFrom) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterFrom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterFrom) GetPassword() string {
	if x != nil {
		return x.Password
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
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
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
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 发送登录信息
type LoginForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginForm) Reset() {
	*x = LoginForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginForm) ProtoMessage() {}

func (x *LoginForm) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginForm.ProtoReflect.Descriptor instead.
func (*LoginForm) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginForm) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginForm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Exist    bool   `protobuf:"varint,3,opt,name=exist,proto3" json:"exist,omitempty"`
	Success  bool   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *LogInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LogInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LogInfo) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *LogInfo) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b,
	0x47, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x1d, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x22,
	0x64, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6f, 0x0a, 0x07, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x4a, 0x0a, 0x0a, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61,
	0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x15, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x22, 0x00, 0x32, 0x59, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x48, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61,
	0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x1a, 0x17, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61,
	0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x00, 0x32, 0x50, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x42, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x61, 0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x17, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x61,
	0x6e, 0x61, 0x67, 0x52, 0x70, 0x63, 0x50, 0x4b, 0x47, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_user_proto_goTypes = []interface{}{
	(*Email)(nil),        // 0: lucianagRpcPKG.Email
	(*Exist)(nil),        // 1: lucianagRpcPKG.Exist
	(*RegisterFrom)(nil), // 2: lucianagRpcPKG.RegisterFrom
	(*Success)(nil),      // 3: lucianagRpcPKG.Success
	(*LoginForm)(nil),    // 4: lucianagRpcPKG.LoginForm
	(*LogInfo)(nil),      // 5: lucianagRpcPKG.LogInfo
}
var file_user_user_proto_depIdxs = []int32{
	0, // 0: lucianagRpcPKG.CheckExist.CheckExist:input_type -> lucianagRpcPKG.Email
	2, // 1: lucianagRpcPKG.RegisterCheck.RegisterCheck:input_type -> lucianagRpcPKG.RegisterFrom
	4, // 2: lucianagRpcPKG.LoginCheck.LoginCheck:input_type -> lucianagRpcPKG.LoginForm
	1, // 3: lucianagRpcPKG.CheckExist.CheckExist:output_type -> lucianagRpcPKG.Exist
	3, // 4: lucianagRpcPKG.RegisterCheck.RegisterCheck:output_type -> lucianagRpcPKG.Success
	5, // 5: lucianagRpcPKG.LoginCheck.LoginCheck:output_type -> lucianagRpcPKG.LogInfo
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exist); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterFrom); i {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginForm); i {
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
		file_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
