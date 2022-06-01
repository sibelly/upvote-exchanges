gen:
	protoc --proto_path=proto proto/*.proto \
	--go_out=:pb --go_opt=paths=source_relative \
	--go-grpc_out=:pb --go-grpc_opt=paths=source_relative

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -addr 0.0.0.0:8080

.PHONY: gen client server
