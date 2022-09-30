package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	//网站名称
	Name string `json:"name,omitempty" validate:"required"`
	//网站Logo
	Logo string `json:"logo,omitempty" validate:"required"`
	//个人邮箱
	Email string `json:"email,omitempty" validate:"required"`
	//个人头像
	Avatar string `json:"avatar,omitempty" validate:"required"`
	//个人名称
	AuthorName string `json:"authorName" validate:"required"`
	//首页文字
	HomeText string `json:"homeText"`
	//关于界面介绍
	AboutContent string `json:"aboutContent,omitempty" gorm:"type:longtext"`
	//Github链接
	GithubUrl string `json:"githubUrl,omitempty"`
}
