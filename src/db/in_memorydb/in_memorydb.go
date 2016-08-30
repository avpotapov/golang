package in_memorydb

import (
	. "db"
	"sync"
)

type inMemoryDb struct {
	m   map[int][]byte
	lck sync.RWMutex
}

func NewInMemoryDb() Db {
	return &inMemoryDb{m: make(map[int][]byte)}
}

func (i *inMemoryDb) Get(key int) ([]byte, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (i *inMemoryDb) Set(key int, val []byte) error {
	i.lck.Lock()
	defer i.lck.Unlock()
	i.m[key] = val
	return nil
}

func (i *inMemoryDb) Del(key int) error {
	i.lck.RLock()
	defer i.lck.RUnlock()
	_, ok := i.m[key]
	if !ok {
		return ErrNotFound
	}
	delete(i.m, key)
	return nil
}

func (i *inMemoryDb) All() ([]([]byte), error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	if len(i.m) == 0 {
		return nil, ErrNotFound
	}
	all := make([]([]byte), 0, len(i.m))
	for key := range i.m {
		all = append(all, i.m[key])
	}
	return all, nil
}

func (i *inMemoryDb) Clear() error {
	i.lck.RLock()
	defer i.lck.RUnlock()
	i.m = make(map[int][]byte)
	return nil
}
