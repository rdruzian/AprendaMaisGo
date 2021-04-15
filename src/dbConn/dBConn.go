package dbConn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dBConn() (db *sql.DB){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "AprendaMais"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}