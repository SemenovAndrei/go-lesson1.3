package main

import (
	"fmt"
	"github.com/i-hit/go-lesson1.3.git/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Id : 1,
		Issuer : "MasterCard",
		Balance : 100_000_00,
		Currency : "RUB",
		Number : "0001",
		Transactions: []*card.Transaction{
			&card.Transaction{
				Id: "1",
				Amount: 1,
				Date: time.Now().Unix(),
				Mcc: "5011",
				Status: "В обработке",
			},
			&card.Transaction{
				Id: "2",
				Amount: 10,
				Date: time.Now().Unix(),
				Mcc: "5012",
				Status: "В обработке",
			},
		},
	}
	master.Transactions = append(master.Transactions,
		&card.Transaction{"3", 100, time.Now().Unix(), "5011", "В обработке"})
	card.AddTransaktion(master,
		&card.Transaction{"4", 1000, time.Now().Unix(), "5011", "Выполнено"})


	fmt.Println(master)
	fmt.Println(master.Transactions)
	fmt.Println(master.Transactions[0])
	fmt.Println(master.Transactions[1])
	fmt.Println(master.Transactions[2])
	fmt.Println(master.Transactions[3])

	var mcc []string
	mcc = append(mcc, "5011", "5014", "5015")
	fmt.Println(mcc)
	fmt.Println(card.SumByMCC(master.Transactions, mcc))

	category := card.TranslateMCC(master.Transactions[0].Mcc)
	fmt.Println(category)
}

