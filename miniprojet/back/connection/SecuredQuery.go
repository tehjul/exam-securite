package connection

import (
	"dataacces"
	"database/sql"
	"log"
)

func (query securedQuery) GetResult() *sql.Rows {
	rows, err := dataacces.ExecPreparedQuery("SELECT ID, USER, PASSWORD FROM USER WHERE USER = ? AND PASSWORD = ?", query.info.User, query.info.Password)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
