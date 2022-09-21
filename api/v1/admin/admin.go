package admin

import (
	"fmt"
	"gblog_api/global"
	"gblog_api/models"
	"gblog_api/pkg"
	"gblog_api/utils"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"strconv"
)

//
// Login
//  @Description: 后台登录
//  @param c
//
func Login(c *gin.Context) {
	var login models.User

	if err := c.ShouldBindJSON(&login); err != nil {
		log.Println("[Login] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}

	//校验参数
	if err := utils.GetValidate().Struct(login); err != nil {
		log.Println("[Login] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询数据库对比账户密码
	var verifyCount int64
	global.GlobalMysql.Model(models.User{}).Where("username = ?  and password = ?", login.Username, login.Password).Count(&verifyCount)

	if verifyCount == 0 {
		//说明校验失败
		log.Println("[Login] User Error")
		pkg.ResponseJsonError(c, pkg.ERROR_USERNAME_PASSWORD_ERROR)
		return
	}

	//先查询是否有还在有效期内的Token

	thirtyMin := 60 * 30
	tokenKey := utils.RedisToken + login.Username
	if ttl, err := utils.GetRedisTTL(tokenKey); err == nil && ttl > thirtyMin {
		//查询成功并且还在30分钟时效内采用旧Token
		if token, err := utils.GetRedisString(tokenKey); err == nil {
			pkg.ResponseJsonOKAndData(c, token)
			return
		}
		//获取Token错误就生成新的
	}

	//校验成功返回token
	if token, err := utils.GenerateToken(login); err != nil {
		//生成失败
		log.Println("[Login] Token Generate Failed")
		pkg.ResponseJsonError(c, pkg.ERROR_TOKEN_GENERATE_ERROR)
	} else {
		utils.SetToken(login.Username, token)
		pkg.ResponseJsonOKAndData(c, token)
	}

}

//
// GetBlogInfo
//  @Description: 获取博客信息
//  @param c
//
func GetBlogInfo(c *gin.Context) {
	var articleCount int64
	var categoryCount int64
	global.GlobalMysql.Model(models.Article{}).Count(&articleCount)
	global.GlobalMysql.Model(models.Category{}).Count(&categoryCount)
	// 构建响应体
	var resp = models.BlogInfo{
		ArticleCount:  strconv.Itoa(int(articleCount)),
		CategoryCount: strconv.Itoa(int(categoryCount)),
		GolangVersion: runtime.Version(),
		SystemInfo:    fmt.Sprintf("%s-%s", runtime.GOARCH, runtime.GOOS),
	}

	pkg.ResponseJsonOKAndData(c, resp)
	return

}

//
// GetArticleInfo
//  @Description: 获取博客文章数据
//  @param c
//
func GetArticleInfo(c *gin.Context) {
	var articles []models.Article
	global.GlobalMysql.Model(models.Article{}).Find(&articles)
	articleLen := len(articles)
	var resp models.ArticleInfo
	if articleLen != 0 {
		//遍历所有的文章
		for i := 0; i < articleLen; i++ {
			resp.WatchNum += articles[i].WatchNum
			resp.ShareNum += articles[i].ShareNum
			resp.CommentNum += articles[i].CommentNum
		}
	}
	pkg.ResponseJsonOKAndData(c, resp)
}
