build-proto:
	protoc \
	--go_out=client_api_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=client_api_proto \
	--go-grpc_opt=paths=source_relative \
	client_api.proto
