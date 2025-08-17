package inits

import (
	"fmt"
	"github.com/GUOSAITONG/MyConfig/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

func InitMysql() {
	var Once sync.Once
	var err error
	mysqls := config.Config.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqls.User, mysqls.Password, mysqls.Host, mysqls.Port, mysqls.Database)
	Once.Do(func() {
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Println("mysql init success")
	})

}
