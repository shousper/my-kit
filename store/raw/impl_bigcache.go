package raw

import (
	"github.com/allegro/bigcache"
	"time"
)

type BigCacheStore struct {
	bigcache.BigCache
}

var _ Store = (*BigCacheStore)(nil)

func NewBigCacheStore() *BigCacheStore {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return &BigCacheStore{
		BigCache: *cache,
	}
}

func (s *BigCacheStore) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	v, err := s.BigCache.Get(key)
	return v, s.normalizeError(err)
}

func (s *BigCacheStore) Set(key string, in []byte) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := s.BigCache.Set(key, in)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := s.BigCache.Delete(key)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Iterate(fn IteratorFunc) error {
	iterator := s.Iterator()
	for iterator.SetNext() {
		entry, err := iterator.Value()
		if err != nil {
			return s.normalizeError(err)
		}

		done, err := fn(entry.Key(), entry.Value())
		if err != nil {
			return s.normalizeError(err)
		}

		if done {
			return nil
		}
	}
	return nil
}

func (s *BigCacheStore) normalizeError(err error) error {
	if err == bigcache.ErrEntryNotFound {
		return ErrNotFound
	}
	return err
}
