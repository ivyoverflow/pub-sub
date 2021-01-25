FROM golang:1.15.6-buster

ENV GOPATH=/

COPY ./ ./

RUN apt-get update

RUN go mod download
RUN go build -o build/publisher ./cmd/publisher/main.go

CMD [ "./build/publisher", "-topic", "news" ]

