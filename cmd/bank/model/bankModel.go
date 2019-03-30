package model

import (
	"go-micro/cmd/bank/proto"
	"go-micro/tool"
	"errors"
	"context"
	"fmt"
)

type BankModel struct {
	Id               int32
	SiteId           int32
	IsDelete         int32
	Account          string
	Card             string
	Name             string
	Tel              string
	Type             int8
	Status           int32
	CreateTime       string
	UpdateTime       string
	FinancialContact string
}

func NewBankModel() *BankModel  {
	return &BankModel{}
}

func (m *BankModel) TableName() string  {
	return "platv5_website_bank"
}

func (m *BankModel) Show(ctx context.Context, req  *bank.Request, rsp *bank.Response) error  {
	if req.Id > 0 {
		tool.Mysql.First(m, "id = ?" ,req.Id)
		fmt.Println(req.Id)
		if m.Id != req.Id {
			return errors.New("empty rows")
		}

		rsp.SiteId = m.SiteId
		rsp.Status = m.Status
		rsp.Id = m.Id
		rsp.Account = m.Account
		rsp.Card = m.Card
		rsp.FianacialContact = m.FinancialContact
		rsp.Tel = m.Tel
		return nil
	}

	query := tool.Mysql
	if req.SiteId > 0 {
		query = query.Where("site_id = ?",req.SiteId)
	}

	if req.Status > 0 {
		query.Where("status = ?", req.Status)
	}

	query.Find(m)

	if m.SiteId != req.SiteId {
		return errors.New("empty rows")
	}

	rsp.SiteId = m.SiteId
	rsp.Status = m.Status
	rsp.Id = m.Id
	rsp.Account = m.Account
	rsp.Card = m.Card
	rsp.FianacialContact = m.FinancialContact
	rsp.Tel = m.Tel


	return nil
}