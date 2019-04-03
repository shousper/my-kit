package object

type Hook int

const (
	HookBeforeGet Hook = iota
	HookAfterGet Hook = iota
	HookBeforeSet
	HookAfterSet
	HookBeforeDelete
	HookAfterDelete
	HookBeforeReset
	HookAfterReset
)

type HookCallback func(key string, value interface{})

type HookedStore struct {
	Store

	signals map[Hook][]HookCallback
}

var _ Store = (*HookedStore)(nil)

func NewHookedStore(store Store) *HookedStore {
	return &HookedStore{
		Store: store,

		signals: make(map[Hook][]HookCallback),
	}
}

func (s *HookedStore) On(signal Hook, fn HookCallback) {
	s.signals[signal] = append(s.signals[signal], fn)
}

func (s *HookedStore) emit(signal Hook, key string, value interface{}) {
	for _, fn := range s.signals[signal] {
		fn(key, value)
	}
}

func (s *HookedStore) Get(key string, out interface{}) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(HookBeforeGet, key, out)
	err := s.Store.Get(key, out)
	s.emit(HookAfterGet, key, out)
	return err
}
func (s *HookedStore) Set(key string, in interface{}) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(HookBeforeSet, key, in)
	err := s.Store.Set(key, in)
	s.emit(HookAfterSet, key, in)
	return err
}
func (s *HookedStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}

	s.emit(HookBeforeDelete, key, nil)
	err := s.Store.Delete(key)
	s.emit(HookAfterDelete, key, nil)
	return err
}
func (s *HookedStore) Reset() error {
	s.emit(HookBeforeReset, "", nil)
	err := s.Store.Reset()
	s.emit(HookAfterReset, "", nil)
	return err
}
