// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	contact.proto
	document.proto
	section.proto

It has these top-level messages:
	UnlockedContact
	UnlockedDocument
	UnlockedSection
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UnlockedContact struct {
	SectionKey map[string][]byte `protobuf:"bytes,1,rep,name=section_key,json=sectionKey" json:"section_key,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *UnlockedContact) Reset()                    { *m = UnlockedContact{} }
func (m *UnlockedContact) String() string            { return proto.CompactTextString(m) }
func (*UnlockedContact) ProtoMessage()               {}
func (*UnlockedContact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UnlockedContact) GetSectionKey() map[string][]byte {
	if m != nil {
		return m.SectionKey
	}
	return nil
}

func init() {
	proto.RegisterType((*UnlockedContact)(nil), "pb.UnlockedContact")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x9a, 0xc6,
	0xc8, 0xc5, 0x1f, 0x9a, 0x97, 0x93, 0x9f, 0x9c, 0x9d, 0x9a, 0xe2, 0x0c, 0x91, 0x15, 0x72, 0xe1,
	0xe2, 0x2e, 0x4e, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x36, 0x52, 0xd6, 0x2b, 0x48, 0xd2, 0x43, 0x53, 0xa9, 0x17, 0x0c, 0x51, 0xe6, 0x9d,
	0x5a, 0xe9, 0x9a, 0x57, 0x52, 0x54, 0x19, 0xc4, 0x55, 0x0c, 0x17, 0x90, 0xb2, 0xe5, 0xe2, 0x47,
	0x93, 0x16, 0x12, 0xe0, 0x62, 0x86, 0x18, 0xc8, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x89, 0x70,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x56,
	0x4c, 0x16, 0x8c, 0x49, 0x6c, 0x60, 0x37, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x3a,
	0x50, 0xfd, 0xb4, 0x00, 0x00, 0x00,
}
