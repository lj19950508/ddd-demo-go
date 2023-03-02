package mysql

import "time"

//TODO 优化config option
// Option -.
type Option func(*Mysql)

// MaxPoolSize -.
func MaxIdleConns(size int) Option {
	return func(c *Mysql) {
		c.MaxIdleConns = size
	}
}

// ConnAttempts -.
func MaxOpenConns(attempts int) Option {
	return func(c *Mysql) {
		c.MaxOpenConns = attempts
	}
}

// ConnTimeout -.
func ConnMaxLifetime(timeout time.Duration) Option {
	return func(c *Mysql) {
		c.ConnMaxLifetime = timeout
	}
}
