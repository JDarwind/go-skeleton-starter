package database

import "fmt"

type Database interface {
	Connect() error
	Close()  error
}

type DatabaseManager struct{
	databaseConnectionsMap map[string]Database
}

var databaseManager *DatabaseManager

func GetDatabaseManager() *DatabaseManager{
	if databaseManager == nil{
		databaseManager = &DatabaseManager{
			databaseConnectionsMap: map[string]Database{
				"default": nil,
			},
		}
	}
	
	return databaseManager
}

func (dbManager *DatabaseManager) AddDatabaseToList(db Database, name string) (*DatabaseManager){
	dbManager.databaseConnectionsMap[name] = db
	return dbManager
}

func (dbManager *DatabaseManager) RemoveDbFromList(name string) (*DatabaseManager){
	delete(dbManager.databaseConnectionsMap, name)
	return dbManager
}


func (dbManager *DatabaseManager) GetDatabaseConnection(name string) (Database, error) {
	if name == "" {
		name = "default"
	}

	db, ok := dbManager.databaseConnectionsMap[name]
	if !ok {
		return nil, fmt.Errorf("database connection with name %s not found", name)
	}
	
	return db, nil
}