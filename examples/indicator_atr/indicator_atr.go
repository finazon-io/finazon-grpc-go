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

	timeSeriesClient := con.GetTimeSeriesClient()
	timeSeriesRequest := finazon.GetTimeSeriesRequest{
		Dataset:  "sip_non_pro",
		Ticker:   "AAPL",
		Interval: "1h",
		PageSize: 5,
	}
	timeSeriesAtrRequest := finazon.GetTimeSeriesAtrRequest{
		TimeSeries: &timeSeriesRequest,
	}
	data, err := timeSeriesClient.GetTimeSeriesAtr(&timeSeriesAtrRequest)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	fmt.Printf("Last ATR value %s\n", data.Result[0].Atr)
}
