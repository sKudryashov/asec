package cache

import (
	"testing"

	model "github.com/sKudryashov/asec/fileserver/internal/platform"
)

func TestCache(t *testing.T) {
	order := []int{
		6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	}
	cache := New()
	for i := 0; i <= 15; i++ {
		m := model.FileInfo{
			ID: i,
		}
		cache.Set(m)
	}
	cacheRecords := cache.GetAll()
	if len(cacheRecords) != 10 {
		t.Fatalf("length should be 10, %d given", len(cacheRecords))
	}
	index := 0
	for _, record := range cacheRecords {
		order := order
		if order[index] != record.ID {
			t.Fatal("wrong id order")
		}
		index++
	}
}
