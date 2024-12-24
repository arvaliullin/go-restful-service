# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# winget install protobuf

protoc -I=api/proto --go_out=. --go-grpc_out=. api/proto/glossary.proto
