package logic

import (
	"context"
	"github.com/Michael-Levitin/Library/LibraryService/internal/database"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type LibraryLogic struct {
	LibraryDB database.LibraryDbI
}

// подключаем интерфейс БД в новую логику
func NewLibraryLogic(LibraryDB database.LibraryDbI) *LibraryLogic {
	return &LibraryLogic{LibraryDB: LibraryDB}
}

func (l LibraryLogic) GetAuthor(ctx context.Context, title string) (*[]ob.BookDB, error) {
	books, err := l.LibraryDB.GetAuthorLike(ctx, title) // делаем запрос в БД
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (l LibraryLogic) GetTitle(ctx context.Context, name string) (*[]ob.BookDB, error) {
	books, err := l.LibraryDB.GetTitleLike(ctx, name)
	if err != nil {
		return nil, err
	}
	return books, nil
}
