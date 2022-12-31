package main

import (
	"net"

	"github.com/lavanet-challenge/internal/proxy"
)

func main() {
	// init grpc server pointed to osmosis endpoint
	p, err := proxy.New("grpc.osmosis.zone:9090")
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	// serve requests
	if err := p.Serve(l); err != nil {
		panic(err)
	}
}
