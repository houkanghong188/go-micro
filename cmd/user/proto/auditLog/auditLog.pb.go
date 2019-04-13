// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auditLog.proto

package AuditLog

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
	Uid                  int32    `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Page                 int32    `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Reason               string   `protobuf:"bytes,12,opt,name=reason,proto3" json:"reason,omitempty"`
	Status               int32    `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	WorksId              string   `protobuf:"bytes,14,opt,name=Works_id,json=WorksId,proto3" json:"Works_id,omitempty"`
	Ids                  []int32  `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_05adb379ee134fa1, []int{0}
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

func (m *Request) GetUid() int32 {
	if m != nil {
		return m.Uid
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

func (m *Request) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Request) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Request) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *Request) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type AuditLogBracket struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	WorksId              string   `protobuf:"bytes,3,opt,name=works_id,json=worksId,proto3" json:"works_id,omitempty"`
	Reason               string   `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	CreateTime           string   `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditLogBracket) Reset()         { *m = AuditLogBracket{} }
func (m *AuditLogBracket) String() string { return proto.CompactTextString(m) }
func (*AuditLogBracket) ProtoMessage()    {}
func (*AuditLogBracket) Descriptor() ([]byte, []int) {
	return fileDescriptor_05adb379ee134fa1, []int{1}
}

func (m *AuditLogBracket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLogBracket.Unmarshal(m, b)
}
func (m *AuditLogBracket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLogBracket.Marshal(b, m, deterministic)
}
func (m *AuditLogBracket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLogBracket.Merge(m, src)
}
func (m *AuditLogBracket) XXX_Size() int {
	return xxx_messageInfo_AuditLogBracket.Size(m)
}
func (m *AuditLogBracket) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLogBracket.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLogBracket proto.InternalMessageInfo

