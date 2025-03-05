
## gRPC Sample code (go)


- 구조
```mermaid
flowchart LR
    subgraph agent["agent"]
        agent_port["50051"]
    end

    subgraph server["server"]
        server_port["8080"]
    end

    request["api request"] 

    request --(HTTP)--> server_port
    server --(gRPC)--> agent_port
```


- 프로젝트 구조
```
.
├── Makefile
├── README.md
├── agent
│   ├── agent.go
│   ├── agent.proto
│   ├── agentpb
│   │   ├── agent.pb.go
│   │   └── agent_grpc.pb.go
│   ├── go.mod
│   └── go.sum
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── server
    ├── go.mod
    ├── go.sum
    └── server.go
```

- pb 코드 생성
```sh
make generate
```

- tidy
```sh
make tidy
```

- 실행
```sh
cd cmd
go run main.py
```

- 테스트
```sh
# AnotherFunction
curl -X POST "http://localhost:8080/execute" \
-H "Content-Type: application/json" \
-d '{"function_name": "AnotherFunction", "payload": "Hello"}' 

# TestFunction
curl -X POST "http://localhost:8080/execute" \
-H "Content-Type: application/json" \
-d '{"function_name": "TestFunction", "payload": "Hello"}'
```

