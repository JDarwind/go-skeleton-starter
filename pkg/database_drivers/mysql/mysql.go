package database_drivers

import (
	"database/sql"
	"fmt"

	"github.com/JDarwind/go-skeleton-starter/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct{
	username string
	password string
	host string
	port string 
	database string
	db *sql.DB
	connectionName string
}

func NewMysqlDB(username,password,host,port,database,connectionName string)(*MysqlDB){
	mysql := &MysqlDB{
		username: username,
		password: password,
		host: host,
		port: port,
		database: database,
		db: nil,
		connectionName: connectionName,
	}
	return mysql
}

func (mysql *MysqlDB) Connect() error{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysql.username, mysql.password, mysql.host, mysql.port, mysql.database)
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return err
    }
    if err:= db.Ping(); err != nil{
   		return err
    }
    mysql.db = db
    
    database.GetDatabaseManager().AddDatabaseToList(mysql, mysql.connectionName)
    return nil
}

func (mysql *MysqlDB) Close() error{
	
	if mysql.db == nil{
		return fmt.Errorf("db not initalized")
	}
	
	err:= mysql.db.Close()
	if err!=nil{
		return err
	}
	database.GetDatabaseManager().RemoveDbFromList(mysql.connectionName)
	return nil
}

func (mysql *MysqlDB) GetDriver() *sql.DB{
	return mysql.db
}