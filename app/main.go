package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Transaction struct {
	UserId		int 	`json:"user_id"`
	Amount		int 	`json:"amount"`
	Category	string 	`json:"category"`
}

type Category struct {
	Name	string
	Count	int
	Sum		int
}

type Report struct {
	UserId		int
	Sum			int
	Categories	[]Category
}

const (
	JSON_TRANSACTION = "transactions.json"
	DIR_APP = "/app"
)

func main()  {

	transaction := ParseJson(JSON_TRANSACTION)

	report := make(map[int]Report)
	panic(report[0].UserId)
	report = append()
	report[0].UserId = transaction[0].UserId
	report[0].Sum = transaction[0].Amount
	report[0].Categories[0].Name = transaction[0].Category
	report[0].Categories[0].Count = 1
	report[0].Categories[0].Sum = transaction[0].Amount

	for i := 1; i < len(transaction); i++ {
		//fmt.Println(transaction[i].UserId)
		for j := 0; j <= len(report); j++ {
			if transaction[i].UserId == report[j].UserId {
				report[j].Sum += transaction[i].Amount
				for a := 0; a <= len(report[j].Categories); a++ {
					if transaction[i].Category == report[j].Categories[a].Name {
						report[j].Categories[a].Count ++
						report[j].Categories[a].Sum += transaction[i].Amount
					} else {
						report[j].Categories[a].Name = transaction[i].Category
						report[j].Categories[a].Count = 1
						report[j].Categories[a].Sum = transaction[i].Amount
					}
				}
			} else {
				report[j].UserId = transaction[i].UserId
				report[j].Sum = transaction[i].Amount
				report[j].Categories[0].Name = transaction[i].Category
				report[j].Categories[0].Count = 1
				report[j].Categories[0].Sum = transaction[i].Amount
			}
		}
	}

	fmt.Println(report)
}

func PrintReport() {
	fmt.Println()
}


func ParseJson(pathFile string) []Transaction {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := os.Chdir(pwd + DIR_APP); err != nil {
		panic(err)
	}

	jsonFile, err := os.Open(JSON_TRANSACTION)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	transaction := make([]Transaction)

	json.Unmarshal([]byte(byteValue), &transaction)

	return transaction
}