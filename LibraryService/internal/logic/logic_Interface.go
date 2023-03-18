package logic

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
)

type LibraryLogicI interface {
	GetAuthor(ctx context.Context, s string) (*[]ob.BookDB, error)
	GetTitle(ctx context.Context, s string) (*[]ob.BookDB, error)
}
