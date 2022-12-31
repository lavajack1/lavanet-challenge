package client

import (
	bank "github.com/lavanet-challenge/internal/pb/cosmos/bank/v1beta1"
	tm "github.com/lavanet-challenge/internal/pb/cosmos/base/tendermint/v1beta1"
	"github.com/lavanet-challenge/internal/pb/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	tm.ServiceClient
	bank.QueryClient
	proxy.ProxyServiceClient
}

func New(endpoint string) (*Client, error) {
	// connect to endpoint
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := &Client{}
	c.ProxyServiceClient = proxy.NewProxyServiceClient(conn)
	c.ServiceClient = tm.NewServiceClient(conn)
	c.QueryClient = bank.NewQueryClient(conn)

	return c, nil
}
