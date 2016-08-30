package in_memorydb

import (
	"bytes"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	db := NewInMemoryDb()
	if db == nil {
		t.Fatal("expected not nil")
	}
	if err := db.Set(1, []byte("Hello ")); err != nil {
		t.Fatal(err)
	}
	buffer, err := db.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal([]byte("Hello "), buffer) {
		t.Fatal("expected true")
	}

}

func TestSetAndDel(t *testing.T) {
	db := NewInMemoryDb()
	if db == nil {
		t.Fatal("expected not nil")
	}
	if err := db.Set(1, []byte("Hello ")); err != nil {
		t.Fatal(err)
	}
	if err := db.Del(1); err != nil {
		t.Fatal(err)
	}
	buffer, err := db.Get(1)
	if err == nil {
		t.Fatal(err)
	}
	if buffer != nil {
		t.Fatal("expected nil")
	}
}

func TestAll(t *testing.T) {
	db := NewInMemoryDb()
	if db == nil {
		t.Fatal("expected not nil")
	}
	for i := 0; i < 100; i++ {
		if err := db.Set(i, []byte("Hello ")); err != nil {
			t.Fatal(err)
		}
	}

	buffer, err := db.All()
	if err != nil {
		t.Fatal(err)
	}
	if buffer == nil {
		t.Fatal("expected not nil")
	}
	for data := range buffer {
		if !bytes.Equal([]byte("Hello "), data) {
			t.Fatal("expected true")
		}
	}
}
