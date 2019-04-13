// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auditConf.proto

package auditConf

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
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Page                 int32    `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_703ab07fe430db0d, []int{0}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Request) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Request) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Request) GetIsDelete() int32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *Request) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Request) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type AuditConfBracket struct {
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditConfBracket) Reset()         { *m = AuditConfBracket{} }
func (m *AuditConfBracket) String() string { return proto.CompactTextString(m) }
func (*AuditConfBracket) ProtoMessage()    {}
func (*AuditConfBracket) Descriptor() ([]byte, []int) {
	return fileDescriptor_703ab07fe430db0d, []int{1}
}

func (m *AuditConfBracket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditConfBracket.Unmarshal(m, b)
}
func (m *AuditConfBracket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditConfBracket.Marshal(b, m, deterministic)
}
func (m *AuditConfBracket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditConfBracket.Merge(m, src)
}
func (m *AuditConfBracket) XXX_Size() int {
	return xxx_messageInfo_AuditConfBracket.Size(m)
}
func (m *AuditConfBracket) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditConfBracket.DiscardUnknown(m)
}

var xxx_messageInfo_AuditConfBracket proto.InternalMessageInfo

func (m *AuditConfBracket) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuditConfBracket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuditConfBracket) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AuditConfBracket) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AuditConfBracket) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AuditConfBracket) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *AuditConfBracket) GetIsDelete() int32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

type Response struct {
	Data                 []*AuditConfBracket `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32               `protobuf:"varint,11,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32               `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_703ab07fe430db0d, []int{2}
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

func (m *Response) GetData() []*AuditConfBracket {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

func (m *Response) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*AuditConfBracket)(nil), "AuditConfBracket")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("auditConf.proto", fileDescriptor_703ab07fe430db0d) }

var fileDescriptor_703ab07fe430db0d = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x2d, 0xb4, 0xd0, 0x0e, 0xf8, 0x6f, 0xe3, 0x61, 0xa3, 0x51, 0x48, 0x13, 0x13, 0x4e,
	0x3d, 0xe0, 0x13, 0x28, 0x5e, 0xbc, 0x16, 0x3d, 0x93, 0x95, 0x1d, 0x74, 0x03, 0x74, 0x2b, 0xbb,
	0x15, 0xe5, 0x19, 0x7c, 0x27, 0x1f, 0xc7, 0xd7, 0x30, 0x3b, 0x2b, 0x24, 0x9a, 0xf4, 0x01, 0x3c,
	0xed, 0xcc, 0xef, 0xfb, 0x76, 0x32, 0x5f, 0x32, 0x70, 0x28, 0x2a, 0xa9, 0xec, 0x48, 0x17, 0xb3,
	0xac, 0x5c, 0x69, 0xab, 0xd3, 0xaf, 0x00, 0xda, 0x39, 0xbe, 0x54, 0x68, 0x2c, 0x3b, 0x80, 0x86,
	0x92, 0x3c, 0xe9, 0x07, 0x83, 0x28, 0x6f, 0x28, 0xc9, 0x18, 0x84, 0x85, 0x58, 0x22, 0x0f, 0xfa,
	0xc1, 0x20, 0xc9, 0xa9, 0x76, 0xcc, 0xbe, 0x97, 0xc8, 0x1b, 0x9e, 0xb9, 0x9a, 0x9d, 0x40, 0xf4,
	0x2a, 0x16, 0x15, 0xf2, 0x26, 0x41, 0xdf, 0xb0, 0x1e, 0x74, 0xa6, 0x2b, 0x14, 0x16, 0x27, 0x56,
	0x2d, 0x91, 0x87, 0xa4, 0x81, 0x47, 0xf7, 0x6a, 0x49, 0x86, 0xaa, 0x94, 0x3b, 0x43, 0xe4, 0x0d,
	0x1e, 0x91, 0xe1, 0x0c, 0x12, 0x65, 0x26, 0x12, 0x17, 0x68, 0x91, 0xb7, 0x68, 0xad, 0x58, 0x99,
	0x5b, 0xea, 0xdd, 0x22, 0xa5, 0x78, 0x42, 0xde, 0x26, 0x4e, 0xb5, 0xfb, 0xe0, 0xde, 0x89, 0x51,
	0x1b, 0xe4, 0xb1, 0xff, 0xe0, 0xc0, 0x58, 0x6d, 0x30, 0xfd, 0x0c, 0xe0, 0xe8, 0x7a, 0x9b, 0xfe,
	0x66, 0x25, 0xa6, 0x73, 0xfc, 0x5f, 0x91, 0xd3, 0x39, 0xc4, 0x39, 0x9a, 0x52, 0x17, 0x06, 0xd9,
	0x25, 0x84, 0x52, 0x58, 0xc1, 0xbb, 0xfd, 0xe6, 0xa0, 0x33, 0x3c, 0xce, 0xfe, 0x26, 0xcb, 0x49,
	0x66, 0x29, 0x74, 0xc5, 0x6c, 0x86, 0x53, 0x8b, 0x32, 0xd7, 0x6b, 0xc3, 0x3b, 0x34, 0xf2, 0x17,
	0x73, 0x59, 0xac, 0xb6, 0x62, 0xc1, 0xf7, 0x49, 0xf4, 0xcd, 0xf0, 0x23, 0x80, 0x64, 0x37, 0x94,
	0xf5, 0xa0, 0x35, 0xa2, 0x18, 0x2c, 0xce, 0x7e, 0xce, 0xe5, 0x34, 0xc9, 0xb6, 0xdb, 0xa4, 0x7b,
	0xce, 0xf0, 0x40, 0x31, 0xea, 0x0c, 0xe7, 0x10, 0x8e, 0x9f, 0xf5, 0xba, 0x4e, 0xbe, 0x80, 0xe8,
	0xae, 0x90, 0xf8, 0x56, 0xa3, 0x3f, 0xb6, 0xe8, 0x5c, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xab, 0xe0, 0x9a, 0xa9, 0xc1, 0x02, 0x00, 0x00,
}