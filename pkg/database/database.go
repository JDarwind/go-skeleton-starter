package database

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
