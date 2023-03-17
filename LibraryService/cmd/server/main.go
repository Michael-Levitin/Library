package main

import (
	"fmt"
	"github.com/Michael-Levitin/Library/LibraryService/internal/database"
	"github.com/Michael-Levitin/Library/LibraryService/internal/library"
	"github.com/Michael-Levitin/Library/LibraryService/internal/logic"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	// подключаемся к базе данных
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3307)/library")
	if err != nil {
		log.Println("error connecting to database: ", err)
		os.Exit(1)
	}
	log.Println("connected to database")
	defer db.Close()

	// готовимся принимать сообщения от клиента на порту 5000
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}
	log.Println("tcp listener started at port: ", port)

	grpcServer := grpc.NewServer()                            // создаем новый grpc сервер
	libraryDB := database.NewLibraryDB(db)                    // подключаем бд
	libraryLogic := logic.NewLibraryLogic(libraryDB)          // подключаем бд к логике...
	libraryServer := library.NewLibraryServer(*libraryLogic)  // ... а логику в библиотеку
	pb.RegisterLibrarySearchServer(grpcServer, libraryServer) // регистрируем сервис библиотеки в grpc

	if err := grpcServer.Serve(lis); err != nil { // передаем полученные от клиента данные
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
