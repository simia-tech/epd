// Code generated by protoc-gen-go. DO NOT EDIT.
// source: document.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UnlockedDocument struct {
	Id        string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PublicKey []byte                      `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Contacts  map[string]*UnlockedContact `protobuf:"bytes,3,rep,name=contacts" json:"contacts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Sections  map[string]*UnlockedSection `protobuf:"bytes,4,rep,name=sections" json:"sections,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UnlockedDocument) Reset()                    { *m = UnlockedDocument{} }
func (m *UnlockedDocument) String() string            { return proto.CompactTextString(m) }
func (*UnlockedDocument) ProtoMessage()               {}
func (*UnlockedDocument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *UnlockedDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnlockedDocument) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *UnlockedDocument) GetContacts() map[string]*UnlockedContact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *UnlockedDocument) GetSections() map[string]*UnlockedSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

func init() {
	proto.RegisterType((*UnlockedDocument)(nil), "pb.UnlockedDocument")
}

func init() { proto.RegisterFile("document.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x4f, 0x2e,
	0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2,
	0x4d, 0xce, 0xcf, 0x2b, 0x49, 0x4c, 0x86, 0x0a, 0x49, 0xf1, 0x16, 0xa7, 0x26, 0x97, 0x64, 0xe6,
	0xe7, 0x41, 0xb8, 0x4a, 0x4f, 0x99, 0xb8, 0x04, 0x42, 0xf3, 0x72, 0xf2, 0x93, 0xb3, 0x53, 0x53,
	0x5c, 0xa0, 0x9a, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x98, 0x32, 0x53, 0x84, 0x64, 0xb9, 0xb8, 0x0a, 0x4a, 0x93, 0x72, 0x32, 0x93, 0xe3, 0xb3, 0x53,
	0x2b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x38, 0x21, 0x22, 0xde, 0xa9, 0x95, 0x42, 0x76,
	0x5c, 0x1c, 0x50, 0x3b, 0x8a, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x94, 0xf4, 0x0a, 0x92,
	0xf4, 0xd0, 0x8d, 0xd5, 0x73, 0x86, 0x2a, 0x72, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xeb, 0x01,
	0xe9, 0x87, 0x3a, 0xaa, 0x58, 0x82, 0x05, 0x8f, 0xfe, 0x60, 0xa8, 0x22, 0xa8, 0x7e, 0x98, 0x1e,
	0xa9, 0x00, 0x2e, 0x5e, 0x14, 0xa3, 0x85, 0x04, 0xb8, 0x98, 0x41, 0x0e, 0x85, 0x78, 0x00, 0xc4,
	0x14, 0xd2, 0xe4, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0x3b, 0x9e, 0xdb, 0x48, 0x18, 0xd9,
	0x7c, 0xa8, 0xde, 0x20, 0x88, 0x0a, 0x2b, 0x26, 0x0b, 0x46, 0x90, 0x89, 0x28, 0x96, 0x91, 0x68,
	0x22, 0x54, 0x2f, 0x92, 0x89, 0x49, 0x6c, 0xe0, 0xe0, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x4a, 0x4d, 0x60, 0xa2, 0x01, 0x00, 0x00,
}
