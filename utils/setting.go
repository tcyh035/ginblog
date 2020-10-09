package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	// AppMode debug或者release
	AppMode string
	// HTTPPort 默认端口号
	HTTPPort string
	// JwtKey 默认JwtKey生成
	JwtKey string
	// Db db used, e.g. mysql, pgsql
	Db string
	// DbHost db hostname
	DbHost string
	// DbPort db port
	DbPort string
	// DbUser username
	DbUser string
	// DbPassword db password
	DbPassword string
	// DbName db name
	DbName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	loadServer(file)
	loadDatabase(file)
}

//
func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HTTPPort = file.Section("server").Key("HTTPPort").MustString("8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("dcg11.RF,!")
}

func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("admin")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
