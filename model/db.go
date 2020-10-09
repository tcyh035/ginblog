package model

import (
	"fmt"
	"ginblog/utils"
	"time"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//InitDb 初始化数据库
func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	// 禁用默认表名的复数形式
	db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	db.DB().SetMaxIdleConns(10)

	db.DB().SetMaxOpenConns(100)

	db.DB().SetConnMaxLifetime(time.Second * 10)

}
