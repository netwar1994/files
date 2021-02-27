package main

import (
	"github.com/netwar1994/files/pkg/card"
	"log"
)

func main()  {
	transactions := card.MakeTransactions(1)
	//transactionsXML := []card.Transactions{{
	//	XMLName:      "transactions",
	//	Transactions: transactions,
	//}}
	fileCSV := "test.csv"
	fileJSON := "test.json"
	fileXML := "test.xml"

	expCSV, err := card.ExportCSV(fileCSV, transactions)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Exported", len(expCSV), "rows to", fileCSV)

	err = card.ExportJson(fileJSON, transactions)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Exported to", fileJSON)

	err = card.ExportXML(fileXML, transactions)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Exported to", fileXML)
}
