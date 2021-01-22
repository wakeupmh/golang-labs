FROM golang:1-alpine AS builder

WORKDIR /root/go/src/github.com/devopmh/golang-labs/minikube-go

RUN apk add --update git

COPY server.go .
RUN go get github.com/google/uuid
RUN go get github.com/kelseyhightower/envconfig

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine
EXPOSE 8080

COPY --from=builder /root/go/src/github.com/devopmh/golang-labs/minikube-go/app .

ENTRYPOINT ./app