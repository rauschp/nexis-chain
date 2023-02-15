NAME=bin/nexis-chain

build:
	@go build -o ${NAME} main.go

run: build
	@./${NAME}

test: 
	@go test ./...

clean: 
	@go clean
	@rm ${NAME}

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
