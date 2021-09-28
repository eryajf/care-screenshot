default: build

build:
	go build -o care-screenshot main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o care-screenshot main.go