package finfo

import (
	"encoding/json"

	cache "github.com/sKudryashov/asec/fileserver/internal/cache"
	model "github.com/sKudryashov/asec/fileserver/internal/platform"
	storage "github.com/sKudryashov/asec/fileserver/internal/platform/sqlite"
)

// Service manage storing and retreiving file info from storage
type Service struct {
	cache *cache.Cache
}

const (
	cacheSize = 10
)

// NewFinfoService return ready for usage FinfoService
func NewFinfoService() *Service {
	s := new(Service)
	s.cache = cache.New()
	return s
}

//SaveFileInfo validates and saves file info in the storage
func (s *Service) SaveFileInfo(data []byte) error {
	// read file info here into byte stream
	finfo := model.FileInfo{}
	err := json.Unmarshal(data, &finfo)
	if err != nil {
		return err
	}
	s.cache.Set(&finfo)
	err = storage.Save(&finfo)
	if err != nil {
		return err
	}

	return nil
}

// GetFileInfo returns file info
func (s *Service) GetFileInfo() ([]byte, error) {
	infocache := s.cache.GetAll()
	return json.Marshal(&infocache)
}
