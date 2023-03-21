package database

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type LibraryDbI interface {
	GetAuthorLike(ctx context.Context, s string) (*[]ob.BookDB, error)
	GetTitleLike(ctx context.Context, s string) (*[]ob.BookDB, error)
}
