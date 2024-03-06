package finazon

type ForexClient struct {
	connection *Connection
	client     ForexServiceClient
}

func NewForexClient(c *Connection) *ForexClient {
	return &ForexClient{connection: c, client: NewForexServiceClient(c.Con)}
}

func (s *ForexClient) GetTimeSeries(request *GetForexTimeSeriesRequest) (*GetForexTimeSeriesResponse, error) {
	return s.client.GetTimeSeries(s.connection.Ctx, request)
}
