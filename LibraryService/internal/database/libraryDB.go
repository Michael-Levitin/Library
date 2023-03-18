package database

import (
	"context"
	"log"

	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	"github.com/jmoiron/sqlx"
)

const (
	authorLikeQuery = `SELECT name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE title like ?;`

	authorExactQuery = `SELECT name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE title = ?;`

	titleLikeQuery = `SELECT name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE name like ?;`

	titleExactQuery = `SELECT name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE name = ?;`
)

type LibraryDB struct {
	db *sqlx.DB
}

type BookDb struct { // структура для получения данный из бд
	Name  string `db:"name"`
	Title string `db:"title"`
}

func NewLibraryDB(db *sqlx.DB) *LibraryDB {
	return &LibraryDB{db: db}
}

func (l LibraryDB) GetAuthorLike(ctx context.Context, name string) (*[]ob.BookDB, error) {
	name = "%" + name + "%" // для поиска совпадений добавляем % около строки
	return l.queryDo(authorLikeQuery, name)
}

func (l LibraryDB) GetAuthorExact(ctx context.Context, name string) (*[]ob.BookDB, error) {
	return l.queryDo(authorExactQuery, name)
}
func (l LibraryDB) GetTitleLike(ctx context.Context, title string) (*[]ob.BookDB, error) {
	title = "%" + title + "%"
	return l.queryDo(titleLikeQuery, title)
}
func (l LibraryDB) GetTitleExact(ctx context.Context, title string) (*[]ob.BookDB, error) {
	return l.queryDo(titleExactQuery, title)
}

func (l LibraryDB) queryDo(query, placeholder string) (*[]ob.BookDB, error) {
	var booksDb []BookDb // создаем переменную для результатов
	err := l.db.Select(&booksDb, query, placeholder)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var books = make([]ob.BookDB, len(booksDb))
	for i := 0; i < len(booksDb); i++ {
		books[i] = ob.BookDB{booksDb[i].Name, booksDb[i].Title}
	}
	return &books, err
}
