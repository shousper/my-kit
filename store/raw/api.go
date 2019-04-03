package raw

import "io"

type IteratorFunc func(key string, value []byte) (bool, error)


type Store interface {
	io.Closer

	Iterate(fn IteratorFunc) error
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
}
