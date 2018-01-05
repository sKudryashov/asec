package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	//postgres driver initialization
	_ "github.com/lib/pq"
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
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("STORAGE_HOST"),
		os.Getenv("STORAGE_PORT"),
		os.Getenv("STORAGE_USER"),
		os.Getenv("STORAGE_DB"),
		os.Getenv("STORAGE_PASSWORD"))
	// connStr := "host=database user=postgres dbname=postgres sslmode=disable"
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	st.database = database
	return st
}

// Save saves data inside the storage
func (s *Storage) Save(m *model.FileInfo) error {
	s.mu.Lock()
	_, err := s.database.Exec("CREATE TABLE IF NOT EXISTS finfo (id SERIAL PRIMARY KEY, name TEXT, mode TEXT, mod_time TEXT, size TEXT)")
	s.mu.Unlock()
	if err != nil {
		println(".err1", err.Error())
		return err
	}
	//fine-grained lock
	s.mu.Lock()
	stmt := "INSERT INTO finfo (name, mode, mod_time, size) VALUES ($1, $2, $3, $4)"
	_, err = s.database.Exec(stmt, m.Name, m.Mode, m.ModTime, m.Size)
	s.mu.Unlock()
	log.Printf("insertion to the database: %s %s %s %d", m.Name, m.Mode, m.ModTime, m.Size)
	if err != nil {
		println(".err3", err.Error())
		return err
	}

	return err
}
