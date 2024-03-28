package model

type Goods struct {
	Id             uint `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primaryKey;"`
	Name           string `json:"name" gorm:"column:name;type:varchar(16) not null;default:'';index:idx_name"`
	Description    string `json:"description" gorm:"column:description;type:varchar(16) not null;default:'';index:idx_description"`
	Price          string `json:"price" gorm:"column:price;type:varchar(16) not null;default:'';index:idx_price"`
	Url            string `json:"url" gorm:"column:url;type:varchar(64) not null;default:'';index:idx_url"`
	CategoryId     int `json:"category_id" gorm:"column:category_id;type:int(10) unsigned not null; default:0;index:idx_category_id"`
	CreatedTime int64   `json:"created_time" gorm:"column:created_time;type:int(11) not null;default:0;index:idx_created_time"`
    UpdatedTime int64   `json:"updated_time" gorm:"column:updated_time;type:int(11) not null;default:0;index:idx_updated_time"`
}