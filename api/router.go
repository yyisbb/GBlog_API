package api

import (
	"gblog_api/api/v1/admin"
	"gblog_api/api/v1/article"
	"gblog_api/api/v1/category"
	"gblog_api/api/v1/setting"
	"gblog_api/pkg"
	"gblog_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func SetRouter(r *gin.Engine) {
	v1Group := r.Group("/v1")

	//无需校验token的
	//文章相关
	v1Group.GET("/getArticle", article.GetArticleByID)
	v1Group.GET("/getArticleTimeProportion", article.GetArticleTimeProportion)
	v1Group.GET("/getArticleByName", article.GetArticleByName)
	v1Group.GET("/getAllArticle", article.GetAllArticle)
	v1Group.GET("/getArticleByCategoryID", article.GetArticleByCategoryID)

	//设置相关
	v1Group.GET("/getSetting", setting.GetSetting)

	//分类相关
	v1Group.GET("/getCategory", category.GetCategory)

	//后台相关
	v1Group.POST("/login", admin.Login)
	v1Group.GET("/getBlogInfo", admin.GetBlogInfo)
	v1Group.GET("/getArticleInfo", admin.GetArticleInfo)
	v1Group.POST("/addWatchNum", article.AddWatchNum)

	//需要校验token的
	//文章相关
	adminGroup := v1Group.Group("/admin")
	adminGroup.Use(JWT())
	adminGroup.POST("/createArticle", article.CreateArticle)
	adminGroup.POST("/deleteArticle", article.DeleteArticle)
	adminGroup.POST("/updateArticle", article.UpdateArticle)

	//设置相关
	adminGroup.POST("/updateSetting", setting.UpdateSetting)

	//分类相关
	adminGroup.POST("/deleteCategory", category.DeleteCategory)
	adminGroup.POST("/createCategory", category.CreateCategory)
	adminGroup.POST("/updateCategory", category.UpdateCategory)
}

func SetMiddleware(r *gin.Engine) {
	r.Use(Cross())
}

func Cross() gin.HandlerFunc {
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

// JWT 自定义JWT中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if strings.Contains(method, "POST") {
			token := c.GetHeader("token")
			if token == "" {
				pkg.ResponseJsonError(c, pkg.ERROR_TOTKEN_NULL)
				c.Abort()
				return
			} else {
				// 解析token
				claims, err := utils.ParseToken(token)
				if err != nil {
					pkg.ResponseJsonError(c, pkg.ERROR_TOTKEN)
					c.Abort()
					return
				} else if time.Now().Unix() > claims.ExpiresAt {
					pkg.ResponseJsonError(c, pkg.ERROR_TOKEN_TIMEOUT)
					c.Abort()
					return
				}
			}
			c.Next()
		}
	}
}
