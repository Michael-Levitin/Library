package logic

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type LibraryLogicMock struct {
}

func NewLibraryLogicMock() *LibraryLogicMock {
	return &LibraryLogicMock{}
}

func (l LibraryLogicMock) GetAuthor(ctx context.Context, title string) (*[]ob.BookDB, error) {
	if title == "Amphibian Man" {
		return &[]ob.BookDB{
			{Name: "Alexander Belyaev", Title: "Amphibian Man"},
		}, nil
	}

	if title == "error" {
		return nil, ob.SomeError
	}

	return &[]ob.BookDB{}, nil
}

func (l LibraryLogicMock) GetTitle(ctx context.Context, name string) (*[]ob.BookDB, error) {
	if name == "Belyaev" {
		return &[]ob.BookDB{
			{Name: "Alexander Belyaev", Title: "Amphibian Man"},
		}, nil
	}

	if name == "error" {
		return nil, ob.SomeError
	}

	return &[]ob.BookDB{}, nil
}
