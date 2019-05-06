package models

import (
	"context"
	"go-micro/cmd/vip/proto/user"
	"go-micro/tool"
)

type CustomerVipsModel struct {
}

func NewCustomerVipsModel() *CustomerVipsModel {
	return &CustomerVipsModel{}
}

func (m *CustomerVipsModel) TableName() string {
	return "platv4_customer_vips"
}

func (m *CustomerVipsModel) UserVipInfo(ctx context.Context, in *userVip.UserVipInfoRequest, out *userVip.UserVipInfoResponse) error {

	return nil
}
