package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //...
)

// Postgres .
type Postgres struct {
	DB *sql.DB
}

// New .
func New(url string) (*Postgres, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres - New - db.Ping: %w", err)
	}

	return &Postgres{db}, nil
}

// Close .
func (p *Postgres) Close() {
}
