package raw

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

type HookCallback func(key string, value []byte)

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

func (s *HookedStore) emit(signal Hook, key string, value []byte) {
	for _, fn := range s.signals[signal] {
		fn(key, value)
	}
}

func (s *HookedStore) Get(key string) (out []byte, err error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	s.emit(HookBeforeGet, key, nil)
	defer s.emit(HookAfterGet, key, out)
	out, err = s.Store.Get(key)
	return out, err
}
func (s *HookedStore) Set(key string, in []byte) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.emit(HookBeforeSet, key, in)
	defer s.emit(HookAfterSet, key, in)
	err := s.Store.Set(key, in)
	return err
}
func (s *HookedStore) Delete(key string) error {
	if key == "" {
		return ErrInvalidKey
	}
	s.emit(HookBeforeDelete, key, nil)
	defer s.emit(HookAfterDelete, key, nil)
	err := s.Store.Delete(key)
	return err
}
func (s *HookedStore) Reset() error {
	s.emit(HookBeforeReset, "", nil)
	defer s.emit(HookAfterReset, "", nil)
	err := s.Store.Reset()
	return err
}
