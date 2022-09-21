package models

type BlogInfo struct {
	ArticleCount  string `json:"article_count,omitempty"`
	CategoryCount string `json:"category_count,omitempty"`
	GolangVersion string `json:"golang_version,omitempty"`
	SystemInfo    string `json:"system_info,omitempty"`
}
