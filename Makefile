test:
	cat test.log | go run main.go

linux:
	env GOARCH=amd64 GOOS=linux go build
