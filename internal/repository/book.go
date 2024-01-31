package repository

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) CreateBook() (string, error) {
	return "create book", nil
}

func (r *BookRepository) GetBookByID() (string, error) {
	return "get book by ID", nil
}

func (r *BookRepository) GetAllBooks() (string, error) {
	return "get All Books", nil
}

func (r *BookRepository) UpdateBook() (string, error) {
	return "update Book", nil
}

func (r *BookRepository) DeleteBook() (string, error) {
	return "delete book", nil
}
