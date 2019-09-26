package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("EXPERIMENTAL CHANGEFEED FOR users")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = rows.Close()
	}()
	type changeFeed struct {
		table string
		key   string
		value []byte
	}
	for rows.Next() {
		changeFeed := &changeFeed{}
		err := rows.Scan(&changeFeed.table, &changeFeed.key, &changeFeed.value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(changeFeed.value))
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
