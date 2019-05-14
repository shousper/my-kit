package raw

type Event int

const (
	EventBeforeGet Event = iota
	EventAfterGet
	EventBeforeSet
	EventAfterSet
	EventBeforeDelete
	EventAfterDelete
	EventBeforeReset
	EventAfterReset
)

type EventCallback func(key string, value Value)

type EventedStore struct {
	Store

	callbacks map[Event][]EventCallback
}

var _ Store = (*EventedStore)(nil)

func NewEventedStore(store Store) *EventedStore {
	return &EventedStore{
		Store: store,

		callbacks: make(map[Event][]EventCallback),
	}
}

func (s *EventedStore) On(event Event, fn EventCallback) {
	s.callbacks[event] = append(s.callbacks[event], fn)
}

func (s *EventedStore) emit(event Event, key string, value Value) {
	for _, fn := range s.callbacks[event] {
		fn(key, value)
	}
}

func (s *EventedStore) Get(key string) (out Value, err error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	s.emit(EventBeforeGet, key, nil)
	defer s.emit(EventAfterGet, key, out)
	out, err = s.Store.Get(key)
	return out, err
}
func (s *EventedStore) Set(key string, in Value) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.emit(EventBeforeSet, key, in)
	defer s.emit(EventAfterSet, key, in)
	err := s.Store.Set(key, in)
	return err
}
func (s *EventedStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.emit(EventBeforeDelete, key, nil)
	defer s.emit(EventAfterDelete, key, nil)
	err := s.Store.Delete(key)
	return err
}
func (s *EventedStore) Reset() error {
	s.emit(EventBeforeReset, "", nil)
	defer s.emit(EventAfterReset, "", nil)
	err := s.Store.Reset()
	return err
}
