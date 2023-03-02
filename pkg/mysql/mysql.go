package mysql

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/lj19950508/ddd-demo-go/config"
	"go.uber.org/fx"
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

func New(lc fx.Lifecycle, cfg *config.Config) (mysql *Mysql) {
	mysql = &Mysql{
		MaxIdleConns:    _defaultMaxIdleConns,
		MaxOpenConns:    _defaultMaxOpenConns,
		ConnMaxLifetime: _defaultConnMaxLifetime,
	}
	lc.Append(fx.Hook{
		//被需要的时候只会执行一次
		OnStart: func(ctx context.Context) error {
			gormDb, err := gorm.Open(
				driver.New(driver.Config{
					DSN: cfg.Mysql.Url,
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
			mysql.GormDb = gormDb
			sqlDB, err := gormDb.DB()
			if err != nil {
				return err
			}
			mysql.SqlDb = sqlDB
			sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
			sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(mysql.ConnMaxLifetime)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			mysql.SqlDb.Close()
			return nil
		},
	})
	return mysql
}

func (mysql *Mysql) Close() {
	mysql.SqlDb.Close()
}
