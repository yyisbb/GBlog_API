package api

import (
	"gblog_api/api/v1/article"
	"gblog_api/api/v1/category"
	"gblog_api/api/v1/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(r *gin.Engine) {
	v1Group := r.Group("/v1")

	//文章相关
	articleGroup := v1Group.Group("/article")
	articleGroup.POST("/createArticle", article.CreateArticle)
	articleGroup.GET("/getArticle", article.GetArticleByID)
	articleGroup.GET("/getArticleTimeProportion", article.GetArticleTimeProportion)
	articleGroup.GET("/getArticleByName", article.GetArticleByName)
	articleGroup.GET("/getAllArticle", article.GetAllArticle)
	articleGroup.GET("/getArticleByCategoryID", article.GetArticleByCategoryID)
	articleGroup.POST("/deleteArticle", article.DeleteArticle)
	articleGroup.POST("/updateArticle", article.UpdateArticle)

	//设置相关
	settingGroup := v1Group.Group("/setting")
	settingGroup.GET("/getSetting", setting.GetSetting)
	settingGroup.POST("/updateSetting", setting.UpdateSetting)

	//分类相关
	categoryGroup := v1Group.Group("/category")
	categoryGroup.GET("/getCategory", category.GetCategory)
}

func SetMiddleware(r *gin.Engine) {
	r.Use(Cors())
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
