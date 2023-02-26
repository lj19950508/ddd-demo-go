package orm

import (
	// "time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initOrm()(*gorm.DB,error){
	return gorm.Open(mysql.New(mysql.Config{
		DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	  }), &gorm.Config{})
// 	  sqlDB, err := db.DB()
// 	  if(err!=nil){
// 		  panic(err)
// 	  }
//   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// 	  sqlDB.SetMaxIdleConns(10)
  
// 	  // SetMaxOpenConns sets the maximum number of open connections to the database.
// 	  sqlDB.SetMaxOpenConns(100)
  
// 	  // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// 	  sqlDB.SetConnMaxLifetime(time.Hour)

}

//只创建一次
func StartOrm(){
	//datasource

	initOrm()


	//使用sql.db维护连接池

}