// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketContentConf.proto

package marketContentConf

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
	Mark                 []string `protobuf:"bytes,1,rep,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f518a3ce46308c3, []int{0}
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

func (m *Request) GetMark() []string {
	if m != nil {
		return m.Mark
	}
	return nil
}

type Response struct {
	ContentConf          []*ContentConf `protobuf:"bytes,1,rep,name=contentConf,proto3" json:"contentConf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f518a3ce46308c3, []int{1}
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

func (m *Response) GetContentConf() []*ContentConf {
	if m != nil {
		return m.ContentConf
	}
	return nil
}

type ContentConf struct {
	Id                   int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string              `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Description          string              `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Sort                 int32               `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	ParentId             int32               `protobuf:"varint,6,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Mark                 string              `protobuf:"bytes,7,opt,name=mark,proto3" json:"mark,omitempty"`
	RelevantCategory     int32               `protobuf:"varint,8,opt,name=relevantCategory,proto3" json:"relevantCategory,omitempty"`
	Recommend            []*ContentRecommend `protobuf:"bytes,9,rep,name=recommend,proto3" json:"recommend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ContentConf) Reset()         { *m = ContentConf{} }
func (m *ContentConf) String() string { return proto.CompactTextString(m) }
func (*ContentConf) ProtoMessage()    {}
func (*ContentConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f518a3ce46308c3, []int{2}
}

func (m *ContentConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentConf.Unmarshal(m, b)
}
func (m *ContentConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentConf.Marshal(b, m, deterministic)
}
func (m *ContentConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentConf.Merge(m, src)
}
func (m *ContentConf) XXX_Size() int {
	return xxx_messageInfo_ContentConf.Size(m)
}
func (m *ContentConf) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentConf.DiscardUnknown(m)
}

var xxx_messageInfo_ContentConf proto.InternalMessageInfo

func (m *ContentConf) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContentConf) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContentConf) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ContentConf) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContentConf) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *ContentConf) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *ContentConf) GetMark() string {
	if m != nil {
		return m.Mark
	}
	return ""
}

func (m *ContentConf) GetRelevantCategory() int32 {
	if m != nil {
		return m.RelevantCategory
	}
	return 0
}

func (m *ContentConf) GetRecommend() []*ContentRecommend {
	if m != nil {
		return m.Recommend
	}
	return nil
}

type ContentRecommend struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ContentId            int32    `protobuf:"varint,3,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Sort                 int32    `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	NexusTemplateId      string   `protobuf:"bytes,5,opt,name=nexusTemplateId,proto3" json:"nexusTemplateId,omitempty"`
	NexusCruxWords       string   `protobuf:"bytes,6,opt,name=nexusCruxWords,proto3" json:"nexusCruxWords,omitempty"`
	NexusScene           string   `protobuf:"bytes,7,opt,name=nexusScene,proto3" json:"nexusScene,omitempty"`
	NexusIndustry        string   `protobuf:"bytes,8,opt,name=nexusIndustry,proto3" json:"nexusIndustry,omitempty"`
	NexusColor           string   `protobuf:"bytes,9,opt,name=nexusColor,proto3" json:"nexusColor,omitempty"`
	NexusStyle           string   `protobuf:"bytes,10,opt,name=nexusStyle,proto3" json:"nexusStyle,omitempty"`
	RelevantCategory     int32    `protobuf:"varint,11,opt,name=relevantCategory,proto3" json:"relevantCategory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentRecommend) Reset()         { *m = ContentRecommend{} }
func (m *ContentRecommend) String() string { return proto.CompactTextString(m) }
func (*ContentRecommend) ProtoMessage()    {}
func (*ContentRecommend) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f518a3ce46308c3, []int{3}
}

func (m *ContentRecommend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentRecommend.Unmarshal(m, b)
}
func (m *ContentRecommend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentRecommend.Marshal(b, m, deterministic)
}
func (m *ContentRecommend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRecommend.Merge(m, src)
}
func (m *ContentRecommend) XXX_Size() int {
	return xxx_messageInfo_ContentRecommend.Size(m)
}
func (m *ContentRecommend) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRecommend.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRecommend proto.InternalMessageInfo

func (m *ContentRecommend) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContentRecommend) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContentRecommend) GetContentId() int32 {
	if m != nil {
		return m.ContentId
	}
	return 0
}

func (m *ContentRecommend) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *ContentRecommend) GetNexusTemplateId() string {
	if m != nil {
		return m.NexusTemplateId
	}
	return ""
}

func (m *ContentRecommend) GetNexusCruxWords() string {
	if m != nil {
		return m.NexusCruxWords
	}
	return ""
}

func (m *ContentRecommend) GetNexusScene() string {
	if m != nil {
		return m.NexusScene
	}
	return ""
}

func (m *ContentRecommend) GetNexusIndustry() string {
	if m != nil {
		return m.NexusIndustry
	}
	return ""
}

func (m *ContentRecommend) GetNexusColor() string {
	if m != nil {
		return m.NexusColor
	}
	return ""
}

func (m *ContentRecommend) GetNexusStyle() string {
	if m != nil {
		return m.NexusStyle
	}
	return ""
}

func (m *ContentRecommend) GetRelevantCategory() int32 {
	if m != nil {
		return m.RelevantCategory
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "marketContentConf.Request")
	proto.RegisterType((*Response)(nil), "marketContentConf.Response")
	proto.RegisterType((*ContentConf)(nil), "marketContentConf.ContentConf")
	proto.RegisterType((*ContentRecommend)(nil), "marketContentConf.ContentRecommend")
}

func init() { proto.RegisterFile("marketContentConf.proto", fileDescriptor_4f518a3ce46308c3) }

var fileDescriptor_4f518a3ce46308c3 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd5, 0x6c, 0xb3, 0x5b, 0x4f, 0xc4, 0xb2, 0x6b, 0x21, 0x61, 0x15, 0x58, 0x55, 0x05,
	0xa1, 0x8a, 0xc3, 0x1e, 0x96, 0x17, 0x58, 0x94, 0x53, 0x24, 0xb8, 0x18, 0x10, 0xe7, 0x10, 0x0f,
	0x28, 0x22, 0xb1, 0x83, 0xed, 0xa0, 0xf6, 0x5d, 0x78, 0x46, 0x9e, 0x01, 0x65, 0x1a, 0x12, 0x93,
	0x06, 0x89, 0x53, 0x67, 0xbe, 0xdf, 0x33, 0xd3, 0xfc, 0xa3, 0x81, 0xc7, 0x75, 0x6e, 0xbf, 0xa1,
	0x4f, 0x8d, 0xf6, 0xa8, 0xbb, 0x9f, 0x2f, 0xb7, 0x8d, 0x35, 0xde, 0xf0, 0xeb, 0x13, 0x61, 0xfb,
	0x0c, 0x2e, 0x24, 0x7e, 0x6f, 0xd1, 0x79, 0xce, 0x61, 0xd9, 0xe9, 0x62, 0xb1, 0x39, 0xdb, 0x31,
	0x49, 0xf1, 0xf6, 0x2d, 0xac, 0x24, 0xba, 0xc6, 0x68, 0x87, 0xfc, 0x1e, 0x92, 0x62, 0xac, 0xa4,
	0x67, 0xc9, 0xdd, 0xcd, 0xed, 0xe9, 0xb0, 0x20, 0x96, 0x61, 0xc9, 0xf6, 0x67, 0x04, 0x49, 0x20,
	0xf2, 0x4b, 0x88, 0x4a, 0x25, 0x16, 0x9b, 0xc5, 0x2e, 0x96, 0x51, 0xa9, 0xf8, 0x23, 0x88, 0x7d,
	0xe9, 0x2b, 0x14, 0xd1, 0x66, 0xb1, 0x63, 0xf2, 0x98, 0x74, 0xff, 0xcb, 0x1f, 0x1a, 0x14, 0x67,
	0x04, 0x29, 0xe6, 0x1b, 0x48, 0x14, 0xba, 0xc2, 0x96, 0x8d, 0x2f, 0x8d, 0x16, 0x4b, 0x92, 0x42,
	0xd4, 0x55, 0x39, 0x63, 0xbd, 0x88, 0xa9, 0x3b, 0xc5, 0x7c, 0x0d, 0xab, 0x26, 0xb7, 0xa8, 0x7d,
	0xa6, 0xc4, 0x39, 0xf1, 0x21, 0x1f, 0xbe, 0xfe, 0xe2, 0x38, 0xa5, 0x8b, 0xf9, 0x2b, 0xb8, 0xb2,
	0x58, 0xe1, 0x8f, 0x5c, 0xfb, 0x34, 0xf7, 0xf8, 0xd5, 0xd8, 0x83, 0x58, 0x51, 0xdd, 0x09, 0xe7,
	0x6f, 0x80, 0x59, 0x2c, 0x4c, 0x5d, 0xa3, 0x56, 0x82, 0x91, 0x37, 0xcf, 0xff, 0xed, 0x8d, 0xfc,
	0xf3, 0x54, 0x8e, 0x55, 0xdb, 0x5f, 0x11, 0x5c, 0x4d, 0xf5, 0xff, 0xf4, 0xe8, 0x29, 0xb0, 0xde,
	0xe8, 0x4c, 0x91, 0x51, 0xb1, 0x1c, 0xc1, 0xe0, 0xc5, 0x32, 0xf0, 0x62, 0x07, 0x0f, 0x35, 0xee,
	0x5b, 0xf7, 0x01, 0xeb, 0xa6, 0xca, 0x3d, 0x66, 0x8a, 0xac, 0x62, 0x72, 0x8a, 0xf9, 0x4b, 0xb8,
	0x24, 0x94, 0xda, 0x76, 0xff, 0xc9, 0x58, 0xe5, 0xc8, 0x3b, 0x26, 0x27, 0x94, 0xdf, 0x00, 0x10,
	0x79, 0x5f, 0xa0, 0xc6, 0xde, 0xc7, 0x80, 0xf0, 0x17, 0xf0, 0x80, 0xb2, 0x4c, 0xab, 0xd6, 0xf9,
	0xde, 0x4a, 0x26, 0xff, 0x86, 0x43, 0x97, 0xd4, 0x54, 0xc6, 0x0a, 0x16, 0x74, 0x21, 0x32, 0x4e,
	0xf1, 0x87, 0x0a, 0x05, 0x84, 0x53, 0x3a, 0x32, 0xbb, 0xb3, 0x64, 0x7e, 0x67, 0x77, 0x1f, 0xe1,
	0xfa, 0xdd, 0x74, 0x43, 0xfc, 0x1e, 0xe2, 0x4c, 0x2b, 0xdc, 0xf3, 0xf5, 0xcc, 0xfa, 0xfa, 0x5b,
	0x59, 0x3f, 0x99, 0xd5, 0x8e, 0x87, 0xf2, 0xf9, 0x9c, 0xae, 0xed, 0xf5, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0x1a, 0x1e, 0x14, 0x88, 0x03, 0x00, 0x00,
}
