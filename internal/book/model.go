package book

import "errors"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var ErrBookNotFound = errors.New("book not found")
