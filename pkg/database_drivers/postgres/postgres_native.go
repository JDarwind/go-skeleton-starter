package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPGXNativeDriver struct {
	postgresConfig PostgresConfig
	pool           *pgxpool.Pool
 	mu 				sync.Mutex

}

func NewPostgresPGXNative(config PostgresConfig) *PostgresPGXNativeDriver {
	return &PostgresPGXNativeDriver{
		postgresConfig: config,
	}
}

func (p *PostgresPGXNativeDriver) Connect() error {
	
	p.mu.Lock()
    defer p.mu.Unlock()
	
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		p.postgresConfig.Username, 
		p.postgresConfig.Password, 
		p.postgresConfig.Host, 
		p.postgresConfig.Port, 
		p.postgresConfig.Database,
		*p.postgresConfig.SslMode,
		p.postgresConfig.Schema,
	)

	ctx, cancel := context.WithTimeout(context.Background(), p.postgresConfig.MaxConnnectionTimeout)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}
	
	cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheStatement
	cfg.ConnConfig.StatementCacheCapacity = 512

	if p.postgresConfig.MaxOpenConns > 0 {
		cfg.MaxConns = int32(p.postgresConfig.MaxOpenConns)
	}
	if p.postgresConfig.MaxIdleConns > 0 {
		cfg.MinConns = int32(p.postgresConfig.MaxIdleConns)
	}
	if p.postgresConfig.MaxIdleTime > 0 {
		cfg.MaxConnIdleTime = p.postgresConfig.MaxIdleTime
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
	if _,err := database.GetDatabaseManager().AddDatabaseToList(p, p.postgresConfig.ConnectionName); err != nil {
    	p.pool.Close()
    	return err
	}
	return nil
}

func (p *PostgresPGXNativeDriver) Close() error {
	
	if p.pool == nil {
		return fmt.Errorf("db not initialized")
	}
	
	p.mu.Lock()
    defer p.mu.Unlock()
	
	p.pool.Close()
	database.GetDatabaseManager().RemoveDbFromList(p.postgresConfig.ConnectionName)
	return nil
}

func (p *PostgresPGXNativeDriver) GetDriver() *pgxpool.Pool {
	return p.pool
}


func (p *PostgresPGXNativeDriver) IsConnected() bool{
	return p.pool != nil
}