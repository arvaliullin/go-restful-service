FROM golang:1.23-alpine AS builder

RUN apk add --no-cache \
    protobuf \
    git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PROTO_DIR=api/proto
ENV GO_OUT_DIR=/app
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY . .

RUN protoc -I=$PROTO_DIR --go_out=$GO_OUT_DIR --go-grpc_out=$GO_OUT_DIR $PROTO_DIR/glossary.proto

RUN go build -o /bin/grpc_service github.com/arvaliullin/go-restful-service/cmd/grpc_service
RUN go build -o /bin/grpc_client github.com/arvaliullin/go-restful-service/cmd/grpc_client

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/grpc_service /bin/grpc_service

CMD ["/bin/grpc_service"]
