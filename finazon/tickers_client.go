package finazon

type TickersClient struct {
	connection *Connection
	client     TickersServiceClient
}

func NewTickersClient(c *Connection) *TickersClient {
	return &TickersClient{connection: c, client: NewTickersServiceClient(c.Con)}
}

func (s *TickersClient) FindTickersStocks(request *FindTickersStocksRequest) (*FindTickersStocksResponse, error) {
	return s.client.FindTickersStocks(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickersCrypto(request *FindTickersCryptoRequest) (*FindTickersCryptoResponse, error) {
	return s.client.FindTickersCrypto(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickersForex(request *FindTickersForexRequest) (*FindTickersForexResponse, error) {
	return s.client.FindTickersForex(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickersUS(request *FindTickersUSRequest) (*FindTickersUSResponse, error) {
	return s.client.FindTickersUS(s.connection.Ctx, request)
}
