// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trial.proto

package trial

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

type CreateRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contract             string   `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38dfa79a07a63357, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *CreateRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type CreateResponse struct {
	AffectedRows         int32    `protobuf:"varint,1,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38dfa79a07a63357, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

type ShowRequest struct {
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowRequest) Reset()         { *m = ShowRequest{} }
func (m *ShowRequest) String() string { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()    {}
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38dfa79a07a63357, []int{2}
}

func (m *ShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRequest.Unmarshal(m, b)
}
func (m *ShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRequest.Marshal(b, m, deterministic)
}
func (m *ShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRequest.Merge(m, src)
}
func (m *ShowRequest) XXX_Size() int {
	return xxx_messageInfo_ShowRequest.Size(m)
}
func (m *ShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRequest proto.InternalMessageInfo

func (m *ShowRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ShowResponse struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contract             string   `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38dfa79a07a63357, []int{3}
}

func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (m *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(m, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ShowResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShowResponse) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ShowResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "trial.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "trial.CreateResponse")
	proto.RegisterType((*ShowRequest)(nil), "trial.ShowRequest")
	proto.RegisterType((*ShowResponse)(nil), "trial.ShowResponse")
}

func init() { proto.RegisterFile("trial.proto", fileDescriptor_38dfa79a07a63357) }

var fileDescriptor_38dfa79a07a63357 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x6d, 0x6d, 0x16, 0x9d, 0xae, 0x8b, 0x8c, 0x2b, 0x94, 0x82, 0xb0, 0xe4, 0xb4, 0xa7,
	0x05, 0xff, 0x80, 0x67, 0xd9, 0x37, 0x88, 0x4f, 0x10, 0xdb, 0xa9, 0x2d, 0xb8, 0x49, 0x6d, 0xa6,
	0x2c, 0xfb, 0xf6, 0xd2, 0xa4, 0x15, 0xeb, 0x7d, 0x6f, 0xf3, 0x7d, 0x93, 0xe4, 0xfb, 0x65, 0x06,
	0x52, 0xee, 0x1a, 0xfd, 0xb5, 0x6b, 0x3b, 0xcb, 0x16, 0x85, 0x17, 0xf2, 0x13, 0x6e, 0xf6, 0x1d,
	0x69, 0x26, 0x45, 0xdf, 0x3d, 0x39, 0xc6, 0x5b, 0xb8, 0xec, 0x9b, 0x32, 0x8b, 0x36, 0xd1, 0x56,
	0xa8, 0xa1, 0x44, 0x84, 0xc4, 0xe8, 0x03, 0x65, 0xf1, 0x26, 0xda, 0x5e, 0x2b, 0x5f, 0x63, 0x0e,
	0x57, 0x85, 0x35, 0xdc, 0xe9, 0x82, 0xb3, 0xc4, 0xfb, 0xbf, 0x1a, 0xd7, 0x20, 0xda, 0xda, 0x1a,
	0xca, 0x84, 0x6f, 0x04, 0x21, 0x5f, 0x60, 0x35, 0x05, 0xb9, 0xd6, 0x1a, 0x47, 0x28, 0x61, 0xa9,
	0xab, 0x8a, 0x0a, 0xa6, 0x52, 0xd9, 0xa3, 0x1b, 0x23, 0x67, 0x9e, 0x7c, 0x80, 0xf4, 0xbd, 0xb6,
	0xc7, 0x09, 0x6e, 0x05, 0x71, 0x53, 0xfa, 0x40, 0xa1, 0xe2, 0xa6, 0x94, 0x15, 0x2c, 0x43, 0x7b,
	0x7c, 0xf2, 0x4c, 0xf0, 0x4f, 0x27, 0x48, 0xdf, 0x7a, 0xae, 0xf7, 0xf6, 0xd0, 0x6a, 0x73, 0xc2,
	0x47, 0x48, 0x86, 0x58, 0xc4, 0x5d, 0x98, 0xe8, 0x1f, 0xc4, 0xfc, 0x6e, 0xe6, 0x05, 0x2e, 0x79,
	0x81, 0xaf, 0xb0, 0x08, 0xdf, 0xc7, 0xf5, 0x78, 0x60, 0x36, 0xf6, 0xfc, 0xfe, 0x9f, 0x3b, 0x5d,
	0xfc, 0x58, 0xf8, 0x75, 0x3d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x80, 0x52, 0xf0, 0xbd,
	0x01, 0x00, 0x00,
}
