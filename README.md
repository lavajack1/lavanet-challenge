# Lavanet Challenge

A gRPC server proxying requests to an Osmosis gRPC endpoint

```bash
# generate the protobuf files
make gen

# start grpc proxy in terminal 1
make proxy

# run client in terminal 2
# client will collect latest blocks and write to a file named 'results.json'
make client

# run tests
# tests compare output received by proxy gRPC compared to directly invoking the endpoint
make test

# grpcurl examples
#
# describe tendermint service
grpcurl -plaintext localhost:5555 describe cosmos.base.tendermint.v1beta1.Service

# get latest block
grpcurl -plaintext localhost:5555 cosmos.base.tendermint.v1beta1.Service/GetLatestBlock

# query bank params
grpcurl -plaintext localhost:5555 cosmos.bank.v1beta1.Query/Params
```

## Details

- `buf` used to import cosmos protobuf files and generate the go grpc files, importing these files makes it easier to implement the proxy and return the exact data as a cosmos server
- test compares my proxy response with the response directly from the gRPC server
- makefile used for ease of use
- `pkg/client` exports a client for convenience to be used by consumers
- for reflection to work properly to support `grpcurl` I had to switch out the usual reflection pkg with `github.com/cosmos/cosmos-sdk/server/grpc/gogoreflection`
- point the proxy to any cosmos-sdk compatible rpc via `proxy.New("endpoint_here")`

### Add a service

To proxy a new cosmos service do these steps

1. add import to `proto/proxy/proxy.proto`
2. run `make gen` to generate the go grpc files
3. create a new file in `internal/proxy/` implementing the service
