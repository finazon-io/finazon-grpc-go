package finazon_grpc_go

type DatasetsClient struct {
	connection *Connection
	client     DatasetsServiceClient
}

func NewDatasetsClient(c *Connection) *DatasetsClient {
	return &DatasetsClient{connection: c, client: NewDatasetsServiceClient(c.Con)}
}

func (s *DatasetsClient) GetDatasets(request *GetDatasetsRequest) (*GetDatasetsResponse, error) {
	return s.client.GetDatasets(s.connection.Ctx, request)
}
