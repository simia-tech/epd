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
	Name      string                      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
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

func (m *UnlockedDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnlockedDocument) GetSections() map[string]*UnlockedSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

type LockedDocument struct {
	Id        string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PublicKey []byte                    `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Name      string                    `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Contacts  map[string]*LockedContact `protobuf:"bytes,4,rep,name=contacts" json:"contacts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Sections  map[string]*LockedSection `protobuf:"bytes,5,rep,name=sections" json:"sections,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Signature []byte                    `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *LockedDocument) Reset()                    { *m = LockedDocument{} }
func (m *LockedDocument) String() string            { return proto.CompactTextString(m) }
func (*LockedDocument) ProtoMessage()               {}
func (*LockedDocument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *LockedDocument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LockedDocument) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *LockedDocument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LockedDocument) GetContacts() map[string]*LockedContact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *LockedDocument) GetSections() map[string]*LockedSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

func (m *LockedDocument) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*UnlockedDocument)(nil), "pb.UnlockedDocument")
	proto.RegisterType((*LockedDocument)(nil), "pb.LockedDocument")
}

func init() { proto.RegisterFile("document.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0x03, 0x6c, 0xcb, 0x78, 0x13, 0x32, 0xeb, 0x85, 0x10, 0x4d, 0x08, 0x17, 0xf1, 0xc2,
	0x61, 0x5e, 0x8c, 0x31, 0x5e, 0xd4, 0x93, 0xc6, 0x18, 0x8c, 0x67, 0x03, 0xa5, 0x31, 0xcd, 0x58,
	0x4b, 0xa0, 0x98, 0xf0, 0x21, 0xfc, 0x9c, 0x7e, 0x0d, 0x53, 0xda, 0xe1, 0x50, 0x89, 0x97, 0xdd,
	0xca, 0xe3, 0xfd, 0xde, 0xff, 0xff, 0xfe, 0x2d, 0xb8, 0x39, 0xc7, 0xcd, 0x86, 0x30, 0x11, 0x97,
	0x15, 0x17, 0x1c, 0x99, 0x65, 0xe6, 0x3b, 0x98, 0x33, 0x91, 0x62, 0x5d, 0xf2, 0x9d, 0x9a, 0x60,
	0x41, 0x39, 0x53, 0x9f, 0xe1, 0xa7, 0x01, 0xcb, 0x17, 0x56, 0x70, 0xbc, 0x26, 0xf9, 0xad, 0x86,
	0x91, 0x0b, 0x26, 0xcd, 0x3d, 0x23, 0x30, 0x22, 0x3b, 0x31, 0x69, 0x8e, 0x4e, 0x00, 0xca, 0x26,
	0x2b, 0x28, 0x7e, 0x5d, 0x93, 0xd6, 0x33, 0x03, 0x23, 0x3a, 0x48, 0x6c, 0x55, 0xb9, 0x27, 0x2d,
	0x42, 0x30, 0x61, 0xe9, 0x86, 0x78, 0x56, 0x07, 0x74, 0x67, 0x74, 0x0d, 0x73, 0x2d, 0x54, 0x7b,
	0x93, 0xc0, 0x8a, 0x16, 0xab, 0x30, 0x2e, 0xb3, 0xf8, 0xa7, 0x54, 0xfc, 0xac, 0x9b, 0xee, 0x98,
	0xa8, 0xda, 0xa4, 0x67, 0xfc, 0x27, 0x70, 0x06, 0xbf, 0xd0, 0x12, 0x2c, 0x29, 0xae, 0x4c, 0xc9,
	0x23, 0x3a, 0x83, 0xe9, 0x7b, 0x5a, 0x34, 0xa4, 0x33, 0xb4, 0x58, 0x1d, 0xed, 0xce, 0xd7, 0x6c,
	0xa2, 0x3a, 0x2e, 0xcd, 0x0b, 0x23, 0xfc, 0xb0, 0xc0, 0x7d, 0xd8, 0xfb, 0x9e, 0x57, 0x30, 0xd7,
	0xf9, 0x6e, 0xf7, 0x0c, 0xa4, 0x8f, 0xa1, 0x50, 0x7c, 0xa3, 0x5b, 0xf4, 0x96, 0x5b, 0x42, 0xd2,
	0x7d, 0x4a, 0xd3, 0x51, 0x7a, 0x24, 0x23, 0x74, 0x0c, 0x76, 0x4d, 0xdf, 0x58, 0x2a, 0x9a, 0x8a,
	0x78, 0x33, 0xe5, 0xb6, 0x2f, 0xf8, 0x8f, 0xe0, 0x0c, 0x64, 0xff, 0x48, 0xf0, 0x74, 0x98, 0xe0,
	0xe1, 0xb7, 0xb6, 0x26, 0x77, 0xf2, 0x93, 0xf3, 0xfe, 0xbb, 0x91, 0xf1, 0x79, 0xbf, 0xef, 0x23,
	0x9b, 0x75, 0x0f, 0xf0, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xda, 0xfb, 0xb0, 0x5c, 0xb4, 0x02,
	0x00, 0x00,
}
