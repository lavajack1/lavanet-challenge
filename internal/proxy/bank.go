package proxy

import (
	"context"

	bank "github.com/lavanet-challenge/internal/pb/cosmos/bank/v1beta1"
)

var _ bank.QueryServer = (*Proxy)(nil)

// implement bank queries
func (p *Proxy) Balance(ctx context.Context, req *bank.QueryBalanceRequest) (*bank.QueryBalanceResponse, error) {
	return p.bankService.Balance(ctx, req)
}

func (p *Proxy) AllBalances(ctx context.Context, req *bank.QueryAllBalancesRequest) (*bank.QueryAllBalancesResponse, error) {
	return p.bankService.AllBalances(ctx, req)
}

func (p *Proxy) SpendableBalances(ctx context.Context, req *bank.QuerySpendableBalancesRequest) (*bank.QuerySpendableBalancesResponse, error) {
	return p.bankService.SpendableBalances(ctx, req)
}

func (p *Proxy) TotalSupply(ctx context.Context, req *bank.QueryTotalSupplyRequest) (*bank.QueryTotalSupplyResponse, error) {
	return p.bankService.TotalSupply(ctx, req)
}

func (p *Proxy) SupplyOf(ctx context.Context, req *bank.QuerySupplyOfRequest) (*bank.QuerySupplyOfResponse, error) {
	return p.bankService.SupplyOf(ctx, req)
}

func (p *Proxy) Params(ctx context.Context, req *bank.QueryParamsRequest) (*bank.QueryParamsResponse, error) {
	return p.bankService.Params(ctx, req)
}

func (p *Proxy) DenomMetadata(ctx context.Context, req *bank.QueryDenomMetadataRequest) (*bank.QueryDenomMetadataResponse, error) {
	return p.bankService.DenomMetadata(ctx, req)
}

func (p *Proxy) DenomsMetadata(ctx context.Context, req *bank.QueryDenomsMetadataRequest) (*bank.QueryDenomsMetadataResponse, error) {
	return p.bankService.DenomsMetadata(ctx, req)
}

func (p *Proxy) DenomOwners(ctx context.Context, req *bank.QueryDenomOwnersRequest) (*bank.QueryDenomOwnersResponse, error) {
	return p.bankService.DenomOwners(ctx, req)
}

func (p *Proxy) SendEnabled(ctx context.Context, req *bank.QuerySendEnabledRequest) (*bank.QuerySendEnabledResponse, error) {
	return p.bankService.SendEnabled(ctx, req)
}
