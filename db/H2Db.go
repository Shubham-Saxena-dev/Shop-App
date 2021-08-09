package db

import (
	"database/sql"
	"log"
)

var conn *sql.DB

func InitializeDatabase(){
	conn, err := sql.Open("h2", "h2://sa@localhost/testdb?mem=true")

	if err !=nil{
		log.Fatal("Unable to connect to H2 Database: {}", err)
	}
	err = conn.Ping()

	if err != nil {
		log.Fatalf("Ping to H2 Database failed: %s", err)
	}
}

func getConnection() *sql.DB {
	return conn
}

func closeConnection() {
	err:=conn.Close()
	if err != nil {
		log.Fatalf("Unable to close connection to H2 Database: %s", err)
	}
}