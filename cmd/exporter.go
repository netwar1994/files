package main

import (
	"github.com/netwar1994/files/pkg/card"
	"log"
)

func main()  {
	transactions := card.MakeTransactions(1)
	filename := "test.csv"
	trans, err := card.ExportCSV(filename, transactions)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Exported", len(trans), "rows to", filename)
}
