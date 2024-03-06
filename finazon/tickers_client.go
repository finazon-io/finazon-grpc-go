package finazon

type TickersClient struct {
	connection *Connection
	client     TickersServiceClient
}

func NewTickersClient(c *Connection) *TickersClient {
	return &TickersClient{connection: c, client: NewTickersServiceClient(c.Con)}
}

func (s *TickersClient) FindTickersStocks(request *FindTickersStocksRequest) (*FindTickerStocksResponse, error) {
	return s.client.FindTickersStocks(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickersCrypto(request *FindTickersCryptoRequest) (*FindTickerCryptoResponse, error) {
	return s.client.FindTickersCrypto(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickersForex(request *FindTickersForexRequest) (*FindTickerForexResponse, error) {
	return s.client.FindTickersForex(s.connection.Ctx, request)
}
func (s *TickersClient) FindTickerUS(request *FindTickerUSRequest) (*FindTickerUSResponse, error) {
	return s.client.FindTickerUS(s.connection.Ctx, request)
}
