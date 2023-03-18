package database

import (
	"context"
	"fmt"
	"log"

	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	"github.com/jmoiron/sqlx"
)

const (
	authorLikeQuery = `select name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE title like '%Man%';`

	authorExactQuery = `select name
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE title = 'Amphibian Man';`

	titleLikeQuery = `select name, title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE name like '%Tols%';`

	titleExactQuery = `select title
FROM library.books
JOIN library.authors a on books.author_id = a.id
WHERE name = 'Chechov';`
)

type LibraryDB struct {
	db *sqlx.DB
}

func NewLibraryDB(db *sqlx.DB) *LibraryDB {
	return &LibraryDB{db: db}
}

func (l LibraryDB) GetAuthorLike(ctx context.Context, name string) (*[]ob.BookDB, error) {
	log.Println("database - getting books for name", name)
	return l.FuncName(authorLikeQuery, name)
}

func (l LibraryDB) GetAuthorExact(ctx context.Context, name string) (*[]ob.BookDB, error) {
	log.Println("database - getting books for name", name)
	return l.FuncName(authorExactQuery, name)
}
func (l LibraryDB) GetTitleLike(ctx context.Context, title string) (*[]ob.BookDB, error) {
	log.Println("database - getting books for title", title)
	return l.FuncName(titleLikeQuery, title)
}
func (l LibraryDB) GetTitleExact(ctx context.Context, title string) (*[]ob.BookDB, error) {
	log.Println("database - getting books for title", title)
	return l.FuncName(titleExactQuery, title)
}

func (l LibraryDB) FuncName(query, placeholder string) (*[]ob.BookDB, error) {
	rows, err := l.db.Queryx(query)
	log.Println(rows, err)
	//if err != sql.ErrNoRows {
	if err != nil {
		fmt.Println("error executing query:", err)
		return &[]ob.BookDB{}, err
	}
	defer rows.Close()

	var books []ob.BookDB
	//var book ob.BookDB

	// обработка результатов запроса
	for rows.Next() {
		var name, title string
		//err = rows.Scan(&book.Name, &book.Title)
		err = rows.Scan(&name, &title)
		if err != nil {
			fmt.Println("error scanning row:", err)
			return &[]ob.BookDB{}, err
		}
		//log.Println(name, title)
		books = append(books, ob.BookDB{name, title})
	}
	if err = rows.Err(); err != nil {
		fmt.Println("error scanning rows:", err)
		return &[]ob.BookDB{}, err
	}
	return &books, err
}
