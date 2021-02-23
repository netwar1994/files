package main

import (
	"github.com/netwar1994/files/pkg/card"
	"log"
)

func main()  {
	fileCSV := "test.csv"
	fileJSON := "test.json"
	impCSV, err := card.ImportCsv(fileCSV)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Imported", len(impCSV), "rows from", fileCSV)

	impJSON, err := card.ImportJSON(fileJSON)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Imported", len(impJSON), "rows from", fileJSON)
}
