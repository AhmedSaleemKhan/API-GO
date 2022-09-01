package controller

import (
	"fmt"
	"log" //using sqlx library

	"github.com/jmoiron/sqlx"
)

//url of databse for connection
const url = "postgres://myuser2:123@localhost:5432/mydb"

//connectioon start
type DBConection struct {
	Conn *sqlx.DB
}

//function for connection
func ConnectDB() *DBConection {
	db, err := sqlx.Connect("postgres", url) //connecting to database
	if err != nil {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
				log.Println("Connect Unsuccessed")
			}
		}()

	}
	return &DBConection{
		Conn: db,
	}
}
