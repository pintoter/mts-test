.PHONY: gen
gen:
	mkdir -p pkg
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative api/order-service/order-service.proto

.PHONY: deps
deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go install github.com/bufbuild/buf/cmd/buf@v0.56.0
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@v0.56.0
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v0.56.0

.PHONY: build
build:
	go build -o order-service ./cmd/order-service/main.go

.PHONY: run
run: build
	./order-service

.PHONY: mock-gen
mock-gen:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock_repository.go
