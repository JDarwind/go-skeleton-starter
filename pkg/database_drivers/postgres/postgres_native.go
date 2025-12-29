package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPGXNativeDriver struct {
	username       string
	password       string
	host           string
	port           string
	database       string
	pool           *pgxpool.Pool
	connectionName string

	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func NewPostgresPGXNative(config PostgresConfig) *PostgresPGXNativeDriver {
	return &PostgresPGXNativeDriver{
		username:       config.Username,
		password:       config.Password,
		host:           config.Host,
		port:           config.Port,
		database:       config.Database,
		connectionName: config.ConnectionName,
		maxOpenConns:   config.MaxOpenConns,
		maxIdleConns:   config.MaxIdleConns,
		maxIdleTime:    config.MaxIdleTime,
	}
}

func (p *PostgresPGXNativeDriver) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.username, p.password, p.host, p.port, p.database,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}

	if p.maxOpenConns > 0 {
		cfg.MaxConns = int32(p.maxOpenConns)
	}
	if p.maxIdleConns > 0 {
		cfg.MinConns = int32(p.maxIdleConns)
	}
	if p.maxIdleTime > 0 {
		cfg.MaxConnIdleTime = p.maxIdleTime
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return err
	}

	p.pool = pool
	database.GetDatabaseManager().AddDatabaseToList(p, p.connectionName)
	return nil
}

func (p *PostgresPGXNativeDriver) Close() error {
	if p.pool == nil {
		return fmt.Errorf("db not initialized")
	}
	p.pool.Close()
	database.GetDatabaseManager().RemoveDbFromList(p.connectionName)
	return nil
}

func (p *PostgresPGXNativeDriver) GetDriver() *pgxpool.Pool {
	return p.pool
}
