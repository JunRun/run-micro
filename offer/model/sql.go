/**
 *
 * @Description: 初始化数据库配置
 * @Version: 1.0.0
 * @Date: 2020-03-27 10:38
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/micro/go-micro/v2/config"
	log "github.com/micro/go-micro/v2/logger"
	"sync"
)

var (
	gDB    *gorm.DB
	single sync.Once
	err    error
)

func init() {
	if err := config.LoadFile("conf/application.yml"); err != nil {
		log.Error("读取配置文件出错: ", err)
	}

	host := config.Get("databases", "host").String("localhost")
	user := config.Get("databases", "user").String("root")
	//port := config.Get("databases","port").String("5432")
	password := config.Get("databases", "password").String("root")
	dbname := config.Get("databases", "dbname").String("123")
	conf := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)

	single.Do(func() {
		if gDB, err = gorm.Open("postgres", conf); err != nil {
			panic(err)
		}
	})
	if !gDB.HasTable(&Offer{}) {
		gDB.CreateTable(&Offer{})
		log.Info("创建表成功，offer")
	}

}

func CloseDB() {
	gDB.Close()
}
