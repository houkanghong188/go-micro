package models

import (
	"context"
	"fmt"
	"go-micro/cmd/template/proto/templateStore"
	"go-micro/tool"
)

/**

 */
type TemplateStoreModel struct {
	Id int32
}

type TemplateStoreResult struct {
	Id                int
	TemplateId        string
	DesignerId        int
	Price             int
	Version           int
	Star              int
	Level             int
	Sort              int
	Status            string
	IsImgCaptured     int
	LatestCaptureTime string
	Discount          int
	Corner            string
	SaleNumber        int
	ShowSaleNumber    int
	Pv                int
	Uv                int
	IsVipFree         int
	OnlineTime        string
	FirstOnlineTime   string
	IsDelete          int
	CreateTime        string
	UpdateTime        string
	Title             string
	Identity          int
	Category          int
	SecondaryCategory int
	Scene             int
	SecondaryScene    int
	Industry          int
	SecondaryIndustry int
	Style             int
	Color             int
	Size              int
	CruxWords         string
	TemplateInfo      string
	TemplateProperty  string
}

func NewTemplateStoreModel() *TemplateStoreModel {
	return &TemplateStoreModel{}
}

func (m *TemplateStoreModel) TableName() string {
	return "platv5_template_store"
}

// 通过 template_ids 获取 template_store  template_store_info 信息
func (m *TemplateStoreModel) Index(ctx context.Context, in *templateStore.Request, out *templateStore.Response) error {

	// 拦截 错误，仅仅报错，并使本次请求出错
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("捕获到的错误：%s\n", p)
		}
	}()

	conn := tool.GetSlavelConn()
	info := NewTemplateStoreInfoModel()
	conn.Table(m.TableName()+" a").
		Select("a.*,b.title,b.identity,b.category,b.secondary_category,b.scene,b.secondary_scene,b.industry,b.secondary_industry,b.style,b.color,b.size,b.crux_words,b.template_info,b.template_property").
		Joins(" left join "+info.TableName()+" b on a.template_id = b.template_id").
		Where("a.template_id in (?)", in.TemplateId).
		Find(&out.Templates)

	return nil
}
