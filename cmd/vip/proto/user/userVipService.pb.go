// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userVipService.proto

package userVip

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

type UserVipInfoRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserVipInfoRequest) Reset()         { *m = UserVipInfoRequest{} }
func (m *UserVipInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserVipInfoRequest) ProtoMessage()    {}
func (*UserVipInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_769955709423b8a5, []int{0}
}

func (m *UserVipInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserVipInfoRequest.Unmarshal(m, b)
}
func (m *UserVipInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserVipInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserVipInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserVipInfoRequest.Merge(m, src)
}
func (m *UserVipInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserVipInfoRequest.Size(m)
}
func (m *UserVipInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserVipInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserVipInfoRequest proto.InternalMessageInfo

func (m *UserVipInfoRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserVipInfoResponse struct {
	Vips                 []*VipInfo `protobuf:"bytes,1,rep,name=vips,proto3" json:"vips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserVipInfoResponse) Reset()         { *m = UserVipInfoResponse{} }
func (m *UserVipInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserVipInfoResponse) ProtoMessage()    {}
func (*UserVipInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_769955709423b8a5, []int{1}
}

func (m *UserVipInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserVipInfoResponse.Unmarshal(m, b)
}
func (m *UserVipInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserVipInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserVipInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserVipInfoResponse.Merge(m, src)
}
func (m *UserVipInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserVipInfoResponse.Size(m)
}
func (m *UserVipInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserVipInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserVipInfoResponse proto.InternalMessageInfo

func (m *UserVipInfoResponse) GetVips() []*VipInfo {
	if m != nil {
		return m.Vips
	}
	return nil
}

// 可以独立到 customer_vips 中去，独立出来作为会员基本信息
type VipInfo struct {
	Id                    int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias                 string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Name                  string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price                 string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Corner                string   `protobuf:"bytes,5,opt,name=corner,proto3" json:"corner,omitempty"`
	CornetText            string   `protobuf:"bytes,6,opt,name=cornetText,proto3" json:"cornetText,omitempty"`
	EnableMaka            int32    `protobuf:"varint,7,opt,name=enableMaka,proto3" json:"enableMaka,omitempty"`
	EnableDanye           int32    `protobuf:"varint,8,opt,name=enableDanye,proto3" json:"enableDanye,omitempty"`
	EnablePoster          int32    `protobuf:"varint,9,opt,name=enablePoster,proto3" json:"enablePoster,omitempty"`
	EnableVideo           int32    `protobuf:"varint,10,opt,name=enableVideo,proto3" json:"enableVideo,omitempty"`
	EnableFont            int32    `protobuf:"varint,11,opt,name=enableFont,proto3" json:"enableFont,omitempty"`
	EnableLogo            int32    `protobuf:"varint,12,opt,name=enableLogo,proto3" json:"enableLogo,omitempty"`
	EnablePosterLogo      int32    `protobuf:"varint,13,opt,name=enablePosterLogo,proto3" json:"enablePosterLogo,omitempty"`
	EnableMaterial        int32    `protobuf:"varint,14,opt,name=enableMaterial,proto3" json:"enableMaterial,omitempty"`
	EnableContract        int32    `protobuf:"varint,15,opt,name=enableContract,proto3" json:"enableContract,omitempty"`
	EnableBottomMenu      int32    `protobuf:"varint,16,opt,name=enableBottomMenu,proto3" json:"enableBottomMenu,omitempty"`
	EnableLottery         int32    `protobuf:"varint,17,opt,name=enableLottery,proto3" json:"enableLottery,omitempty"`
	EnableRelay           int32    `protobuf:"varint,18,opt,name=enableRelay,proto3" json:"enableRelay,omitempty"`
	EnableFormLamp        int32    `protobuf:"varint,19,opt,name=enableFormLamp,proto3" json:"enableFormLamp,omitempty"`
	EnableEndPageStyle    int32    `protobuf:"varint,20,opt,name=enableEndPageStyle,proto3" json:"enableEndPageStyle,omitempty"`
	EnableEndPageReadNum  int32    `protobuf:"varint,21,opt,name=enableEndPageReadNum,proto3" json:"enableEndPageReadNum,omitempty"`
	EnablePublicBall      int32    `protobuf:"varint,22,opt,name=enablePublicBall,proto3" json:"enablePublicBall,omitempty"`
	EnableOpenPassword    int32    `protobuf:"varint,23,opt,name=enableOpenPassword,proto3" json:"enableOpenPassword,omitempty"`
	EnablePvuvData        int32    `protobuf:"varint,24,opt,name=enablePvuvData,proto3" json:"enablePvuvData,omitempty"`
	EnableSharedata       int32    `protobuf:"varint,25,opt,name=enableSharedata,proto3" json:"enableSharedata,omitempty"`
	EnableReadTime        int32    `protobuf:"varint,26,opt,name=enableReadTime,proto3" json:"enableReadTime,omitempty"`
	EnableReadTimePerPage int32    `protobuf:"varint,27,opt,name=enableReadTimePerPage,proto3" json:"enableReadTimePerPage,omitempty"`
	EnablePagesReadTime   int32    `protobuf:"varint,28,opt,name=enablePagesReadTime,proto3" json:"enablePagesReadTime,omitempty"`
	EnablePagesPvuvData   int32    `protobuf:"varint,29,opt,name=enablePagesPvuvData,proto3" json:"enablePagesPvuvData,omitempty"`
	EnableVisitorSource   int32    `protobuf:"varint,30,opt,name=enableVisitorSource,proto3" json:"enableVisitorSource,omitempty"`
	EnableVisitorDevice   int32    `protobuf:"varint,31,opt,name=enableVisitorDevice,proto3" json:"enableVisitorDevice,omitempty"`
	EnableVisitorLocation int32    `protobuf:"varint,32,opt,name=enableVisitorLocation,proto3" json:"enableVisitorLocation,omitempty"`
	EnableVisitorBrand    int32    `protobuf:"varint,33,opt,name=enableVisitorBrand,proto3" json:"enableVisitorBrand,omitempty"`
	Sort                  int32    `protobuf:"varint,34,opt,name=sort,proto3" json:"sort,omitempty"`
	Status                int32    `protobuf:"varint,35,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *VipInfo) Reset()         { *m = VipInfo{} }
func (m *VipInfo) String() string { return proto.CompactTextString(m) }
func (*VipInfo) ProtoMessage()    {}
func (*VipInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_769955709423b8a5, []int{2}
}

func (m *VipInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VipInfo.Unmarshal(m, b)
}
func (m *VipInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VipInfo.Marshal(b, m, deterministic)
}
func (m *VipInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VipInfo.Merge(m, src)
}
func (m *VipInfo) XXX_Size() int {
	return xxx_messageInfo_VipInfo.Size(m)
}
func (m *VipInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VipInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VipInfo proto.InternalMessageInfo

func (m *VipInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VipInfo) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *VipInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VipInfo) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *VipInfo) GetCorner() string {
	if m != nil {
		return m.Corner
	}
	return ""
}

func (m *VipInfo) GetCornetText() string {
	if m != nil {
		return m.CornetText
	}
	return ""
}

func (m *VipInfo) GetEnableMaka() int32 {
	if m != nil {
		return m.EnableMaka
	}
	return 0
}

func (m *VipInfo) GetEnableDanye() int32 {
	if m != nil {
		return m.EnableDanye
	}
	return 0
}

func (m *VipInfo) GetEnablePoster() int32 {
	if m != nil {
		return m.EnablePoster
	}
	return 0
}

func (m *VipInfo) GetEnableVideo() int32 {
	if m != nil {
		return m.EnableVideo
	}
	return 0
}

func (m *VipInfo) GetEnableFont() int32 {
	if m != nil {
		return m.EnableFont
	}
	return 0
}

func (m *VipInfo) GetEnableLogo() int32 {
	if m != nil {
		return m.EnableLogo
	}
	return 0
}

func (m *VipInfo) GetEnablePosterLogo() int32 {
	if m != nil {
		return m.EnablePosterLogo
	}
	return 0
}

func (m *VipInfo) GetEnableMaterial() int32 {
	if m != nil {
		return m.EnableMaterial
	}
	return 0
}

func (m *VipInfo) GetEnableContract() int32 {
	if m != nil {
		return m.EnableContract
	}
	return 0
}

func (m *VipInfo) GetEnableBottomMenu() int32 {
	if m != nil {
		return m.EnableBottomMenu
	}
	return 0
}

func (m *VipInfo) GetEnableLottery() int32 {
	if m != nil {
		return m.EnableLottery
	}
	return 0
}

func (m *VipInfo) GetEnableRelay() int32 {
	if m != nil {
		return m.EnableRelay
	}
	return 0
}

func (m *VipInfo) GetEnableFormLamp() int32 {
	if m != nil {
		return m.EnableFormLamp
	}
	return 0
}

func (m *VipInfo) GetEnableEndPageStyle() int32 {
	if m != nil {
		return m.EnableEndPageStyle
	}
	return 0
}

func (m *VipInfo) GetEnableEndPageReadNum() int32 {
	if m != nil {
		return m.EnableEndPageReadNum
	}
	return 0
}

func (m *VipInfo) GetEnablePublicBall() int32 {
	if m != nil {
		return m.EnablePublicBall
	}
	return 0
}

func (m *VipInfo) GetEnableOpenPassword() int32 {
	if m != nil {
		return m.EnableOpenPassword
	}
	return 0
}

func (m *VipInfo) GetEnablePvuvData() int32 {
	if m != nil {
		return m.EnablePvuvData
	}
	return 0
}

func (m *VipInfo) GetEnableSharedata() int32 {
	if m != nil {
		return m.EnableSharedata
	}
	return 0
}

func (m *VipInfo) GetEnableReadTime() int32 {
	if m != nil {
		return m.EnableReadTime
	}
	return 0
}

func (m *VipInfo) GetEnableReadTimePerPage() int32 {
	if m != nil {
		return m.EnableReadTimePerPage
	}
	return 0
}

func (m *VipInfo) GetEnablePagesReadTime() int32 {
	if m != nil {
		return m.EnablePagesReadTime
	}
	return 0
}

func (m *VipInfo) GetEnablePagesPvuvData() int32 {
	if m != nil {
		return m.EnablePagesPvuvData
	}
	return 0
}

func (m *VipInfo) GetEnableVisitorSource() int32 {
	if m != nil {
		return m.EnableVisitorSource
	}
	return 0
}

func (m *VipInfo) GetEnableVisitorDevice() int32 {
	if m != nil {
		return m.EnableVisitorDevice
	}
	return 0
}

func (m *VipInfo) GetEnableVisitorLocation() int32 {
	if m != nil {
		return m.EnableVisitorLocation
	}
	return 0
}

func (m *VipInfo) GetEnableVisitorBrand() int32 {
	if m != nil {
		return m.EnableVisitorBrand
	}
	return 0
}

func (m *VipInfo) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *VipInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*UserVipInfoRequest)(nil), "userVip.userVipInfoRequest")
	proto.RegisterType((*UserVipInfoResponse)(nil), "userVip.userVipInfoResponse")
	proto.RegisterType((*VipInfo)(nil), "userVip.vipInfo")
}

