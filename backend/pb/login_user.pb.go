// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: login_user.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginForm) Reset() {
	*x = LoginForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginForm) ProtoMessage() {}

func (x *LoginForm) ProtoReflect() protoreflect.Message {
	mi := &file_login_user_proto_msgTypes[0]
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
	return file_login_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginForm) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginForm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_login_user_proto protoreflect.FileDescriptor

var file_login_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x58, 0x0a, 0x10,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x64, 0x74, 0x34, 0x35, 0x36,
	0x2f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_user_proto_rawDescOnce sync.Once
	file_login_user_proto_rawDescData = file_login_user_proto_rawDesc
)

func file_login_user_proto_rawDescGZIP() []byte {
	file_login_user_proto_rawDescOnce.Do(func() {
		file_login_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_user_proto_rawDescData)
	})
	return file_login_user_proto_rawDescData
}

var file_login_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_login_user_proto_goTypes = []interface{}{
	(*LoginForm)(nil), // 0: proto.LoginForm
	(*Response)(nil),  // 1: proto.Response
}
var file_login_user_proto_depIdxs = []int32{
	0, // 0: proto.LoginUserService.LoginUser:input_type -> proto.LoginForm
	1, // 1: proto.LoginUserService.LoginUser:output_type -> proto.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_user_proto_init() }
func file_login_user_proto_init() {
	if File_login_user_proto != nil {
		return
	}
	file_database_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_login_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_user_proto_goTypes,
		DependencyIndexes: file_login_user_proto_depIdxs,
		MessageInfos:      file_login_user_proto_msgTypes,
	}.Build()
	File_login_user_proto = out.File
	file_login_user_proto_rawDesc = nil
	file_login_user_proto_goTypes = nil
	file_login_user_proto_depIdxs = nil
}
