.PHONY: client server test evans format_proto gen_proto

gen_proto:
	rm -f pb/*.go
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

format_proto:
	clang-format -i proto/*.proto

server:
	go run cmd/server/main.go --port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover ./...

evans:
	evans -r repl -p 8080
