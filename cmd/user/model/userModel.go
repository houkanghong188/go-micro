package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-micro/cmd/works/model"
	"strconv"

	//"encoding/json"
	"errors"
	//"fmt"
	"go-micro/cmd/user/proto"
	"go-micro/tool"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type UserModel struct {
	ID                 int32          `gorm:"column:id;primary_key"`
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
	Status             int32          `gorm:"column:status"`
	CheckEmail         int            `gorm:"column:check_email"`
	UserBalance        int32          `gorm:"column:user_balance"`
	GiveBalance        int32          `gorm:"column:give_balance"`
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

type AuditLogModel struct {
	Id         int32
	WorksId    string
	Uid        int32
	Reason     string
	Type       string
	CreateTime string
	UpdateTime string
}

//type works struct {
//	WorksId string
//}

// TableName sets the insert table name for this struct type
func (p *AuditLogModel) TableName() string {
	return "platv5_works_audit_log"
}

// TableName sets the insert table name for this struct type
func (p *UserModel) TableName() string {
	return "platv4_user"
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) Show(ctx context.Context, req *user.Request, rsp *user.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	data := []*user.UserBracket{}

	if req.Uid != 0 {
		query = query.Where("id = ?", req.Uid)
	}

	query.Table("platv4_user").First(&data)

	rsp.Data = data
	return nil
}

func (m *UserModel) Closure(ctx context.Context, req *user.Request, rsp *user.Response) error {
	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.Uid <= 0 {
		return errors.New("empty rows")
	}

	var User = UserModel{}
	query.Where("id = ?", req.Uid).First(&User)
	User.Status = req.Status
	query.Table(User.TableName()).Where("id = ?", req.Uid).Updates(&User)
	query.Debug().Table(User.TableName()).Where("id = ?", req.Uid).Updates(&User)

	// 添加日志
	var LogType string
	if req.Status == 0 {
		LogType = "block"
		blockWorks(req.Uid)
	} else {
		LogType = "deblock"
		deBlockWorks(req.Uid)
	}
	auditLog := AuditLogModel{Uid: req.Uid, Reason: req.Reason, Type: LogType}
	query.Create(&auditLog)

	// 如果有封禁日期
	if req.ClosureDay != "" {
		c, _ := tool.GetRedis()
		c.Do("hset", "user_closure_day"+req.ClosureDay, req.Uid)
	}
	// 设置redis
	return nil
}

func blockWorks(uid int32) {
	query := tool.GetMasterConn()
	worksName := "platv5_works_" + strconv.Itoa(int(uid%16))
	worksIds := []model.Works{}
	query.Table(worksName).Select("works_id").Where("uid = ?", uid).Scan(&worksIds)
	//设置对应信息
	b, err := json.Marshal(worksIds)
	if err != nil {
	}
	c, _ := tool.GetRedis()
	_, berr := c.Do("hset", "user_closure", uid, b)
	if berr != nil {
		fmt.Println(berr)
	}
	for _, v := range worksIds {
		if v.WorksID == "" {
			continue
		}
		newQuery := tool.GetMasterConn()
		newQuery.Table(v.TableName(v.UID)).Where("works_id = ?", v.WorksID).Update("status", -3)
		newQuery.Debug().Table(v.TableName(v.UID)).Where("works_id = ?", v.WorksID).Update("status", -3)
	}
}

func deBlockWorks(uid int32) {
	c, _ := tool.GetRedis()

	redisWorksIds, err := redis.String(c.Do("hget", "user_closure", uid))
	if err != nil {
		return
	}
	data := []model.Works{}
	json.Unmarshal([]byte(redisWorksIds), &data)

	for _, v := range data {
		fmt.Println(v)
		if v.WorksID == "" {
			continue
		}
		newQuery := tool.GetMasterConn()
		newQuery.Table(v.TableName(v.UID)).Where("works_id = ?", v.WorksID).Update("status", 0)
	}
}
