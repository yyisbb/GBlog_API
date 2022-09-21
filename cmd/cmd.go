package cmd

import (
	"gblog_api/global"
	"gblog_api/models"
	"gblog_api/utils"
	"log"
)

func Run(configCenter func()) {
	go timerTask()
	go configCenter()
	select {}
}

func timerTask() {
	//从0小时开始,每小时执行一次
	_, err := utils.CronUtil.AddFunc("0 0 0/1 * * ?", func() {
		//查询所有文章统计次数
		var articles []models.Article
		global.GlobalMysql.Model(models.Article{}).Find(&articles)
		articleLen := len(articles)
		if articleLen != 0 {
			//遍历所有的文章
			for i := 0; i < articleLen; i++ {

			}
		}
	})
	utils.CronUtil.Start()
	if err != nil {
		log.Println("timerTask Error", err)
	}
}
