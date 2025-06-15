package main

import (
	"context"
	"ipc/db"
	"ipc/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	db.MigrateDatabase()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	log.Println("Starting mail service...")
	// 	mail.SendEmail()
	// }()
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting server...")
		go server.Server()
	}()

	<-sigs
	log.Println("shutting down gracefully...")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdownCancel()

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		log.Println("All services stopped gracefully.")
	case <-shutdownCtx.Done():
		log.Println("Shutdown timeout reached, forcing exit.")
	}
}
