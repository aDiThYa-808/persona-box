## HOW TO BUILD GO BINARY AND ZIP FILE FOR LAMBDA
```bash
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./cmd/server
zip function.zip bootstrap
```