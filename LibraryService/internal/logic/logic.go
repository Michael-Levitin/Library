package logic

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type LibraryLogic struct {
	LibraryDB LibraryDbI
}

func NewLibraryLogic(LibraryDB LibraryDbI) *LibraryLogic {
	return &LibraryLogic{LibraryDB: LibraryDB}
}

func (l LibraryLogic) GetAuthor(ctx context.Context, title string) (*[]ob.BookDB, error) {
	books, err := l.LibraryDB.GetAuthorExact(ctx, title) // аналогично GetTitle
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
	books, err := l.LibraryDB.GetTitleExact(ctx, name) // сначала пробуем найти полное совпадение
	if err != nil {                                    // если есть ошибка возвращаем ее
		return nil, err
	}

	if len(*books) > 0 { // если длина слайса книг > 0 - возвращаем его
		return books, err
	}
	// если длина слайса книг = 0 - пробуем найти частичное совпадение
	books, err = l.LibraryDB.GetTitleLike(ctx, name)
	if err != nil {
		return nil, err
	}
	return books, nil
}
