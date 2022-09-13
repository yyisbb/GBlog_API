package config

import (
	"database/sql"
	"fmt"
	"gblog_api/api"
	"gblog_api/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var GlobalConfigCenter = &Config{}

type InitConfig interface {
	InitViper()
	InitRouter()
	InitDataBase()
}

type Config struct {
}

func RegisterConfigCenter() {

	//初始化配置文件
	GlobalConfigCenter.InitViper()
	//初始化数据库
	GlobalConfigCenter.InitDataBase()
	//初始化路由
	GlobalConfigCenter.InitRouter()
}

//InitViper 读取配置文件
func (c *Config) InitViper() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("toml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("[InitViper] init viper success")
}

//InitRouter 初始化路由
func (c *Config) InitRouter() {
	gin.SetMode(viper.GetString("Server.Mode"))
	r := gin.Default()
	api.SetMiddleware(r)
	api.SetRouter(r)
	err := r.Run(fmt.Sprintf(":%s", viper.GetString("Server.Port")))
	if err != nil {
		panic(fmt.Errorf("init router error , %s", err.Error()))
	}
	log.Println("[InitRouter] init router success")
}

func (c *Config) InitDataBase() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("Mysql.DBUser"),
		viper.GetString("Mysql.DBPassWord"),
		viper.GetString("Mysql.DBHost"),
		viper.GetString("Mysql.DBPort"),
		viper.GetString("Mysql.DBName"),
	)
	dbTemp, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(fmt.Errorf("init mysql conn error , %s", err.Error()))
	}
	//设置DB连接池参数
	dbTemp.SetMaxIdleConns(10)
	dbTemp.SetMaxOpenConns(100)
	dbTemp.SetConnMaxLifetime(time.Hour)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbTemp, // DSN data source name
		DefaultStringSize:         256,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		//数据库初始化失败
		panic(fmt.Errorf("init database error , %s", err.Error()))
	}
	if viper.GetString("Server.Mode") == "debug" {
		global.GlobalMysql = db.Debug()
	} else {
		global.GlobalMysql = db
	}
	DBAutoCreate()
	log.Println("[InitDataBase] init database success")
}
