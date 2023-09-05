package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/didi/gendry/scanner"
)

var (
	// db *gorm.DB
	db *sql.DB
)

// TODO(Cat): Exec 方法和 Query 方法的区别是什么？

func query2InterfaceMap(sqlStr string) (res []map[string]interface{}, err error) {
	var (
		rows *sql.Rows
		// rawDB *sql.DB
	)

	rows, err = db.Query(sqlStr)
	defer rows.Close()

	if err != nil {
		logger.Error("query fail||err=%v||sql=%s", err, sqlStr)
		return
	}

	// res, err = scanner.ScanMap(rows)
	res, err = scanner.ScanMapDecode(rows)
	if err != nil {
		logger.Error("scan to map fail||err=%v||sql=%s", err, sqlStr)
		return
	}

	return
}

func main() {
	logger.InitLoggerDefault()
	defer logger.Close()

	// var err error

	// // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:12345@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// // 获取当前数据库
	// fmt.Println(db.Migrator().CurrentDatabase())
	// res, _ := query2InterfaceMap("select * from time_zone limit 3;")
	// fmt.Println(res)

	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// TODO(cat): 必须通过 ping 测试是否连接上数据库
	// https://github.com/jhouyang/usermm/blob/master/initdb.go
	if err := db.Ping(); err != nil {
		logger.Fatal("ping mysql db fail with err [%v]", err)
	}

	if db == nil {
		panic("db is nil")
	}

	// _, err = db.Exec("SHOW DATABASES")
	// if err != nil {
	// 	panic(err)
	// }

	type Database struct {
		Database string
	}

	var res Database

	db.QueryRow("SHOW DATABASES like 'mysql'").Scan(&res.Database)
	fmt.Println(res)

	// res, _ := query2InterfaceMap("show databases")
	// fmt.Println(res)

	// _, err = db.Exec("USE" + name)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	// if err != nil {
	// 	panic(err)
	// }
}
