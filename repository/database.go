package repository

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	return db
}
