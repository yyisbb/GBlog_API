package models

type ArticleInfo struct {
	//分享数
	ShareNum int `json:"shareNum"`
	//观看数
	WatchNum int `json:"watchNum"`
	//评论数
	CommentNum int `json:"commentNum"`
}
