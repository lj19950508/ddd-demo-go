package mysql

import "time"

// Option -.
type Option func(*Mysql)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Mysql) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *Mysql) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Mysql) {
		c.connTimeout = timeout
	}
}
