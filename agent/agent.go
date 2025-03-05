package agent

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-grpc-sample/agent/agentpb"

	"google.golang.org/grpc"
)

// AgentServer 구현
type agentServer struct {
	pb.UnimplementedAgentServiceServer
}

// 실행 가능한 함수 목록을 저장하는 맵
var functionMap = map[string]func(string) string{
	"TestFunction":    TestFunction,
	"AnotherFunction": AnotherFunction, // 다른 함수도 추가 가능
}

// TestFunction - 예제 함수
func TestFunction(payload string) string {
	log.Println("Executing TestFunction: Hello, World!")
	return fmt.Sprintf("TestFunction executed with payload: %s", payload)
}

// AnotherFunction - 또 다른 예제 함수
func AnotherFunction(payload string) string {
	log.Println("Executing AnotherFunction")
	return fmt.Sprintf("AnotherFunction executed with payload: %s", payload)
}

// ExecuteFunction - gRPC 요청 처리 (맵을 이용한 동적 함수 실행)
func (s *agentServer) ExecuteFunction(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Printf("Received function request: %s", req.FunctionName)

	// 요청된 함수명이 functionMap에 있는지 확인
	if function, exists := functionMap[req.FunctionName]; exists {
		result := function(req.Payload) // 동적으로 함수 실행
		return &pb.ResponseMessage{Result: result}, nil
	}

	// 함수가 존재하지 않는 경우 예외 처리
	errMsg := fmt.Sprintf("Error: Function '%s' not found", req.FunctionName)
	log.Println(errMsg)
	return nil, fmt.Errorf(errMsg)
}

// StartAgentServer - gRPC 서버 시작
func StartAgentServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAgentServiceServer(grpcServer, &agentServer{})

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
