package main

import (
	"fmt"

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

	// 创建数据库
	createDatabase(kDatabaseName)
}

func createDatabase(dbName string) {
	var (
		err error
	)

	// 初始化数据库
	if err = db.PtrMysqlClient.InitWithoutDatabase(); err != nil {
		logger.Fatal("init mysql client fail with err [%v]", err)
	}
	defer db.PtrMysqlClient.Close()
	logger.Info("init mysql client without database successfully!")

	// 检查数据库是否存在
	if isDatabaseExist(dbName) {
		logger.Info("database [%s] already exist!", dbName)
		return
	}
	logger.Info("database [%s] not exist, create it now!", dbName)

	// 创建数据库
	if execResult := db.PtrMysqlClient.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)); execResult.Error != nil {
		logger.Fatal("crate database [%s] fail with err [%v]", dbName, execResult.Error)
	} else if execResult.RowsAffected != 1 {
		logger.Fatal("create database [%s] fail with RowsAffected [%d]", dbName, execResult.RowsAffected)
	}

	// 再次检查数据库是否创建成功
	if !isDatabaseExist(dbName) {
		logger.Fatal("quit due to create database [%s] fail!")
	}

	logger.Info("create database [%s] successfully", dbName)
}

func isDatabaseExist(dbName string) bool {
	var (
		err error
		res []map[string]string
	)

	// 检查数据库是否存在
	// SQL 中预处理语句的参数化不能用于表名、列名和数据库名等, 这些名字在 SQL 执行前就需要确定下来
	res, err = db.PtrMysqlClient.Query2StringMap(fmt.Sprintf("SHOW DATABASES like '%s';", dbName))
	if err != nil {
		logger.Fatal("show database [%s] fail with err [%v]", dbName, err)
	}
	if len(res) == 1 {
		return true
	}
	return false
}
