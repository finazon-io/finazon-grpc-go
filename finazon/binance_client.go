package finazon

type BinanceClient struct {
	connection *Connection
	client     BinanceServiceClient
}

func NewBinanceClient(c *Connection) *BinanceClient {
	return &BinanceClient{connection: c, client: NewBinanceServiceClient(c.Con)}
}

func (s *BinanceClient) GetTimeSeries(request *GetBinanceTimeSeriesRequest) (*GetBinanceTimeSeriesResponse, error) {
	return s.client.GetTimeSeries(s.connection.Ctx, request)
}
