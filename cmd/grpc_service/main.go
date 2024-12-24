package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/arvaliullin/go-restful-service/internal/glossary/delivery"
	"github.com/arvaliullin/go-restful-service/internal/glossary/repository"
	"github.com/arvaliullin/go-restful-service/internal/glossary/usecase"
	"github.com/arvaliullin/go-restful-service/pkg/glossary"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"os"
)

func main() {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewGlossaryRepository(db)
	uc := usecase.NewGlossaryUsecase(repo)

	grpcServer := grpc.NewServer()
	glossaryServer := delivery.NewGlossaryServer(uc)

	glossary.RegisterGlossaryServiceServer(grpcServer, glossaryServer)

	listener, err := net.Listen("tcp", ":50051") // gRPC default port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
