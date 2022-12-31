
# start the proxy server
proxy:
	go run cmd/proxy/main.go

# start the client
client:
	rm -rf results.json && go run cmd/client/main.go

# generate protobuf files for the proxy
gen:
	pushd proto && buf build && buf generate --include-imports && popd

# run integration tests
test:
	go test -v ./tests/...

.PHONY: gen test proxy client
