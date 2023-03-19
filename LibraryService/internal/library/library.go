package library

import (
	"context"
	"github.com/Michael-Levitin/Library/LibraryService/internal/logic"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"log"
)

type LibraryServer struct {
	pb.UnimplementedLibrarySearchServer
	logic logic.LibraryLogicI
}

func NewLibraryServer(logic logic.LibraryLogicI) *LibraryServer {
	return &LibraryServer{logic: logic}
}

func (s *LibraryServer) GetAuthor(ctx context.Context, in *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	log.Println("getting author for", in)
	title := in.GetTitle()
	books, err := s.logic.GetAuthor(ctx, title)
	return &pb.GetAuthorResponse{Books: transferBooks(books)}, err
}

func (s *LibraryServer) GetBooks(ctx context.Context, in *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	log.Println("getting books for", in)
	author := in.GetName()
	books, err := s.logic.GetTitle(ctx, author)
	return &pb.GetBooksResponse{Books: transferBooks(books)}, err
}

func transferBooks(books *[]ob.BookDB) []*pb.Book {
	booksPB := make([]*pb.Book, len(*books))
	for i := 0; i < len(*books); i++ {
		booksPB[i] = &pb.Book{
			Name:  (*books)[i].Name,
			Title: (*books)[i].Title,
		}
	}
	return booksPB
}
