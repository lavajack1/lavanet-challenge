package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/lavanet-challenge/internal/pb/proxy"
	"github.com/lavanet-challenge/pkg/client"
)

type Results struct {
	TestResult []TestResult `json:"test_result"`
}

type TestResult struct {
	Height int64  `json:"height"`
	Hash   []byte `json:"hash"`
}

func main() {

	client, err := client.New("localhost:5555")
	if err != nil {
		panic(err)
	}

	// init a long lived connection to server to receive latest blocks streams
	stream, err := client.GetLatestBlocks(context.TODO(), &proxy.GetLatestBlocksRequest{Count: 5})
	if err != nil {
		panic(err)
	}

	// wait for server to send blocks
	var testResults []TestResult
	for {
		res, err := stream.Recv()
		// EOF means server is done sending us blocks, exit loop
		if err == io.EOF {
			fmt.Println("stream closed, exiting loop")
			break
		}
		if err != nil {
			panic(err)
		}
		testResult := TestResult{
			Height: res.Height,
			Hash:   res.Hash,
		}

		testResults = append(testResults, testResult)
	}

	// encode to json and write to file
	results := Results{
		TestResult: testResults,
	}

	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("results.json", b, 0755); err != nil {
		panic(err)
	}
}
