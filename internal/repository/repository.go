package repository

import (
    "fmt"
    "my-go-server/config"
)

// Book struct represents a book entity
type Book struct {
    BookID   int    `json:"bookID"`
    BookName string `json:"bookName"`
    Cost     int    `json:"cost"`
}

// GetBooks retrieves all books from the database
func GetUsers() ([]Book, error) {
    query := "SELECT bookID,bookName,cost FROM books"
    rows, err := config.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying books: %v", err)
    }
    defer rows.Close()

    var books []Book
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.BookID, &book.BookName, &book.Cost); err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    return books, nil // Return the books and nil if no error
}
