BUILD_DIR=./build
GO_FLAGS= -ldflags "-s -w"
GOOS=linux
GOARCH=amd64

PROTO_DIR := ./internal/proto
GEN_DIR := ./internal/gen

.PHONY: docs run

all: gen main

main: gen
	GOARCH=${GOARCH} GOOS=${GOOS} CGO_ENABLED=0 go build ${GO_FLAGS} ${GO_FLAGS} -o ${BUILD_DIR}/${GOOS}/app.${GOARCH}.bin ./cmd/app/main.go 

lint:  
	golangci-lint run 

test: 
	go test -v 


docs: 
	echo "Running docs on :3000 port"  
	cd docs && python -m http.server 3000

gen:
	@protoc --go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--proto_path=$(PROTO_DIR) \
		$(PROTO_DIR)/*.proto

run: main
	@${BUILD_DIR}/${GOOS}/app.${GOARCH}.bin
