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

type Platv4User struct {
	ID                 int            `gorm:"column:id;primary_key"`
	Pid                int            `gorm:"column:pid"`
	Nickname           string         `gorm:"column:nickname"`
	Password           string         `gorm:"column:password"`
	Password2          string         `gorm:"column:password2"`
	Username           string         `gorm:"column:username"`
	ChildaccountRemark string         `gorm:"column:childaccount_remark"`
	ChildaccountMax    int            `gorm:"column:childaccount_max"`
	Date               time.Time      `gorm:"column:date"`
	Vip                int            `gorm:"column:vip"`
	Designer           int            `gorm:"column:designer"`
	TemplateTax        int            `gorm:"column:template_tax"`
	OrgName            string         `gorm:"column:org_name"`
	Author             string         `gorm:"column:author"`
	Tel                string         `gorm:"column:tel"`
	TeamTel            string         `gorm:"column:team_tel"`
	Email              sql.NullString `gorm:"column:email"`
	URL                string         `gorm:"column:url"`
	Thumb              string         `gorm:"column:thumb"`
	Province           string         `gorm:"column:province"`
	City               string         `gorm:"column:city"`
	Industry           string         `gorm:"column:industry"`
	Weixin             string         `gorm:"column:weixin"`
	Qq                 string         `gorm:"column:qq"`
	QqUnionid          string         `gorm:"column:qq_unionid"`
	Sina               string         `gorm:"column:sina"`
	Facebook           string         `gorm:"column:facebook"`
	Twitter            string         `gorm:"column:twitter"`
	CreateType         string         `gorm:"column:create_type"`
	CreateSource       string         `gorm:"column:create_source"`
	Truename           string         `gorm:"column:truename"`
	OpenID             string         `gorm:"column:open_id"`
	Status             int            `gorm:"column:status"`
	CheckEmail         int            `gorm:"column:check_email"`
	UserBalance        int            `gorm:"column:user_balance"`
	GiveBalance        int            `gorm:"column:give_balance"`
	FrozenBalance      int            `gorm:"column:frozen_balance"`
	DesignerBalance    int            `gorm:"column:designer_balance"`
	LoginTime          time.Time      `gorm:"column:login_time"`
	Identity           int            `gorm:"column:identity"`
	CheckTel           int            `gorm:"column:check_tel"`
	IsSuiteApp         int            `gorm:"column:is_suite_app"`
	SuiteAppDeviceNo   string         `gorm:"column:suite_app_device_no"`
	SuiteAppType       int            `gorm:"column:suite_app_type"`
	SuiteAppDeviceType string         `gorm:"column:suite_app_device_type"`
	IsLiteUser         int            `gorm:"column:is_lite_user"`
	LiteGoldCount      int            `gorm:"column:lite_gold_count"`
	Team               int            `gorm:"column:team"`
	Position           string         `gorm:"column:position"`
	IosBalance         sql.NullInt64  `gorm:"column:ios_balance"`
}

// TableName sets the insert table name for this struct type
func (p *Platv4User) TableName() string {
	return "platv4_user"
}
