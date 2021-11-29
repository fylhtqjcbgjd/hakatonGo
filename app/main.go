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
	CategoryName	string 	`json:"category"`
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
//var categories := make([]Category)
//	categories
	report := make(map[int]Report)
	//report[transaction[0].UserId] = Report{
	//	transaction[0].UserId,
	//	transaction[0].Amount,
	//	[]Category{
	//		Category{
	//			transaction[0].CategoryName,
	//			1,
	//			transaction[0].Amount,
	//		},
	//	},
	//}
	for i := 0; i < len(transaction); i++ {
		//fmt.Println(transaction[i].UserId)
		if elemReport, ok := report[transaction[i].UserId]; ok {
			elemReport.Sum += transaction[i].Amount
			fmt.Println(elemReport)
			for j := 0; j < len(elemReport.Categories); j++ {
				if  elemReport.Categories[j].Name == transaction[i].CategoryName {
					elemReport.Categories[j].Count += 1
					elemReport.Categories[j].Sum += transaction[i].Amount
				} else {
					elemReport.Categories = append(
						elemReport.Categories,
						Category{
							transaction[i].CategoryName,
							1,
							transaction[i].Amount,
						})
				}
			}
			//fmt.Println(elemReport)
			report[transaction[i].UserId] = elemReport
		} else {
			report[transaction[i].UserId] = Report{
				transaction[i].UserId,
				transaction[i].Amount,
				[]Category{
					Category{
						transaction[i].CategoryName,
						1,
						transaction[i].Amount,
					},
				},
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
	transaction := []Transaction{}

	json.Unmarshal([]byte(byteValue), &transaction)

	return transaction
}