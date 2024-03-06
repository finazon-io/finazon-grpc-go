package finazon

import (
	"context"
	"crypto/tls"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"sync"
)

type Connection struct {
	Con *grpc.ClientConn
	Ctx context.Context
}

var lock = &sync.Mutex{}
var apiKeyToConnection map[string]*Connection

func GetConnection(apiKey string) (*Connection, error) {
	if apiKeyToConnection == nil {
		apiKeyToConnection = make(map[string]*Connection, 10)
	}
	if apiKeyToConnection[apiKey] == nil {
		lock.Lock()
		defer lock.Unlock()
		if apiKeyToConnection[apiKey] == nil {
			conn, err := newConnection(apiKey)
			if err != nil {
				return nil, err
			}
			apiKeyToConnection[apiKey] = conn
			return apiKeyToConnection[apiKey], nil
		}
	}

	return apiKeyToConnection[apiKey], nil
}

func newConnection(apiKey string) (*Connection, error) {
	con := &Connection{}
	return con, con.connect(apiKey)
}

func (c *Connection) connect(apiKey string) error {
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
	conn, err := grpc.Dial(FINAZON_GRPC_HOST, grpc.WithTransportCredentials(creds))
	if err != nil {
		return errors.New("Unable to connect to grpc server")
	}
	c.Con = conn
	c.Ctx = metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+apiKey)

	return nil
}

func (c *Connection) Close() error {
	return c.Con.Close()
}

// getters

func (c *Connection) GetApiUsageClient() *ApiUsageClient {
  return NewApiUsageClient(c)
}
func (c *Connection) GetBenzingaClient() *BenzingaClient {
  return NewBenzingaClient(c)
}
func (c *Connection) GetBinanceClient() *BinanceClient {
  return NewBinanceClient(c)
}
func (c *Connection) GetCryptoClient() *CryptoClient {
  return NewCryptoClient(c)
}
func (c *Connection) GetDatasetsClient() *DatasetsClient {
  return NewDatasetsClient(c)
}
func (c *Connection) GetExchangeClient() *ExchangeClient {
  return NewExchangeClient(c)
}
func (c *Connection) GetForexClient() *ForexClient {
  return NewForexClient(c)
}
func (c *Connection) GetPublisherClient() *PublisherClient {
  return NewPublisherClient(c)
}
func (c *Connection) GetSecClient() *SecClient {
  return NewSecClient(c)
}
func (c *Connection) GetSipClient() *SipClient {
  return NewSipClient(c)
}
func (c *Connection) GetSnapshotClient() *SnapshotClient {
  return NewSnapshotClient(c)
}
func (c *Connection) GetTickersClient() *TickersClient {
  return NewTickersClient(c)
}
func (c *Connection) GetTimeSeriesClient() *TimeSeriesClient {
  return NewTimeSeriesClient(c)
}
func (c *Connection) GetTradeClient() *TradeClient {
  return NewTradeClient(c)
}

