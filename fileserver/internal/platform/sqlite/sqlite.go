package sqlite

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	model "github.com/sKudryashov/asec/fileserver/internal/platform"
)

// Storage is wrapper for db storage
type Storage struct {
	database *sql.DB
	mu       sync.RWMutex
}

// NewStorage returns storage wrapper
func NewStorage() *Storage {
	st := new(Storage)
	database, err := sql.Open("sqlite3", "./nraboy.db")
	if err != nil {
		panic(err)
	}
	st.database = database
	return st
}

// Save saves data inside the storage
func (s *Storage) Save(m *model.FileInfo) error {
	statement, err := s.database.Prepare("CREATE TABLE IF NOT EXISTS finfo (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		return err
	}
	s.mu.Lock()
	statement.Exec()
	s.mu.Unlock()
	statement, err = s.database.Prepare("INSERT INTO finfo (name) VALUES (?)")
	if err != nil {
		return err
	}
	s.mu.Lock()
	statement.Exec(m.Name)
	s.mu.Unlock()

	return nil
}

// Fetch returns data from DB
func (s *Storage) Fetch() (*model.FileInfo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rows, err := s.database.Query("SELECT id, name FROM finfo")
	if err != nil {
		return nil, err
	}
	r := &model.FileInfo{}
	for rows.Next() {
		rows.Scan(&r.ID, &r.Name)
	}

	return r, nil
}
