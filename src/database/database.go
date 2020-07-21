package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "tcd_amazon"
)

func Connection(query string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var QueryResultShow string

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	queryResult, err := db.Query(query)

	for queryResult.Next() {
		var name string
		if err := queryResult.Scan(&name); err != nil {
			log.Fatal(err)
		}
		QueryResultShow = fmt.Sprintf("%s\n", name)
	}
	if err := queryResult.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return QueryResultShow
}
