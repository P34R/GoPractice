package src

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	url string
	db  *sql.DB
}

func NewStore() *Store {
	return &Store{
		url: "host=localhost dbname=projectdb",
	}
}
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.url)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
