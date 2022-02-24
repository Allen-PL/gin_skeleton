package dao

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/driver/postgres"
	"log"
	"novel_learning/global/my_errors"
	zerolog "novel_learning/log"
	"strings"
)

var (
	db *gorm.DB
)

//  当前支持的数据库
const (
	DriverMysql = "mysql"
	DriverSqlServer = "sqlserver"
	DriverPostgresql = "postgresql"
)

// 初始化数据库配置
func SetUp() {
	if dbDriver, err := greSqlDriver(); err != nil {
		log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
	} else {
		db = dbDriver
	}
}

// 获取sql驱动
func greSqlDriver() (*gorm.DB, error) {
	var dbDialect gorm.Dialector
	sqlType := viper.GetString("database.sqlType")
	if dialect, err := getDbDialect(sqlType, "write"); err != nil {
		zerolog.Error(my_errors.ErrorsDialectorDbInitFail + err.Error())
	} else {
		dbDialect = dialect
	}
	// gorm配置，获取数据库句柄
	gorm, err := gorm.Open(dbDialect, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	if viper.
}

// 获取一个数据库方言(Dialector)，通俗的将就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialect(sqlType, readWrite string) (gorm.Dialector, error) {
	var dbDialect gorm.Dialector
	dsn := getDsn(sqlType, readWrite)
	switch strings.ToLower(sqlType) {
	case DriverMysql:
		dbDialect = mysql.Open(dsn)
	case DriverSqlServer:
		dbDialect = sqlserver.Open(dsn)
	case DriverPostgresql:
		dbDialect = postgres.Open(dsn)
	default:
		return nil, errors.New(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return dbDialect, nil
}

func getDsn(sqlType, readWrite string) string {
	/*
	sqlType: 根据数据库的类型确定数据库连接方式
	readWrite: 读写分离
	*/
	host := viper.GetString(fmt.Sprintf("database.%s.%s.host", sqlType, readWrite))  // yml的使用.取嵌套的值
	user := viper.GetString(fmt.Sprintf("database.%s.%s.user", sqlType, readWrite))
	password := viper.GetString(fmt.Sprintf("database.%s.%s.password", sqlType, readWrite))
	name := viper.GetString(fmt.Sprintf("database.%s.%s.name", sqlType, readWrite))
	port := viper.GetInt(fmt.Sprintf("database.%s.%s.port", sqlType, readWrite))
	charset := viper.GetString(fmt.Sprintf("database.%s.%s.charset", sqlType, readWrite))

	switch strings.ToLower(sqlType) {
	case DriverMysql:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, name, charset)
	case DriverSqlServer:
		return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, name, user, password)
	case DriverPostgresql:
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, name, user, password)
	}
	return ""
}

func GetDb() *gorm.DB {
	return db
}
