package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConection() (*sql.DB, error) {
	db, err := sql.Open("mqsql", "root:Sri@7272@tcp(127.0.0.1:3306)/golangdb")
	if err != nil {
		fmt.Println("Error in connectiong the DB", err)
		return nil, err
	}
	fmt.Println("Database is connected successfully")
	return db, nil
}
