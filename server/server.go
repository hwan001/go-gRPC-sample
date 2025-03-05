package server

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "go-grpc-sample/agent/agentpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// StartHTTPServer 함수 추가 → cmd/main.go에서 실행
func StartHTTPServer() {
	router := gin.Default()

	// gRPC 클라이언트 설정
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgentServiceClient(conn)

	router.POST("/execute", func(c *gin.Context) {
		var request struct {
			FunctionName string `json:"function_name"`
			Payload      string `json:"payload"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// gRPC 요청 보내기
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.ExecuteFunction(ctx, &pb.RequestMessage{
			FunctionName: request.FunctionName,
			Payload:      request.Payload,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": resp.Result})
	})

	log.Println("Gin HTTP server is running on port 8080...")
	router.Run(":8080")
}
