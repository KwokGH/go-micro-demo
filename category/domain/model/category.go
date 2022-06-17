package model

type Category struct {
	ID       int64  `json:"id" gorm:"id;primaryKey;autoIncrement;not null"`
	Name     string `json:"name" gorm:"name"`
	Level    int32  `json:"level" gorm:"name"`
	ParentID int64  `json:"parent_id" gorm:"name"`
	Image    string `json:"image" gorm:"name"`
	Desc     string `json:"desc" gorm:"name"`
}
