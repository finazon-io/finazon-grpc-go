package finazon_grpc_go

type SipClient struct {
	connection *Connection
	client     SipServiceClient
}

func NewSipClient(c *Connection) *SipClient {
	return &SipClient{connection: c, client: NewSipServiceClient(c.Con)}
}

func (s *SipClient) GetTrades(request *GetSipTradesRequest) (*GetSipTradesResponse, error) {
	return s.client.GetTrades(s.connection.Ctx, request)
}
func (s *SipClient) GetMarketCenter(request *GetSipMarketCenterRequest) (*GetSipMarketCenterResponse, error) {
	return s.client.GetMarketCenter(s.connection.Ctx, request)
}
