lint:
	golangci-lint run

proto:
	protoc --go_opt= --go_out=plugins=grpc:. ./delivery/pb/*.proto
