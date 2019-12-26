// Copyright © 2019 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

/*
Package safestringmap implements a thread-safe string map for Go using RWLock to optimise for concurrent reads
*/
package safestringmap

import "sync"

// SafeStringMap defines a map structure of strings as well as a RWLock to
// ensure safe Read-Write access in concurrent environments.
// All read operations utilse the RLock.
// All write operations utilise the Lock.
type SafeStringMap struct {
	sync.RWMutex
	internal map[string]string
}

// NewSafeStringMap creates a new instance of a SafeStringMap
func NewSafeStringMap() *SafeStringMap {
	return &SafeStringMap{
		internal: make(map[string]string),
	}
}

// Load will return the string value associated with the requested
// string key if it exists in the map. If not an emptry string will be
// returned and the ok bool will be set to false.
func (ssm *SafeStringMap) Load(key string) (value string, ok bool) {
	ssm.RLock()
	result, ok := ssm.internal[key]
	ssm.RUnlock()
	return result, ok
}

// Delete will removed a value from the map that is associated with the
// specified key value - if it exists.
func (ssm *SafeStringMap) Delete(key string) {
	ssm.Lock()
	delete(ssm.internal, key)
	ssm.Unlock()
}

// Store will add a new key value pair to the map if it doesnt already exist.
// If the key already exists, the provided string value will be used to update
// the value contained within the map with the new provided value
func (ssm *SafeStringMap) Store(key, value string) {
	ssm.Lock()
	ssm.internal[key] = value
	ssm.Unlock()
}

// Count will return the count of the number of objects stored within the map as an integer
func (ssm *SafeStringMap) Count() int {
	ssm.RLock()
	result := len(ssm.internal)
	ssm.RUnlock()
	return result
}
