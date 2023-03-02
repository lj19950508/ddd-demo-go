package db

import "time"

// Option -.
type Option func(*DB)

// MaxPoolSize -.
func MaxIdleConns(size int) Option {
	return func(c *DB) {
		c.MaxIdleConns = size
	}
}

// ConnAttempts -.
func MaxOpenConns(attempts int) Option {
	return func(c *DB) {
		c.MaxOpenConns = attempts
	}
}

// ConnTimeout -.
func ConnMaxLifetime(timeout time.Duration) Option {
	return func(c *DB) {
		c.ConnMaxLifetime = timeout
	}
}
