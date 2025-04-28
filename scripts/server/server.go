package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tokopedia/gripmock/stub"
)

type serverParam struct {
	address string
	port    int64
}

func main() {
	grpcParam := serverParam{}
	flag.StringVar(&grpcParam.address, "grpc-listen", "0.0.0.0", "Address the gRPC server will bind to. Default to localhost, set to 0.0.0.0 to use from another machine")
	flag.Int64Var(&grpcParam.port, "grpc-port", 4770, "BindPort of gRPC tcp server")

	stubOptions := stub.Options{}
	flag.StringVar(&stubOptions.BindAddr, "admin-listen", "0.0.0.0", "Address the admin server will bind to. Default to localhost, set to 0.0.0.0 to use from another machine")
	flag.Int64Var(&stubOptions.BindPort, "admin-port", 4771, "BindPort of stub admin server")
	flag.StringVar(&stubOptions.StubPath, "stubs", "/stubs", "Path where the stub files are (Optional)")

	flag.Parse()

	// run admin stub server
	stub.RunStubServer(stubOptions)

	tcpAddress := fmt.Sprintf("%s:%d", grpcParam.address, grpcParam.port)
	lis, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	register(s)

	reflection.Register(s)
	fmt.Println("Serving gRPC on tcp://" + tcpAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
