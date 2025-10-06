package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DRIVER   = "mysql"
	DBNAME   = "apicrud"
	USER     = "root"
	PASSWORD = "root"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", USER, PASSWORD, DBNAME)
	fmt.Println("CONECTADO NA DATABASE")
	con, err := sql.Open(DRIVER, URL)
	if err != nil {
		fmt.Println("NÃO CONECTADO NA DATABASE")
		log.Fatal(err)
		return nil
	}
	return con
}
