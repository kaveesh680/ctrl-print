package db_connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
)

var DBConnection *sql.DB

func InitConnection() {

	connStr := "user=postgres password=123 dbname=ctrl sslmode=disable"
	var err error

	DBConnection, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	go func() {

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		select {
		case <-signals:
			fmt.Println("Close DB Connection!")
			DBConnection.Close()
			return
		}
	}()

	err = DBConnection.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL database!")

}
