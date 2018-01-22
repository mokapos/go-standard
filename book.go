package root

import "time"

// Book is an entity of a book inside a library.
type Book struct {
	ID          uint64    `db:"id"`
	Title       string    `db:"title"`
	PublishedAt time.Time `db:"published_at"`
	Pages       uint16    `db:"pages"`
	Timestampable
}

// BookInteractor is a collection of method that can be done to Book object.
type BookInteractor struct {
	bookRepo BookRepository
}

// BookRepository is a collection of method that can call Book table in the database.
type BookRepository interface {
	FindAll() ([]*Book, error)
	FindByTitle(title string) ([]*Book, error)
	Create(book *Book) error
	Delete(id uint64) error
}

// NewBookInteractor returns an instance of BookInteractor.
func NewBookInteractor(bookRepo BookRepository) *BookInteractor {
	return &BookInteractor{bookRepo: bookRepo}
}

// GroupByPublishedYear returns a report of how many book published at each year between 'start' and 'end' inclusive.
func (itrc *BookInteractor) GroupByPublishedYear(start uint16, end uint16) ([]*PublishedYearReport, error) {
	allBooks, err := itrc.bookRepo.FindAll()
	if err != nil {
		return nil, err
	}
	// Use 'make' initializer on map, because map is nullable
	// However, array doesn't need to be initialized because array is initialized implicitly as an empty array
	bookPerYear := make(map[uint16][]string)
	for _, v := range allBooks {
		publishedYear := uint16(v.PublishedAt.Year())
		bookPerYear[publishedYear] = append(bookPerYear[publishedYear], v.Title)
	}
	var result []*PublishedYearReport
	for k, v := range bookPerYear {
		result = append(result, &PublishedYearReport{Year: k, Titles: v})
	}
	return result, nil
}

// FindByTitle returns any book whose title match the given string.
func (itrc *BookInteractor) FindByTitle(title string) ([]*Book, error) {
	return itrc.bookRepo.FindByTitle(title)
}

// Create function inserts a new book entity to the database.
func (itrc *BookInteractor) Create(book *Book) error {
	return itrc.bookRepo.Create(book)
}

// Delete function soft deletes a book with corresponding id.
func (itrc *BookInteractor) Delete(id uint64) error {
	return itrc.bookRepo.Delete(id)
}
