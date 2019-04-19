package models

import (
	"context"
	"go-micro/cmd/vip/proto/user"
	"go-micro/tool"
)

const (
	STATUS_EXPIRED = iota
	STATUS_NORMAL
)

type UserVipModel struct {
}

func NewUserVipModel() *UserVipModel {
	return &UserVipModel{}
}

func (m *UserVipModel) TableName() string {
	return "platv4_user_to_customer_vip"
}

func (m *UserVipModel) UserVipInfo(ctx context.Context, in *userVip.UserVipInfoRequest, out *userVip.UserVipInfoResponse) error {

	conn := tool.GetSlavelConn()
	var customerVipIds []string
	conn.Table(m.TableName()).Select("customer_vip_id").Where("uid in (?) and status = ?", in.Uid, STATUS_NORMAL).Find(&customerVipIds)

	conn = tool.GetSlavelConn()
	customerVipModel := NewCustomerVipsModel()
	conn.Table(customerVipModel.TableName()).Where("id in (?)", customerVipIds).Find(&out.Vips)

	// api/v5/micro/$1
	return nil
}
