package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Opendb() {
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_database := os.Getenv("MYSQL_DATABASE")
	container_database_ports := os.Getenv("CONTAINER_DATABASE_PORTS")
	db_name, err := sql.Open("mysql", mysql_user+":"+mysql_password+"@tcp(mysql_practice:"+container_database_ports+")/"+mysql_database+"?")

	if err != nil {
		panic(err.Error())
	}

	db = db_name
}

func Register(username string, password string) {

	Opendb()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Students(username,password) VALUES(?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	insert.Exec(username, password)
}
