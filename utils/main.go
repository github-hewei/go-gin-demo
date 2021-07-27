package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Db() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:13301)/test")
	return
}
