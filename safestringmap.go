// Copyright © 2019 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

/*
Package safestringmap implements a thread-safe string map for Go using RWLock to optimise for concurrent reads
*/
package safestringmap

import "sync"

// SafeStringMap structure. Defines a map of strings that are managed by a `sync.RWLock` to
// ensure safe Read-Write access in concurrent environments.
// All read operations utilse `RLock`. All write operations utilise `Lock`.
type SafeStringMap struct {
	sync.RWMutex
	internal map[string]string
}

// NewSafeStringMap function creates a new instance of a `SafeStringMap` structure.
func NewSafeStringMap() *SafeStringMap {
	return &SafeStringMap{
		internal: make(map[string]string),
	}
}

// Load returns the string value associated with the requested key if it exists in the map.
// If not an emptry string is returned and the functions `bool` result will be set to `false`.
func (ssm *SafeStringMap) Load(key string) (value string, ok bool) {
	ssm.RLock()
	result, ok := ssm.internal[key]
	ssm.RUnlock()
	return result, ok
}

// Delete removes a value from the map that is associated with the specified key value - if it exists.
func (ssm *SafeStringMap) Delete(key string) {
	ssm.Lock()
	delete(ssm.internal, key)
	ssm.Unlock()
}

// Store adds a new key value pair to the map if it doesnt already exist.
// If the key already exists, the provided string value will be used to update (overwrite) the value associated with the key.
func (ssm *SafeStringMap) Store(key, value string) {
	ssm.Lock()
	ssm.internal[key] = value
	ssm.Unlock()
}

// Count returns the count of strings stored within the map (as an integer)
func (ssm *SafeStringMap) Count() int {
	ssm.RLock()
	result := len(ssm.internal)
	ssm.RUnlock()
	return result
}
