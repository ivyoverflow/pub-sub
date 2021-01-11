export ADDR=localhost
export PORT=8080

cd client/ && \
go build -o build/client ./cmd/client/main.go && \
./build/client
