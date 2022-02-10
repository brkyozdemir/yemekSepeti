package main

import (
	"github.com/tidwall/buntdb"
	"log"
)

func main() {
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
