// Code will replace with generated
package main

import (
	"os"

	"google.golang.org/grpc"

	"github.com/tokopedia/gripmock/protogen"
	"github.com/tokopedia/gripmock/stub"
)

// Use imports to prevent them from being removed by go mod tidy
var _ = protogen.ProtoGen
var _ = stub.FindStub

func register(_ *grpc.Server) {
	os.Exit(0)
}
