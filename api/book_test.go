package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title:"American Pyscho", Author:"Brett Easton Ellis", ISBN:"9781407443591"}
	json := book.toJSON()
	assert.Equal(t, `{"title":"American Pyscho","author":"Brett Easton Ellis","isbn":"9781407443591"}`, string(json),
		"Book JSON marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"American Pyscho","author":"Brett Easton Ellis","isbn":"9781407443591"}`)
	book := fromJSON(json)

	assert.Equal(t, Book{Title:"American Pyscho", Author:"Brett Easton Ellis", ISBN:"9781407443591"},
		book, "Book JSON unmarshalling wrong.")
}