package api

import (
	"encoding/json"
	"net/http"
)

//Book with name, author and isbn
type Book struct {
	Title 	string 		`json:"title"`
	Author 	string		`json:"author"`
	ISBN 	string		`json:"isbn"`
}

func (b Book) toJSON() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func fromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	{
		Title:  "American Pyscho",
		Author: "Brett Easton Ellis",
		ISBN:   "9781407443591",
	},
	{
		Title:  "My Book",
		Author: "Dave Rowntree",
		ISBN:   "123456789",
	},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
