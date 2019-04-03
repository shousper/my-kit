package object

import (
	"my-kit/store/raw"
)

type LazyStore struct {
	Store

	getter func(key string) (interface{}, error)
}

var _ Store = (*LazyStore)(nil)

func NewLazyStore(store Store, getter func(key string) (interface{}, error)) *LazyStore {
	return &LazyStore{
		Store:  store,
		getter: getter,
	}
}

func (s *LazyStore) Set(key string, in interface{}) error {
	return s.Store.Set(key, in)
}

func (s *LazyStore) Get(key string, out interface{}) error {
	err := s.Store.Get(key, out)
	if err == raw.ErrNotFound {
		entry, err := s.getter(key)
		if err != nil {
			return err
		}
		if err := s.Store.Set(key, entry); err != nil {
			return err
		}
		err = s.Store.Get(key, out)
	}
	return err
}

