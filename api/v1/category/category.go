package category

import (
	"gblog_api/global"
	"gblog_api/models"
	"gblog_api/pkg"
	"gblog_api/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func GetCategory(c *gin.Context) {
	//查询所有分类
	var categories []models.Category
	global.GlobalMysql.Model(models.Category{}).Find(&categories)

	//分类存在
	pkg.ResponseJsonOKAndData(c, categories)
}

//
// GetCategoryByID
//  @Description: 根据ID查找分类
//  @param c
//
func GetCategoryByID(c *gin.Context) {
	//获取分类ID
	id := c.Query("id")

	if len(id) == 0 {
		//id为空
		log.Println("[GetCategoryByID] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询是否存在该分类
	var category models.Category
	global.GlobalMysql.Model(models.Category{}).Where("id = ?", id).First(&category)
	if category.ID == 0 {
		log.Println("[GetCategoryByID] Category Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	//分类存在
	pkg.ResponseJsonOKAndData(c, category)
}

//
// CreateCategory
//  @Description: 新增分类
//  @param c
//
func CreateCategory(c *gin.Context) {
	//接收参数
	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("[CreateCategory] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}

	//校验参数
	err = utils.GetValidate().Struct(category)
	if err != nil {
		log.Println("[CreateCategory] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//新增分类
	if err := global.GlobalMysql.Model(models.Category{}).Create(&category).Error; err != nil {
		log.Println("[CreateCategory] Create Category Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)
}

//
// DeleteCategory
//  @Description: 删除分类
//  @param c
//
func DeleteCategory(c *gin.Context) {
	//获取文章id
	var temp struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Println("[DeleteCategory] Json Parse Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}
	if temp.Id == 0 {
		//id为空
		log.Println("[DeleteCategory] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	//查询是否存在该分类
	var category models.Category
	global.GlobalMysql.Model(models.Category{}).Where("id = ?", temp.Id).First(&category)
	if category.ID == 0 {
		log.Println("[DeleteCategory] Category Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	//删除分类,如果该分类下有文章,则不允许删除
	//查询该分类下所有的文章
	var article []models.Article
	global.GlobalMysql.Model(models.Article{}).Where("category_id = ?", temp.Id).Find(&article)
	if len(article) != 0 {
		log.Println("[DeleteCategory] Category Article Is Not Null")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_EXIST)
		return
	}

	//删除
	if err := global.GlobalMysql.Model(models.Category{}).Delete(&models.Category{}, temp.Id).Error; err != nil {
		log.Println("[DeleteCategory] Category Delete Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)

}

//
// UpdateCategory
//  @Description: 修改分类信息
//  @param c
//
func UpdateCategory(c *gin.Context) {
	//接收参数
	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("[UpdateCategory] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}
	//校验参数
	err = utils.GetValidate().Struct(category)
	if err != nil {
		log.Println("[UpdateCategory] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	if category.ID == 0 {
		//id为空
		log.Println("[UpdateCategory] Category Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	var oldCategory models.Category
	//判断是否有该分类
	global.GlobalMysql.Model(models.Category{}).Where("id = ?", category.ID).First(&oldCategory)
	if oldCategory.ID == 0 {
		//id为空
		log.Println("[UpdateCategory] Category Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	if err := global.GlobalMysql.Model(models.Category{}).Where("id = ?", category.ID).Updates(category).Error; err != nil {
		log.Println("[UpdateCategory] Save Category Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)
}
