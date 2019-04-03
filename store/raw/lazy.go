package raw

type LazyStore struct {
	Store

	getter func(key string) ([]byte, error)
}

var _ Store = (*LazyStore)(nil)

func NewLazyStore(store Store, getter func(key string) ([]byte, error)) *LazyStore {
	return &LazyStore{
		Store:  store,
		getter: getter,
	}
}

func (s *LazyStore) Get(key string) ([]byte, error) {
	out, err := s.Store.Get(key)
	if err == ErrNotFound {
		entry, err := s.getter(key)
		if err != nil {
			return nil, err
		}
		if err := s.Store.Set(key, entry); err != nil {
			return nil, err
		}
		out, err = s.Store.Get(key)
	}
	return out, err
}

