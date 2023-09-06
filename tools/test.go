package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/db"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
)

const (
	// 配置文件路径: 用于获取数据库账号密码等
	kConfPath = "./conf/config-dev.toml"
)

func main() {
	if err := config.ParseConfig(kConfPath); err != nil {
		logger.Error("parse config fail with err [%v]", err)
		panic(err)
	}

	logger.InitLoggerDefault()
	defer logger.Close()

	if err := db.PtrMysqlClient.Init(); err != nil {
		logger.Fatal("init mysql redis fail with err [%v]", err)
	}
	defer db.PtrMysqlClient.Close()

	res, err := db.PtrMysqlClient.Query2InterfaceMap("SELECT * FROM user_infos WHERE id > ? and uid > ?", 0, 10)
	if err != nil {
		logger.Fatal("query fail with err [%v]", err)
	}
	fmt.Println(res)
}
