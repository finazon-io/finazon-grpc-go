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
		PageSize: 10,
	}
	timeSeriesMaRequest := finazon.GetTimeSeriesMaRequest{
		TimeSeries: &timeSeriesRequest,
	}
	data, err := timeSeriesClient.GetTimeSeriesMa(&timeSeriesMaRequest)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	lastItem := data.Result[0]
	fmt.Printf("Last MA value at time %d is %s\n", lastItem.Timestamp, lastItem.Ma)
}
