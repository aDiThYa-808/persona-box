#!/bin/bash

set -e

echo "Removing old build..."
rm -f bootstrap function.zip
echo "Removed old build"

echo "Building Go Lambda..."
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./cmd/server
echo "Built Go Lambda"

echo "Zipping the binary..."
zip function.zip bootstrap
echo "Zipped the binary"

echo "Deploying to lambda..."
aws lambda update-function-code --function-name persona-box-backend-server --zip-file fileb://function.zip 
echo "Deployment complete"

