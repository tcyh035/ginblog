package model

import "github.com/jinzhu/gorm"

//Catagory 目录model
type Catagory struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"catagory"`
}
