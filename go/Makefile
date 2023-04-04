gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

switch:
	export PATH=$PATH:$(go env GOPATH)/bin

run-server:
	cd server && go run main.go

run-client:
	cd client && go run main.go