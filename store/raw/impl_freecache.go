package raw

import (
	"github.com/coocood/freecache"
	"time"
)

type FreeCacheStore struct {
	cache *freecache.Cache

	expiration time.Duration
}

var _ Store = (*FreeCacheStore)(nil)

func NewFreeCacheStore(size int, expiration time.Duration) *FreeCacheStore {
	return &FreeCacheStore{
		freecache.NewCache(size),
		expiration,
	}
}

func (s FreeCacheStore) Get(key string) ([]byte, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	v, err := s.cache.Get([]byte(key))
	return v, s.normalizeError(err)
}

func (s FreeCacheStore) Set(key string, in []byte) error {
	if key == "" {
		return ErrInvalidKey
	}
	err := s.cache.Set([]byte(key), in, int(s.expiration.Seconds()))
	return s.normalizeError(err)
}

func (s FreeCacheStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.cache.Del([]byte(key))
	return nil
}

func (s FreeCacheStore) Iterate(fn IteratorFunc) error {
	iterator := s.cache.NewIterator()
	for {
		entry := iterator.Next()
		if entry == nil {
			return nil
		}

		d, err := fn(string(entry.Key), entry.Value)
		if err != nil {
			return s.normalizeError(err)
		}

		if d {
			return nil
		}
	}
}

func (s *FreeCacheStore) Close() error {
	return nil
}

func (s *FreeCacheStore) Reset() error {
	s.cache.Clear()
	return nil
}

func (s *FreeCacheStore) Len() int {
	return int(s.cache.EntryCount())
}

func (s *FreeCacheStore) Capacity() int {
	return -1
}

func (s FreeCacheStore) normalizeError(err error) error {
	if err == freecache.ErrNotFound {
		return ErrNotFound
	}
	return err
}

