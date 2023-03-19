package objects

import "errors"

type BookDB struct {
	Name  string
	Title string
}

var SomeError = errors.New("something wrong")
