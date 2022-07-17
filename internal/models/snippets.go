package models

import (
	"database/sql"
	"time"
)

// Snippet type holds the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type wraps a sql.Db connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert return a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
