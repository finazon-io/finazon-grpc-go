package finazon

type ApiUsageClient struct {
	connection *Connection
	client     ApiUsageServiceClient
}

func NewApiUsageClient(c *Connection) *ApiUsageClient {
	return &ApiUsageClient{connection: c, client: NewApiUsageServiceClient(c.Con)}
}

func (s *ApiUsageClient) GetApiUsage(request *GetApiUsageRequest) (*GetApiUsageResponse, error) {
	return s.client.GetApiUsage(s.connection.Ctx, request)
}
