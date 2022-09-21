package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	//文章标题
	Title string `json:"title,omitempty" validate:"required"`
	//文章首页大图
	BackGround string `json:"backGround,omitempty" validate:"required"`
	//文章内容
	Content string `json:"content,omitempty" validate:"required" gorm:"type:longtext"`
	//文章分类ID
	CategoryID int `json:"categoryID,omitempty" validate:"required"`
	//分类名
	CategoryName string `json:"categoryName,omitempty"`
	//分享数
	ShareNum int `json:"shareNum"`
	//观看数
	WatchNum int `json:"watchNum"`
	//评论数
	CommentNum int `json:"commentNum"`
}
