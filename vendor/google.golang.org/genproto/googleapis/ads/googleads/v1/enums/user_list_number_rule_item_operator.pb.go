// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_number_rule_item_operator.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible user list number rule item operators.
type UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator int32

const (
	// Not specified.
	UserListNumberRuleItemOperatorEnum_UNSPECIFIED UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListNumberRuleItemOperatorEnum_UNKNOWN UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 1
	// Greater than.
	UserListNumberRuleItemOperatorEnum_GREATER_THAN UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 2
	// Greater than or equal.
	UserListNumberRuleItemOperatorEnum_GREATER_THAN_OR_EQUAL UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 3
	// Equals.
	UserListNumberRuleItemOperatorEnum_EQUALS UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 4
	// Not equals.
	UserListNumberRuleItemOperatorEnum_NOT_EQUALS UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 5
	// Less than.
	UserListNumberRuleItemOperatorEnum_LESS_THAN UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 6
	// Less than or equal.
	UserListNumberRuleItemOperatorEnum_LESS_THAN_OR_EQUAL UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator = 7
)

var UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GREATER_THAN",
	3: "GREATER_THAN_OR_EQUAL",
	4: "EQUALS",
	5: "NOT_EQUALS",
	6: "LESS_THAN",
	7: "LESS_THAN_OR_EQUAL",
}

var UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator_value = map[string]int32{
	"UNSPECIFIED":           0,
	"UNKNOWN":               1,
	"GREATER_THAN":          2,
	"GREATER_THAN_OR_EQUAL": 3,
	"EQUALS":                4,
	"NOT_EQUALS":            5,
	"LESS_THAN":             6,
	"LESS_THAN_OR_EQUAL":    7,
}

func (x UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator) String() string {
	return proto.EnumName(UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator_name, int32(x))
}

func (UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33b46e553d0fd0da, []int{0, 0}
}

// Supported rule operator for number type.
type UserListNumberRuleItemOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListNumberRuleItemOperatorEnum) Reset()         { *m = UserListNumberRuleItemOperatorEnum{} }
func (m *UserListNumberRuleItemOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListNumberRuleItemOperatorEnum) ProtoMessage()    {}
func (*UserListNumberRuleItemOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_33b46e553d0fd0da, []int{0}
}

func (m *UserListNumberRuleItemOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListNumberRuleItemOperatorEnum.Unmarshal(m, b)
}
func (m *UserListNumberRuleItemOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListNumberRuleItemOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListNumberRuleItemOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListNumberRuleItemOperatorEnum.Merge(m, src)
}
func (m *UserListNumberRuleItemOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListNumberRuleItemOperatorEnum.Size(m)
}
func (m *UserListNumberRuleItemOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListNumberRuleItemOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListNumberRuleItemOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator", UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator_name, UserListNumberRuleItemOperatorEnum_UserListNumberRuleItemOperator_value)
	proto.RegisterType((*UserListNumberRuleItemOperatorEnum)(nil), "google.ads.googleads.v1.enums.UserListNumberRuleItemOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_number_rule_item_operator.proto", fileDescriptor_33b46e553d0fd0da)
}

