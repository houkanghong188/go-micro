package model

import (
	"context"
	"database/sql"
	"errors"
	"go-micro/cmd/works/proto"
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

type Works struct {
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

type AuditLogModel struct {
	Id         int32
	WorksId    int32
	Uid        int32
	Reason     string
	Type       string
	CreateTime string
	UpdateTime string
}

// TableName sets the insert table name for this struct type
func (p *AuditLogModel) TableName() string {
	return "platv5_works_audit_log"
}

// TableName sets the insert table name for this struct type
func (p *Works) TableName(uid int32) string {
	return "platv5_works_" + strconv.Itoa(int(uid%16))
}

func NewWorksModel() *Works {
	return &Works{}
}

func NewWorksAuditModel() *WorksAuditModel {
	return &WorksAuditModel{}
}

func (m *WorksAuditModel) TableName() string {
	return "platv5_works_audit"
}

func (m *WorksAuditModel) Show(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.ShowResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	data := worksAudit.AuditBracket{}

	if req.Uid != 0 {
		query = query.Where("uid = ?", req.Uid)
	}

	if req.WorksId != "" {
		query = query.Where("event_id = ?", req.WorksId)
	}

	query.Table(m.TableName()).First(&data)

	tableName := "platv5_works_" + strconv.Itoa(int(data.Uid%16))
	work := Works{}
	newQuery := tool.GetMasterConn()
	newQuery.Table(tableName).Where("works_id = ?", data.EventId).First(&work)

	if work == (Works{}) {
		data.Content = " "
		data.Title = " "
	} else {
		data.Title = work.Title
		data.Content = work.Content
	}

	rsp.Data = &data

	return nil
}

func (m *WorksAuditModel) Index(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.PageSize == 0 {
		return errors.New("empty rows")
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	if req.Uid != 0 {
		query = query.Where("uid = ?", req.Uid)
	}

	if req.WorksId != "" {
		query = query.Where("event_id = ?", req.WorksId)
	}

	if len(req.AuditStatus) > 0 {
		query = query.Where("status in (?)", req.AuditStatus)
	}

	query.Table(m.TableName()).Count(&rsp.Total)

	data := []*worksAudit.AuditBracket{}

	query.Table(m.TableName()).Limit(req.PageSize).Offset(req.PageSize * req.Page).Find(&data)

	newQuery := tool.GetMasterConn()
	for k, v := range data {
		tableName := "platv5_works_" + strconv.Itoa(int(v.Uid%16))
		work := Works{}
		newQuery.Table(tableName).Where("works_id = ?", v.EventId).First(&work)

		if work == (Works{}) {
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

	works := Works{}

	query.Table(works.TableName(req.Uid)).Where("works_id = ?", req.WorksId).First(&works)

	// 添加日志
	var LogType string
	if req.Status == 1 {
		LogType = "block"
	} else {
		LogType = "deblock"
	}
	auditLog := AuditLogModel{Uid: req.Uid, Reason: req.Reason, Type: LogType}
	query.Create(&auditLog)

	return nil
}

func (m *Works) WorksDetail(ctx context.Context, req *worksAudit.Request, rsp *worksAudit.WorksResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.WorksId == "" || req.Uid == 0 {
		return errors.New("empty rows")
	}

	work := []*worksAudit.WorksBracket{}

	query.Table(m.TableName(req.Uid)).Where("works_id = ?", req.WorksId).First(&work)

	rsp.Data = work

	return nil
}
