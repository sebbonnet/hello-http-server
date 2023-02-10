FROM golang:1.19-alpine

WORKDIR /go/build

COPY main.go go.mod go.sum /go/build/

RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

RUN go build -o /simple-app

ENTRYPOINT ["/simple-app"]