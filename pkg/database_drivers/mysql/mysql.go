package mysql

import (
	"database/sql"
	"fmt"
	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MysqlDriver struct {
	mu  sync.Mutex
	cfg MysqlConfig
	db  *sql.DB
}

func NewMysqlDriver(cfg MysqlConfig) *MysqlDriver {
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = defaultMysqlMaxOpenConns
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = defaultMysqlMaxIdleConns
	}
	if cfg.ConnMaxIdle <= 0 {
		cfg.ConnMaxIdle = defaultMysqlIdleTime
	}
	if cfg.ConnMaxLife <= 0 {
		cfg.ConnMaxLife = defaultMysqlLifeTime
	}

	return &MysqlDriver{cfg: cfg}
}

func (m *MysqlDriver) Connect() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.db != nil {
		return fmt.Errorf("mysql database already connected")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&interpolateParams=false",
		m.cfg.Username,
		m.cfg.Password,
		m.cfg.Host,
		m.cfg.Port,
		m.cfg.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(m.cfg.MaxOpenConns)
	db.SetMaxIdleConns(m.cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(m.cfg.ConnMaxIdle)
	db.SetConnMaxLifetime(m.cfg.ConnMaxLife)

	if err := db.Ping(); err != nil {
		db.Close()
		return err
	}

	if _, err := database.GetDatabaseManager().AddDatabaseToList(m, m.cfg.ConnectionName); err != nil {
		db.Close()
		return err
	}

	m.db = db
	return nil
}

func (m *MysqlDriver) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.db == nil {
		return fmt.Errorf("mysql database not initialized")
	}

	if err := m.db.Close(); err != nil {
		return err
	}

	m.db = nil
	database.GetDatabaseManager().RemoveDbFromList(m.cfg.ConnectionName)
	return nil
}

func (m *MysqlDriver) IsConnected() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.db != nil
}

func (m *MysqlDriver) GetDriver() *sql.DB {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.db
}
