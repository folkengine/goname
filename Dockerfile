FROM golang:1.8.3-alpine3.6
RUN mkdir -p /go/src/github.com/folkengine/goname
WORKDIR /go/src/github.com/folkengine/goname
COPY . /go/src/github.com/folkengine/goname
RUN go build -o main cmd/goname/main.go
CMD ["/go/src/github.com/folkengine/goname/main"]

