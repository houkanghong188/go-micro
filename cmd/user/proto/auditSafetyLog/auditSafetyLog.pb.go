// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auditSafetyLog.proto

package auditSafetyLog

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

type Request struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WorksId              string   `protobuf:"bytes,2,opt,name=works_id,json=worksId,proto3" json:"works_id,omitempty"`
	Scene                string   `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Uid                  int32    `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	Limit                int32    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5fe151c65caf41, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *Request) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Request) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Request) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type AuditSafetyLogBracket struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	WorksId              string   `protobuf:"bytes,3,opt,name=works_id,json=worksId,proto3" json:"works_id,omitempty"`
	Reason               string   `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	CreateTime           string   `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Scene                string   `protobuf:"bytes,7,opt,name=scene,proto3" json:"scene,omitempty"`
	Content              string   `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditSafetyLogBracket) Reset()         { *m = AuditSafetyLogBracket{} }
func (m *AuditSafetyLogBracket) String() string { return proto.CompactTextString(m) }
func (*AuditSafetyLogBracket) ProtoMessage()    {}
func (*AuditSafetyLogBracket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5fe151c65caf41, []int{1}
}

func (m *AuditSafetyLogBracket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditSafetyLogBracket.Unmarshal(m, b)
}
func (m *AuditSafetyLogBracket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditSafetyLogBracket.Marshal(b, m, deterministic)
}
func (m *AuditSafetyLogBracket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditSafetyLogBracket.Merge(m, src)
}
func (m *AuditSafetyLogBracket) XXX_Size() int {
	return xxx_messageInfo_AuditSafetyLogBracket.Size(m)
}
func (m *AuditSafetyLogBracket) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditSafetyLogBracket.DiscardUnknown(m)
}

var xxx_messageInfo_AuditSafetyLogBracket proto.InternalMessageInfo

func (m *AuditSafetyLogBracket) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuditSafetyLogBracket) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuditSafetyLogBracket) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *AuditSafetyLogBracket) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *AuditSafetyLogBracket) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AuditSafetyLogBracket) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AuditSafetyLogBracket) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *AuditSafetyLogBracket) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Response struct {
	Data                 []*AuditSafetyLogBracket `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	Total                int32                    `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5fe151c65caf41, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetData() []*AuditSafetyLogBracket {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "auditSafetyLog.Request")
	proto.RegisterType((*AuditSafetyLogBracket)(nil), "auditSafetyLog.AuditSafetyLogBracket")
	proto.RegisterType((*Response)(nil), "auditSafetyLog.Response")
}

func init() { proto.RegisterFile("auditSafetyLog.proto", fileDescriptor_6c5fe151c65caf41) }

var fileDescriptor_6c5fe151c65caf41 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xb5, 0x2d, 0x6d, 0x61, 0x50, 0x62, 0x36, 0x88, 0xa3, 0x17, 0x49, 0x13, 0x13, 0x4e, 0x1c,
	0xf0, 0x64, 0xe2, 0x45, 0x6f, 0x24, 0xc6, 0x43, 0xf5, 0xe6, 0x81, 0xac, 0xdd, 0xc1, 0x6c, 0x80,
	0x2e, 0x76, 0x87, 0x28, 0x3f, 0xe3, 0x3f, 0xf9, 0x47, 0x86, 0x6d, 0x23, 0x14, 0xbc, 0xcd, 0x7b,
	0x9b, 0x99, 0xf7, 0xde, 0xcc, 0x42, 0x57, 0xae, 0x94, 0xe6, 0x67, 0x39, 0x25, 0x5e, 0x3f, 0x9a,
	0xf7, 0xe1, 0xb2, 0x30, 0x6c, 0x44, 0xa7, 0xce, 0x26, 0xdf, 0x1e, 0xc4, 0x29, 0x7d, 0xac, 0xc8,
	0xb2, 0xe8, 0x80, 0xaf, 0x15, 0x7a, 0x7d, 0x6f, 0x10, 0xa6, 0xbe, 0x56, 0xe2, 0x02, 0x9a, 0x9f,
	0xa6, 0x98, 0xd9, 0x89, 0x56, 0xe8, 0xf7, 0xbd, 0x41, 0x2b, 0x8d, 0x1d, 0x1e, 0x2b, 0xd1, 0x85,
	0xd0, 0x66, 0x94, 0x13, 0x06, 0x8e, 0x2f, 0x81, 0x10, 0xd0, 0xe0, 0xf5, 0x92, 0xb0, 0xe1, 0x48,
	0x57, 0x8b, 0x53, 0x08, 0x56, 0x5a, 0x61, 0xe8, 0xa6, 0x6e, 0xca, 0x4d, 0xef, 0x5c, 0x2f, 0x34,
	0x63, 0xe4, 0xb8, 0x12, 0x88, 0x1e, 0x44, 0x66, 0x3a, 0xb5, 0xc4, 0x18, 0x3b, 0xba, 0x42, 0xc9,
	0x8f, 0x07, 0x67, 0xf7, 0x35, 0xcf, 0x0f, 0x85, 0xcc, 0x66, 0x74, 0x68, 0xb7, 0x52, 0xf2, 0xb7,
	0x4a, 0xbb, 0x01, 0x82, 0x7a, 0x80, 0x1e, 0x44, 0x05, 0x49, 0x6b, 0xf2, 0xca, 0x6c, 0x85, 0xfe,
	0x22, 0x84, 0x3b, 0x11, 0xae, 0xa0, 0x9d, 0x15, 0x24, 0x99, 0x26, 0xac, 0x17, 0xe4, 0x6c, 0xb7,
	0x52, 0x28, 0xa9, 0x17, 0xbd, 0xa0, 0xed, 0x36, 0xe2, 0xdd, 0x6d, 0x20, 0xc4, 0x99, 0xc9, 0x99,
	0x72, 0xc6, 0x66, 0x29, 0x5e, 0xc1, 0xe4, 0x15, 0x9a, 0x29, 0xd9, 0xa5, 0xc9, 0x2d, 0x89, 0x5b,
	0x68, 0x28, 0xc9, 0x12, 0x8f, 0xfb, 0xc1, 0xa0, 0x3d, 0xba, 0x1e, 0xee, 0x5d, 0xed, 0xdf, 0xe8,
	0xa9, 0x6b, 0xd9, 0xc8, 0xb2, 0x61, 0x39, 0xc7, 0x93, 0x72, 0x91, 0x0e, 0x8c, 0x9e, 0xa0, 0x53,
	0x6f, 0x12, 0x77, 0x10, 0x8e, 0x73, 0x45, 0x5f, 0xe2, 0x7c, 0x7f, 0x7a, 0x75, 0xf9, 0x4b, 0x3c,
	0x7c, 0x28, 0xed, 0x25, 0x47, 0x6f, 0x91, 0xfb, 0x38, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x63, 0x26, 0xa2, 0x50, 0x02, 0x00, 0x00,
}
