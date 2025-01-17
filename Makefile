.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/connect_to_rds connect_to_rds/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
