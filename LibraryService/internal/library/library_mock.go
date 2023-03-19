package library

import (
	"context"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
)

type LibraryMockServer struct {
	pb.UnimplementedLibrarySearchServer
}

func NewLibraryMockServer() *LibraryMockServer {
	return &LibraryMockServer{}
}

func (s *LibraryMockServer) GetAuthor(ctx context.Context, in *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	return &pb.GetAuthorResponse{Books: []*pb.Book{
		{Name: "Alexander Belyaev", Title: "Amphibian Man"},
	}}, nil
}

func (s *LibraryMockServer) GetBooks(ctx context.Context, in *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	return &pb.GetBooksResponse{Books: []*pb.Book{
		{Name: "Erich Maria Remarque", Title: "Three comrades"},
		{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"},
		{Name: "Erich Maria Remarque", Title: "Black Obelisk"},
	}}, nil
}
