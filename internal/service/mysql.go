package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"xanthing/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlSqlList = map[string]*sql.DB{}

func InitMysql(name string) {
	conf := config.GetConfig(name).(map[string]any)
	dsnT := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnT, []any{
		conf["username"],
		conf["password"],
		conf["hostname"],
		conf["port"],
		conf["database"],
		conf["charset"],
	}...)
	sqlDB, err := sql.Open("mysql", dsn)

	if err != nil {
		panic("failed to connect mysql " + err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping " + err.Error())
	}

	MysqlSqlList[name] = sqlDB
}

func GetDb(name string) (*gorm.DB, error) {
	sqlDB, ok := MysqlSqlList[name]
	if !ok {
		return nil, errors.New("该db不存在")
	}
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return db, err
}
