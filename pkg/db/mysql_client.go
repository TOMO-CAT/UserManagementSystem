package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/didi/gendry/scanner"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var PtrMysqlClient *mysqlClient

type mysqlClient struct {
	*gorm.DB
}

func (c *mysqlClient) Init() (err error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.GlobalUmsConfig.Mysql.User,
		config.GlobalUmsConfig.Mysql.Password,
		config.GlobalUmsConfig.Mysql.IP,
		config.GlobalUmsConfig.Mysql.Port,
		config.GlobalUmsConfig.Mysql.Database,
	)

	if c.DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{}); err != nil {
		return
	}

	// 设置连接池
	var sqlDB *sql.DB
	if sqlDB, err = c.DB.DB(); err != nil {
		logger.Error("fetch mysql db fail with err [%v]", err)
		return
	}
	sqlDB.SetMaxOpenConns(config.GlobalUmsConfig.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.GlobalUmsConfig.Mysql.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func (c *mysqlClient) InitWithoutDatabase() (err error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.GlobalUmsConfig.Mysql.User,
		config.GlobalUmsConfig.Mysql.Password,
		config.GlobalUmsConfig.Mysql.IP,
		config.GlobalUmsConfig.Mysql.Port,
	)

	if c.DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{}); err != nil {
		return
	}

	// 设置连接池
	var sqlDB *sql.DB
	if sqlDB, err = c.DB.DB(); err != nil {
		logger.Error("fetch mysql db fail with err [%v]", err)
		return
	}
	sqlDB.SetMaxOpenConns(config.GlobalUmsConfig.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.GlobalUmsConfig.Mysql.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func (c *mysqlClient) Close() {
	if sqlDB, err := c.DB.DB(); err != nil {
		logger.Error("fetch mysql db fail with err [%v]", err)
		return
	} else {
		sqlDB.Close()
	}
}

func (c *mysqlClient) RawDB() (rawDB *sql.DB) {
	var err error
	if rawDB, err = c.DB.DB(); err != nil {
		logger.Error("fetch mysql raw db fail with err [%v]", err)
		panic(err)
	}
	return
}

func (c *mysqlClient) Query2InterfaceMap(sqlStr string) (res []map[string]interface{}, err error) {
	var rows *sql.Rows
	rows, err = c.RawDB().Query(sqlStr)
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

func init() {
	PtrMysqlClient = &mysqlClient{}
}
