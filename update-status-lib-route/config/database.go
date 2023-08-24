package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB()  {
	db, err := sql.Open("mysql", "lib:(Password!123)@tcp(127.0.0.1:4000)/lib?parseTime=true")
	if err != nil{
		panic(err)
	}

	log.Println("Database Connected")

	DB = db
}