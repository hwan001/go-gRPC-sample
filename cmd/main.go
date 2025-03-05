package main

import (
	"go-grpc-sample/agent"
	"go-grpc-sample/server"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// gRPC 서버 실행
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting gRPC Agent Server...")
		agent.StartAgentServer()
	}()

	// HTTP 서버 실행
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting HTTP Gin Server...")
		server.StartHTTPServer()
	}()

	wg.Wait() // 두 서버가 종료되지 않도록 대기
}
