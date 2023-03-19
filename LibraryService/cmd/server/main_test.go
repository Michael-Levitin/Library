package main

import (
	"context"
	"github.com/Michael-Levitin/Library/LibraryService/internal/library"
	"log"
	"net"
	"testing"

	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	srv := library.NewLibraryMockServer()
	pb.RegisterLibrarySearchServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("LibraryServer exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func TestServer_GetAuthor(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibrarySearchClient(conn)
	in := &pb.GetAuthorRequest{Title: "Amphibian Man"}
	resp, err := client.GetAuthor(ctx, in)
	assert.NoError(t, err)
	books := []*pb.Book{
		{Name: "Alexander Belyaev", Title: "Amphibian Man"},
	}
	assert.Equal(t, books, resp.Books)
}

func TestServer_GetBooks(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewLibrarySearchClient(conn)
	in := &pb.GetBooksRequest{Name: "Erich Maria Remarque"}
	resp, err := client.GetBooks(ctx, in)
	assert.NoError(t, err)
	books := []*pb.Book{
		{Name: "Erich Maria Remarque", Title: "Three comrades"},
		{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"},
		{Name: "Erich Maria Remarque", Title: "Black Obelisk"},
	}
	assert.Equal(t, books, resp.Books)
}