var fileDescriptor_33b46e553d0fd0da = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0xae, 0x93, 0x40,
	0x14, 0xc7, 0x85, 0x6a, 0x1b, 0xa7, 0x7e, 0x90, 0x49, 0x34, 0xd1, 0x58, 0x93, 0xd6, 0xfd, 0x10,
	0xe2, 0x6e, 0x5c, 0x4d, 0x15, 0xb1, 0xb1, 0x81, 0x0a, 0xa5, 0x26, 0x86, 0x64, 0x42, 0x65, 0x42,
	0x48, 0x60, 0x86, 0xcc, 0x0c, 0x7d, 0x20, 0x97, 0x2e, 0x7c, 0x10, 0xf7, 0xbe, 0x84, 0x4f, 0xe0,
	0xd2, 0x30, 0xb4, 0xdc, 0xbb, 0xb9, 0xdd, 0x90, 0xff, 0xf9, 0xfa, 0x1d, 0xe6, 0x7f, 0x40, 0x50,
	0x0a, 0x51, 0xd6, 0xcc, 0xcd, 0x0b, 0xe5, 0x0e, 0xb2, 0x57, 0x27, 0xcf, 0x65, 0xbc, 0x6b, 0x94,
	0xdb, 0x29, 0x26, 0x69, 0x5d, 0x29, 0x4d, 0x79, 0xd7, 0x1c, 0x99, 0xa4, 0xb2, 0xab, 0x19, 0xad,
	0x34, 0x6b, 0xa8, 0x68, 0x99, 0xcc, 0xb5, 0x90, 0xa8, 0x95, 0x42, 0x0b, 0xb8, 0x18, 0xa6, 0x51,
	0x5e, 0x28, 0x34, 0x82, 0xd0, 0xc9, 0x43, 0x06, 0xf4, 0xf2, 0xd5, 0x65, 0x4f, 0x5b, 0xb9, 0x39,
	0xe7, 0x42, 0xe7, 0xba, 0x12, 0x5c, 0x0d, 0xc3, 0xab, 0x3f, 0x16, 0x58, 0xa5, 0x8a, 0xc9, 0x6d,
	0xa5, 0x74, 0x68, 0x16, 0xc5, 0x5d, 0xcd, 0x36, 0x9a, 0x35, 0xd1, 0x79, 0x8b, 0xcf, 0xbb, 0x66,
	0xf5, 0xcb, 0x02, 0xaf, 0xaf, 0xb7, 0xc1, 0xa7, 0x60, 0x9e, 0x86, 0xc9, 0xce, 0x7f, 0xbf, 0xf9,
	0xb8, 0xf1, 0x3f, 0x38, 0xf7, 0xe0, 0x1c, 0xcc, 0xd2, 0xf0, 0x73, 0x18, 0x7d, 0x0d, 0x1d, 0x0b,
	0x3a, 0xe0, 0x51, 0x10, 0xfb, 0x64, 0xef, 0xc7, 0x74, 0xff, 0x89, 0x84, 0x8e, 0x0d, 0x5f, 0x80,
	0x67, 0xb7, 0x33, 0x34, 0x8a, 0xa9, 0xff, 0x25, 0x25, 0x5b, 0x67, 0x02, 0x01, 0x98, 0x1a, 0x99,
	0x38, 0xf7, 0xe1, 0x13, 0x00, 0xc2, 0x68, 0x4f, 0xcf, 0xf1, 0x03, 0xf8, 0x18, 0x3c, 0xdc, 0xfa,
	0x49, 0x32, 0x50, 0xa6, 0xf0, 0x39, 0x80, 0x63, 0x78, 0x83, 0x98, 0xad, 0xff, 0x59, 0x60, 0xf9,
	0x5d, 0x34, 0xe8, 0xaa, 0x37, 0xeb, 0x37, 0xd7, 0xdf, 0xb4, 0xeb, 0x2d, 0xda, 0x59, 0xdf, 0xd6,
	0x67, 0x4a, 0x29, 0xea, 0x9c, 0x97, 0x48, 0xc8, 0xd2, 0x2d, 0x19, 0x37, 0x06, 0x5e, 0x4e, 0xd7,
	0x56, 0xea, 0x8e, 0x4b, 0xbe, 0x33, 0xdf, 0x1f, 0xf6, 0x24, 0x20, 0xe4, 0xa7, 0xbd, 0x08, 0x06,
	0x14, 0x29, 0x14, 0x1a, 0x64, 0xaf, 0x0e, 0x1e, 0xea, 0x6d, 0x56, 0xbf, 0x2f, 0xf5, 0x8c, 0x14,
	0x2a, 0x1b, 0xeb, 0xd9, 0xc1, 0xcb, 0x4c, 0xfd, 0xaf, 0xbd, 0x1c, 0x92, 0x18, 0x93, 0x42, 0x61,
	0x3c, 0x76, 0x60, 0x7c, 0xf0, 0x30, 0x36, 0x3d, 0xc7, 0xa9, 0xf9, 0xb1, 0xb7, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x0a, 0x8a, 0x67, 0x61, 0x02, 0x00, 0x00,
}
