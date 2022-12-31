package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	bank "github.com/lavanet-challenge/internal/pb/cosmos/bank/v1beta1"
	tmservice "github.com/lavanet-challenge/internal/pb/cosmos/base/tendermint/v1beta1"
	"github.com/lavanet-challenge/internal/proxy"
	"github.com/lavanet-challenge/pkg/client"
	"github.com/stretchr/testify/assert"
)

// start our grpc proxy
func setupServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	p, err := proxy.New("grpc.osmosis.zone:9090")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ready on port %d\n", port)

	go func() {
		if err := p.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

func TestGetLatestBlock(t *testing.T) {
	setupServer(5555)

	proxyClient, err := client.New("localhost:5555")
	assert.Nil(t, err)

	directClient, err := client.New("grpc.osmosis.zone:9090")
	assert.Nil(t, err)

	res, err := proxyClient.GetLatestBlock(context.TODO(), &tmservice.GetLatestBlockRequest{})
	assert.Nil(t, err)

	directRes, err := directClient.GetLatestBlock(context.TODO(), &tmservice.GetLatestBlockRequest{})
	assert.Nil(t, err)

	// check that the response received from our proxy and the osmosis endpoint is the same
	assert.True(t, reflect.DeepEqual(res, directRes))
}

func TestBankParams(t *testing.T) {
	setupServer(5556)

	proxyClient, err := client.New("localhost:5555")
	assert.Nil(t, err)

	directClient, err := client.New("grpc.osmosis.zone:9090")
	assert.Nil(t, err)

	res, err := proxyClient.Params(context.TODO(), &bank.QueryParamsRequest{})
	assert.Nil(t, err)

	directRes, err := directClient.Params(context.TODO(), &bank.QueryParamsRequest{})
	assert.Nil(t, err)

	// check that the response received from our proxy and the osmosis endpoint is the same
	assert.True(t, reflect.DeepEqual(res, directRes))
}
