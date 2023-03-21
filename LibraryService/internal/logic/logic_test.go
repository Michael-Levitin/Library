package logic

import (
	"context"
	db "github.com/Michael-Levitin/Library/LibraryService/internal/database"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	"reflect"
	"testing"
)

func TestLibraryLogic_GetAuthor(t *testing.T) {
	type fields struct {
		LibraryDB db.LibraryDbI
	}
	type args struct {
		ctx   context.Context
		title string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name:    "Full Title",
			args:    args{ctx: context.Background(), title: "Amphibian Man"},
			want:    &[]ob.BookDB{{Name: "Alexander Belyaev", Title: "Amphibian Man"}},
			wantErr: nil,
		},
		{
			name: "Partial Title",
			args: args{ctx: context.Background(), title: "Man"},
			want: &[]ob.BookDB{{"Alexander Belyaev", "Amphibian Man"},
				{Name: "Alexander Pushkin", Title: "The Bronze Horseman"},
				{Name: "Mikhail Sholokhov", Title: "The fate of man"},
				{Name: "Ernest Hemingway", Title: "The Old Man and the Sea"}},
			wantErr: nil,
		},
		{
			name:    "Error",
			args:    args{ctx: context.Background(), title: ""},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}
	l := LibraryLogic{
		LibraryDB: db.NewLibraryDBMock(),
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := l.GetAuthor(tt.args.ctx, tt.args.title)
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

func TestLibraryLogic_GetTitle(t *testing.T) {
	type fields struct {
		LibraryDB db.LibraryDbI
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name:    "Full Name",
			args:    args{ctx: context.Background(), name: "Alexander Belyaev"},
			want:    &[]ob.BookDB{{Name: "Alexander Belyaev", Title: "Amphibian Man"}},
			wantErr: nil,
		},
		{
			name: "Partial Name",
			args: args{ctx: context.Background(), name: "Chehov"},
			want: &[]ob.BookDB{{Name: "Anton Chekhov", Title: "The Cherry Orchard"},
				{Name: "Anton Chekhov", Title: "Hunting Drama"},
				{Name: "Anton Chekhov", Title: "Uncle Vanya"},
				{Name: "Anton Chekhov", Title: "Ward No. 6"},
				{Name: "Anton Chekhov", Title: "Stories"},
				{Name: "Anton Chekhov", Title: "Three Sisters"},
				{Name: "Anton Chekhov", Title: "Seagull"}},
			wantErr: nil,
		},
		{
			name: "Partial Name - Many authors",
			args: args{ctx: context.Background(), name: "Tolstoy"},
			want: &[]ob.BookDB{{Name: "Alexey Tolstoy", Title: "Peter the Great"},
				{Name: "Alexey Tolstoy", Title: "Going through the throes"},
				{Name: "Lev Tolstoy", Title: "Anna Karenina"},
				{Name: "Lev Tolstoy", Title: "War and Peace"},
				{Name: "Lev Tolstoy", Title: "Resurrection"},
				{Name: "Lev Tolstoy", Title: "Childhood. Adolescence. Youth"},
				{Name: "Lev Tolstoy", Title: "Prisoner of the Caucasus"},
				{Name: "Lev Tolstoy", Title: "Hadji Murad"}},
			wantErr: nil,
		},
		{
			name:    "Error",
			args:    args{ctx: context.Background(), name: ""},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	l := LibraryLogic{
		LibraryDB: db.NewLibraryDBMock(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := l.GetTitle(tt.args.ctx, tt.args.name)
			if err != tt.wantErr {
				t.Errorf("GetTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
