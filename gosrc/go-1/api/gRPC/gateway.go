package grpc

import (
	"io"
	"ipc/go-1/api/gRPC/grpc-gateway/exchange_gatewaypb"
	"log"
)

func (s *server) StreamOrders(stream exchange_gatewaypb.ExchangeGateway_StreamOrdersServer) error {
	for {
		req,err := stream.Recv()
		if err !=io.EOF {
			return nil
		}
		if req == nil {
			return err
		}
		log.Printf("Received OrderRequest: %v", req)
	}

}
