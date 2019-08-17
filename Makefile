SCALESIZE=4
build:
	dep ensure
	GOOS=linux GOARCH=amd64 go build -o bin/cryptocurrency cmd/myblockchain.go
test:
	go test ./... -v
run:
	go run cmd/myblockchain.go
dev: build
	docker-compose up --build --scale blockchain=$(SCALESIZE)
