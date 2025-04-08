package connection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8163)/api")
	if err != nil {
		return nil, err
	}
	return db, nil
}
