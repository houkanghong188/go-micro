// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worksAudit.proto

package worksAudit

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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Page                 int32    `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	WorksId              string   `protobuf:"bytes,10,opt,name=works_id,json=worksId,proto3" json:"works_id,omitempty"`
	Status               int32    `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Uid                  int32    `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	Reason               string   `protobuf:"bytes,13,opt,name=reason,proto3" json:"reason,omitempty"`
	AuditStatus          []int32  `protobuf:"varint,14,rep,packed,name=AuditStatus,proto3" json:"AuditStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{0}
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

func (m *Request) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Request) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Request) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Request) GetAuditStatus() []int32 {
	if m != nil {
		return m.AuditStatus
	}
	return nil
}

type Response struct {
	Data                 []*Response_Notify `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32              `protobuf:"varint,11,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32              `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{1}
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

func (m *Response) GetData() []*Response_Notify {
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

type Response_Notify struct {
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventId              string   `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Status               int32    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Title                string   `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Notify) Reset()         { *m = Response_Notify{} }
func (m *Response_Notify) String() string { return proto.CompactTextString(m) }
func (*Response_Notify) ProtoMessage()    {}
func (*Response_Notify) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{1, 0}
}

func (m *Response_Notify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Notify.Unmarshal(m, b)
}
func (m *Response_Notify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Notify.Marshal(b, m, deterministic)
}
func (m *Response_Notify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Notify.Merge(m, src)
}
func (m *Response_Notify) XXX_Size() int {
	return xxx_messageInfo_Response_Notify.Size(m)
}
func (m *Response_Notify) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Notify.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Notify proto.InternalMessageInfo

func (m *Response_Notify) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response_Notify) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Response_Notify) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Response_Notify) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response_Notify) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Response_Notify) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Response_Notify) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Response_Notify) GetIsDelete() int32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *Response_Notify) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Response_Notify) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type WorksResponse struct {
	Data                 []*WorksResponse_Notify `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32                   `protobuf:"varint,11,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32                   `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *WorksResponse) Reset()         { *m = WorksResponse{} }
func (m *WorksResponse) String() string { return proto.CompactTextString(m) }
func (*WorksResponse) ProtoMessage()    {}
func (*WorksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{2}
}

func (m *WorksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorksResponse.Unmarshal(m, b)
}
func (m *WorksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorksResponse.Marshal(b, m, deterministic)
}
func (m *WorksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorksResponse.Merge(m, src)
}
func (m *WorksResponse) XXX_Size() int {
	return xxx_messageInfo_WorksResponse.Size(m)
}
func (m *WorksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorksResponse proto.InternalMessageInfo

func (m *WorksResponse) GetData() []*WorksResponse_Notify {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WorksResponse) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

func (m *WorksResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type WorksResponse_Notify struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WorksId              string   `protobuf:"bytes,2,opt,name=worksId,proto3" json:"worksId,omitempty"`
	Uid                  int32    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Comment              string   `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	Thumb                string   `protobuf:"bytes,7,opt,name=thumb,proto3" json:"thumb,omitempty"`
	FirstImg             string   `protobuf:"bytes,8,opt,name=firstImg,proto3" json:"firstImg,omitempty"`
	Version              string   `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	EditorVersion        string   `protobuf:"bytes,10,opt,name=editorVersion,proto3" json:"editorVersion,omitempty"`
	TemplateId           string   `protobuf:"bytes,11,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ParentEventId        string   `protobuf:"bytes,12,opt,name=parent_event_id,json=parentEventId,proto3" json:"parent_event_id,omitempty"`
	Identity             int32    `protobuf:"varint,13,opt,name=identity,proto3" json:"identity,omitempty"`
	Category             int32    `protobuf:"varint,14,opt,name=category,proto3" json:"category,omitempty"`
	SizeId               int32    `protobuf:"varint,15,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	SecondaryCategory    int32    `protobuf:"varint,16,opt,name=secondary_category,json=secondaryCategory,proto3" json:"secondary_category,omitempty"`
	PageWidth            int32    `protobuf:"varint,20,opt,name=page_width,json=pageWidth,proto3" json:"page_width,omitempty"`
	PageHeight           int32    `protobuf:"varint,17,opt,name=page_height,json=pageHeight,proto3" json:"page_height,omitempty"`
	OtherProperty        string   `protobuf:"bytes,18,opt,name=other_property,json=otherProperty,proto3" json:"other_property,omitempty"`
	SuiteAppType         int32    `protobuf:"varint,21,opt,name=SuiteAppType,proto3" json:"SuiteAppType,omitempty"`
	CreateDevice         int32    `protobuf:"varint,22,opt,name=CreateDevice,proto3" json:"CreateDevice,omitempty"`
	UpdateDevice         int32    `protobuf:"varint,23,opt,name=UpdateDevice,proto3" json:"UpdateDevice,omitempty"`
	CreateSite           int32    `protobuf:"varint,24,opt,name=CreateSite,proto3" json:"CreateSite,omitempty"`
	CreateTime           int32    `protobuf:"varint,25,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime           int32    `protobuf:"varint,26,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	IsDelete             string   `protobuf:"bytes,27,opt,name=IsDelete,proto3" json:"IsDelete,omitempty"`
	IsBoughtTemplate     string   `protobuf:"bytes,28,opt,name=IsBoughtTemplate,proto3" json:"IsBoughtTemplate,omitempty"`
	IsUsedLocalFonts     string   `protobuf:"bytes,29,opt,name=IsUsedLocalFonts,proto3" json:"IsUsedLocalFonts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorksResponse_Notify) Reset()         { *m = WorksResponse_Notify{} }
func (m *WorksResponse_Notify) String() string { return proto.CompactTextString(m) }
func (*WorksResponse_Notify) ProtoMessage()    {}
func (*WorksResponse_Notify) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{2, 0}
}

func (m *WorksResponse_Notify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorksResponse_Notify.Unmarshal(m, b)
}
func (m *WorksResponse_Notify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorksResponse_Notify.Marshal(b, m, deterministic)
}
func (m *WorksResponse_Notify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorksResponse_Notify.Merge(m, src)
}
func (m *WorksResponse_Notify) XXX_Size() int {
	return xxx_messageInfo_WorksResponse_Notify.Size(m)
}
func (m *WorksResponse_Notify) XXX_DiscardUnknown() {
	xxx_messageInfo_WorksResponse_Notify.DiscardUnknown(m)
}

var xxx_messageInfo_WorksResponse_Notify proto.InternalMessageInfo

func (m *WorksResponse_Notify) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WorksResponse_Notify) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *WorksResponse_Notify) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *WorksResponse_Notify) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *WorksResponse_Notify) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *WorksResponse_Notify) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *WorksResponse_Notify) GetThumb() string {
	if m != nil {
		return m.Thumb
	}
	return ""
}

func (m *WorksResponse_Notify) GetFirstImg() string {
	if m != nil {
		return m.FirstImg
	}
	return ""
}

func (m *WorksResponse_Notify) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *WorksResponse_Notify) GetEditorVersion() string {
	if m != nil {
		return m.EditorVersion
	}
	return ""
}

func (m *WorksResponse_Notify) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *WorksResponse_Notify) GetParentEventId() string {
	if m != nil {
		return m.ParentEventId
	}
	return ""
}

func (m *WorksResponse_Notify) GetIdentity() int32 {
	if m != nil {
		return m.Identity
	}
	return 0
}

func (m *WorksResponse_Notify) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *WorksResponse_Notify) GetSizeId() int32 {
	if m != nil {
		return m.SizeId
	}
	return 0
}

func (m *WorksResponse_Notify) GetSecondaryCategory() int32 {
	if m != nil {
		return m.SecondaryCategory
	}
	return 0
}

func (m *WorksResponse_Notify) GetPageWidth() int32 {
	if m != nil {
		return m.PageWidth
	}
	return 0
}

func (m *WorksResponse_Notify) GetPageHeight() int32 {
	if m != nil {
		return m.PageHeight
	}
	return 0
}

func (m *WorksResponse_Notify) GetOtherProperty() string {
	if m != nil {
		return m.OtherProperty
	}
	return ""
}

func (m *WorksResponse_Notify) GetSuiteAppType() int32 {
	if m != nil {
		return m.SuiteAppType
	}
	return 0
}

func (m *WorksResponse_Notify) GetCreateDevice() int32 {
	if m != nil {
		return m.CreateDevice
	}
	return 0
}

func (m *WorksResponse_Notify) GetUpdateDevice() int32 {
	if m != nil {
		return m.UpdateDevice
	}
	return 0
}

func (m *WorksResponse_Notify) GetCreateSite() int32 {
	if m != nil {
		return m.CreateSite
	}
	return 0
}

func (m *WorksResponse_Notify) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *WorksResponse_Notify) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *WorksResponse_Notify) GetIsDelete() string {
	if m != nil {
		return m.IsDelete
	}
	return ""
}

func (m *WorksResponse_Notify) GetIsBoughtTemplate() string {
	if m != nil {
		return m.IsBoughtTemplate
	}
	return ""
}

func (m *WorksResponse_Notify) GetIsUsedLocalFonts() string {
	if m != nil {
		return m.IsUsedLocalFonts
	}
	return ""
}

type DailyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	IsDelete             int32    `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Page                 int32    `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	WorksId              string   `protobuf:"bytes,10,opt,name=works_id,json=worksId,proto3" json:"works_id,omitempty"`
	Status               int32    `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Uids                 []int32  `protobuf:"varint,11,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	WorksIds             []string `protobuf:"bytes,14,rep,name=WorksIds,proto3" json:"WorksIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DailyRequest) Reset()         { *m = DailyRequest{} }
func (m *DailyRequest) String() string { return proto.CompactTextString(m) }
func (*DailyRequest) ProtoMessage()    {}
func (*DailyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{3}
}

func (m *DailyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyRequest.Unmarshal(m, b)
}
func (m *DailyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyRequest.Marshal(b, m, deterministic)
}
func (m *DailyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyRequest.Merge(m, src)
}
func (m *DailyRequest) XXX_Size() int {
	return xxx_messageInfo_DailyRequest.Size(m)
}
func (m *DailyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DailyRequest proto.InternalMessageInfo

func (m *DailyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DailyRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DailyRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DailyRequest) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *DailyRequest) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *DailyRequest) GetIsDelete() int32 {
	if m != nil {
		return m.IsDelete
	}
	return 0
}

func (m *DailyRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DailyRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *DailyRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DailyRequest) GetWorksId() string {
	if m != nil {
		return m.WorksId
	}
	return ""
}

func (m *DailyRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DailyRequest) GetUids() []int32 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *DailyRequest) GetWorksIds() []string {
	if m != nil {
		return m.WorksIds
	}
	return nil
}

type DailyResponse struct {
	Data                 []*DailyResponse_Notify `protobuf:"bytes,12,rep,name=data,proto3" json:"data,omitempty"`
	AffectedRows         int32                   `protobuf:"varint,11,opt,name=affectedRows,proto3" json:"affectedRows,omitempty"`
	Total                int32                   `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DailyResponse) Reset()         { *m = DailyResponse{} }
func (m *DailyResponse) String() string { return proto.CompactTextString(m) }
func (*DailyResponse) ProtoMessage()    {}
func (*DailyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{4}
}

func (m *DailyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyResponse.Unmarshal(m, b)
}
func (m *DailyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyResponse.Marshal(b, m, deterministic)
}
func (m *DailyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyResponse.Merge(m, src)
}
func (m *DailyResponse) XXX_Size() int {
	return xxx_messageInfo_DailyResponse.Size(m)
}
func (m *DailyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DailyResponse proto.InternalMessageInfo

func (m *DailyResponse) GetData() []*DailyResponse_Notify {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DailyResponse) GetAffectedRows() int32 {
	if m != nil {
		return m.AffectedRows
	}
	return 0
}

func (m *DailyResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type DailyResponse_Notify struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid                  string   `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	AccessDate           string   `protobuf:"bytes,3,opt,name=access_date,json=accessDate,proto3" json:"access_date,omitempty"`
	EventId              string   `protobuf:"bytes,13,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Pv                   int32    `protobuf:"varint,4,opt,name=pv,proto3" json:"pv,omitempty"`
	Uv                   int32    `protobuf:"varint,5,opt,name=uv,proto3" json:"uv,omitempty"`
	TailAd               int32    `protobuf:"varint,6,opt,name=tail_ad,json=tailAd,proto3" json:"tail_ad,omitempty"`
	Share                int32    `protobuf:"varint,7,opt,name=share,proto3" json:"share,omitempty"`
	TotalShare           int32    `protobuf:"varint,8,opt,name=total_share,json=totalShare,proto3" json:"total_share,omitempty"`
	TotalTailAd          int32    `protobuf:"varint,9,opt,name=total_tail_ad,json=totalTailAd,proto3" json:"total_tail_ad,omitempty"`
	TotalPv              int32    `protobuf:"varint,10,opt,name=total_pv,json=totalPv,proto3" json:"total_pv,omitempty"`
	TotalUv              int32    `protobuf:"varint,11,opt,name=total_uv,json=totalUv,proto3" json:"total_uv,omitempty"`
	Uid                  int32    `protobuf:"varint,12,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DailyResponse_Notify) Reset()         { *m = DailyResponse_Notify{} }
func (m *DailyResponse_Notify) String() string { return proto.CompactTextString(m) }
func (*DailyResponse_Notify) ProtoMessage()    {}
func (*DailyResponse_Notify) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de297fef73caeca, []int{4, 0}
}

func (m *DailyResponse_Notify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyResponse_Notify.Unmarshal(m, b)
}
func (m *DailyResponse_Notify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyResponse_Notify.Marshal(b, m, deterministic)
}
func (m *DailyResponse_Notify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyResponse_Notify.Merge(m, src)
}
func (m *DailyResponse_Notify) XXX_Size() int {
	return xxx_messageInfo_DailyResponse_Notify.Size(m)
}
func (m *DailyResponse_Notify) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyResponse_Notify.DiscardUnknown(m)
}

