package database

import (
	"context"
	"errors"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

var SomeError = errors.New("something wrong")

type LibraryMockDB struct {
}

func NewLibraryMockDB() *LibraryMockDB {
	return &LibraryMockDB{}
}

func (l LibraryMockDB) GetAuthorLike(ctx context.Context, name string) (*[]ob.BookDB, error) {
	if name == "Man" {
		return &[]ob.BookDB{
			{"Alexander Belyaev", "Amphibian Man"},
			{"Alexander Pushkin", "The Bronze Horseman"},
			{"Mikhail Sholokhov", "The fate of man"},
			{"Ernest Hemingway", "The Old Man and the Sea"},
		}, nil
	}
	return &[]ob.BookDB{}, nil
}

func (l LibraryMockDB) GetAuthorExact(ctx context.Context, title string) (*[]ob.BookDB, error) {
	if title == "Man" { // бд не содержит книг с точным названием "Man"
		return &[]ob.BookDB{}, nil
	}
	if title == "Amphibian Man" { // а такое название есть - возвращаем
		return &[]ob.BookDB{
			{"Alexander Belyaev", "Amphibian Man"},
		}, nil
	}
	return nil, SomeError
}

func (l LibraryMockDB) GetTitleLike(ctx context.Context, author string) (*[]ob.BookDB, error) {
	if author == "Chehov" { // "ищем" и возвращаем по частичному значение
		return &[]ob.BookDB{
			{"Anton Chekhov", "The Cherry Orchard"},
			{"Anton Chekhov", "Hunting Drama"},
			{"Anton Chekhov", "Uncle Vanya"},
			{"Anton Chekhov", "Ward No. 6"},
			{"Anton Chekhov", "Stories"},
			{"Anton Chekhov", "Three Sisters"},
			{"Anton Chekhov", "Seagull"}}, nil
	}

	if author == "Tolstoy" { // "ищем" и возвращаем по частичному значение
		return &[]ob.BookDB{
			{"Alexey Tolstoy", "Peter the Great"},
			{"Alexey Tolstoy", "Going through the throes"},
			{"Lev Tolstoy", "Anna Karenina"},
			{"Lev Tolstoy", "War and Peace"},
			{"Lev Tolstoy", "Resurrection"},
			{"Lev Tolstoy", "Childhood. Adolescence. Youth"},
			{"Lev Tolstoy", "Prisoner of the Caucasus"},
			{"Lev Tolstoy", "Hadji Murad"}}, nil
	}

	return &[]ob.BookDB{}, nil
}

func (l LibraryMockDB) GetTitleExact(ctx context.Context, author string) (*[]ob.BookDB, error) {
	if author == "Chehov" { // бд не содержит книг с точным автором "Chehov"
		return &[]ob.BookDB{}, nil
	}

	if author == "Alexander Belyaev" { // а такое значение есть
		return &[]ob.BookDB{
			{"Alexander Belyaev", "Amphibian Man"},
		}, nil
	}

	if author == "Tolstoy" { // бд не содержит книг с точным автором "Tolstoy"
		return &[]ob.BookDB{}, nil
	}

	if author == "Anton Chekhov" { // а такое значение есть
		return &[]ob.BookDB{
			{"Anton Chekhov", "The Cherry Orchard"},
			{"Anton Chekhov", "Hunting Drama"},
			{"Anton Chekhov", "Uncle Vanya"},
			{"Anton Chekhov", "Ward No. 6"},
			{"Anton Chekhov", "Stories"},
			{"Anton Chekhov", "Three Sisters"},
			{"Anton Chekhov", "Seagull"},
		}, nil
	}

	return nil, SomeError
}
