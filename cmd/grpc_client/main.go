package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/arvaliullin/go-restful-service/pkg/glossary"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	command := flag.String("command", "", "Command to execute: list, get, create")
	term := flag.String("term", "", "Term to interact with (for get and create commands)")
	description := flag.String("description", "", "Description for the create command")

	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGlossaryServiceClient(conn)

	switch *command {
	case "list":
		getAllTerms(client)
	case "get":
		if *term == "" {
			fmt.Println("Please provide a term using -term flag.")
			os.Exit(1)
		}
		getTermByTerm(client, *term)
	case "create":
		if *term == "" || *description == "" {
			fmt.Println("Please provide both term and description using -term and -description flags.")
			os.Exit(1)
		}
		createTerm(client, *term, *description)
	default:
		fmt.Println("Invalid command. Use one of: list, get, create")
		os.Exit(1)
	}
}

func getAllTerms(client pb.GlossaryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetAllTerms(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Could not get terms: %v", err)
	}

	for _, term := range r.Terms {
		fmt.Printf("%s: %s - %s\n", term.Id, term.Term, term.Description)
	}
}

func getTermByTerm(client pb.GlossaryServiceClient, term string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetTermByTerm(ctx, &pb.GlossaryTermRequest{Term: term})
	if err != nil {
		log.Fatalf("Could not get term: %v", err)
	}

	fmt.Printf("%s: %s - %s\n", r.Id, r.Term, r.Description)
}

func createTerm(client pb.GlossaryServiceClient, term, description string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.CreateTerm(ctx, &pb.GlossaryTerm{
		Term:        term,
		Description: description,
	})

	if err != nil {
		log.Fatalf("Could not create term: %v", err)
	}

	fmt.Printf("Term created: %s: %s - %s\n", r.Id, r.Term, r.Description)
}