var xxx_messageInfo_DailyResponse_Notify proto.InternalMessageInfo

func (m *DailyResponse_Notify) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DailyResponse_Notify) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *DailyResponse_Notify) GetAccessDate() string {
	if m != nil {
		return m.AccessDate
	}
	return ""
}

func (m *DailyResponse_Notify) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *DailyResponse_Notify) GetPv() int32 {
	if m != nil {
		return m.Pv
	}
	return 0
}

func (m *DailyResponse_Notify) GetUv() int32 {
	if m != nil {
		return m.Uv
	}
	return 0
}

func (m *DailyResponse_Notify) GetTailAd() int32 {
	if m != nil {
		return m.TailAd
	}
	return 0
}

func (m *DailyResponse_Notify) GetShare() int32 {
	if m != nil {
		return m.Share
	}
	return 0
}

func (m *DailyResponse_Notify) GetTotalShare() int32 {
	if m != nil {
		return m.TotalShare
	}
	return 0
}

func (m *DailyResponse_Notify) GetTotalTailAd() int32 {
	if m != nil {
		return m.TotalTailAd
	}
	return 0
}

func (m *DailyResponse_Notify) GetTotalPv() int32 {
	if m != nil {
		return m.TotalPv
	}
	return 0
}

func (m *DailyResponse_Notify) GetTotalUv() int32 {
	if m != nil {
		return m.TotalUv
	}
	return 0
}

