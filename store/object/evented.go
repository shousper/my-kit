package object

import "github.com/shousper/my-kit/store/raw"

type EventCallback func(key string, value Value)

type EventedStore struct {
	Store

	callbacks map[raw.Event][]EventCallback
}

var _ Store = (*EventedStore)(nil)

func NewEventedStore(store Store) *EventedStore {
	return &EventedStore{
		Store: store,

		callbacks: make(map[raw.Event][]EventCallback),
	}
}

func (s *EventedStore) On(event raw.Event, fn EventCallback) {
	s.callbacks[event] = append(s.callbacks[event], fn)
}

func (s *EventedStore) emit(event raw.Event, key string, value Value) {
	for _, fn := range s.callbacks[event] {
		fn(key, value)
	}
}

func (s *EventedStore) Get(key string, out Value) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(raw.EventBeforeGet, key, out)
	err := s.Store.Get(key, out)
	s.emit(raw.EventAfterGet, key, out)
	return err
}
func (s *EventedStore) Set(key string, in Value) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(raw.EventBeforeSet, key, in)
	err := s.Store.Set(key, in)
	s.emit(raw.EventAfterSet, key, in)
	return err
}
func (s *EventedStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(raw.EventBeforeDelete, key, nil)
	err := s.Store.Delete(key)
	s.emit(raw.EventAfterDelete, key, nil)
	return err
}
func (s *EventedStore) Reset() error {
	s.emit(raw.EventBeforeReset, "", nil)
	err := s.Store.Reset()
	s.emit(raw.EventAfterReset, "", nil)
	return err
}
