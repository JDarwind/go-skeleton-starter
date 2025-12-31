package mysql

import "time"


type MysqlConfig struct {
	Username       string
	Password       string
	Host           string
	Port           int
	Database       string
	ConnectionName string

	MaxOpenConns int
	MaxIdleConns int
	ConnMaxIdle  time.Duration
	ConnMaxLife  time.Duration
}

const (
	defaultMysqlMaxOpenConns = 25
	defaultMysqlMaxIdleConns = 5
	defaultMysqlIdleTime     = 5 * time.Minute
	defaultMysqlLifeTime     = 30 * time.Minute
)