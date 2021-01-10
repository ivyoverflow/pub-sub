.PHONY: server
server:
	cd server/ && \
	go build -o build/server ./cmd/server/main.go && \
	./build/server

.PHONY: client
client:
	cd client/ && \
	go build -o build/client ./cmd/client/main.go && \
	./build/client

.PHONY: clear
clear:
	rm -rf server/build/ && \
	rm -rf client/build/
