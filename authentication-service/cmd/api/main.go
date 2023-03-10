package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("starting authentication service")

	//Connect to database
	conn := connectToDb()
	if conn == nil {
		log.Fatalf("can't connect to postgres")
	}

	//Setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error while starting service: %s", err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func connectToDb() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready..")
			count++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds..")

		//Chilling for 2 seconds before trying the forloop again
		time.Sleep(2 * time.Second)
		continue

	}

}
