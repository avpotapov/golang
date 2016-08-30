package db

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Db interface {
	All() ([]([]byte), error)
	Get(key int) ([]byte, error)
	Set(key int, value []byte) error
	Del(key int) error
	Clear() error
}
