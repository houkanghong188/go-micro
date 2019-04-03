package model

type DailyPvUvModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

func NewDailyPvUvModel() *DailyPvUvModel {
	return &DailyPvUvModel{}
}

func (m *DailyPvUvModel) TableName() string {
	return "platv4_datastory_daily_pvuv_2019"
}
