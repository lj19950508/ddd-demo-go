package middleware

import (
	"gorm.io/gorm"
	"time"
)
import "gorm.io/driver/mysql"

type GormResource struct {
	db *gorm.DB
}

// NewGormResource 返回数据源
func NewGormResource() *GormResource {

	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDb, _ := db.DB()
	//defer sqlDb.Close()
	//if err !=nil{
	//	panic("faild to connect database")
	//}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return &GormResource{
		db: db,
	}
}

func (this GormResource) DB() *gorm.DB {
	return this.db
}
