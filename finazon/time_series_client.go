package finazon

type TimeSeriesClient struct {
	connection *Connection
	client     TimeSeriesServiceClient
}

func NewTimeSeriesClient(c *Connection) *TimeSeriesClient {
	return &TimeSeriesClient{connection: c, client: NewTimeSeriesServiceClient(c.Con)}
}

func (s *TimeSeriesClient) GetTimeSeries(request *GetTimeSeriesRequest) (*GetTimeSeriesResponse, error) {
	return s.client.GetTimeSeries(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesAtr(request *GetTimeSeriesAtrRequest) (*GetTimeSeriesAtrResponse, error) {
	return s.client.GetTimeSeriesAtr(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesBBands(request *GetTimeSeriesBBandsRequest) (*GetTimeSeriesBBandsResponse, error) {
	return s.client.GetTimeSeriesBBands(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesIchimoku(request *GetTimeSeriesIchimokuRequest) (*GetTimeSeriesIchimokuResponse, error) {
	return s.client.GetTimeSeriesIchimoku(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesMa(request *GetTimeSeriesMaRequest) (*GetTimeSeriesMaResponse, error) {
	return s.client.GetTimeSeriesMa(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesMacd(request *GetTimeSeriesMacdRequest) (*GetTimeSeriesMacdResponse, error) {
	return s.client.GetTimeSeriesMacd(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesObv(request *GetTimeSeriesObvRequest) (*GetTimeSeriesObvResponse, error) {
	return s.client.GetTimeSeriesObv(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesRsi(request *GetTimeSeriesRsiRequest) (*GetTimeSeriesRsiResponse, error) {
	return s.client.GetTimeSeriesRsi(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesSar(request *GetTimeSeriesSarRequest) (*GetTimeSeriesSarResponse, error) {
	return s.client.GetTimeSeriesSar(s.connection.Ctx, request)
}
func (s *TimeSeriesClient) GetTimeSeriesStoch(request *GetTimeSeriesStochRequest) (*GetTimeSeriesStochResponse, error) {
	return s.client.GetTimeSeriesStoch(s.connection.Ctx, request)
}
