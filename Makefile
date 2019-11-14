build:
	go test -cover ./...
	GOARCH=amd64 GOOS=linux go build .

deploy: build
	npx serverless deploy