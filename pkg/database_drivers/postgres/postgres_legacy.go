package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresPGXStdDriver struct {
	postgresConfig PostgresConfig
	db             *sql.DB
	mu 			   sync.Mutex
}

func NewPostgresPGXStdlib(config PostgresConfig) *PostgresPGXStdDriver {
	return &PostgresPGXStdDriver{
		postgresConfig: config,
	}
}

func (p *PostgresPGXStdDriver) Connect() error {
	
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

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	if p.postgresConfig.MaxOpenConns > 0 {
		db.SetMaxOpenConns(p.postgresConfig.MaxOpenConns)
	}
	if p.postgresConfig.MaxIdleConns > 0 {
		db.SetMaxIdleConns(p.postgresConfig.MaxIdleConns)
	}
	if p.postgresConfig.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(p.postgresConfig.MaxIdleTime)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return err
	}

	p.db = db
	if _,err := database.GetDatabaseManager().AddDatabaseToList(p, p.postgresConfig.ConnectionName); err != nil {
    	p.db.Close()
    	return err
	}
	
	return nil
}

func (p *PostgresPGXStdDriver) Close() error {
	if p.db == nil {
		return fmt.Errorf("db not initialized")
	}
	
	p.mu.Lock()
    defer p.mu.Unlock()
	if err := p.db.Close(); err != nil {
		return err
	}
	database.GetDatabaseManager().RemoveDbFromList(p.postgresConfig.ConnectionName)
	return nil
}

func (p *PostgresPGXStdDriver) GetDriver() *sql.DB {
	return p.db
}

func (p *PostgresPGXStdDriver) IsConnected() bool{
	return p.db != nil
}