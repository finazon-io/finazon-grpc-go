package finazon

type SnapshotClient struct {
	connection *Connection
	client     SnapshotServiceClient
}

func NewSnapshotClient(c *Connection) *SnapshotClient {
	return &SnapshotClient{connection: c, client: NewSnapshotServiceClient(c.Con)}
}

func (s *SnapshotClient) GetSnapshot(request *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	return s.client.GetSnapshot(s.connection.Ctx, request)
}
