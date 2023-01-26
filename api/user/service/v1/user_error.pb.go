// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: api/user/service/v1/user_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 用户名错误
	ErrorReason_ACCOUNT_ERROR ErrorReason = 0
	// 密码错误
	ErrorReason_PASSWORD_ERROR ErrorReason = 1
	// 无此sessionID
	ErrorReason_SESSION_ID_NOT_FOUND ErrorReason = 2
	// 账号已注册
	ErrorReason_ACCOUNT_REGISTERED ErrorReason = 3
	// 非法账号名
	ErrorReason_ACCOUNT_INVALID ErrorReason = 4
	// 非法密码
	ErrorReason_PASSWORD_INVALID ErrorReason = 5
	// 微信已注册
	ErrorReason_WECHAT_REGISTERED ErrorReason = 6
	// 微信code错误
	ErrorReason_WECHAT_CODE_ERROR ErrorReason = 7
	// 无此用户
	ErrorReason_USER_NOT_FOUND ErrorReason = 8
	// 微信已注册
	ErrorReason_PHONE_NUMBER_REGISTERED ErrorReason = 9
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "ACCOUNT_ERROR",
		1: "PASSWORD_ERROR",
		2: "SESSION_ID_NOT_FOUND",
		3: "ACCOUNT_REGISTERED",
		4: "ACCOUNT_INVALID",
		5: "PASSWORD_INVALID",
		6: "WECHAT_REGISTERED",
		7: "WECHAT_CODE_ERROR",
		8: "USER_NOT_FOUND",
		9: "PHONE_NUMBER_REGISTERED",
	}
	ErrorReason_value = map[string]int32{
		"ACCOUNT_ERROR":           0,
		"PASSWORD_ERROR":          1,
		"SESSION_ID_NOT_FOUND":    2,
		"ACCOUNT_REGISTERED":      3,
		"ACCOUNT_INVALID":         4,
		"PASSWORD_INVALID":        5,
		"WECHAT_REGISTERED":       6,
		"WECHAT_CODE_ERROR":       7,
		"USER_NOT_FOUND":          8,
		"PHONE_NUMBER_REGISTERED": 9,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_service_v1_user_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_user_service_v1_user_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_user_service_v1_user_error_proto_rawDescGZIP(), []int{0}
}

var File_api_user_service_v1_user_error_proto protoreflect.FileDescriptor

var file_api_user_service_v1_user_error_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb2, 0x02, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x91,
	0x03, 0x12, 0x18, 0x0a, 0x0e, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x1b, 0x0a, 0x11, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1b, 0x0a,
	0x11, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x08, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x21, 0x0a, 0x17, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55,
	0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x09, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x21, 0x5a,
	0x1f, 0x73, 0x64, 0x2d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_service_v1_user_error_proto_rawDescOnce sync.Once
	file_api_user_service_v1_user_error_proto_rawDescData = file_api_user_service_v1_user_error_proto_rawDesc
)

func file_api_user_service_v1_user_error_proto_rawDescGZIP() []byte {
	file_api_user_service_v1_user_error_proto_rawDescOnce.Do(func() {
		file_api_user_service_v1_user_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_service_v1_user_error_proto_rawDescData)
	})
	return file_api_user_service_v1_user_error_proto_rawDescData
}

var file_api_user_service_v1_user_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_service_v1_user_error_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.user.v1.ErrorReason
}
var file_api_user_service_v1_user_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_service_v1_user_error_proto_init() }
func file_api_user_service_v1_user_error_proto_init() {
	if File_api_user_service_v1_user_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_service_v1_user_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_service_v1_user_error_proto_goTypes,
		DependencyIndexes: file_api_user_service_v1_user_error_proto_depIdxs,
		EnumInfos:         file_api_user_service_v1_user_error_proto_enumTypes,
	}.Build()
	File_api_user_service_v1_user_error_proto = out.File
	file_api_user_service_v1_user_error_proto_rawDesc = nil
	file_api_user_service_v1_user_error_proto_goTypes = nil
	file_api_user_service_v1_user_error_proto_depIdxs = nil
}
