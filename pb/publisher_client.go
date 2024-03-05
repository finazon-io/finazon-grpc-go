package finazon_grpc_go

type PublisherClient struct {
	connection *Connection
	client     PublisherServiceClient
}

func NewPublisherClient(c *Connection) *PublisherClient {
	return &PublisherClient{connection: c, client: NewPublisherServiceClient(c.Con)}
}

func (s *PublisherClient) GetPublishers(request *GetPublishersRequest) (*GetPublishersResponse, error) {
	return s.client.GetPublishers(s.connection.Ctx, request)
}