func (m *DailyResponse_Notify) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Response_Notify)(nil), "Response.Notify")
	proto.RegisterType((*WorksResponse)(nil), "WorksResponse")
	proto.RegisterType((*WorksResponse_Notify)(nil), "WorksResponse.Notify")
	proto.RegisterType((*DailyRequest)(nil), "DailyRequest")
	proto.RegisterType((*DailyResponse)(nil), "DailyResponse")
	proto.RegisterType((*DailyResponse_Notify)(nil), "DailyResponse.Notify")
}

func init() { proto.RegisterFile("worksAudit.proto", fileDescriptor_2de297fef73caeca) }

var fileDescriptor_2de297fef73caeca = []byte{
	// 1040 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0x4e, 0x64, 0xcb, 0x3f, 0x1d, 0xdb, 0xeb, 0x9d, 0xda, 0x9f, 0x89, 0x97, 0x84, 0x94, 0x2a,
	0x40, 0x96, 0x2a, 0x7c, 0x08, 0xbc, 0x40, 0xd8, 0x40, 0xe1, 0x2a, 0x8a, 0x4a, 0xc9, 0x09, 0x39,
	0xba, 0xb4, 0x56, 0x27, 0x9e, 0xc2, 0x96, 0xb4, 0xd2, 0x48, 0xc1, 0x39, 0xf1, 0x5a, 0x5c, 0x79,
	0x01, 0xde, 0x02, 0x4e, 0x54, 0xf1, 0x08, 0x54, 0xf7, 0x8c, 0x64, 0x8b, 0x2d, 0xdf, 0x96, 0x1b,
	0x27, 0x4f, 0x7f, 0xfd, 0x4d, 0x7b, 0xba, 0xbf, 0xee, 0x19, 0xc1, 0xf0, 0x21, 0x4e, 0x7f, 0xca,
	0x2e, 0xf2, 0x50, 0xe9, 0x71, 0x92, 0xc6, 0x3a, 0xf6, 0xfe, 0x70, 0xa0, 0xed, 0xe3, 0xbb, 0x1c,
	0x33, 0x2d, 0x04, 0x34, 0xa3, 0x60, 0x85, 0x72, 0xff, 0x64, 0xff, 0xac, 0xeb, 0xf3, 0x9a, 0x30,
	0xbd, 0x4e, 0x50, 0x3a, 0x06, 0xa3, 0xb5, 0x78, 0x06, 0x6e, 0x11, 0x2c, 0x73, 0x94, 0x0d, 0x06,
	0x8d, 0x21, 0x3e, 0x86, 0x83, 0x79, 0x8a, 0x81, 0xc6, 0x99, 0x56, 0x2b, 0x94, 0x4d, 0xf6, 0x81,
	0x81, 0xae, 0xd5, 0x8a, 0x09, 0x79, 0x12, 0x56, 0x04, 0xd7, 0x10, 0x0c, 0xc4, 0x84, 0x57, 0xd0,
	0x55, 0xd9, 0x2c, 0xc4, 0x25, 0x6a, 0x94, 0xad, 0x93, 0xfd, 0x33, 0xd7, 0xef, 0xa8, 0xec, 0x92,
	0x6d, 0x3a, 0x48, 0x12, 0xdc, 0xa3, 0x6c, 0x33, 0xce, 0x6b, 0xda, 0x40, 0xbf, 0xb3, 0x4c, 0x3d,
	0xa2, 0xec, 0x98, 0x0d, 0x04, 0x4c, 0xd5, 0x23, 0x8a, 0x01, 0x38, 0x2a, 0x94, 0x5d, 0x46, 0x1d,
	0x15, 0x8a, 0x43, 0xe8, 0x70, 0xf6, 0x33, 0x15, 0x4a, 0xe0, 0xff, 0x6e, 0xb3, 0x3d, 0x09, 0xc5,
	0x0b, 0x68, 0x65, 0x3a, 0xd0, 0x79, 0x26, 0x7b, 0x4c, 0xb7, 0x96, 0x18, 0x42, 0x23, 0x57, 0xa1,
	0x3c, 0x60, 0x90, 0x96, 0xc4, 0x4c, 0x31, 0xc8, 0xe2, 0x48, 0xf6, 0x39, 0x84, 0xb5, 0xc4, 0x09,
	0x1c, 0x70, 0x55, 0xa7, 0x26, 0xcc, 0xe0, 0xa4, 0x71, 0xe6, 0xfa, 0xdb, 0x90, 0xf7, 0x97, 0x03,
	0x1d, 0x1f, 0xb3, 0x24, 0x8e, 0x32, 0x14, 0xa7, 0xd0, 0x0c, 0x03, 0x1d, 0xc8, 0xde, 0x49, 0xe3,
	0xec, 0xe0, 0x7c, 0x38, 0x2e, 0x1d, 0xe3, 0x1f, 0x62, 0xad, 0xee, 0xd6, 0x3e, 0x7b, 0x85, 0x07,
	0xbd, 0xe0, 0xee, 0x0e, 0xe7, 0x1a, 0x43, 0x3f, 0x7e, 0xc8, 0xec, 0x39, 0x6a, 0x18, 0x69, 0xa1,
	0x63, 0x1d, 0x2c, 0xf9, 0x3c, 0xae, 0x6f, 0x8c, 0xd1, 0x2f, 0x0e, 0xb4, 0x4c, 0xa8, 0xf7, 0xca,
	0x60, 0x73, 0xda, 0xdf, 0xe4, 0x74, 0x08, 0x1d, 0x2c, 0x30, 0xd2, 0x54, 0x18, 0x23, 0x73, 0x9b,
	0xed, 0x5a, 0x61, 0xda, 0xb5, 0xc2, 0x94, 0x5d, 0xd1, 0xd8, 0xea, 0x8a, 0xff, 0x58, 0x7f, 0x4a,
	0x54, 0xe9, 0x25, 0x5a, 0xed, 0x8c, 0x21, 0x24, 0xb4, 0xe7, 0x71, 0xa4, 0x31, 0xd2, 0x5c, 0x9d,
	0xae, 0x5f, 0x9a, 0xde, 0x9f, 0x6d, 0xe8, 0xdf, 0x92, 0xbe, 0x55, 0xd1, 0x5f, 0xd7, 0x8a, 0xfe,
	0x7c, 0x5c, 0xf3, 0x7e, 0xa8, 0xca, 0xff, 0xdd, 0xfa, 0x57, 0xe5, 0xf7, 0xab, 0xca, 0x4b, 0x28,
	0x1b, 0xae, 0x2c, 0x73, 0xd9, 0x7f, 0x56, 0x93, 0xc6, 0x46, 0x93, 0x2a, 0xdb, 0xe6, 0x8e, 0x6c,
	0xdd, 0x5a, 0xb6, 0xc6, 0xb3, 0x5a, 0x91, 0xa7, 0x55, 0x7a, 0xd8, 0xe4, 0x48, 0x8b, 0x7c, 0xf5,
	0x96, 0x15, 0xa4, 0x48, 0x64, 0x88, 0x11, 0x74, 0xee, 0x54, 0x9a, 0xe9, 0xc9, 0xea, 0x9e, 0x07,
	0xa7, 0xeb, 0x57, 0x36, 0xc5, 0x2a, 0x30, 0xcd, 0x54, 0x1c, 0x71, 0xdb, 0x74, 0xfd, 0xd2, 0x14,
	0xa7, 0xd0, 0xc7, 0x50, 0xe9, 0x38, 0xfd, 0xd1, 0xfa, 0x8d, 0x16, 0x75, 0x90, 0x74, 0xd6, 0xb8,
	0x4a, 0x96, 0xa4, 0xb4, 0x9d, 0x9e, 0xae, 0x0f, 0x25, 0x34, 0x09, 0xc5, 0xa7, 0xf0, 0x24, 0x09,
	0x52, 0xea, 0xb8, 0xaa, 0xef, 0x7a, 0x26, 0x90, 0x81, 0xbf, 0xb1, 0xdd, 0x37, 0x82, 0x8e, 0x0a,
	0x31, 0xd2, 0x4a, 0xaf, 0x6d, 0x91, 0x2b, 0x9b, 0x7c, 0xf3, 0x40, 0xe3, 0x7d, 0x9c, 0xae, 0xe5,
	0xc0, 0xf8, 0x4a, 0x5b, 0xbc, 0x84, 0x36, 0xdd, 0x08, 0x14, 0xf7, 0x89, 0x6d, 0x5b, 0xf5, 0x48,
	0x7f, 0xfc, 0x05, 0x88, 0x0c, 0xe7, 0x71, 0x14, 0x06, 0xe9, 0x7a, 0x56, 0x6d, 0x1f, 0x32, 0xe7,
	0x69, 0xe5, 0x79, 0x53, 0xc6, 0x39, 0x02, 0xe0, 0xeb, 0xe5, 0x41, 0x85, 0x7a, 0x21, 0x9f, 0x31,
	0x8d, 0x2f, 0x9c, 0x5b, 0x02, 0x28, 0x4f, 0x76, 0x2f, 0x50, 0xdd, 0x2f, 0xb4, 0x7c, 0xca, 0x7e,
	0xde, 0xf1, 0x1d, 0x23, 0xe2, 0x13, 0x18, 0xc4, 0x7a, 0x81, 0xe9, 0x2c, 0x49, 0xe3, 0x04, 0x53,
	0xbd, 0x96, 0xc2, 0xa4, 0xc9, 0xe8, 0x95, 0x05, 0xa9, 0xd9, 0xa6, 0xb9, 0xd2, 0x78, 0x91, 0x24,
	0xd7, 0x34, 0x54, 0xcf, 0x4d, 0xb3, 0x6d, 0x63, 0xc4, 0x79, 0xc3, 0x93, 0x74, 0x89, 0x85, 0x9a,
	0xa3, 0x7c, 0x61, 0x38, 0xdb, 0x18, 0x71, 0x6e, 0x78, 0x98, 0x2c, 0xe7, 0xa5, 0xe1, 0x6c, 0x63,
	0xe2, 0x18, 0xc0, 0xec, 0x99, 0x2a, 0x8d, 0x52, 0x9a, 0x23, 0x6f, 0x90, 0x8d, 0x9f, 0x06, 0x52,
	0x1e, 0x6e, 0xfb, 0x79, 0x44, 0x8f, 0x01, 0x6e, 0xaa, 0x81, 0x95, 0x23, 0xe3, 0xdf, 0x20, 0x24,
	0xcb, 0xc4, 0x4e, 0xac, 0x7c, 0x65, 0xfa, 0xaa, 0xb4, 0xc5, 0xe7, 0x30, 0x9c, 0x64, 0x5f, 0xc7,
	0xf9, 0xfd, 0x42, 0x5f, 0xdb, 0x66, 0x90, 0x1f, 0x31, 0xe7, 0x3d, 0xdc, 0x70, 0x6f, 0x32, 0x0c,
	0xbf, 0x8f, 0xe7, 0xc1, 0xf2, 0xdb, 0x38, 0xd2, 0x99, 0x3c, 0x2a, 0xb9, 0x75, 0xdc, 0xfb, 0xdd,
	0x81, 0xde, 0x65, 0xa0, 0x96, 0xeb, 0xff, 0xdf, 0xb1, 0x1d, 0xef, 0x98, 0x80, 0x66, 0xae, 0x42,
	0xba, 0xc6, 0xe8, 0x59, 0xe2, 0x35, 0x29, 0x75, 0x6b, 0xb6, 0x99, 0xe7, 0xaa, 0xeb, 0x57, 0xb6,
	0xf7, 0x6b, 0x03, 0xfa, 0xb6, 0xa2, 0x3b, 0xee, 0xce, 0x9a, 0xf7, 0x43, 0xdd, 0x9d, 0xbf, 0x39,
	0x3b, 0xef, 0xce, 0x21, 0x34, 0x92, 0xea, 0x79, 0xa2, 0x25, 0xa9, 0x10, 0xcc, 0xe7, 0x98, 0x65,
	0x33, 0xaa, 0xbb, 0x95, 0x10, 0x0c, 0x74, 0x49, 0x2d, 0xb4, 0xfd, 0xac, 0xf5, 0xeb, 0xcf, 0xda,
	0x00, 0x9c, 0xa4, 0x60, 0x65, 0x5d, 0xdf, 0x49, 0x0a, 0xb2, 0xf3, 0x82, 0x85, 0x74, 0x7d, 0x27,
	0x2f, 0xe8, 0x02, 0xd1, 0x81, 0x5a, 0xce, 0x82, 0xd0, 0xca, 0xd7, 0x22, 0xf3, 0x82, 0xaf, 0xe5,
	0x6c, 0x11, 0xa4, 0xa5, 0x7a, 0xc6, 0xe0, 0x0b, 0x8f, 0x12, 0x98, 0x19, 0x9f, 0x11, 0x10, 0x18,
	0x9a, 0x32, 0xc1, 0x83, 0xbe, 0x21, 0x94, 0x51, 0x8d, 0x9a, 0x66, 0xd7, 0xb5, 0x09, 0x7d, 0x08,
	0x1d, 0xc3, 0x49, 0x0a, 0x96, 0xd5, 0xf5, 0xdb, 0x6c, 0x5f, 0x15, 0x1b, 0x57, 0x5e, 0xd8, 0x6a,
	0x1a, 0xd7, 0x4d, 0x51, 0xbe, 0x1c, 0xbd, 0xea, 0xe5, 0x38, 0x7f, 0x07, 0x70, 0x5b, 0x7d, 0xe4,
	0x89, 0x23, 0x68, 0x4e, 0x17, 0xf1, 0x83, 0xe8, 0x8c, 0xed, 0x70, 0x8c, 0xba, 0xd5, 0xc7, 0x86,
	0xb7, 0x27, 0x8e, 0xc1, 0x9d, 0x44, 0x21, 0xfe, 0xbc, 0xcb, 0x7f, 0x0a, 0x07, 0x1c, 0xcc, 0x4c,
	0xf8, 0x0e, 0xd6, 0xf9, 0x39, 0xb8, 0xcc, 0x12, 0xaf, 0x2d, 0xfd, 0x12, 0x29, 0xd1, 0x2d, 0xfa,
	0xa0, 0xfe, 0xd8, 0x7a, 0x7b, 0xe7, 0x5f, 0x41, 0x97, 0x7b, 0xe8, 0xaa, 0xb8, 0x29, 0xc4, 0x67,
	0xf6, 0x94, 0xfd, 0xf1, 0xf6, 0x1c, 0x8f, 0x06, 0xf5, 0x36, 0xf3, 0xf6, 0xde, 0xb6, 0xf8, 0xa3,
	0xf5, 0xcb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x6f, 0xf2, 0xcf, 0xc8, 0x0a, 0x00, 0x00,
}