func init() { proto.RegisterFile("userVipService.proto", fileDescriptor_769955709423b8a5) }

var fileDescriptor_769955709423b8a5 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0xcf, 0x6f, 0x13, 0x3b,
	0x10, 0xc7, 0x95, 0xfe, 0x7c, 0x9d, 0xb4, 0x69, 0x9e, 0x9b, 0xf6, 0xcd, 0x6b, 0x4b, 0x09, 0xa1,
	0xaa, 0x22, 0x0e, 0x11, 0x2a, 0xdc, 0xb8, 0x95, 0x52, 0x81, 0xd4, 0x42, 0x94, 0x54, 0x3d, 0x70,
	0x73, 0x77, 0x87, 0x62, 0xb1, 0xb1, 0x17, 0xdb, 0x1b, 0xc8, 0xdf, 0xc6, 0x3f, 0x87, 0xd6, 0xbb,
	0x59, 0x7b, 0x93, 0xdc, 0x3c, 0x9f, 0xef, 0x77, 0xc6, 0x33, 0x5e, 0x69, 0x16, 0x3a, 0x99, 0x21,
	0xfd, 0x20, 0xd2, 0x31, 0xe9, 0xa9, 0x88, 0x68, 0x90, 0x6a, 0x65, 0x15, 0xdb, 0x2e, 0x69, 0xef,
	0x02, 0x58, 0x79, 0xfc, 0x24, 0xbf, 0xa9, 0x11, 0xfd, 0xcc, 0xc8, 0x58, 0xd6, 0x86, 0xf5, 0x4c,
	0xc4, 0xd8, 0xe8, 0x36, 0xfa, 0x9b, 0xa3, 0xfc, 0xd8, 0x7b, 0x07, 0x07, 0x35, 0x9f, 0x49, 0x95,
	0x34, 0xc4, 0xce, 0x61, 0x63, 0x2a, 0x52, 0x83, 0x8d, 0xee, 0x7a, 0xbf, 0x79, 0xd9, 0x1e, 0x94,
	0x9e, 0xc1, 0xb4, 0xf4, 0x39, 0xb5, 0xf7, 0x07, 0x60, 0xbb, 0x24, 0xac, 0x05, 0x6b, 0x55, 0xe5,
	0x35, 0x11, 0xb3, 0x0e, 0x6c, 0xf2, 0x44, 0x70, 0x83, 0x6b, 0xdd, 0x46, 0x7f, 0x67, 0x54, 0x04,
	0x8c, 0xc1, 0x86, 0xe4, 0x13, 0xc2, 0x75, 0x07, 0xdd, 0x39, 0x77, 0xa6, 0x5a, 0x44, 0x84, 0x1b,
	0x85, 0xd3, 0x05, 0xec, 0x08, 0xb6, 0x22, 0xa5, 0x25, 0x69, 0xdc, 0x74, 0xb8, 0x8c, 0xd8, 0x19,
	0x80, 0x3b, 0xd9, 0x7b, 0xfa, 0x6d, 0x71, 0xcb, 0x69, 0x01, 0xc9, 0x75, 0x92, 0xfc, 0x31, 0xa1,
	0x3b, 0xfe, 0x83, 0xe3, 0xb6, 0xeb, 0x27, 0x20, 0xac, 0x0b, 0xcd, 0x22, 0xba, 0xe6, 0x72, 0x46,
	0xf8, 0x8f, 0x33, 0x84, 0x88, 0xf5, 0x60, 0xb7, 0x08, 0x87, 0xca, 0x58, 0xd2, 0xb8, 0xe3, 0x2c,
	0x35, 0xe6, 0xab, 0x3c, 0x88, 0x98, 0x14, 0x42, 0x58, 0xc5, 0x21, 0xdf, 0xc7, 0x8d, 0x92, 0x16,
	0x9b, 0x61, 0x1f, 0x39, 0xf1, 0xfa, 0xad, 0x7a, 0x52, 0xb8, 0x1b, 0xea, 0x39, 0x61, 0xaf, 0xa0,
	0x1d, 0xde, 0xe8, 0x5c, 0x7b, 0xce, 0xb5, 0xc4, 0xd9, 0x05, 0xb4, 0xe6, 0x13, 0x5a, 0xd2, 0x82,
	0x27, 0xd8, 0x72, 0xce, 0x05, 0xea, 0x7d, 0xef, 0x95, 0xb4, 0x9a, 0x47, 0x16, 0xf7, 0x43, 0xdf,
	0x9c, 0xfa, 0xbb, 0xaf, 0x94, 0xb5, 0x6a, 0x72, 0x47, 0x32, 0xc3, 0x76, 0x78, 0xb7, 0xe7, 0xec,
	0x1c, 0xf6, 0xe6, 0x5d, 0x5b, 0x4b, 0x7a, 0x86, 0xff, 0x3a, 0x63, 0x1d, 0xfa, 0xf7, 0x1a, 0x51,
	0xc2, 0x67, 0xc8, 0xc2, 0xf7, 0x72, 0xc8, 0xf7, 0x76, 0xa3, 0xf4, 0xe4, 0x96, 0x4f, 0x52, 0x3c,
	0x08, 0x7b, 0x9b, 0x53, 0x36, 0x00, 0x56, 0x90, 0x0f, 0x32, 0x1e, 0xf2, 0x27, 0x1a, 0xdb, 0x59,
	0x42, 0xd8, 0x71, 0xde, 0x15, 0x0a, 0xbb, 0x84, 0x4e, 0x8d, 0x8e, 0x88, 0xc7, 0x9f, 0xb3, 0x09,
	0x1e, 0xba, 0x8c, 0x95, 0x5a, 0xf0, 0xf6, 0xd9, 0x63, 0x22, 0xa2, 0x2b, 0x9e, 0x24, 0x78, 0x54,
	0x7b, 0xfb, 0x8a, 0xfb, 0x7e, 0xbe, 0xa4, 0x24, 0x87, 0xdc, 0x98, 0x5f, 0x4a, 0xc7, 0xf8, 0x5f,
	0xd8, 0x4f, 0xa8, 0xf8, 0x39, 0x87, 0xd3, 0x6c, 0x7a, 0xcd, 0x2d, 0x47, 0x0c, 0xe7, 0x9c, 0x53,
	0xd6, 0x87, 0xfd, 0x82, 0x8c, 0xbf, 0x73, 0x4d, 0x71, 0x6e, 0xfc, 0xdf, 0x19, 0x17, 0xb1, 0xaf,
	0x98, 0xb7, 0x7f, 0x2f, 0x26, 0x84, 0xc7, 0x61, 0xc5, 0x39, 0x65, 0x6f, 0xe1, 0xb0, 0x4e, 0x86,
	0xa4, 0xf3, 0xa9, 0xf1, 0xc4, 0xd9, 0x57, 0x8b, 0xec, 0x35, 0x1c, 0x94, 0x9d, 0xf1, 0x27, 0x32,
	0xd5, 0x15, 0xa7, 0x2e, 0x67, 0x95, 0xb4, 0x90, 0x51, 0x8d, 0xf9, 0x6c, 0x29, 0xa3, 0x9a, 0xb5,
	0xca, 0x78, 0x10, 0x46, 0x58, 0xa5, 0xc7, 0x2a, 0xd3, 0x11, 0xe1, 0x59, 0x98, 0x51, 0x93, 0x96,
	0x32, 0xae, 0x29, 0x5f, 0x82, 0xf8, 0x7c, 0x45, 0x46, 0x21, 0xf9, 0xe9, 0x4b, 0x7c, 0xab, 0x22,
	0x6e, 0x85, 0x92, 0xd8, 0x0d, 0xa7, 0x5f, 0x10, 0xfd, 0xd7, 0x2d, 0x85, 0x2b, 0xcd, 0x65, 0x8c,
	0x2f, 0xc2, 0xaf, 0x1b, 0x2a, 0xf9, 0x7e, 0x33, 0x4a, 0x5b, 0xec, 0x39, 0x87, 0x3b, 0xe7, 0x9b,
	0xcc, 0x58, 0x6e, 0x33, 0x83, 0x2f, 0x1d, 0x2d, 0xa3, 0xcb, 0xaf, 0xd0, 0xaa, 0xef, 0x70, 0xf6,
	0x11, 0x9a, 0xc1, 0x32, 0x66, 0x27, 0xd5, 0xda, 0x5d, 0x5e, 0xe5, 0xc7, 0xa7, 0xab, 0xc5, 0x62,
	0x7f, 0x3f, 0x6e, 0xb9, 0xdf, 0xc1, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x1b, 0xe5,
	0x45, 0x26, 0x06, 0x00, 0x00,
}
