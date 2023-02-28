package mysql

import (
	"database/sql"
	"time"

	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	_defaultMaxIdleConns    = 10
	_defaultMaxOpenConns    = 100
	_defaultConnMaxLifetime = time.Hour
)

type Mysql struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	GormDb          *gorm.DB
	SqlDb           *sql.DB
}

func New(url string, opts ...Option) (mysql *Mysql, err error) {
	mysql = &Mysql{
		MaxIdleConns:    _defaultMaxIdleConns,
		MaxOpenConns:    _defaultMaxOpenConns,
		ConnMaxLifetime: _defaultConnMaxLifetime,
	}
	// Custom options
	for _, opt := range opts {
		opt(mysql)
	}

	gormDb, err := gorm.Open(
		driver.New(driver.Config{
			DSN:url,
		}), &gorm.Config{})
	if err != nil {
		return
	}
	mysql.GormDb = gormDb

	sqlDB, err := gormDb.DB()
	if err != nil {
		return
	}
	mysql.SqlDb = sqlDB
	sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(mysql.ConnMaxLifetime)

	return mysql, nil

}

func (mysql *Mysql) Close() {
	mysql.SqlDb.Close()
}
