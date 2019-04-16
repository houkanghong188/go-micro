package models

/**

 */
type TemplateStoreInfoModel struct {
	Id int32
}

func NewTemplateStoreInfoModel() *TemplateStoreInfoModel {
	return &TemplateStoreInfoModel{}
}

func (m *TemplateStoreInfoModel) TableName() string {
	return "platv5_template_store_info"
}
