package main

import (
	"github.com/netwar1994/files/pkg/card"
	"log"
)

func main()  {
	filename := "test.csv"
	trans, err := card.ImportCsv(filename)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Imported", len(trans), "rows from", filename)
}
