// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: partner/service/v1/partner_error.proto

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

type PartnerErrorReason int32

const (
	//  Get_Account_Failed = 1 [(errors.code) = 401];
	PartnerErrorReason_ADD_TEAM_FAILED      PartnerErrorReason = 0
	PartnerErrorReason_DELETE_TEAM_FAILED   PartnerErrorReason = 1
	PartnerErrorReason_UPDATE_TEAM_FAILED   PartnerErrorReason = 2
	PartnerErrorReason_GET_TEAM_FAILED      PartnerErrorReason = 3
	PartnerErrorReason_GET_TEAM_LIST_FAILED PartnerErrorReason = 4
	PartnerErrorReason_ADD_USER_TEAM_FAILED PartnerErrorReason = 5
	PartnerErrorReason_JOIN_TEAM_FAILED     PartnerErrorReason = 6
	PartnerErrorReason_Quit_TEAM_FAILED     PartnerErrorReason = 7
)

// Enum value maps for PartnerErrorReason.
var (
	PartnerErrorReason_name = map[int32]string{
		0: "ADD_TEAM_FAILED",
		1: "DELETE_TEAM_FAILED",
		2: "UPDATE_TEAM_FAILED",
		3: "GET_TEAM_FAILED",
		4: "GET_TEAM_LIST_FAILED",
		5: "ADD_USER_TEAM_FAILED",
		6: "JOIN_TEAM_FAILED",
		7: "Quit_TEAM_FAILED",
	}
	PartnerErrorReason_value = map[string]int32{
		"ADD_TEAM_FAILED":      0,
		"DELETE_TEAM_FAILED":   1,
		"UPDATE_TEAM_FAILED":   2,
		"GET_TEAM_FAILED":      3,
		"GET_TEAM_LIST_FAILED": 4,
		"ADD_USER_TEAM_FAILED": 5,
		"JOIN_TEAM_FAILED":     6,
		"Quit_TEAM_FAILED":     7,
	}
)

func (x PartnerErrorReason) Enum() *PartnerErrorReason {
	p := new(PartnerErrorReason)
	*p = x
	return p
}

func (x PartnerErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartnerErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_partner_service_v1_partner_error_proto_enumTypes[0].Descriptor()
}

func (PartnerErrorReason) Type() protoreflect.EnumType {
	return &file_partner_service_v1_partner_error_proto_enumTypes[0]
}

func (x PartnerErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartnerErrorReason.Descriptor instead.
func (PartnerErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_partner_service_v1_partner_error_proto_rawDescGZIP(), []int{0}
}

var File_partner_service_v1_partner_error_proto protoreflect.FileDescriptor

var file_partner_service_v1_partner_error_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xd4, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x45,
	0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45,
	0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x47,
	0x45, 0x54, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x44,
	0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x75,
	0x69, 0x74, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x62, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_partner_service_v1_partner_error_proto_rawDescOnce sync.Once
	file_partner_service_v1_partner_error_proto_rawDescData = file_partner_service_v1_partner_error_proto_rawDesc
)

func file_partner_service_v1_partner_error_proto_rawDescGZIP() []byte {
	file_partner_service_v1_partner_error_proto_rawDescOnce.Do(func() {
		file_partner_service_v1_partner_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_partner_service_v1_partner_error_proto_rawDescData)
	})
	return file_partner_service_v1_partner_error_proto_rawDescData
}

var file_partner_service_v1_partner_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_partner_service_v1_partner_error_proto_goTypes = []interface{}{
	(PartnerErrorReason)(0), // 0: partner_error.v1.PartnerErrorReason
}
var file_partner_service_v1_partner_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_partner_service_v1_partner_error_proto_init() }
func file_partner_service_v1_partner_error_proto_init() {
	if File_partner_service_v1_partner_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_partner_service_v1_partner_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_partner_service_v1_partner_error_proto_goTypes,
		DependencyIndexes: file_partner_service_v1_partner_error_proto_depIdxs,
		EnumInfos:         file_partner_service_v1_partner_error_proto_enumTypes,
	}.Build()
	File_partner_service_v1_partner_error_proto = out.File
	file_partner_service_v1_partner_error_proto_rawDesc = nil
	file_partner_service_v1_partner_error_proto_goTypes = nil
	file_partner_service_v1_partner_error_proto_depIdxs = nil
}
