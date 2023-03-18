package logic

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	"log"
)

type LibraryLogic struct {
	LibraryDB LibraryDbI
}

func NewLibraryLogic(LibraryDB LibraryDbI) *LibraryLogic {
	return &LibraryLogic{LibraryDB: LibraryDB}
}

func (l LibraryLogic) GetAuthor(ctx context.Context, title string) (*[]ob.BookDB, error) {
	log.Println("logic - getting books for title:", title)
	books, err := l.LibraryDB.GetAuthorExact(ctx, title)
	if err != nil {
		return &[]ob.BookDB{}, err
	}
	if len(*books) > 0 {
		return books, err
	}

	books, err = l.LibraryDB.GetAuthorLike(ctx, title)
	if err != nil {
		return &[]ob.BookDB{}, err
	}
	return books, err
}

func (l LibraryLogic) GetTitle(ctx context.Context, name string) (*[]ob.BookDB, error) {
	log.Println("logic - getting books for name", name)
	books, err := l.LibraryDB.GetTitleExact(ctx, name)
	if err != nil {
		return &[]ob.BookDB{}, err
	}

	if len(*books) > 0 {
		return books, err
	}

	books, err = l.LibraryDB.GetTitleLike(ctx, name)
	if err != nil {
		return &[]ob.BookDB{}, err
	}
	return books, err
}
