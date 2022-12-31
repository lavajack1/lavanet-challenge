package proxy

import (
	"context"
	"fmt"
	"time"

	reflection "github.com/cosmos/cosmos-sdk/server/grpc/gogoreflection"
	bank "github.com/lavanet-challenge/internal/pb/cosmos/bank/v1beta1"
	tm "github.com/lavanet-challenge/internal/pb/cosmos/base/tendermint/v1beta1"
	"github.com/lavanet-challenge/internal/pb/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Proxy struct {
	*grpc.Server
	tmService   tm.ServiceClient
	bankService bank.QueryClient
}

func New(endpoint string) (*Proxy, error) {

	p := &Proxy{}

	// init gRPC server and register services
	s := grpc.NewServer()
	tm.RegisterServiceServer(s, p)
	bank.RegisterQueryServer(s, p)
	proxy.RegisterProxyServiceServer(s, p)

	// add reflection for clients to retrieve schema dynamically
	reflection.Register(s)

	// connect to endpoint
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// init clients
	tmService := tm.NewServiceClient(conn)
	bankService := bank.NewQueryClient(conn)

	p.Server = s
	p.tmService = tmService
	p.bankService = bankService

	return p, nil
}

// stream N blocks to the client starting from the latest block
func (p *Proxy) GetLatestBlocks(req *proxy.GetLatestBlocksRequest, streamer proxy.ProxyService_GetLatestBlocksServer) error {
	// if Count is 0, error as we wouldnt send any blocks
	if req.Count == 0 {
		return fmt.Errorf("invalid count argument: %d", req.Count)
	}

	// get latest block in chain
	firstBlockResponse, err := p.tmService.GetLatestBlock(context.TODO(), &tm.GetLatestBlockRequest{})
	if err != nil {
		return err
	}

	startingHeight := firstBlockResponse.GetBlock().GetHeader().Height

	// send first block to client
	if err := streamer.Send(&proxy.GetLatestBlocksResponse{
		Height: startingHeight,
		Hash:   firstBlockResponse.GetBlockId().GetHash(),
	}); err != nil {
		return err
	}

	currentHeight := startingHeight + 1
	for currentHeight < req.Count+startingHeight {
		// attempt to get next block
		nextBlockResponse, err := p.tmService.GetBlockByHeight(context.TODO(), &tm.GetBlockByHeightRequest{Height: currentHeight})
		if err != nil {
			// if there's an error, assume it's because the next block hasn't been produced yet
			// print the error and sleep as to not spam the endpoint
			// in a real system we would want more robust error checking here
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}

		if err := streamer.Send(&proxy.GetLatestBlocksResponse{
			Height: currentHeight,
			Hash:   nextBlockResponse.GetBlockId().GetHash(),
		}); err != nil {
			return err
		}

		// increment next height to query for
		currentHeight++
	}

	return nil
}
