package postgres

import "time"

type SSLMode string

const (
	SSLDisable   SSLMode = "disable"
	SSLRequire   SSLMode = "require"
	SSLVerifyCa  SSLMode = "verify-ca"
	SSLVerifyAll SSLMode = "verify-full"
)

type PostgresDriverType string

const (
	PostgresPGXNative PostgresDriverType = "pgx_native"
	PostgresPGXStdlib PostgresDriverType = "pgx_stdlib"
)

const (
	defaultSchema string = "public"
)

const (
	defaultMaxOpenConns = 25
	defaultMaxIdleConns = 5
	defaultMaxIdleTime  = 5 * time.Minute
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Schema   string
	SslMode  *SSLMode

	ConnectionName string
	DriverType     PostgresDriverType

	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration

	MaxConnnectionTimeout time.Duration //Supported only on potgress native mode
}

func normalizeConfig(config *PostgresConfig) {

	if config.Schema == "" {
		config.Schema = defaultSchema
	}

	if config.SslMode == nil {
		sslDisabled := SSLDisable

		config.SslMode = &sslDisabled
	}

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
		config.MaxIdleConns = config.MaxOpenConns - 1
	}

	if config.MaxConnnectionTimeout <= 0 {
		config.MaxConnnectionTimeout = 20 * time.Second
	}
}
