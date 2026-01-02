package database

import (
	"fmt"
	"sync"
)

type Database interface {
	Connect() error
	Close() error
	IsConnected() bool
}

type DatabaseManager struct {
	mu                     sync.RWMutex
	databaseConnectionsMap map[string]Database
}

var (
	databaseManager *DatabaseManager
	once            sync.Once
)

func GetDatabaseManager() *DatabaseManager {
	once.Do(func() {
		databaseManager = &DatabaseManager{
			databaseConnectionsMap: map[string]Database{},
		}
	})

	return databaseManager
}

func (dbManager *DatabaseManager) AddDatabaseToList(db Database, name string) (*DatabaseManager, error) {

	if err := validateConnectionName(name); err != nil {
		return dbManager, err
	}

	dbManager.mu.Lock()
	defer dbManager.mu.Unlock()

	if _, exists := dbManager.databaseConnectionsMap[name]; exists {
		return dbManager, fmt.Errorf("database %s already registered", name)
	}

	if !db.IsConnected() {
		return nil, fmt.Errorf("database %s not connected", name)
	}

	dbManager.databaseConnectionsMap[name] = db
	return dbManager, nil
}

func validateConnectionName(name string) error {
	if name == "" {
		return fmt.Errorf("connection name cannot be empty")
	}
	return nil
}

func (dbManager *DatabaseManager) RemoveDbFromList(name string) *DatabaseManager {
	dbManager.mu.Lock()
	defer dbManager.mu.Unlock()
	delete(dbManager.databaseConnectionsMap, name)
	return dbManager
}

func (dbManager *DatabaseManager) GetDatabaseConnection(name string) (Database, error) {
	if name == "" {
		name = "default"
	}

	dbManager.mu.RLock()
	defer dbManager.mu.RUnlock()

	db, ok := dbManager.databaseConnectionsMap[name]
	if !ok {
		return nil, fmt.Errorf("database connection with name %s not found", name)
	}

	return db, nil
}
