package card

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Transaction struct {
	UserId int64 `json:"user_id"`
	Sum    int64 `json:"sum"`
	MCC    string `json:"mcc"`
}


func MakeTransactions(userId int64) []Transaction {
	const usersCount = 5
	const transactionsCount = 2
	const transactionAmount = 1_00
	transactions := make([]Transaction, usersCount*transactionsCount)
	x := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5411",
	}
	y := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5812",
	}
	z := Transaction{
		UserId: 2,
		Sum:    transactionAmount,
		MCC:    "5812",
	}

	for index := range transactions {
		switch index % 100 {
		case 0:
			transactions[index] = x
		case 20:
			transactions[index] = y
		default:
			transactions[index] = z
		}
	}
	return transactions
}

func MapRowToTransaction(m [][]string) []Transaction {
	transaction := make([]Transaction, 0)
	for _, v := range m {
		userId, _ := strconv.ParseInt(v[0], 10, 64)
		sum, _ := strconv.ParseInt(v[1], 10, 64)

		tmpTrans := Transaction{
			UserId: userId,
			Sum:    sum,
			MCC:    v[2],
		}

		transaction = append(transaction, tmpTrans)
	}
	return transaction
}

func ExportCSV(filename string, transactions []Transaction) ([]Transaction, error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}(file)

	var records [][]string
	for _, t := range transactions {
		record := []string{
		    strconv.FormatInt(t.UserId, 10),
			strconv.FormatInt(t.Sum, 10),
			TranslateMCC(t.MCC),
		}
		records = append(records, record)
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	recordConv := MapRowToTransaction(records)
	return recordConv, nil
}

func ImportCsv(filename string) ([]Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(err)
			err = cerr
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	recordConv := MapRowToTransaction(records)
	return recordConv, nil
}

func ExportJson(filename string, transactions []Transaction) error {
	encoded, err := json.Marshal(transactions)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ioutil.WriteFile(filename, encoded, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func ImportJSON(filename string) ([]Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
		}
	}(file)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var decoded []Transaction
	err = json.Unmarshal(content, &decoded)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(reflect.TypeOf(decoded), decoded)
	return decoded, nil
}