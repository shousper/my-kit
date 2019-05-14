package raw

type Iterator interface {
	Next() (Value, bool)
	Err() error
}

type DefaultIterator struct {
	Store Store

	keys []string
	idx  int
	err  error
}

var _ Iterator = (*DefaultIterator)(nil)

func NewIterator(s Store) *DefaultIterator {
	return &DefaultIterator{Store: s}
}

func (i *DefaultIterator) Next() (out Value, ok bool) {
	if i.err != nil {
		return nil, false
	}
	if i.keys == nil {
		i.keys, i.err = i.Store.Keys()
		if i.err != nil {
			return nil, false
		}
	}
	if l := len(i.keys); l == 0 || i.idx == l-1 {
		return nil, false
	}
	if out, i.err = i.Store.Get(i.keys[i.idx]); i.err != nil {
		return nil, false
	}
	i.idx++
	return out, true
}

func (i *DefaultIterator) Err() error {
	return i.err
}
