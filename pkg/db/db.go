package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"
	"go.uber.org/fx"
	driver "gorm.io/driver/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_defaultMaxIdleConns    = 10
	_defaultMaxOpenConns    = 100
	_defaultConnMaxLifetime = time.Hour
)

type DB struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	GormDb          *gorm.DB
	SqlDb           *sql.DB
}

func New(opts ...Option) (db *DB) {
	//TODO --
	db = &DB{
		MaxIdleConns:    _defaultMaxIdleConns,
		MaxOpenConns:    _defaultMaxOpenConns,
		ConnMaxLifetime: _defaultConnMaxLifetime,
	}
	for _, opt := range opts {
		opt(db)
	}
	

	lc.Append(fx.Hook{
		//被需要的时候只会执行一次
		OnStart: func(ctx context.Context) error {
			gormDb, err := gorm.Open(
				driver.New(driver.Config{
					DSN: cfgDB.Url,
				}), &gorm.Config{
					Logger: logger.New(
						log.New(os.Stdout, "", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
						logger.Config{
							SlowThreshold:             time.Second, // 慢 SQL 阈值
							LogLevel:                  logger.Info, // 日志级别
							IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
							Colorful:                  false,       // 禁用彩色打印
						},
					),
				})
			if err != nil {
				return err
			}
			db.GormDb = gormDb
			sqlDB, err := gormDb.DB()
			if err != nil {
				return err
			}
			db.SqlDb = sqlDB
			sqlDB.SetMaxIdleConns(db.MaxIdleConns)
			sqlDB.SetMaxOpenConns(db.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(db.ConnMaxLifetime)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			db.SqlDb.Close()
			return nil
		},
	})
	return db
}

func (db  DB) Close() {
	db.SqlDb.Close()
}
