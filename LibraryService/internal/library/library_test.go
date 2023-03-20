package library

import (
	"context"
	"github.com/Michael-Levitin/Library/LibraryService/internal/logic"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	pb "github.com/Michael-Levitin/Library/LibraryService/proto"
	"reflect"
	"testing"
)

func TestLibraryServer_GetAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.GetAuthorRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetAuthorResponse
		wantErr error
	}{
		{
			name: "GetAuthor",
			args: args{
				ctx: context.Background(),
				in:  &pb.GetAuthorRequest{Title: "Amphibian Man"},
			},
			want: &pb.GetAuthorResponse{
				Books: []*pb.Book{
					{Name: "Alexander Belyaev", Title: "Amphibian Man"},
				}},
			wantErr: nil,
		},
		{
			name: "Error",
			args: args{
				ctx: context.Background(),
				in:  &pb.GetAuthorRequest{Title: "error"},
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	s := NewLibraryServer(logic.NewLibraryLogicMock())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetAuthor(tt.args.ctx, tt.args.in)
			if err != tt.wantErr {
				t.Errorf("GetAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLibraryServer_GetBooks(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.GetBooksRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.GetBooksResponse
		wantErr error
	}{
		{
			name: "GetAuthor",
			args: args{
				ctx: context.Background(),
				in:  &pb.GetBooksRequest{Name: "Belyaev"},
			},
			want: &pb.GetBooksResponse{
				Books: []*pb.Book{
					{Name: "Alexander Belyaev", Title: "Amphibian Man"},
				}},
			wantErr: nil,
		},
		{
			name: "Error",
			args: args{
				ctx: context.Background(),
				in:  &pb.GetBooksRequest{Name: "error"},
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	s := NewLibraryServer(logic.NewLibraryLogicMock())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetBooks(tt.args.ctx, tt.args.in)
			if err != tt.wantErr {
				t.Errorf("GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transferBooks(t *testing.T) {
	tests := []struct {
		name  string
		books *[]ob.BookDB
		want  []*pb.Book
	}{
		{
			name: "Data",
			books: &[]ob.BookDB{
				{Name: "Erich Maria Remarque", Title: "Three comrades"},
				{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			want: []*pb.Book{
				{Name: "Erich Maria Remarque", Title: "Three comrades"},
				{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"},
			},
		},
		{
			name:  "Empty",
			books: &[]ob.BookDB{},
			want:  []*pb.Book{},
		},
		{
			name:  "Nil",
			books: nil,
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transferBooks(tt.books); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transferBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
