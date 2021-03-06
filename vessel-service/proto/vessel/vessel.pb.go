// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel.proto

package vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_1cb34d421c9f13bf, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_1cb34d421c9f13bf, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_1cb34d421c9f13bf, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("vessel.proto", fileDescriptor_vessel_1cb34d421c9f13bf) }

var fileDescriptor_vessel_1cb34d421c9f13bf = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x36, 0x4d, 0x46, 0x53, 0x64, 0x40, 0x58, 0x8b, 0x42, 0xc8, 0x41, 0x72, 0x90,
	0x1e, 0xea, 0xcd, 0x9b, 0x08, 0x82, 0x1e, 0x53, 0xd0, 0x8b, 0x50, 0xb6, 0xc9, 0xa8, 0x03, 0xf9,
	0x22, 0x59, 0xd2, 0xf6, 0xdf, 0xf8, 0x53, 0x85, 0x4d, 0x52, 0x69, 0x29, 0xbd, 0xcd, 0x7b, 0x6f,
	0x66, 0xf6, 0xcd, 0x5b, 0xb8, 0x68, 0xa8, 0xae, 0x29, 0x9d, 0x95, 0x55, 0xa1, 0x0b, 0xb4, 0x5b,
	0x14, 0xfc, 0x0a, 0xb0, 0xdf, 0x4d, 0x89, 0x13, 0xb0, 0x38, 0x91, 0xc2, 0x17, 0xa1, 0x1b, 0x59,
	0x9c, 0xe0, 0x14, 0x9c, 0x58, 0x95, 0x2a, 0x66, 0xbd, 0x95, 0x96, 0x2f, 0xc2, 0x51, 0xb4, 0xc3,
	0x78, 0x0b, 0x90, 0xa9, 0xcd, 0x72, 0x4d, 0xfc, 0xfd, 0xa3, 0xe5, 0xc0, 0xa8, 0x6e, 0xa6, 0x36,
	0x1f, 0x86, 0x40, 0x84, 0x61, 0xae, 0x32, 0x92, 0x43, 0xb3, 0xcc, 0xd4, 0x78, 0x03, 0xae, 0x6a,
	0x14, 0xa7, 0x6a, 0x95, 0x92, 0x1c, 0xf9, 0x22, 0x74, 0xa2, 0x7f, 0x02, 0xaf, 0xc1, 0x29, 0xd6,
	0x39, 0x55, 0x4b, 0x4e, 0xa4, 0x6d, 0xa6, 0xc6, 0x06, 0xbf, 0x26, 0xc1, 0x1b, 0x78, 0x8b, 0x92,
	0x62, 0xfe, 0xe2, 0x58, 0x69, 0x2e, 0xf2, 0x3d, 0x63, 0xe2, 0xa4, 0x31, 0xeb, 0xc0, 0x58, 0xf0,
	0x09, 0x4e, 0x44, 0x75, 0x59, 0xe4, 0x35, 0xe1, 0x1d, 0x74, 0x21, 0x98, 0x25, 0xe7, 0xf3, 0xc9,
	0xac, 0x4b, 0xa8, 0xcd, 0x23, 0xea, 0x54, 0x0c, 0x61, 0xdc, 0x56, 0xb5, 0xb4, 0xfc, 0xc1, 0x91,
	0xc6, 0x5e, 0x9e, 0x6f, 0xc1, 0x6b, 0xa9, 0x05, 0x55, 0x0d, 0xc7, 0x84, 0x8f, 0xe0, 0xbd, 0x70,
	0x9e, 0x3c, 0xed, 0xce, 0xbc, 0xea, 0x47, 0xf7, 0x2e, 0x9a, 0x5e, 0xf6, 0x74, 0x6f, 0x2e, 0x38,
	0xc3, 0x7b, 0xb0, 0x9f, 0x2b, 0x52, 0x9a, 0xf0, 0xe0, 0xbd, 0x63, 0xdd, 0x2b, 0xdb, 0x7c, 0xeb,
	0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xfb, 0xd3, 0xe1, 0xe6, 0x01, 0x00, 0x00,
}
