package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
	loadErr error
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type TestConfig struct {
	Postgres Postgres
	Mysql    Mysql
}

var cfg TestConfig

func Load() (TestConfig, error) {
	once.Do(func() {
		wd, err := os.Getwd()
		if err != nil {
			loadErr = fmt.Errorf("failed to get working directory: %w", err)
			return
		}

		var envPath string
		dir := wd

		for {
			candidate := filepath.Join(dir, "tests", ".env.test")
			if _, err := os.Stat(candidate); err == nil {
				envPath = candidate
				break
			}

			parent := filepath.Dir(dir)
			if parent == dir {
				loadErr = fmt.Errorf("could not find tests/.env.test starting from %s", wd)
				return
			}
			dir = parent
		}

		if err := godotenv.Load(envPath); err != nil {
			loadErr = fmt.Errorf("failed to load test env file (%s): %w", envPath, err)
			return
		}

		cfg.Postgres = Postgres{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			Schema:   os.Getenv("POSTGRES_SCHEMA"),
		}

		if cfg.Postgres.Host == "" ||
			cfg.Postgres.Port == "" ||
			cfg.Postgres.User == "" ||
			cfg.Postgres.Database == "" {
			loadErr = fmt.Errorf("postgres test configuration incomplete")
			return
		}

		mysqlPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
		if err != nil {
			loadErr = fmt.Errorf("invalid MYSQL_PORT: %w", err)
			return
		}

		cfg.Mysql = Mysql{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     mysqlPort,
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DB"),
		}

		if cfg.Mysql.Host == "" ||
			cfg.Mysql.User == "" ||
			cfg.Mysql.Database == "" {
			loadErr = fmt.Errorf("mysql test configuration incomplete")
			return
		}
	})

	return cfg, loadErr
}
