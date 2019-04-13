package models

/**
`id` int(11) NOT NULL AUTO_INCREMENT,
 `content_id` int(11) DEFAULT NULL COMMENT '内容id',
 `title` varchar(32) DEFAULT NULL COMMENT '标签名称',
 `type` enum('h5','video','poster') NOT NULL COMMENT '类型',
 `description` varchar(255) DEFAULT NULL COMMENT '描述',
 `sort` smallint(6) DEFAULT NULL COMMENT '排序',
 `nexus_template_id` varchar(255) NOT NULL COMMENT '关联模版',
 `nexus_crux_words` varchar(255) DEFAULT NULL COMMENT '关联关键词',
 `nexus_scene` varchar(255) DEFAULT NULL COMMENT '关联场景',
 `nexus_industry` varchar(255) DEFAULT NULL COMMENT '关联行业',
 `nexus_color` varchar(255) DEFAULT NULL COMMENT '关联色调',
 `nexus_style` varchar(255) DEFAULT NULL COMMENT '关联风格',
 `create_time` datetime DEFAULT NULL COMMENT '创建时间',
 `update_time` datetime DEFAULT NULL COMMENT '更新时间',
 `is_delete` smallint(6) DEFAULT '0' COMMENT '是否删除',
 `relevant_category` int(11) DEFAULT '0' COMMENT '关联的分类',
*/
type MarketContentRecommendModel struct {
	Id int32
}

func NewMarketContentRecommendModel() *MarketContentRecommendModel {
	return &MarketContentRecommendModel{}
}

func (m *MarketContentRecommendModel) TableName() string {
	return "platv5_market_content_Recommend"
}
