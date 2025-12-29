package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresPGXStdDriver struct {
	username       string
	password       string
	host           string
	port           string
	database       string
	db             *sql.DB
	connectionName string

	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func NewPostgresPGXStdlib(config PostgresConfig) *PostgresPGXStdDriver {
	return &PostgresPGXStdDriver{
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

func (p *PostgresPGXStdDriver) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.username, p.password, p.host, p.port, p.database,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	if p.maxOpenConns > 0 {
		db.SetMaxOpenConns(p.maxOpenConns)
	}
	if p.maxIdleConns > 0 {
		db.SetMaxIdleConns(p.maxIdleConns)
	}
	if p.maxIdleTime > 0 {
		db.SetConnMaxIdleTime(p.maxIdleTime)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return err
	}

	p.db = db
	database.GetDatabaseManager().AddDatabaseToList(p, p.connectionName)
	return nil
}

func (p *PostgresPGXStdDriver) Close() error {
	if p.db == nil {
		return fmt.Errorf("db not initialized")
	}
	if err := p.db.Close(); err != nil {
		return err
	}
	database.GetDatabaseManager().RemoveDbFromList(p.connectionName)
	return nil
}

func (p *PostgresPGXStdDriver) GetDriver() *sql.DB {
	return p.db
}
