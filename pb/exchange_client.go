package finazon_grpc_go

type ExchangeClient struct {
	connection *Connection
	client     ExchangeServiceClient
}

func NewExchangeClient(c *Connection) *ExchangeClient {
	return &ExchangeClient{connection: c, client: NewExchangeServiceClient(c.Con)}
}

func (s *ExchangeClient) GetMarketsCrypto(request *GetMarketsCryptoRequest) (*GetMarketsCryptoResponse, error) {
	return s.client.GetMarketsCrypto(s.connection.Ctx, request)
}
func (s *ExchangeClient) GetMarketsStocks(request *GetMarketsStocksRequest) (*GetMarketsStocksResponse, error) {
	return s.client.GetMarketsStocks(s.connection.Ctx, request)
}
