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
		LibraryDB LibraryDbI
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
			want:    &[]ob.BookDB{{"Alexander Belyaev", "Amphibian Man"}},
			wantErr: nil,
		},
		{
			name: "Partial Title",
			args: args{ctx: context.Background(), title: "Man"},
			want: &[]ob.BookDB{{"Alexander Belyaev", "Amphibian Man"},
				{"Alexander Pushkin", "The Bronze Horseman"},
				{"Mikhail Sholokhov", "The fate of man"},
				{"Ernest Hemingway", "The Old Man and the Sea"}},
			wantErr: nil,
		},
		{
			name:    "Error",
			args:    args{ctx: context.Background(), title: ""},
			want:    nil,
			wantErr: db.SomeError,
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
		LibraryDB LibraryDbI
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
			want:    &[]ob.BookDB{{"Alexander Belyaev", "Amphibian Man"}},
			wantErr: nil,
		},
		{
			name: "Partial Name",
			args: args{ctx: context.Background(), name: "Chehov"},
			want: &[]ob.BookDB{{"Anton Chekhov", "The Cherry Orchard"},
				{"Anton Chekhov", "Hunting Drama"},
				{"Anton Chekhov", "Uncle Vanya"},
				{"Anton Chekhov", "Ward No. 6"},
				{"Anton Chekhov", "Stories"},
				{"Anton Chekhov", "Three Sisters"},
				{"Anton Chekhov", "Seagull"}},
			wantErr: nil,
		},
		{
			name: "Partial Name - Many authors",
			args: args{ctx: context.Background(), name: "Tolstoy"},
			want: &[]ob.BookDB{{"Alexey Tolstoy", "Peter the Great"},
				{"Alexey Tolstoy", "Going through the throes"},
				{"Lev Tolstoy", "Anna Karenina"},
				{"Lev Tolstoy", "War and Peace"},
				{"Lev Tolstoy", "Resurrection"},
				{"Lev Tolstoy", "Childhood. Adolescence. Youth"},
				{"Lev Tolstoy", "Prisoner of the Caucasus"},
				{"Lev Tolstoy", "Hadji Murad"}},
			wantErr: nil,
		},
		{
			name:    "Error",
			args:    args{ctx: context.Background(), name: ""},
			want:    nil,
			wantErr: db.SomeError,
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
