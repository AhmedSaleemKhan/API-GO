package models

////defining structure
type Book struct { //json &databse varibales name
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
