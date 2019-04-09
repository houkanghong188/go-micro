package model

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-micro/cmd/font/proto/trial"
	"go-micro/tool"
	"time"
)

type AuthCompanyModel struct {
	gorm.Model
	Id         int32
	Uid        int32
	Name       string
	Contract   string
	Phone      string
	IsDelete   int32
	CreateTime string
	UpdateTime string
}

func (m *AuthCompanyModel) TableName() string {
	return "platv5_font_auth_company"
}

func NewAuthTrialModel() *AuthCompanyModel {
	return &AuthCompanyModel{}
}

func (m *AuthCompanyModel) Show(ctx context.Context, in *trial.ShowRequest, out *trial.ShowResponse) error {
	conn := tool.GetSlavelConn()
	conn.Table(m.TableName()).First(out)

	return nil
}

func (m *AuthCompanyModel) Create(ctx context.Context, in *trial.CreateRequest, out *trial.CreateResponse) error {

	conn := tool.GetMasterConn()

	authCompany := AuthCompanyModel{
		Uid:        in.Uid,
		Name:       in.Name,
		Contract:   in.Contract,
		Phone:      in.Phone,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	conn.Create(&authCompany)

	return nil
}
