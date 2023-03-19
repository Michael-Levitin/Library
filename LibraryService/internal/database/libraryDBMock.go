package database

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type libraryDBMock struct {
}

func NewLibraryDBMock() *libraryDBMock {
	return &libraryDBMock{}
}

func (l libraryDBMock) GetAuthorLike(ctx context.Context, name string) (*[]ob.BookDB, error) {
	if name == "Man" {
		return &[]ob.BookDB{
			{Name: "Alexander Belyaev", Title: "Amphibian Man"},
			{Name: "Alexander Pushkin", Title: "The Bronze Horseman"},
			{Name: "Mikhail Sholokhov", Title: "The fate of man"},
			{Name: "Ernest Hemingway", Title: "The Old Man and the Sea"},
		}, nil
	}
	return &[]ob.BookDB{}, nil
}

func (l libraryDBMock) GetAuthorExact(ctx context.Context, title string) (*[]ob.BookDB, error) {
	if title == "Man" { // бд не содержит книг с точным названием "Man"
		return &[]ob.BookDB{}, nil
	}
	if title == "Amphibian Man" { // а такое название есть - возвращаем
		return &[]ob.BookDB{
			{Name: "Alexander Belyaev", Title: "Amphibian Man"},
		}, nil
	}
	return nil, ob.SomeError
}

func (l libraryDBMock) GetTitleLike(ctx context.Context, author string) (*[]ob.BookDB, error) {
	if author == "Chehov" { // "ищем" и возвращаем по частичному значение
		return &[]ob.BookDB{
			{Name: "Anton Chekhov", Title: "The Cherry Orchard"},
			{Name: "Anton Chekhov", Title: "Hunting Drama"},
			{Name: "Anton Chekhov", Title: "Uncle Vanya"},
			{Name: "Anton Chekhov", Title: "Ward No. 6"},
			{Name: "Anton Chekhov", Title: "Stories"},
			{Name: "Anton Chekhov", Title: "Three Sisters"},
			{Name: "Anton Chekhov", Title: "Seagull"}}, nil
	}

	if author == "Tolstoy" { // "ищем" и возвращаем по частичному значение
		return &[]ob.BookDB{
			{Name: "Alexey Tolstoy", Title: "Peter the Great"},
			{Name: "Alexey Tolstoy", Title: "Going through the throes"},
			{Name: "Lev Tolstoy", Title: "Anna Karenina"},
			{Name: "Lev Tolstoy", Title: "War and Peace"},
			{Name: "Lev Tolstoy", Title: "Resurrection"},
			{Name: "Lev Tolstoy", Title: "Childhood. Adolescence. Youth"},
			{Name: "Lev Tolstoy", Title: "Prisoner of the Caucasus"},
			{Name: "Lev Tolstoy", Title: "Hadji Murad"}}, nil
	}

	return &[]ob.BookDB{}, nil
}

func (l libraryDBMock) GetTitleExact(ctx context.Context, author string) (*[]ob.BookDB, error) {
	if author == "Chehov" { // бд не содержит книг с точным автором "Chehov"
		return &[]ob.BookDB{}, nil
	}

	if author == "Alexander Belyaev" { // а такое значение есть
		return &[]ob.BookDB{
			{Name: "Alexander Belyaev", Title: "Amphibian Man"},
		}, nil
	}

	if author == "Tolstoy" { // бд не содержит книг с точным автором "Tolstoy"
		return &[]ob.BookDB{}, nil
	}

	if author == "Anton Chekhov" { // а такое значение есть
		return &[]ob.BookDB{
			{Name: "Anton Chekhov", Title: "The Cherry Orchard"},
			{Name: "Anton Chekhov", Title: "Hunting Drama"},
			{Name: "Anton Chekhov", Title: "Uncle Vanya"},
			{Name: "Anton Chekhov", Title: "Ward No. 6"},
			{Name: "Anton Chekhov", Title: "Stories"},
			{Name: "Anton Chekhov", Title: "Three Sisters"},
			{Name: "Anton Chekhov", Title: "Seagull"},
		}, nil
	}

	return nil, ob.SomeError
}
