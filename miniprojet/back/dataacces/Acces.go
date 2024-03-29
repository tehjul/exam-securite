package dataacces

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dataSource = "./sqlite-database.db"
	driver     = "sqlite3"
)

func UpdateData(query string, args ...interface{}) error {
	db, _ := sql.Open(driver, dataSource)
	defer db.Close()
	_, err := db.Exec(query, args...)
	return err
}

func ExecQuery(query string, args ...interface{}) (*sql.Rows, error) {
	db, _ := sql.Open(driver, dataSource)
	defer db.Close()
	return db.Query(query, args...)
}

func ExecPreparedQuery(query string, user string, password string) (*sql.Rows, error) {
	db, _ := sql.Open(driver, dataSource)
	defer db.Close()
	statement, error := db.Prepare(query)
	if error != nil {
		return nil, error
	}
	return statement.Query(user, password)
}
