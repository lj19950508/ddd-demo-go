package mysql

import (
	"database/sql"
	"log"
	"os"
	"time"

	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		}), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout,"",log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
				logger.Config{
				  SlowThreshold: time.Second,   // 慢 SQL 阈值
				  LogLevel:      logger.Info, // 日志级别
				  IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
				  Colorful:      false,         // 禁用彩色打印
				},
			  ),
		})
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
