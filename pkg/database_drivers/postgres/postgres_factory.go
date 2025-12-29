package postgres

import (
	"fmt"
	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	"time"
)

type PostgresDriverType string

const (
	PostgresPGXNative  PostgresDriverType = "pgx_native"
	PostgresPGXStdlib  PostgresDriverType = "pgx_stdlib"
)

const (
	defaultMaxOpenConns = 25
	defaultMaxIdleConns = 5
	defaultMaxIdleTime  = 5 * time.Minute
)

type PostgresConfig struct {
	Username       string
	Password       string
	Host           string
	Port           string
	Database       string
	ConnectionName string
	DriverType     PostgresDriverType

	MaxOpenConns     int           
	MaxIdleConns     int           
	MaxIdleTime      time.Duration 
}

func normalizeConfig(config *PostgresConfig) {
	if config.MaxOpenConns <= 0 {
		config.MaxOpenConns = defaultMaxOpenConns
	}

	if config.MaxIdleConns <= 0 {
		config.MaxIdleConns = defaultMaxIdleConns
	}

	if config.MaxIdleTime <= 0 {
		config.MaxIdleTime = defaultMaxIdleTime
	}

	if config.MaxIdleConns > config.MaxOpenConns {
		config.MaxIdleConns = config.MaxOpenConns -1
	}
}

func NewPostgresDriver(driverType PostgresDriverType, cfg PostgresConfig) (database.Database, error) {
	normalizeConfig(&cfg)
	switch driverType {
		case PostgresPGXNative:
			return NewPostgresPGXNative(cfg), nil
		case PostgresPGXStdlib:
			return NewPostgresPGXStdlib(cfg), nil
		default:
			return nil, fmt.Errorf("unsupported postgres driver: %s", driverType)
	}
}