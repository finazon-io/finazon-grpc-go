package finazon

type CryptoClient struct {
	connection *Connection
	client     CryptoServiceClient
}

func NewCryptoClient(c *Connection) *CryptoClient {
	return &CryptoClient{connection: c, client: NewCryptoServiceClient(c.Con)}
}

func (s *CryptoClient) GetTimeSeries(request *GetCryptoTimeSeriesRequest) (*GetCryptoTimeSeriesResponse, error) {
	return s.client.GetTimeSeries(s.connection.Ctx, request)
}
