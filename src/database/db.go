package database

import (
	"fmt"
	"momonga_blog/config"
	"momonga_blog/internal/logging"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    db   *gorm.DB
    once sync.Once
)


func New() (*gorm.DB, error) {
	var e error = nil
	once.Do(func() {
		cnf, err := config.GetConfig()
		if err != nil {
			e = err
			return
		}

		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
			cnf.DbUser, cnf.DbPassword, cnf.DbHost, cnf.DbPort, cnf.DbName,
		)
		logging.AppLogger.Info("DB接続情報", "dns", dns)
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			e = err
			return
		}
		sqlDb, err := db.DB()
		if err != nil {
			e = err
			return
		}

		// アイドル接続プールの最大接続数を設定
		sqlDb.SetMaxIdleConns(cnf.DbMaxIdleConn)
		// オープン接続の最大数を設定
		sqlDb.SetMaxOpenConns(cnf.DbMaxOpenConn)
		// 接続が再使用される最大時間を設定
		sqlDb.SetConnMaxLifetime(time.Duration(cnf.DbConnMaxLifetime) * time.Hour)
	})

	return db, e
}

func HealthCheck() error {
	if db == nil {
		return fmt.Errorf("db is not connected")
	}

	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDb.Ping()
}


func Close() error {
	db, err := db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func GetDB()(*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("db is not connected")
	}
	return db, nil
}