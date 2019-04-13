package models

import (
	"context"
	"go-micro/cmd/template/proto/templateStore"
	"go-micro/tool"
)

/**

 */
type TemplateStoreModel struct {
	Id int32
}

func NewTemplateStoreModel() *TemplateStoreModel {
	return &TemplateStoreModel{}
}

func (m *TemplateStoreModel) TableName() string {
	return "platv5_template_store"
}

func (m *TemplateStoreModel) Index(ctx context.Context, in *templateStore.Request, out *templateStore.Response) error {

	conn := tool.GetSlavelConn()
	conn.Table(m.TableName()).Where("template_id in (?)", in.TemplateId).Find(out.Templates)

	return nil
}
