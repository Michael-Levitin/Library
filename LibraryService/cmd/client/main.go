package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	// подключаемся к grpc серверу
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewLibrarySearchClient(conn)

	booksByAuthor, err := grpcClient.GetBooks(context.TODO(), &pb.GetBooksRequest{
		Name: "Tolstoy",
	})
	printAnswer(booksByAuthor.Books, err)

	booksByAuthor, err = grpcClient.GetBooks(context.TODO(), &pb.GetBooksRequest{
		Name: "Lev Tolstoy",
	})
	printAnswer(booksByAuthor.Books, err)

	booksByAuthor, err = grpcClient.GetBooks(context.TODO(), &pb.GetBooksRequest{
		Name: "Chekhov",
	})
	printAnswer(booksByAuthor.Books, err)

	booksByAuthor, err = grpcClient.GetBooks(context.TODO(), &pb.GetBooksRequest{
		Name: "1235",
	})
	printAnswer(booksByAuthor.Books, err)

	booksByTitle, err := grpcClient.GetAuthor(context.TODO(), &pb.GetAuthorRequest{
		Title: "Man",
	})
	printAnswer(booksByTitle.Books, err)

	booksByTitle, err = grpcClient.GetAuthor(context.TODO(), &pb.GetAuthorRequest{
		Title: "Amphibian Man",
	})
	printAnswer(booksByTitle.Books, err)

	booksByTitle, err = grpcClient.GetAuthor(context.TODO(), &pb.GetAuthorRequest{
		Title: "45554",
	})
	printAnswer(booksByTitle.Books, err)
}

func printAnswer(books []*pb.Book, err error) {
	fmt.Println("==============================================")
	if err != nil {
		log.Println("failed to execute request: ", err)
	}
	if books == nil || len(books) == 0 {
		return
	}
	fmt.Println("\nAuthor\t\tTitle")
	for _, book := range books {
		fmt.Println(book.Name, "-", book.Title)
	}
}
