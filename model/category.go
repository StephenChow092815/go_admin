package model

type Category struct {
	Id          uint   `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primaryKey;"`
	Name        string `json:"name" gorm:"column:name;type:varchar(16) not null;default:'';index:idx_name"`
	Tag         string `json:"tag" gorm:"column:tag;type:varchar(16) not null;default:'';index:idx_tag"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time;type:int(11) not null;default:0;index:idx_created_time"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time;type:int(11) not null;default:0;index:idx_updated_time"`
}
