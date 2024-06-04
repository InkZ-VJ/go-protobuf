gen_proto:
	rm -f pb/*.go
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

format_proto:
	clang-format -i proto/*.proto

run:
	go run main.go

test:
	go test ./...
