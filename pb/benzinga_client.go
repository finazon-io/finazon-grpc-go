package finazon_grpc_go

type BenzingaClient struct {
	connection *Connection
	client     BenzingaServiceClient
}

func NewBenzingaClient(c *Connection) *BenzingaClient {
	return &BenzingaClient{connection: c, client: NewBenzingaServiceClient(c.Con)}
}

func (s *BenzingaClient) GetDividentsCalendar(request *GetDividentsCalendarRequest) (*GetDividentsCalendarResponse, error) {
	return s.client.GetDividentsCalendar(s.connection.Ctx, request)
}
func (s *BenzingaClient) GetEarningsCalendar(request *GetEarningsCalendarRequest) (*GetEarningsCalendarResponse, error) {
	return s.client.GetEarningsCalendar(s.connection.Ctx, request)
}
func (s *BenzingaClient) GetNews(request *GetNewsRequest) (*GetNewsResponse, error) {
	return s.client.GetNews(s.connection.Ctx, request)
}
func (s *BenzingaClient) GetIPO(request *GetIPORequest) (*GetIPOResponse, error) {
	return s.client.GetIPO(s.connection.Ctx, request)
}
