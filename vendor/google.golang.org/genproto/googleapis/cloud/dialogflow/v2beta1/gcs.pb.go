// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/gcs.proto

package dialogflow

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

// Google Cloud Storage location for single input.
type GcsSource struct {
	// Required. The Google Cloud Storage URIs for the inputs. A URI is of the
	// form:
	//   gs://bucket/object-prefix-or-name
	// Whether a prefix or name is used depends on the use case.
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsSource) Reset()         { *m = GcsSource{} }
func (m *GcsSource) String() string { return proto.CompactTextString(m) }
func (*GcsSource) ProtoMessage()    {}
func (*GcsSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e29ad1480788a5, []int{0}
}

func (m *GcsSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsSource.Unmarshal(m, b)
}
func (m *GcsSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsSource.Marshal(b, m, deterministic)
}
func (m *GcsSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsSource.Merge(m, src)
}
func (m *GcsSource) XXX_Size() int {
	return xxx_messageInfo_GcsSource.Size(m)
}
func (m *GcsSource) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsSource.DiscardUnknown(m)
}

var xxx_messageInfo_GcsSource proto.InternalMessageInfo

func (m *GcsSource) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*GcsSource)(nil), "google.cloud.dialogflow.v2beta1.GcsSource")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/gcs.proto", fileDescriptor_d9e29ad1480788a5)
}

var fileDescriptor_d9e29ad1480788a5 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xc9, 0x4c, 0xcc, 0xc9, 0x4f, 0x4f,
	0xcb, 0xc9, 0x2f, 0xd7, 0x2f, 0x33, 0x4a, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0x4f, 0x2e, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x87, 0x28, 0xd5, 0x03, 0x2b, 0xd5, 0x43, 0x28, 0xd5,
	0x83, 0x2a, 0x95, 0x92, 0x81, 0x9a, 0x95, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92,
	0x58, 0x92, 0x99, 0x9f, 0x07, 0xd5, 0xae, 0x24, 0xcb, 0xc5, 0xe9, 0x9e, 0x5c, 0x1c, 0x9c, 0x5f,
	0x5a, 0x94, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0x62, 0x3a, 0x2d, 0x63, 0xe4, 0x52, 0x4e, 0xce, 0xcf, 0xd5, 0x23, 0x60, 0x89, 0x13,
	0x87, 0x7b, 0x72, 0x71, 0x00, 0xc8, 0xc0, 0x00, 0xc6, 0x28, 0x4f, 0xa8, 0xe2, 0xf4, 0xfc, 0x9c,
	0xc4, 0xbc, 0x74, 0xbd, 0xfc, 0xa2, 0x74, 0xfd, 0xf4, 0xd4, 0x3c, 0xb0, 0x75, 0xfa, 0x10, 0xa9,
	0xc4, 0x82, 0xcc, 0x62, 0x9c, 0x7e, 0xb3, 0x46, 0x08, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xe4,
	0xe2, 0xb6, 0x8a, 0x49, 0xde, 0x1d, 0x62, 0xa6, 0x33, 0xd8, 0x01, 0x2e, 0x08, 0x07, 0x84, 0x41,
	0x34, 0x25, 0xb1, 0x81, 0xcd, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0xdd, 0xc3, 0xb4,
	0x3a, 0x01, 0x00, 0x00,
}
