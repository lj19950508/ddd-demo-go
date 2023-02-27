package mysql

import (
	"time"

	"gorm.io/gorm"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Mysql struct {
	maxPoolSize   int
	connAttempts int
	connTimeout   time.Duration
	GormDb        *gorm.DB

}

func New(url string,opts ...Option)(*Mysql,error){
	mysql :=&Mysql{}
		// Custom options
	for _, opt := range opts {
		opt(mysql)
	}
	return mysql,nil

}

func (mysql *Mysql) Close(){

}