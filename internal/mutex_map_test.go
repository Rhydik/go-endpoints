package internal

import (
	"testing"
)

func TestNewMutexMap(t *testing.T) {
	mutexMap := NewMutexMap()

	if mutexMap.M == nil {
		t.Errorf("NewMutexMap() should return a map, but returned nil")
	}
}

func TestSetMutexMap(t *testing.T) {
	mutexMap := NewMutexMap()

	// Test that we can set a key and value
	mutexMap.SetMutexMap("key", []string{"val"})
	if mutexMap.M["key"][0] != "val" {
		t.Errorf("SetMutexMap() should set key to val, but key is %s", mutexMap.M["key"][0])
	}

	// Test that we can set a key and value
	mutexMap.SetMutexMap("key2", []string{"val2"})
	if mutexMap.M["key2"][0] != "val2" {
		t.Errorf("SetMutexMap() should set key to val, but key is %s", mutexMap.M["key2"][0])
	}
}

func TestGetMutexMap(t *testing.T) {
	mutexMap := NewMutexMap()

	mutexMap.SetMutexMap("key", []string{"val"})
	val, ok := mutexMap.GetMutexMap("key")
	if !ok {
		t.Errorf("GetMutexMap() should return true, but returned false")
	}
	if val[0] != "val" {
		t.Errorf("GetMutexMap() should return val, but returned %s", val[0])
	}

	// Test that we get false when we try to get a key that doesn't exist
	_, ok = mutexMap.GetMutexMap("key2")
	if ok {
		t.Errorf("GetMutexMap() should return false, but returned true")
	}
}