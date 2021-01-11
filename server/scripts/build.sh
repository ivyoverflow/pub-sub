export ADDR=localhost
export PORT=8080

cd server/ && \
go build -o build/server ./cmd/server/main.go && \
./build/server
