package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	//分类名
	Name string `json:"name,omitempty"`
	//分类描述
	Description string `json:"description,omitempty"`
	//分类图
	Banner string `json:"banner,omitempty"`
}
