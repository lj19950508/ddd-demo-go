package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_defaultMaxIdleConns    = 10
	_defaultMaxOpenConns    = 100
	_defaultConnMaxLifetime = time.Hour
)

//DB 由发言和orm组成


//方言

//orm

type DB struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	GormDb          *gorm.DB
	SqlDb           *sql.DB
	dialector       gorm.Dialector
}

func New(url string, opts ...Option) (db *DB) {
	db = &DB{
		MaxIdleConns:    _defaultMaxIdleConns,
		MaxOpenConns:    _defaultMaxOpenConns,
		ConnMaxLifetime: _defaultConnMaxLifetime,
	}
	db.dialector= mysql.New(mysql.Config{
		DSN: url,
	})

	for _, opt := range opts {
		opt(db)
	}

	return db
}

func (t *DB) Close() {
	t.SqlDb.Close()
}

func (t *DB) Open() error {
	
	gormDB, err := gorm.Open(t.dialector,&gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		),
	},)
	
	if err != nil {
		return nil
	}
	t.GormDb = gormDB
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(t.MaxIdleConns)
	sqlDB.SetMaxOpenConns(t.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(t.ConnMaxLifetime)
	t.SqlDb=sqlDB
	return nil
}
