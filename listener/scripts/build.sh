export ADDR=localhost
export PORT=8080

cd listener/ && \
go build -o build/listener ./cmd/listener/main.go && \
./build/listener
