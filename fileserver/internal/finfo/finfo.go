package finfo

import (
	"encoding/json"

	cache "github.com/bluele/gcache"
	model "github.com/sKudryashov/asec/fileserver/internal/platform"
	storage "github.com/sKudryashov/asec/fileserver/internal/platform/sqlite"
)

// Service manage storing and retreiving file info from storage
type Service struct {
	cache cache.Cache
}

const (
	cacheSize = 10
)

// NewFinfoService return ready for usage FinfoService
func NewFinfoService() *Service {
	s := new(Service)
	//TODO: wrap with my struct
	s.cache = cache.New(cacheSize).Build()
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

	infocache, err := s.cache.GetIFPresent("finfos")
	if err != nil {
		// return
		//TODO: add first data set
	}
	_, ok := infocache.([]*model.FileInfo)
	if !ok {
		panic("wrong cached data type")
	}
	err = storage.Save(&finfo)
	if err != nil {
		return err
	}

	return nil
}

// GetFileInfo returns file info
func (s *Service) GetFileInfo() ([]byte, error) {
	infocache, err := s.cache.GetALL()
	for key, value := range infocache {

	}
	info, ok := infocache.([]*model.FileInfo)
	if !ok {
		panic("wrong cached data type")
	}
	json.Marshal(&info)

}
