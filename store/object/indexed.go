package object

type IndexerFunc func(entity interface{}) (key string)

type IndexedStore struct {
	Store

	indexers map[string]IndexerFunc
}

var _ Store = (*IndexedStore)(nil)

func NewIndexedStore(store Store) *IndexedStore {
	return &IndexedStore{
		Store:    store,
		indexers: make(map[string]IndexerFunc),
	}
}

func (s *IndexedStore) Index(name string, indexer IndexerFunc) *IndexedStore {
	s.indexers[name] = indexer
	return s
}

func (s *IndexedStore) Set(key string, entry interface{}) error {
	if err := s.Store.Set(key, entry); err != nil {
		return err
	}

	for _, i := range s.indexers {
		indexKey := i(entry)
		_ = s.Store.Set(indexKey, key)
	}
	return nil
}

func (s *IndexedStore) GetBy(index string, inAndOut interface{}) error {
	indexKey := s.indexers[index](inAndOut)
	var key string
	if err := s.Store.Get(indexKey, &key); err != nil {
		return err
	}
	return s.Store.Get(key, inAndOut)
}

