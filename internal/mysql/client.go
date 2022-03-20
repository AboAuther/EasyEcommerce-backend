package mysql

import (
	"EasyEcommerce-backend/internal/utils"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func MysqlDSNFromENV(prefix string, params ...string) string {
	dsn := utils.MustGetenv("DSN")
	if len(params) > 0 {
		dsn = dsn + "?" + strings.Join(params, "&")
	}
	return dsn
}

func InitDB() error {
	var err error
	dsn := MysqlDSNFromENV("", "parseTime=true")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return err
	}

	if sqlDB, err := DB.DB(); err != nil {
		return err
	} else {
		sqlDB.SetMaxIdleConns(10)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)
		return nil
	}
}
