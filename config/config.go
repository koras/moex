package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ProductModel struct {
	Db *sql.DB
}

/**
* connect in base
 */
func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("USER")
	dbPass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
