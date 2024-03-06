package finazon

type SecClient struct {
	connection *Connection
	client     SecServiceClient
}

func NewSecClient(c *Connection) *SecClient {
	return &SecClient{connection: c, client: NewSecServiceClient(c.Con)}
}

func (s *SecClient) GetFilings(request *GetSecFillingsRequest) (*GetSecFillingsResponse, error) {
	return s.client.GetFilings(s.connection.Ctx, request)
}
