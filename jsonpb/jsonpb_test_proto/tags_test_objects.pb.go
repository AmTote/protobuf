// ####### THIS CODE WAS MODIFIED BY ADDING THE "jsonpb" tags to the JsonpbTags struct
//  To recreate this file rename 'tags-test-objects.proto.skip' and remove the '.skip' extention, then run 'make regenerate'
// Then add the:
//  jsonpb:"emitdefault"
//  tag to the tags on the EmitDefault field
// and add the:
//  jsonpb:"noemitdefault"
//  tag to the tags on the NoEmitDefault field
//
// source: tags_test_objects.proto

package jsonpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type JsonpbTags struct {
	EmitDefault          int32    `protobuf:"varint,1,opt,name=emit_default,json=emitDefault,proto3" json:"emit_default,omitempty" jsonpb:"emitdefault"`
	NoEmitDefault        int32    `protobuf:"varint,2,opt,name=no_emit_default,json=noEmitDefault,proto3" json:"no_emit_default,omitempty" jsonpb:"noemitdefault"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JsonpbTags) Reset()         { *m = JsonpbTags{} }
func (m *JsonpbTags) String() string { return proto.CompactTextString(m) }
func (*JsonpbTags) ProtoMessage()    {}
func (*JsonpbTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b1996bb5747bf9e, []int{0}
}

func (m *JsonpbTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JsonpbTags.Unmarshal(m, b)
}
func (m *JsonpbTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JsonpbTags.Marshal(b, m, deterministic)
}
func (m *JsonpbTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JsonpbTags.Merge(m, src)
}
func (m *JsonpbTags) XXX_Size() int {
	return xxx_messageInfo_JsonpbTags.Size(m)
}
func (m *JsonpbTags) XXX_DiscardUnknown() {
	xxx_messageInfo_JsonpbTags.DiscardUnknown(m)
}

var xxx_messageInfo_JsonpbTags proto.InternalMessageInfo

func (m *JsonpbTags) GetEmitDefault() int32 {
	if m != nil {
		return m.EmitDefault
	}
	return 0
}

func (m *JsonpbTags) GetNoEmitDefault() int32 {
	if m != nil {
		return m.NoEmitDefault
	}
	return 0
}

func init() {
	proto.RegisterType((*JsonpbTags)(nil), "jsonpb.JsonpbTags")
}

func init() {
	proto.RegisterFile("tags_test_objects.proto", fileDescriptor_2b1996bb5747bf9e)
}

var fileDescriptor_2b1996bb5747bf9e = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x49, 0x4c, 0x2f,
	0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x2a, 0xce, 0xcf, 0x2b, 0x48, 0x52, 0x0a, 0xe7, 0xe2, 0xf2,
	0x02, 0xb3, 0x42, 0x12, 0xd3, 0x8b, 0x85, 0x14, 0xb9, 0x78, 0x52, 0x73, 0x33, 0x4b, 0xe2, 0x53,
	0x52, 0xd3, 0x12, 0x4b, 0x73, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xb8, 0x41, 0x62,
	0x2e, 0x10, 0x21, 0x21, 0x35, 0x2e, 0xfe, 0xbc, 0xfc, 0x78, 0x14, 0x55, 0x4c, 0x60, 0x55, 0xbc,
	0x79, 0xf9, 0xae, 0x08, 0x75, 0x49, 0x6c, 0x60, 0x7b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x54, 0x49, 0x6e, 0x1f, 0x82, 0x00, 0x00, 0x00,
}
