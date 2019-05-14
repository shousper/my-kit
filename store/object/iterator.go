package object

type Iterator interface {
	Next(out Value) bool
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

func (i *DefaultIterator) Next(out Value) bool {
	if i.err != nil {
		return false
	}
	if i.keys == nil {
		i.keys, i.err = i.Store.Keys()
		if i.err != nil {
			return false
		}
	}
	if l := len(i.keys); l == 0 || i.idx == l-1 {
		return false
	}
	if i.err = i.Store.Get(i.keys[i.idx], out); i.err != nil {
		return false
	}
	i.idx++
	return true
}

func (i *DefaultIterator) Err() error {
	return i.err
}
