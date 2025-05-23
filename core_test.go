package main

import (
	"errors"
	"testing"
)

func TestPut(t *testing.T) {
	const key = "create-key"
	const value = "create-value"
	var contains bool
	var val interface{}

	defer delete(store, key)

	_, contains = store[key]
	if contains {
		t.Error("key/value already exists")
	}

	err := Put(key, value)

	if err != nil {
		t.Error(err)
	}

	val, contains = store[key]
	if !contains {
		t.Error("create failed")
	}

	if val != value {
		t.Error("val/value mismatch")
	}

}

func TestGet(t *testing.T) {
	const key = "read-key"
	const value = "read-value"
	var val interface{}
	var err error

	defer delete(store, key)

	val, err = Get(key)

	if err == nil {
		t.Error("expected an error")
	}
	if !errors.Is(err, ErrorNoSuchKey) {
		t.Error("unexpected err:", err)

	}

	store[key] = value

	val, err = Get(key)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	if val != value {
		t.Error("val/value mismatch")
	}

}

func TestDelete(t *testing.T) {
	const key = "delete-key"
	const value = "delete-value"
	var contains bool

	defer delete(store, key)

	store[key] = value

	_, contains = store[key]
	if !contains {
		t.Error("key/value does not exist")
	}

	Delete(key)
	_, contains = store[key]

	if contains {
		t.Error("Delete failed")
	}

}
