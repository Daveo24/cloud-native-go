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

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}
