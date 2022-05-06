package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"help_center/internal/conf"
	"log"
)

var db *gorm.DB

func init() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetCfg().Mysql.User,
		conf.GetCfg().Mysql.Password,
		conf.GetCfg().Mysql.Host,
		conf.GetCfg().Mysql.Port,
		conf.GetCfg().Mysql.Database)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	tmp, err := db.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	tmp.SetMaxIdleConns(conf.GetCfg().Mysql.MaxIdleConn) //设置最大连接数
	tmp.SetMaxOpenConns(conf.GetCfg().Mysql.MaxOpenConn) //设置最大的空闲连接数
}
func GetDbCli() *gorm.DB {
	return db
}
