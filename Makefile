BIN_DIR = bin
PROTO_DIR = proto

migration-up:
	@echo "migration-up";
migration-down:
	@echo "migration-down";
build-app:
	@go build -o ${BIN_DIR}/app ./cmd/api/main.go
build-api:
	@protoc -I${PROTO_DIR} --go_out=. --go-grpc_out=. ${PROTO_DIR}/*.proto
	@echo "build api: DONE"
build-all:
	@echo "build-all";
