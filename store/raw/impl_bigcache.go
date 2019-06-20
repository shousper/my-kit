package raw

import (
	"github.com/allegro/bigcache"
	"time"
)

type BigCacheStore struct {
	cache *bigcache.BigCache
}

var _ Store = (*BigCacheStore)(nil)

func NewBigCacheStore(eviction time.Duration) *BigCacheStore {
	cfg := bigcache.DefaultConfig(eviction)
	cfg.Verbose = false
	cache, _ := bigcache.NewBigCache(cfg)
	return &BigCacheStore{
		cache,
	}
}

func (s *BigCacheStore) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	v, err := s.cache.Get(key)
	return v, s.normalizeError(err)
}

func (s *BigCacheStore) Set(key string, in []byte) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := s.cache.Set(key, in)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := s.cache.Delete(key)
	return s.normalizeError(err)
}

func (s *BigCacheStore) Iterate(fn IteratorFunc) error {
	iterator := s.cache.Iterator()
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

func (s *BigCacheStore) Close() error {
	return s.cache.Close()
}

func (s *BigCacheStore) Reset() error {
	return s.cache.Reset()
}

func (s *BigCacheStore) Len() int {
	return s.cache.Len()
}

func (s *BigCacheStore) Capacity() int {
	return s.cache.Capacity()
}

func (s *BigCacheStore) normalizeError(err error) error {
	if err == bigcache.ErrEntryNotFound {
		return ErrNotFound
	}
	return err
}
