package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Platv5Works0 struct {
	ID                int            `gorm:"column:id;primary_key"`
	WorksID           string         `gorm:"column:works_id"`
	UID               sql.NullInt64  `gorm:"column:uid"`
	Title             string         `gorm:"column:title"`
	Content           sql.NullString `gorm:"column:content"`
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
	Status            int            `gorm:"column:status"`
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
func (p *Platv5Works0) TableName() string {
	return "platv5_works_0"
}
