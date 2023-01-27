// Naive generic storage for basic functionality
// TODO(mlee): Replace this with real, go-routine safe storage. Preferably something with an external datastore that can be persisted to disk
package storage

import "sync"

// DataStore is a generic in-memory store for protobuf messages
type DataStore[T any] struct {
	items []T
	m     sync.Mutex
}

type Comparator[T any] func(value T) bool

// Add a new item to DataStore
func (ds *DataStore[T]) Add(item T) {
	ds.m.Lock()
	defer ds.m.Unlock()
	ds.items = append(ds.items, item)
}

// Adds a new item as a pointer by copying the value
func (ds *DataStore[T]) AddRef(item *T) {
	ds.m.Lock()
	defer ds.m.Unlock()
	ds.items = append(ds.items, *item)
}

// List copies all the items in ds.items into a new slice and returns it
func (ds *DataStore[T]) List() []T {
	ds.m.Lock()
	defer ds.m.Unlock()
	result := make([]T, len(ds.items))
	copy(result, ds.items)
	return result
}

// ListAsRef returns a new slice of pointers for all items in ds.items
func (ds *DataStore[T]) ListAsRef() []*T {
	// Copy first to avoid passing out internal references
	// NOTE(mlee): No mutex lock needed here - it's being handled in ds.List()
	items := ds.List()
	// Pre-allocate the pointer list
	result := make([]*T, len(items))
	for idx, item := range items {
		result[idx] = &item
	}
	return result
}

// FindOne value that matches a comparator
func (ds *DataStore[T]) FindOne(cmp Comparator[T]) T {
	ds.m.Lock()
	defer ds.m.Unlock()
	var result T
	for _, item := range ds.items {
		if cmp(item) {
			result = item
		}
	}
	return result
}

// FindAll values that match a comparator
func (ds *DataStore[T]) FindAll(cmp Comparator[T]) []T {
	ds.m.Lock()
	defer ds.m.Unlock()
	var results []T
	for _, item := range ds.items {
		if cmp(item) {
			results = append(results, item)
		}
	}
	return results
}

// Size tells us how many items are in the DataStore
func (ds *DataStore[T]) Size() int {
	ds.m.Lock()
	defer ds.m.Unlock()
	return len(ds.items)
}

// Clear the datastore
func (ds *DataStore[T]) Clear() {
	ds.m.Lock()
	defer ds.m.Unlock()
	ds.items = []T{}
}