func (m *AuditLogBracket) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuditLogBracket) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuditLogBracket) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *AuditLogBracket) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *AuditLogBracket) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AuditLogBracket) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type LogShowResponse struct {
	Data                 *AuditLogBracket `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32            `protobuf:"varint,2,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32            `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogShowResponse) Reset()         { *m = LogShowResponse{} }
func (m *LogShowResponse) String() string { return proto.CompactTextString(m) }
func (*LogShowResponse) ProtoMessage()    {}
func (*LogShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05adb379ee134fa1, []int{2}
}

func (m *LogShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogShowResponse.Unmarshal(m, b)
}
func (m *LogShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogShowResponse.Marshal(b, m, deterministic)
}
func (m *LogShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogShowResponse.Merge(m, src)
}
func (m *LogShowResponse) XXX_Size() int {
	return xxx_messageInfo_LogShowResponse.Size(m)
}
func (m *LogShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogShowResponse proto.InternalMessageInfo

func (m *LogShowResponse) GetData() *AuditLogBracket {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LogShowResponse) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

func (m *LogShowResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type AuditLogResponse struct {
	Data                 []*AuditLogBracket `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32              `protobuf:"varint,11,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32              `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuditLogResponse) Reset()         { *m = AuditLogResponse{} }
func (m *AuditLogResponse) String() string { return proto.CompactTextString(m) }
func (*AuditLogResponse) ProtoMessage()    {}
func (*AuditLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05adb379ee134fa1, []int{3}
}

func (m *AuditLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLogResponse.Unmarshal(m, b)
}
func (m *AuditLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLogResponse.Marshal(b, m, deterministic)
}
func (m *AuditLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLogResponse.Merge(m, src)
}
func (m *AuditLogResponse) XXX_Size() int {
	return xxx_messageInfo_AuditLogResponse.Size(m)
}
func (m *AuditLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLogResponse proto.InternalMessageInfo

func (m *AuditLogResponse) GetData() []*AuditLogBracket {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AuditLogResponse) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

func (m *AuditLogResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "AuditLog.Request")
	proto.RegisterType((*AuditLogBracket)(nil), "AuditLog.AuditLogBracket")
	proto.RegisterType((*LogShowResponse)(nil), "AuditLog.LogShowResponse")
	proto.RegisterType((*AuditLogResponse)(nil), "AuditLog.AuditLogResponse")
}

func init() { proto.RegisterFile("auditLog.proto", fileDescriptor_05adb379ee134fa1) }

var fileDescriptor_05adb379ee134fa1 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0x38, 0x71, 0x26, 0x69, 0x5a, 0x56, 0x08, 0x6d, 0xca, 0x81, 0xc8, 0xa7, 0x5e,
	0xc8, 0xa1, 0xdc, 0x91, 0x8a, 0x7a, 0x41, 0xea, 0xc9, 0x45, 0xe2, 0x18, 0x2d, 0xd9, 0x69, 0x58,
	0xd5, 0xc9, 0x1a, 0xef, 0x1a, 0x8b, 0xf0, 0x0d, 0x7c, 0x01, 0xff, 0xc5, 0xf7, 0xa0, 0x9d, 0xf5,
	0x36, 0xa9, 0x83, 0x38, 0xf4, 0xe4, 0x79, 0xcf, 0x33, 0xfb, 0xde, 0x3c, 0xed, 0xc2, 0x54, 0xd4,
	0x52, 0xd9, 0x5b, 0xbd, 0x5e, 0x94, 0x95, 0xb6, 0x9a, 0xa5, 0xd7, 0x2d, 0xce, 0xfe, 0xc4, 0x30,
	0xcc, 0xf1, 0x5b, 0x8d, 0xc6, 0xb2, 0x29, 0xc4, 0x4a, 0xf2, 0xd1, 0x3c, 0xba, 0x4c, 0xf2, 0x58,
	0x49, 0x76, 0x0e, 0xbd, 0x5a, 0x49, 0x3e, 0x26, 0xc2, 0x95, 0x8c, 0x41, 0x7f, 0x2b, 0x36, 0xc8,
	0xa3, 0x79, 0x74, 0x39, 0xca, 0xa9, 0x76, 0x9c, 0xfd, 0x51, 0x22, 0x8f, 0x3d, 0xe7, 0x6a, 0xf6,
	0x12, 0x92, 0xef, 0xa2, 0xa8, 0x91, 0xf7, 0x88, 0xf4, 0x80, 0xbd, 0x81, 0xf1, 0xaa, 0x42, 0x61,
	0x71, 0x69, 0xd5, 0x06, 0x79, 0x9f, 0xfe, 0x81, 0xa7, 0x3e, 0xa9, 0x0d, 0x35, 0xd4, 0xa5, 0x7c,
	0x6c, 0x48, 0x7c, 0x83, 0xa7, 0xa8, 0xe1, 0x35, 0x8c, 0x94, 0x59, 0x4a, 0x2c, 0xd0, 0x22, 0x1f,
	0x90, 0xaf, 0x54, 0x99, 0x1b, 0xc2, 0xce, 0x48, 0x29, 0xd6, 0xc8, 0x87, 0xc4, 0x53, 0xed, 0x06,
	0xdc, 0x77, 0x69, 0xd4, 0x0e, 0x79, 0xea, 0x07, 0x1c, 0x71, 0xa7, 0x76, 0xc8, 0x5e, 0xc1, 0xa0,
	0x42, 0x61, 0xf4, 0x96, 0x4f, 0x48, 0xa9, 0x45, 0x8e, 0x37, 0x56, 0xd8, 0xda, 0xf0, 0x53, 0x9a,
	0x68, 0x11, 0x9b, 0x41, 0xfa, 0x59, 0x57, 0x0f, 0x66, 0xa9, 0x24, 0x9f, 0xd2, 0xc4, 0x90, 0xf0,
	0x47, 0x8a, 0x4a, 0x49, 0xc3, 0x61, 0xde, 0x73, 0x51, 0x29, 0x69, 0xb2, 0xdf, 0x11, 0x9c, 0x85,
	0x94, 0x3f, 0x54, 0x62, 0xf5, 0x80, 0x21, 0xe0, 0xa8, 0x1b, 0x70, 0xbc, 0x0f, 0x78, 0x06, 0x69,
	0x13, 0x24, 0x7c, 0x76, 0xc3, 0xa6, 0x95, 0xd8, 0xbb, 0xed, 0x3f, 0x71, 0x1b, 0xf2, 0x4f, 0x0e,
	0xf2, 0xef, 0x24, 0x3d, 0xe8, 0x26, 0x9d, 0xed, 0xe0, 0xec, 0x56, 0xaf, 0xef, 0xbe, 0xea, 0x26,
	0x47, 0x53, 0xea, 0xad, 0x41, 0xf6, 0x16, 0xfa, 0x52, 0x58, 0x41, 0xf6, 0xc6, 0x57, 0xb3, 0x45,
	0x30, 0xbf, 0xe8, 0x6c, 0x91, 0x53, 0x1b, 0xcb, 0x60, 0x22, 0xee, 0xef, 0x71, 0x65, 0x51, 0xe6,
	0xba, 0x31, 0xed, 0x12, 0x4f, 0x38, 0x77, 0x0d, 0xac, 0xb6, 0xa2, 0xa0, 0x55, 0x92, 0xdc, 0x83,
	0xec, 0x27, 0x9c, 0x87, 0x23, 0x8f, 0xc4, 0x27, 0xf3, 0xde, 0x73, 0xc4, 0xc7, 0xff, 0x13, 0x3f,
	0x3d, 0x10, 0xbf, 0xfa, 0x15, 0xc1, 0xe3, 0xe5, 0x67, 0xd7, 0x30, 0x0d, 0xf5, 0x0d, 0x5a, 0xa1,
	0x0a, 0xf6, 0x62, 0xaf, 0xdc, 0xbe, 0x8a, 0x8b, 0x8b, 0x63, 0x33, 0xc1, 0x76, 0x76, 0xc2, 0xde,
	0xc3, 0x24, 0xb0, 0x2e, 0xcd, 0x7f, 0x1d, 0x70, 0xb0, 0x4d, 0x27, 0xf3, 0xec, 0xe4, 0xcb, 0x80,
	0x1e, 0xe4, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0xaf, 0x88, 0x4f, 0xa2, 0x03, 0x00,
	0x00,
}