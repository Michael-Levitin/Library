package library

import (
	"context"
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
	srv := NewServer()
	pb.RegisterLibrarySearchServer(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
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
	in := &pb.GetAuthorRequest{Title: "Человек-амфибия"}
	resp, err := client.GetAuthor(ctx, in)
	assert.NoError(t, err)
	books := []*pb.Book{
		{Name: "Александр Беляев", Title: "Человек-амфибия"},
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
	in := &pb.GetBooksRequest{Name: "Виктор Гюго"}
	resp, err := client.GetBooks(ctx, in)
	assert.NoError(t, err)
	books := []*pb.Book{
		{Name: "Виктор Гюго", Title: "Отверженные"},
		{Name: "Виктор Гюго", Title: "Собор Парижской Богоматери"},
	}
	assert.Equal(t, books, resp.Books)
}
