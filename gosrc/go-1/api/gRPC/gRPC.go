package grpc

import (
	"fmt"
	"ipc/go-1/api/gRPC/grpc-gateway/exchange_gatewaypb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	exchange_gatewaypb.UnimplementedExchangeGatewayServer
}

func StartGRPC() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Error starting gRPC server: " + err.Error())
	}
	grpcServer := grpc.NewServer()
	exchange_gatewaypb.RegisterExchangeGatewayServer(grpcServer, &server{})
	fmt.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic("Error serving gRPC server: " + err.Error())
	}

}
