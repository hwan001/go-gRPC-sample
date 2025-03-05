PROTO_DIR=agent
PROTO_FILE=agent.proto
OUT_DIR=agent

generate:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) --proto_path=$(PROTO_DIR) $(PROTO_FILE)

tidy:
	go mod tidy
	cd server && go mod tidy
	cd agent && go mod tidy