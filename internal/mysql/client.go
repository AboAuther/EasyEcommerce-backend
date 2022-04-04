package mysql

import (
	"EasyEcommerce-backend/internal/utils"
	"errors"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func mysqlDSNFromENV(prefix string, params ...string) string {
	dsn := utils.MustGetenv("DSN")
	if len(params) > 0 {
		dsn = dsn + "?" + strings.Join(params, "&")
	}
	return dsn
}

func InitDB() error {
	var err error
	dsn := mysqlDSNFromENV("", "parseTime=true")
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

func IsMissing(tx *gorm.DB) bool {
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return true
	} else if tx.Error == nil {
		return false
	} else {
		panic(tx.Error)
	}
}
