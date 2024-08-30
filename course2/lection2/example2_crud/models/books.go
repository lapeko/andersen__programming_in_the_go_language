package models

type Book struct {
	Name      string
	Publisher string
	Author    *Author
}

type Author struct {
	FirstName   string
	LastName    string
	YearOfBirth uint16
}

var DB = map[uint]*Book{}
var counter = uint(0)

func init() {
	counter++
	DB[counter] = &Book{
		Name:      "Read People Like a Book",
		Publisher: "Independently published",
		Author: &Author{
			FirstName:   "Patrick",
			LastName:    "King",
			YearOfBirth: 1981,
		},
	}
}

func GetBookById(id uint) (book *Book, ok bool) {
	book, ok = DB[id]
	return
}

func GetAllBooks() (books []*Book) {
	for _, book := range DB {
		books = append(books, book)
	}
	return
}

func CreateBook(book *Book) {
	counter++
	for _, exists := DB[counter]; exists; counter++ {
	}
	DB[counter] = book
}

func PutBook(book *Book, id uint) {
	DB[id] = book
}

func DeleteBook(id uint) {
	delete(DB, id)
}
