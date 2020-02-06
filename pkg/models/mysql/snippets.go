package mysql

import (
	"database/sql"

	"github.com/jacobboykin/snippetbox/pkg/models"
)

// SnippetModel : wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert : inserts a new snippet into the db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get : returns a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest : returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
