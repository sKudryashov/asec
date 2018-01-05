package finfo

import (
	"encoding/json"
	"encoding/xml"
	"io"

	cache "github.com/sKudryashov/asec/fileserver/internal/cache"
	model "github.com/sKudryashov/asec/fileserver/internal/platform"
	storage "github.com/sKudryashov/asec/fileserver/internal/platform/sqlite"
	"github.com/sKudryashov/asec/fileserver/internal/stat"
	"gopkg.in/go-playground/validator.v9"
)

const (
	cacheSize = 10
	// TypeXML short-declaration of xml content-type
	TypeXML = "xml"
	// TypeJSON short-declaration of json content-type
	TypeJSON = "json"
)

// Service manage storing and retreiving file info from storage
type Service struct {
	cache   *cache.Cache
	storage *storage.Storage
	stat    *stat.Stat
}

var getStorage = func() *storage.Storage {
	return storage.NewStorage()
}

// NewFinfoService return ready for usage FinfoService
func NewFinfoService() *Service {
	s := new(Service)
	s.cache = cache.New()
	s.storage = getStorage()
	s.stat = stat.NewStat()
	return s
}

//GetStat returns stat service
func (s *Service) GetStat() *stat.Stat {
	return s.stat
}

// SaveFileInfo validates and saves file info in the storage
func (s *Service) SaveFileInfo(data io.Reader, cType string) error {
	// read file info here into byte stream
	finfo := model.FileInfo{}
	if cType == TypeJSON {
		if err := json.NewDecoder(data).Decode(&finfo); err != nil {
			return err
		}
	} else if cType == TypeXML {
		if err := xml.NewDecoder(data).Decode(&finfo); err != nil {
			return err
		}
	}
	val := validator.New()
	val.SetTagName("vls")
	if err := val.Struct(finfo); err != nil {
		return err
	}
	s.cache.Set(finfo)
	if err := s.storage.Save(&finfo); err != nil {
		return err
	}
	s.stat.Hit()

	return nil
}

// GetFileInfo returns file info
func (s *Service) GetFileInfo() ([]byte, error) {
	infocache := s.cache.GetAll()
	data, err := json.Marshal(&infocache)
	println("infocache returns:", string(data))
	return data, err
}
