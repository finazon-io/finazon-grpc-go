package finazon

type TradeClient struct {
	connection *Connection
	client     TradeServiceClient
}

func NewTradeClient(c *Connection) *TradeClient {
	return &TradeClient{connection: c, client: NewTradeServiceClient(c.Con)}
}

func (s *TradeClient) GetTrades(request *GetTradesRequest) (*GetTradesResponse, error) {
	return s.client.GetTrades(s.connection.Ctx, request)
}
