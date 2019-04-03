package object

import (
	"bytes"
	"encoding/gob"
	"my-kit/store/raw"
)

func NewGobStore(store raw.Store) *DefaultStore {
	return NewStore(store, gobMarshal, gobUnmarshal)
}

func gobMarshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := gob.NewEncoder(buf).Encode(v)
	return buf.Bytes(), err
}

func gobUnmarshal(data []byte, out interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(out)
}