package encoded

import (
	"github.com/shousper/my-kit/store/object"
	"github.com/shousper/my-kit/store/raw"
)

type Marshaller func(v interface{}) ([]byte, error)
type Unmarshaller func(data []byte, out interface{}) error

type DefaultStore struct {
	raw.Store

	marshaller   Marshaller
	unmarshaller Unmarshaller
}

var _ object.Store = (*DefaultStore)(nil)

func NewStore(
	store raw.Store,
	marshaller Marshaller,
	unmarshaller Unmarshaller,
) *DefaultStore {
	return &DefaultStore{
		Store:        store,
		marshaller:   marshaller,
		unmarshaller: unmarshaller,
	}
}

func (s *DefaultStore) Iterate() object.Iterator {
	return &iterator{
		iter:         s.Store.Iterate(),
		unmarshaller: s.unmarshaller,
	}
}

func (s *DefaultStore) Get(key string, out object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}
	if out == nil {
		return object.ErrInvalidOutput
	}

	v, err := s.Store.Get(key)
	if err != nil {
		return err
	}

	return s.unmarshaller(v, out)
}

func (s *DefaultStore) Set(key string, in object.Value) error {
	if key == "" {
		return object.ErrInvalidKey
	}
	if in == nil {
		return object.ErrInvalidValue
	}

	v, err := s.marshaller(in)
	if err != nil {
		return err
	}
	return s.Store.Set(key, v)
}

type iterator struct {
	iter         raw.Iterator
	unmarshaller Unmarshaller

	err error
}

func (i *iterator) Next(out object.Value) bool {
	if i.iter.Err() != nil {
		i.err = i.iter.Err()
		return false
	}
	v, ok := i.iter.Next()
	if !ok {
		return false
	}
	if i.err = i.unmarshaller(v, out); i.err != nil {
		return false
	}
	return true
}

func (i *iterator) Err() error {
	return i.err
}
