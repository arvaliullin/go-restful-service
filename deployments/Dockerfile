FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/go-restful-service github.com/arvaliullin/go-restful-service/cmd/service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /bin/go-restful-service /app/go-restful-service
EXPOSE 8080
CMD ["/app/go-restful-service"]
