# Finazon Go gRPC Client

This is the official Go library for Finazon, offering access to:
- Lists of datasets, publishers, markets, and tickers.
- Market data: ticker snapshots, time series, trades, and technical indicators.
- Data from specific datasets such as Benzinga, Binance, Crypto, Forex, SEC, and SIP.
- Full compatibility with both JavaScript and TypeScript.

🔑 **API key** is essential. If you haven't got one yet, [sign up here](https://finazon.io/).

## Requirements

Ensure you have Go 1.18 or later.

## Installation

```bash
go get github.com/finazon-io/finazon-grpc-go
```

## Updating to last version

```bash
go get -u github.com/finazon-io/finazon-grpc-go
```

## Quick start

### 1. Set up a new project
```bash
mkdir hello-finazon && cd hello-finazon
go mod init example/hello
go get github.com/finazon-io/finazon-grpc-go
```

### 2. Create `hello-world.go` script
```go
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
		Dataset: "sip_non_pro",
		Ticker:  "AAPL",
	}
	data, err := timeSeriesClient.GetTimeSeries(&timeSeriesRequest)
	if err != nil {
		fmt.Println("%s", err)
		return
	}
	fmt.Println("%s", data.String())
}
```

### 3. Download dependencies 
```bash
go mod tidy
go mod vendor
```

### 4. Input your API key
Replace `'your_api_key'` with your actual key.

### 5. Run the example
```bash
go run hello-world.go
```

📝 Expected output:
```
result:{timestamp:1709326680  open:179.655  close:179.695  high:179.73  low:179.64  volume:563082}
result:{timestamp:1709326620  open:179.67  close:179.66  high:179.672  low:179.6  volume:242826}
...
result:{timestamp:1709324940  open:179.71  close:179.68  high:179.715  low:179.64  volume:107197}
```

## RPC support

The following table outlines the supported rpc calls:
<!--rpc_table_boundary-->
| Service              | rpc                       | Description                                        |
|----------------------|---------------------------|----------------------------------------------------|
| ApiUsageService      | GetApiUsage               | Get a list of products with quota limit/usage      |
| BenzingaService      | GetDividentsCalendar      | Returns the dividends calendar from Benzinga       |
| BenzingaService      | GetEarningsCalendar       | Returns the earnings calendar from Benzinga        |
| BenzingaService      | GetNews                   | Returns the news articles from Benzinga            |
| BenzingaService      | GetIPO                    | Returns IPO data from Benzinga                     |
| BinanceService       | GetTimeSeries             | Get time series data without technical indicators  |
| CryptoService        | GetTimeSeries             | Get time series data for any given ticker          |
| DatasetsService      | GetDatasets               | Get a list of all datasets available at Finazon    |
| ExchangeService      | GetMarketsCrypto          | Returns a list of supported crypto markets         |
| ExchangeService      | GetMarketsStocks          | Returns a list of supported stock markets          |
| ForexService         | GetTimeSeries             | Get time series data for any given ticker          |
| PublisherService     | GetPublishers             | Get a list of all publishers available at Finazon  |
| SecService           | GetFilings                | Real-time and historical access to all forms, filings, and exhibits directly from the SEC's EDGAR system |
| SipService           | GetTrades                 | Returns detailed information on trades executed through the Securities Information Processor (SIP) |
| SipService           | GetMarketCenter           | Returns a list of market centers                   |
| SnapshotService      | GetSnapshot               | This endpoint returns a combination of different data points, such as daily performance, last quote, last trade, minute data, and previous day performance |
| TickersService       | FindTickersStocks         | This API call returns an array of stocks tickers available at Finazon Data API. This list is updated daily |
| TickersService       | FindTickersCrypto         | This API call returns an array of crypto tickers available at Finazon Data API. This list is updated daily |
| TickersService       | FindTickersForex          | This API call returns an array of forex tickers available at Finazon Data API. This list is updated daily |
| TickersService       | FindTickerUS              | This API call returns an array of US tickers available at Finazon Data API. This list is updated daily |
| TimeSeriesService    | GetTimeSeries             | Get time series data without technical indicators  |
| TimeSeriesService    | GetTimeSeriesAtr          | Get time series data for ATR technical indicator   |
| TimeSeriesService    | GetTimeSeriesBBands       | Get time series data for BBands technical indicator |
| TimeSeriesService    | GetTimeSeriesIchimoku     | Get time series data for Ichimoku technical indicator |
| TimeSeriesService    | GetTimeSeriesMa           | Get time series data for Ma technical indicator    |
| TimeSeriesService    | GetTimeSeriesMacd         | Get time series data for Macd technical indicator  |
| TimeSeriesService    | GetTimeSeriesObv          | Get time series data for Obv technical indicator   |
| TimeSeriesService    | GetTimeSeriesRsi          | Get time series data for Rsi technical indicator   |
| TimeSeriesService    | GetTimeSeriesSar          | Get time series data for Sar technical indicator   |
| TimeSeriesService    | GetTimeSeriesStoch        | Get time series data for Stoch technical indicator |
| TradeService         | GetTrades                 | Returns general information on executed trades     |
<!--rpc_table_boundary-->
Here's how you can import `client` and `request` objects:
```go
import "github.com/finazon-io/finazon-grpc-go/finazon"

con, _ := finazon.GetConnection(API_KEY)

client := con.GetServiceNameClient()

request := finazon.GetServiceNameRequest{}
data, err := client.RpcName(&request)
```

## Documentation
Delve deeper with our [official documentation](https://finazon.io/docs).

## Support
- 🌐 Visit our [contact page](https://finazon.io/contact-sales).
- 🛠 Or our [support center](https://support.finazon.io/en/).

## Stay updated
- Follow us on [𝕏](https://twitter.com/finazon_io).
- Join our community on [Discord](https://discord.gg/D5u4ZpB7w7).
- Follow us on [LinkedIn](https://www.linkedin.com/company/finazon).

## Contributing
1. Fork and clone: `$ git clone https://github.com/finazon-io/finazon-grpc-go.git`.
2. Branch out: `$ cd finazon-grpc-go && git checkout -b my-feature`.
3. Commit changes and test.
4. Push your branch and open a pull request with a comprehensive description.

For more guidance on contributing, see the [GitHub Docs](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) on GitHub.

## License

This project is licensed under the MIT License. See the `LICENSE.txt` file in this repository for more details.
