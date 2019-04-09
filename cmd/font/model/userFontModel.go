package model

type UserFontModel struct {
	Id         int32
	Uid        int32
	FontIdNo   string
	Type       int8
	Value      string
	Status     int8
	IsDelete   int32
	CreateTime string
	UpdateTime string
}

func (m *UserFontModel) TableName() string {
	return "platv4_font"
}

func NewUserFontModel() *UserFontModel {
	return &UserFontModel{}
}

func (m *UserFontModel) Show() {

}
