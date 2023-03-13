package main

import (
	"fmt"
	"log"
	"net"
	"os"

	libServer "github.com/Michael-Levitin/Library/LibraryService/internal/library"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"google.golang.org/grpc"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}
	log.Println("tcp listener started at port: ", port)
	grpcServer := grpc.NewServer()
	libServiceServer := libServer.NewServer()

	// registering library service server into grpc server
	pb.RegisterLibrarySearchServer(grpcServer, libServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
