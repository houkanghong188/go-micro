package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-micro/cmd/Works/proto"
	"go-micro/tool"
	"strconv"
	"time"
)

type WorksAuditModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type works struct {
	ID                int            `gorm:"column:id;primary_key"`
	WorksID           string         `gorm:"column:works_id"`
	UID               sql.NullInt64  `gorm:"column:uid"`
	Title             string         `gorm:"column:title"`
	Content           string         `gorm:"column:content"`
	Comment           string         `gorm:"column:comment"`
	Thumb             sql.NullString `gorm:"column:thumb"`
	FirstImg          string         `gorm:"column:first_img"`
	Version           int            `gorm:"column:version"`
	EditorVersion     sql.NullInt64  `gorm:"column:editor_version"`
	TemplateID        sql.NullString `gorm:"column:template_id"`
	ParentEventID     sql.NullString `gorm:"column:parent_event_id"`
	Identity          int            `gorm:"column:identity"`
	Category          int            `gorm:"column:category"`
	SecondaryCategory int            `gorm:"column:secondary_category"`
	SizeID            sql.NullInt64  `gorm:"column:size_id"`
	PageWidth         int            `gorm:"column:page_width"`
	PageHeight        int            `gorm:"column:page_height"`
	OtherProperty     sql.NullString `gorm:"column:other_property"`
	SuiteAppType      int            `gorm:"column:suite_app_type"`
	Status            int32          `gorm:"column:status"`
	CreateDevice      string         `gorm:"column:create_device"`
	UpdateDevice      string         `gorm:"column:update_device"`
	CreateSite        sql.NullInt64  `gorm:"column:create_site"`
	CreateTime        time.Time      `gorm:"column:create_time"`
	UpdateTime        time.Time      `gorm:"column:update_time"`
	IsDelete          sql.NullInt64  `gorm:"column:is_delete"`
	IsBoughtTemplate  sql.NullInt64  `gorm:"column:is_bought_template"`
	IsUsedLocalFonts  sql.NullInt64  `gorm:"column:is_used_local_fonts"`
}

// TableName sets the insert table name for this struct type
func (p *works) TableName(uid int32) string {
	return "platv5_works_" + strconv.Itoa(int(uid%16))
}

func NewWorksModel() *works {
	return &works{}
}

func NewWorksAuditModel() *WorksAuditModel {
	return &WorksAuditModel{}
}

func (m *WorksAuditModel) TableName() string {
	return "platv5_works_audit"
}

func (m *WorksAuditModel) Show(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	data := []*worksAudit.Response_Notify{}

	query.Table(m.TableName()).Find(&data)

	for k, v := range data {
		tableName := "platv5_works_" + strconv.Itoa(int(v.Uid%16))
		work := works{}
		query.Table(tableName).Where("works_id = ?", v.EventId).First(&work)

		fmt.Println(work)

		fmt.Println(work == (works{}))
		if work == (works{}) {
			fmt.Println(work)
			data[k].Content = " "
			data[k].Title = " "
		} else {
			data[k].Title = work.Title
			data[k].Content = work.Content
		}
	}

	rsp.Data = data

	return nil
}

func (m *WorksAuditModel) Index(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	query.Table(m.TableName()).Count(&rsp.Total)

	if req.PageSize == 0 {
		return errors.New("empty rows")
	}

	data := []*worksAudit.Response_Notify{}

	query.Table(m.TableName()).Limit(req.PageSize).Offset(req.PageSize * req.Page).Find(&data)

	for k, v := range data {
		tableName := "platv5_works_" + strconv.Itoa(int(v.Uid%16))
		work := works{}
		query.Table(tableName).Where("works_id = ?", v.EventId).First(&work)

		fmt.Println(work)

		fmt.Println(work == (works{}))
		if work == (works{}) {
			fmt.Println(work)
			data[k].Content = " "
			data[k].Title = " "
		} else {
			data[k].Title = work.Title
			data[k].Content = work.Content
		}
	}

	rsp.Data = data

	return nil
}

func (m *WorksAuditModel) WorksUpdate(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.WorksId == "" || req.Uid == 0 {
		return errors.New("empty rows works_id | uid")
	}

	works := works{}

	query.Table(works.TableName(req.Uid)).Where("works_id = ?", req.WorksId).First(&works)

	return nil
}

func (m *works) WorksDetail(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.WorksResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.WorksId == "" || req.Uid == 0 {
		return errors.New("empty rows")
	}

	work := []*worksAudit.WorksResponse_Notify{}

	query.Table(m.TableName(req.Uid)).Where("works_id = ?", req.WorksId).First(&work)

	rsp.Data = work

	return nil
}
