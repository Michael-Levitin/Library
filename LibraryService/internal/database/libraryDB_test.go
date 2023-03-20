package database

import (
	"context"
	ob "github.com/Michael-Levitin/Library/LibraryService/internal/objects"
	"reflect"
	"testing"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

// Тестируем функции БД используя SQLX Mock

func TestLibraryDB_GetAuthorLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx() // подготавливаем мок
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewLibraryDB(db) // подключаем смоканую ДБ

	type args struct { // цтруктура принимаемых аргументов
		ctx  context.Context
		name string
	}

	tests := []struct { // структура теста
		name    string       // имя теста
		s       *LibraryDB   // ссылка не репозиторий
		arg     args         // аргументы функции
		mock    func()       // функция мока
		want    *[]ob.BookDB // ответ, который хотим получить
		wantErr error        // ошибка, которую хотим получить
	}{
		{
			name: "OK",
			s:    s,
			arg: args{
				ctx:  context.Background(),
				name: "Remarque",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"}).
					AddRow("Erich Maria Remarque", "Three comrades").
					AddRow("Erich Maria Remarque", "Arc de Triomphe")
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want: &[]ob.BookDB{{Name: "Erich Maria Remarque", Title: "Three comrades"},
				{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			wantErr: nil,
		},

		{
			name: "Not found",
			s:    s,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"})
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{},
			wantErr: nil,
		},
		{
			name: "Error",
			s:    s,
			mock: func() {
				mock.ExpectQuery("SELECT").
					WillReturnError(ob.SomeError)
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()                                               // Запускаем мок
			got, err := tt.s.GetAuthorLike(tt.arg.ctx, tt.arg.name) // Запускаем тестируемую функцию
			if err != tt.wantErr {                                  // сравниваем ошибку
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) { // сравниваем ошибку полученные данные
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLibraryDB_GetAuthorExact(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewLibraryDB(db)

	type args struct {
		ctx  context.Context
		name string
	}

	tests := []struct {
		name    string
		s       *LibraryDB
		arg     args
		mock    func()
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name: "OK",
			s:    s,
			arg: args{
				ctx:  context.Background(),
				name: "Erich Maria Remarque",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"}).
					AddRow("Erich Maria Remarque", "Three comrades").
					AddRow("Erich Maria Remarque", "Arc de Triomphe")
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want: &[]ob.BookDB{{Name: "Erich Maria Remarque", Title: "Three comrades"},
				{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			wantErr: nil,
		},

		{
			name: "Not found",
			s:    s,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"})
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{},
			wantErr: nil,
		},
		{
			name: "Error",
			s:    s,
			mock: func() {
				mock.ExpectQuery("SELECT").
					WillReturnError(ob.SomeError)
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAuthorExact(tt.arg.ctx, tt.arg.name)
			if err != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLibraryDB_GetTitleLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewLibraryDB(db)

	type args struct {
		ctx   context.Context
		title string
	}

	tests := []struct {
		name    string
		s       *LibraryDB
		arg     args
		mock    func()
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name: "OK",
			s:    s,
			arg: args{
				ctx:   context.Background(),
				title: "Triomphe",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"}).
					AddRow("Erich Maria Remarque", "Arc de Triomphe")
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			wantErr: nil,
		},

		{
			name: "Not found",
			s:    s,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"})
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{},
			wantErr: nil,
		},
		{
			name: "Error",
			s:    s,
			mock: func() {
				mock.ExpectQuery("SELECT").
					WillReturnError(ob.SomeError)
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAuthorExact(tt.arg.ctx, tt.arg.title)
			if err != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLibraryDB_GetTitleExact(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewLibraryDB(db)

	type args struct {
		ctx   context.Context
		title string
	}

	tests := []struct {
		name    string
		s       *LibraryDB
		arg     args
		mock    func()
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name: "OK",
			s:    s,
			arg: args{
				ctx:   context.Background(),
				title: "Arc de Triomphe",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"}).
					AddRow("Erich Maria Remarque", "Arc de Triomphe")
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			wantErr: nil,
		},

		{
			name: "Not found",
			s:    s,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"})
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{},
			wantErr: nil,
		},
		{
			name: "Error",
			s:    s,
			mock: func() {
				mock.ExpectQuery("SELECT").
					WillReturnError(ob.SomeError)
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAuthorExact(tt.arg.ctx, tt.arg.title)
			if err != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLibraryDB_queryDo(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := NewLibraryDB(db)

	type args struct {
		query       string
		placeholder string
	}

	tests := []struct {
		name    string
		s       *LibraryDB
		arg     args
		mock    func()
		want    *[]ob.BookDB
		wantErr error
	}{
		{
			name: "OK",
			s:    s,
			arg: args{
				query:       "SELECT",
				placeholder: "Arc de Triomphe",
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"}).
					AddRow("Erich Maria Remarque", "Arc de Triomphe")
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{{Name: "Erich Maria Remarque", Title: "Arc de Triomphe"}},
			wantErr: nil,
		},

		{
			name: "Not found",
			s:    s,
			mock: func() {
				rows := sqlxmock.NewRows([]string{"name", "title"})
				mock.ExpectQuery("").
					WillReturnRows(rows)
			},
			want:    &[]ob.BookDB{},
			wantErr: nil,
		},
		{
			name: "Error",
			s:    s,
			mock: func() {
				mock.ExpectQuery("").
					WillReturnError(ob.SomeError)
			},
			want:    nil,
			wantErr: ob.SomeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.queryDo(tt.arg.query, tt.arg.placeholder)
			if err != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %+v, want %+v", got, tt.want)
			}
		})
	}
	//}
}
