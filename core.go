package main

import "errors"

var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func Delete(key string) error {
	delete(store, key)
	return nil
}

func Put(key, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, contains := store[key]

	if contains {
		return value, nil
	} else {
		return "", ErrorNoSuchKey
	}

}
