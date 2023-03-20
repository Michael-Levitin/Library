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

type bSearch struct {
	sel, query string
}

func main() {
	// подключаемся к grpc серверу
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewLibrarySearchClient(conn) // передаем подключение в сервис

	booksSearch := []bSearch{ // готовим запросы
		{sel: "A", query: "Tolstoy"},
		{sel: "a", query: "Lev Tolstoy"},
		{sel: "A", query: "Chekhov"},
		{sel: "A", query: "12312"},
		{sel: "t", query: "Man"},
		{sel: "T", query: "Amphibian Man"},
		{sel: "T", query: "45554"},
	}

	for _, b := range booksSearch { // запускаем
		selector(b.sel, b.query, grpcClient)
	}
}

// выбираем нужную функцию на основании запроса
func selector(sel, query string, grpcClient pb.LibrarySearchClient) {
	switch sel {
	case "t":
		fallthrough
	case "T":
		books, err := grpcClient.GetAuthor(context.TODO(), &pb.GetAuthorRequest{
			Title: query,
		})
		fmt.Println("Looking for books with title:", query)
		printAnswer(books.Books, err)
	case "a":
		fallthrough
	case "A":
		books, err := grpcClient.GetBooks(context.TODO(), &pb.GetBooksRequest{
			Name: query,
		})
		fmt.Println("Looking for books by author: ", query)
		printAnswer(books.Books, err)
	default:
		fmt.Println("Unknown selector ", sel)
	}
}

// печатаем результат
func printAnswer(books []*pb.Book, err error) {
	if err != nil {
		log.Println("failed to execute request: ", err)
	}
	if books != nil && len(books) != 0 {
		fmt.Println("Author\t\tTitle")
		fmt.Println("---------------------------------------")
		for _, book := range books {
			fmt.Println(book.Name, "-", book.Title)
		}
	} else {
		fmt.Println(" - Nothing's found")
	}
	fmt.Println("==============================================")
}
