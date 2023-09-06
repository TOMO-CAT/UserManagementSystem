package db

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/config"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
	"github.com/didi/gendry/scanner"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var PtrMysqlClient *mysqlClient

type mysqlClient struct {
	*gorm.DB         // gorm sql 接口, 可以直接继承其方法
	rawDB    *sql.DB // 原生的 sql 接口
}

func (c *mysqlClient) Init() (err error) {
	return c.init(config.GlobalUmsConfig.Mysql.Database)
}

func (c *mysqlClient) InitWithoutDatabase() (err error) {
	return c.init("")
}

func (c *mysqlClient) Close() {
	if sqlDB, err := c.DB.DB(); err != nil {
		logger.Error("fetch mysql db fail with err [%v]", err)
		return
	} else {
		sqlDB.Close()
	}
}

// Query2InterfaceMap
// TODO(cat): 让 Query2InterfaceMap 入参支持 Query2InterfaceMap("SELECT user_name from USER where id = ?", 100); 的 ? 写法
func (c *mysqlClient) Query2InterfaceMap(sqlStr string) (res []map[string]interface{}, err error) {
	var rows *sql.Rows
	rows, err = c.Raw(sqlStr).Rows()
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

func (c *mysqlClient) Query2StringMap(sqlStr string) (res []map[string]string, err error) {
	var interfaceRes []map[string]interface{}
	interfaceRes, err = c.Query2InterfaceMap(sqlStr)
	if err != nil {
		return
	}

	for _, record := range interfaceRes {
		tmpMap := make(map[string]string)
		for k, v := range record {
			switch v.(type) {
			case string:
				tmpMap[k] = v.(string)
			case []byte:
				tmpMap[k] = string(v.([]byte))
			case nil:
				tmpMap[k] = ""
			case int:
				tmpMap[k] = strconv.Itoa(v.(int))
			case float64:
				tmpMap[k] = fmt.Sprintf("%f", v.(float64))
			default:
				logger.Error("Query2StringMap fail with invalid type||k=%s||v=%v||type=%T", k, v, v)
				tmpMap[k] = ""
				continue
			}
		}
		res = append(res, tmpMap)
	}
	return
}

// Query2StructArray 查询 sql, 并将结果映射到 resPtr 指针指向的结构体中
// @param sqlStr 待查询的 sql
// @param resPtr目标结构体指针的数组 &[]struct
// @return err
//
// eg:
// var StudentInfo []struct {
//     Name  string  `ddb:"name"`
//     Score float64 `ddb:"score"`
// }
// sqlStr := "SELECT name, score FROM student"
// err = PtrMySqlClient.Query2StructArray(sql, &StudentInfo)
func (c *mysqlClient) Query2StructArray(sqlStr string, resPtr interface{}) (err error) {
	var (
		rows *sql.Rows
	)
	rows, err = c.Raw(sqlStr).Rows()
	defer rows.Close()

	if err != nil {
		logger.Error("query fail||err=%v||sql=%s", err, sqlStr)
		return
	}

	err = scanner.Scan(rows, resPtr)
	if err != nil {
		logger.Error("scan fail||err=%v||sql=%s", err, sqlStr)
		return
	}
	return
}

// Query2InsertStatement 将查询结果转化成 insert 语句, 用于拷贝线上数据库
// @param tableName 数据表名
// @param sqlStr 查询 sql
// @param outputFile 保存 insert 语句的文件
// @return err
//
// eg:
// @param sqlStr SELECT name, score FROM student WHERE id < 100;
// @param tableName student
// @param outputFile "./output.txt"
//
// 写入文件的内容:
// INSERT INFO `student` (`name`, `score`) VALUES ("cat", 78), ("dog", 21);
func (c *mysqlClient) Query2InsertStatement(tableName, sqlStr, outputFile string) (err error) {
	// 判断文件是否存在, 不存在时创建
	_, err = os.Stat(outputFile)
	if os.IsNotExist(err) {
		var f *os.File
		if f, err = os.Create(outputFile); err != nil {
			return
		}
		defer f.Close()
	}

	var queryRes []map[string]string
	if queryRes, err = c.Query2StringMap(sqlStr); err != nil {
		return
	}
	if len(queryRes) <= 0 {
		return errors.New("empty res")
	}

	var f *os.File
	if f, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	// 获取所有字段名
	var fields []string
	var fieldsWithBackquote []string
	for field := range queryRes[0] {
		fields = append(fields, field)
		fieldsWithBackquote = append(fieldsWithBackquote, "`"+field+"`")
	}

	// 构造字段子句
	// eg: (`name`, `score`)
	fieldClause := strings.Join(fieldsWithBackquote, ",")

	// 构造值子句列表
	// eg: [("cat", 78), ("dog", 21)]
	var valueClauseArray []string
	for _, record := range queryRes {
		var values []string
		for _, field := range fields {
			value := "\"" + record[field] + "\""
			values = append(values, value)
		}
		valueClause := "(" + strings.Join(values, ",") + ")"
		valueClauseArray = append(valueClauseArray, valueClause)
	}
	valueClauseArrayStr := strings.Join(valueClauseArray, ",\n")

	// 组成 insert statement
	// eg: INSERT INFO `student` (`name`, `score`) VALUES ("cat", 78), ("dog", 21);
	insertStatementFmt := "INSERT INTO \n `%s` (%s) \n VALUES \n %s;"
	insertStatement := fmt.Sprintf(insertStatementFmt, tableName, fieldClause, valueClauseArrayStr)
	writer.WriteString(insertStatement)
	writer.WriteString("\n")
	writer.Flush()

	return nil
}

func (c *mysqlClient) init(databaseName string) (err error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.GlobalUmsConfig.Mysql.User,
		config.GlobalUmsConfig.Mysql.Password,
		config.GlobalUmsConfig.Mysql.IP,
		config.GlobalUmsConfig.Mysql.Port,
		databaseName,
	)

	if c.DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{}); err != nil {
		return
	}

	if c.rawDB, err = c.DB.DB(); err != nil {
		logger.Error("fetch mysql db fail with err [%v]", err)
		return
	}

	// 判断是否能 ping 通
	if err = c.rawDB.Ping(); err != nil {
		logger.Error("ping mysql fail with err [%v]", err)
		return
	}

	// 设置连接池
	c.rawDB.SetMaxOpenConns(config.GlobalUmsConfig.Mysql.MaxOpenConns)
	c.rawDB.SetMaxIdleConns(config.GlobalUmsConfig.Mysql.MaxIdleConns)
	c.rawDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func init() {
	PtrMysqlClient = &mysqlClient{}
}
