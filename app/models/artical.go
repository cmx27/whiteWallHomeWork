package models

type Artical struct {
	ArticalID uint   `json:"artical_id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Artical   string `json:"artical"`
	Namestate bool   `json:"-"` //`gorm:"default:'1'"` //默认匿名
	ImageAddr string `json:"-"`
}

//var Num uint = 0
