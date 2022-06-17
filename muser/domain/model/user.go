package model

type User struct {
	ID       int64  `gorm:"id;primaryKey;autoIncrement;not null"`
	Name     string `gorm:"name;type:string;size:256;uniqueIndex:idx_name"`
	Password string `gorm:"password"`
}
