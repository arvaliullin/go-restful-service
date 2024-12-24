FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN bash ./scripts/proto-gen.sh
RUN go build -o /bin/grpc_service github.com/arvaliullin/go-restful-service/cmd/grpc_service

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /bin/grpc_service/ app/grpc_service

CMD ["/app/grpc_service"]
