.PHONY: all build run

all:build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build:
	go build -o ./bin/main ./infrastructure/cmd/main.go

run:
	go run -race infrastructure/cmd/main.go

test:
	go test -tags testing ./...

up:
	docker-compose up -d --force-recreate

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/gomockgen
	@~/go/bin/mockgen -source=domain/post.go -destination=domain/mock/post.go packge=mock