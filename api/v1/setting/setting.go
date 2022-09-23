package setting

import (
	"gblog_api/global"
	"gblog_api/models"
	"gblog_api/pkg"
	"gblog_api/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func InitSetting() models.Setting {
	var setting models.Setting
	global.GlobalMysql.Model(models.Setting{}).First(&setting)
	if setting.ID != 0 {
		return setting
	}

	//初始化个人信息
	if err := global.GlobalMysql.Model(models.Setting{}).Create(&models.Setting{
		Name:       "WebName",
		Logo:       "LogoUrl",
		Email:      "Email@Host.com",
		Avatar:     "https://img1.baidu.com/it/u=873106765,2587410047&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662656400&t=5cfd0b69720c1147364aa4104b1101e8",
		AuthorName: "AuthorName",
	}).Error; err != nil {
		log.Printf("[InitSetting] Init Setting Error %s ", err.Error())
	}

	return InitSetting()
}

func GetSetting(c *gin.Context) {
	//获取到设置信息
	setting := InitSetting()
	pkg.ResponseJsonOKAndData(c, setting)
}

func UpdateSetting(c *gin.Context) {
	//接收参数
	var setting models.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		log.Println("[UpdateSetting] Parse JSON Error")
		pkg.ResponseJsonError(c, pkg.ERROR_JSONPARSE)
		return
	}
	//校验参数
	err = utils.GetValidate().Struct(setting)
	if err != nil {
		log.Println("[UpdateSetting] Param Error")
		pkg.ResponseJsonError(c, pkg.ERROR_PARAM)
		return
	}

	if setting.ID == 0 {
		//id为空
		log.Println("[UpdateSetting] Setting Not Found")
		pkg.ResponseJsonError(c, pkg.ERROR_DATA_NOT_FUOUND)
		return
	}

	if err := global.GlobalMysql.Model(models.Setting{}).Where("id = ?", setting.ID).Save(&setting).Error; err != nil {
		log.Println("[UpdateSetting] Save Setting Error")
		pkg.ResponseJsonError(c, pkg.ERROR_SQL)
		return
	}

	pkg.ResponseJsonOK(c)
}
