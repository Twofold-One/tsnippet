package models

import (
	"database/sql"
	"errors"
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

// TODO: Add custom expiration duration; right now default is 365 days
// There is some problem during execution of the statement, when 3d placeholder
// couldn't be seen.
// Insert return a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES ($1, $2,	timezone('utc', now()),	timezone('utc', now()) + interval '365 day') 
	RETURNING id;`

	id := 0

	err := m.DB.QueryRow(stmt, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("Something went wrong, iserted id is equal to zero")
	}
	return id, nil
}

// Get return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
