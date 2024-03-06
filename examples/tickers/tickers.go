package main

import (
	"fmt"
	"github.com/finazon-io/finazon-grpc-go/finazon"
)

const API_KEY = "your_api_key"

func main() {
	con, err := finazon.GetConnection(API_KEY)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	tickersClient := con.GetTickersClient()
	tickersRequest := finazon.FindTickersCryptoRequest{
		Publisher: "binance",
		PageSize:  10,
	}
	data, err := tickersClient.FindTickersCrypto((&tickersRequest))
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	for _, item := range data.Result {
		fmt.Printf("%s\n", item.Ticker)
	}
}
