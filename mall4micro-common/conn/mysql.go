package conn

import (
	"fmt"
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//
// Conn
// @Description: 连接到mysql数据库
// @return *gorm.DB
// @return error
//
func Conn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Settings.Mysql.User, conf.Settings.Mysql.Password, conf.Settings.Mysql.Host, conf.Settings.Mysql.Port, conf.Settings.Mysql.Database)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.Settings.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.Settings.Mysql.MaxConn)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	return db, nil
}
