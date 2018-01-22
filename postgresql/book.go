package postgresql

import (
	"github.com/mokapos/go-standard"
)

// BookRepository is a repository for admin table.
type BookRepository struct {
	Master *DB
	Slave  *DB
}

// NewBookRepository creates a new repository for admin table.
func NewBookRepository(master *DB, slave *DB) *BookRepository {
	return &BookRepository{
		Master: master,
		Slave:  slave,
	}
}

// FindAll returns all book.
func (repo *BookRepository) FindAll() ([]*standard.Book, error) {
	var books []*standard.Book
	err := repo.Slave.Select(books, "SELECT * FROM books WHERE deleted_at IS NOT NULL")
	if err != nil {
		return nil, err
	}
	return books, nil
}

// FindByTitle returns a book whose title is matching given title case insensitively
func (repo *BookRepository) FindByTitle(title string) ([]*standard.Book, error) {
	var books []*standard.Book
	// https://stackoverflow.com/questions/25214459/go-postgresql-like-query
	err := repo.Slave.Select(books, "SELECT * FROM books WHERE title ILIKE '%'||$1||'%' AND deleted_at IS NOT NULL", title)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Create adds a new row in books table.
func (repo *BookRepository) Create(book *standard.Book) error {
	_, err := repo.Master.Query("INSERT INTO books(title, published_at, pages, created_at) VALUES ($1, $2, $3, NOW())",
		book.Title, book.PublishedAt, book.Pages)
	return err
}

// Delete soft deletes a book with given id.
func (repo *BookRepository) Delete(id uint64) error {
	_, err := repo.Master.Query("UPDATE books SET updated_at=NOW(), deleted_at=NOW() WHERE id=$1", id)
	return err
}
