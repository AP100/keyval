package main

import (
	"errors"
	"sync"
)

type LockableMap struct {
	sync.RWMutex
	m map[string]string
}

var store = LockableMap{m: make(map[string]string)}
var ErrorNoSuchKey = errors.New("no such key")

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return nil
}

func Put(key, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	value, contains := store.m[key]
	store.RUnlock()

	if contains {
		return value, nil
	} else {
		return "", ErrorNoSuchKey
	}

}
