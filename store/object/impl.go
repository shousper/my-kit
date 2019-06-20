package object

import (
	"my-kit/store/raw"
)

type Marshaller func(v interface{}) ([]byte, error)
type Unmarshaller func(data []byte, out interface{}) error

type DefaultStore struct {
	raw.Store

	marshaller   Marshaller
	unmarshaller Unmarshaller
}

var _ Store = (*DefaultStore)(nil)

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

func (s *DefaultStore) Iterate(referenceType interface{}, fn IteratorFunc) error {
	return s.Store.Iterate(func(key string, value []byte) (bool, error) {
		if err := s.unmarshaller(value, referenceType); err != nil {
			return false, err
		}
		return fn(key, referenceType)
	})
}

func (s *DefaultStore) Get(key string, out interface{}) error {
	if key == "" {
		return ErrInvalidKey
	}
	if out == nil {
		return ErrInvalidOutput
	}

	v, err := s.Store.Get(key)
	if err != nil {
		return err
	}

	return s.unmarshaller(v, out)
}

func (s *DefaultStore) Set(key string, in interface{}) error {
	if key == "" {
		return ErrInvalidKey
	}
	if in == nil {
		return ErrInvalidValue
	}

	v, err := s.marshaller(in)
	if err != nil {
		return err
	}
	return s.Store.Set(key, v)
}

