package main

import (
	"fmt"
	"log"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/db"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
)

const (
	// 配置文件路径: 用于获取数据库账号密码等
	kConfPath = "./conf/config-dev.toml"
	// 数据库名
	kDatabaseName = "ums"
)

func main() {
	var (
		err error
	)

	// 初始化 logger
	if err = logger.InitLoggerDefault(); err != nil {
		panic(err)
	}
	defer logger.Close()

	// 解析配置
	if err = config.ParseConfig(kConfPath); err != nil {
		logger.Error("parse config fail with err [%v]", err)
		panic(err)
	}

	createDatabase()
}

func createDatabase() {

	// 检查数据库是否存在
	if err := db.PtrMysqlClient.InitWithoutDatabase(); err != nil {
		logger.Error("init mysql client fail with err [%v]", err)
		panic(err)
	}
	db.PtrMysqlClient.Close()
	logger.Info("init mysql client without database successfully!")

	if rawDB, err := db.PtrMysqlClient.RawDB(); err != nil {
		logger.Error("get raw db from mysql client fail with err [%v]", err)
		panic(err)
	} else {
		_, err = rawDB.Exec("USE ums;")
		logger.Info("[use ums] %v", err)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Database [ums] exists")
	}

	fmt.Println(db.PtrMysqlClient.Raw("SHOW DATABASES LIKE mysql"))

	// db.PtrMysqlClient.DB.DB().Raw("SHOW DATABASES LIKE ?", "mysql").Scan(&showDatabaseRes)
	// fmt.Println(util.ToString(showDatabaseRes))

}
