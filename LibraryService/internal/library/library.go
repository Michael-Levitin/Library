package library

import (
	"context"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"log"
)

type Library interface {
	GetAuthor(context.Context, *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error)
	GetBooks(context.Context, *pb.GetBooksRequest) (*pb.GetBooksResponse, error)
}

type Server struct {
	pb.LibrarySearchServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetAuthor(ctx context.Context, in *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	log.Println("getting author for", in)
	return &pb.GetAuthorResponse{Books: []*pb.Book{
		{Name: "Александр Беляев", Title: "Человек-амфибия"},
	}}, nil
}

func (s *Server) GetBooks(ctx context.Context, in *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	log.Println("getting books for", in)
	return &pb.GetBooksResponse{Books: []*pb.Book{
		{Name: "Виктор Гюго", Title: "Отверженные"},
		{Name: "Виктор Гюго", Title: "Собор Парижской Богоматери"}},
	}, nil
}
