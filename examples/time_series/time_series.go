package main

import (
	"fmt"
	"github.com/finazon-io/finazon-grpc-go/finazon"
	"time"
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
		Dataset: "sip_non_pro",
		Ticker:  "AAPL",
	}
	data, err := timeSeriesClient.GetTimeSeries(&timeSeriesRequest)
	if err != nil {
		fmt.Println("%s", err)
		return
	}

	fmt.Printf("Ticker: %s\n", timeSeriesRequest.Ticker)
	fmt.Printf("%-24s %12s %12s %12s %12s %12s\n", "Time", "Open", "High", "Low", "Close", "Volume")
	for _, item := range data.Result {

		fmt.Printf("%-24s %12f %12f %12f %12f %12.2f\n", time.Unix(item.Timestamp, 0).Format("2006-01-02T15:04:05 MST"), item.Open, item.High, item.Low, item.Close, item.Volume)
	}
}
