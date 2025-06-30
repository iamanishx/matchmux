package server

import (
	grpc "ipc/go-1/api/gRPC"
	restapi "ipc/go-1/api/rest"
	"sync"
)

func Server() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		go restapi.StartRest()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		go grpc.StartGRPC()
	}()
	wg.Wait()
}
