package article

import (
	"gblog_api/global"
	"gblog_api/models"
	"gblog_api/pkg"
	"gblog_api/utils"
	"github.com/gin-gonic/gin"
	"log"
)

//
// CreateArticle
//  @Description: 创建文章
//  @param c
//
func CreateArticle(c *gin.Context) {
	//接收参数
	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		log.Println("[CreateArticle] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}
	//校验参数
	err = utils.GetValidate().Struct(article)
	if err != nil {
		log.Println("[CreateArticle] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//添加作者
	var setting models.Setting
	global.GlobalMysql.Model(models.Setting{}).First(&setting)

	//创建文章
	if err := global.GlobalMysql.Create(&article).Error; err != nil {
		log.Println("[CreateArticle] Create Article Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)
}

//
// GetArticleByID
//  @Description: 根据ID获取文章
//  @param c
//
func GetArticleByID(c *gin.Context) {
	//获取文章id
	id := c.Query("id")

	if len(id) == 0 {
		//id为空
		log.Println("[GetArticleByID] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询对应文章进行展示
	var article models.Article
	global.GlobalMysql.Model(models.Article{}).Where("id = ?", id).First(&article)

	if article.ID == 0 {
		//文章不存在
		log.Println("[GetArticleByID] Article Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	//文章存在
	pkg.ResponseJsonOKAndData(c, article)
}

func GetArticleTimeProportion(c *gin.Context) {
	//查询所有文章进行展示
	var timeProportion []struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	}
	if err := global.GlobalMysql.Model(models.Article{}).Select("YEAR(created_at) as label", "count(1) as value").Group("YEAR(created_at)").Find(&timeProportion).Error; err != nil {
		log.Println("[GetArticleTimeProportion] Article TimeProportion Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}
	//文章存在
	pkg.ResponseJsonOKAndData(c, timeProportion)
}

//
// GetArticleByName
//  @Description: 根据ID获取文章
//  @param c
//
func GetArticleByName(c *gin.Context) {
	//获取文章id
	title := c.Query("title")
	//查询对应文章进行展示
	var article []models.Article

	if len(title) == 0 {
		//查询所有
		global.GlobalMysql.Model(models.Article{}).Find(&article)
	}

	global.GlobalMysql.Model(models.Article{}).Where("title like ?", "%"+title+"%").Find(&article)

	//文章存在
	pkg.ResponseJsonOKAndData(c, article)
}

//
// GetAllArticle
//  @Description: 获取所有文章
//  @param c
//
func GetAllArticle(c *gin.Context) {
	//查询所有文章进行展示
	var articles []models.Article
	if err := global.GlobalMysql.Model(models.Article{}).Order("created_at desc").Find(&articles).Error; err != nil {
		log.Println("[DeleteArticle] Articles Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}
	//文章存在
	pkg.ResponseJsonOKAndData(c, articles)
}

//
// GetArticleByCategoryID
//  @Description: 查询所有该分类下的文章列表
//  @param c
//
func GetArticleByCategoryID(c *gin.Context) {
	//获取分类ID
	id := c.Query("category_id")

	if len(id) == 0 {
		//id为空
		log.Println("[DeleteArticle] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询对应文章进行展示
	var articles []models.Article
	global.GlobalMysql.Model(models.Article{}).Where("category_id = ?", id).Find(&articles)

	//文章存在
	pkg.ResponseJsonOKAndData(c, articles)
}

//
// DeleteArticle
//  @Description: 删除指定文章
//  @param c
//
func DeleteArticle(c *gin.Context) {
	//获取文章id
	id := c.Query("id")

	if len(id) == 0 {
		//id为空
		log.Println("[DeleteArticle] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询对应文章进行展示
	var article models.Article
	global.GlobalMysql.Model(models.Article{}).Where("id = ?", id).First(&article)

	if article.ID == 0 {
		//文章不存在
		log.Println("[DeleteArticle] Article Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	//文章存在就删除
	global.GlobalMysql.Model(models.Article{}).Delete(&models.Article{}, id)
	pkg.ResponseJsonOK(c)
}

//
// UpdateArticle
//  @Description: 更新指定文章
//  @param c
//
func UpdateArticle(c *gin.Context) {
	//接收参数
	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		log.Println("[UpdateArticle] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}
	//校验参数
	err = utils.GetValidate().Struct(article)
	if err != nil {
		log.Println("[UpdateArticle] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	if article.ID == 0 {
		//id为空
		log.Println("[UpdateArticle] Article Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	var oldArticle models.Setting
	//判断是否有该条文章
	global.GlobalMysql.Model(models.Setting{}).Where("id = ?", article.ID).First(&oldArticle)
	if oldArticle.ID == 0 {
		//id为空
		log.Println("[UpdateArticle] Article Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	if err := global.GlobalMysql.Model(models.Article{}).Where("id = ?", article.ID).Save(&article).Error; err != nil {
		log.Println("[UpdateArticle] Save Article Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)
}
