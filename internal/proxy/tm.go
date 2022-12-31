package proxy

import (
	"context"

	tm "github.com/lavanet-challenge/internal/pb/cosmos/base/tendermint/v1beta1"
)

var _ tm.ServiceServer = (*Proxy)(nil)

// directly implement all Service gRPC endpoints, should work with any cosmos chain
func (p *Proxy) GetNodeInfo(ctx context.Context, req *tm.GetNodeInfoRequest) (*tm.GetNodeInfoResponse, error) {
	return p.tmService.GetNodeInfo(ctx, req)
}

func (p *Proxy) GetSyncing(ctx context.Context, req *tm.GetSyncingRequest) (*tm.GetSyncingResponse, error) {
	return p.tmService.GetSyncing(ctx, req)
}

func (p *Proxy) GetLatestBlock(ctx context.Context, req *tm.GetLatestBlockRequest) (*tm.GetLatestBlockResponse, error) {
	return p.tmService.GetLatestBlock(ctx, req)
}

func (p *Proxy) GetBlockByHeight(ctx context.Context, req *tm.GetBlockByHeightRequest) (*tm.GetBlockByHeightResponse, error) {
	return p.tmService.GetBlockByHeight(ctx, req)
}

func (p *Proxy) GetLatestValidatorSet(ctx context.Context, req *tm.GetLatestValidatorSetRequest) (*tm.GetLatestValidatorSetResponse, error) {
	return p.tmService.GetLatestValidatorSet(ctx, req)
}

func (p *Proxy) GetValidatorSetByHeight(ctx context.Context, req *tm.GetValidatorSetByHeightRequest) (*tm.GetValidatorSetByHeightResponse, error) {
	return p.tmService.GetValidatorSetByHeight(ctx, req)
}

func (p *Proxy) ABCIQuery(ctx context.Context, req *tm.ABCIQueryRequest) (*tm.ABCIQueryResponse, error) {
	return p.tmService.ABCIQuery(ctx, req)
}
