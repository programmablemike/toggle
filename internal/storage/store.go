// Naive generic storage for basic functionality
// TODO(mlee): Replace this with real, go-routine safe storage. Preferably something with an external datastore that can be persisted to disk
package storage

// DataStore is a generic in-memory store for protobuf messages
type DataStore[T any] struct {
	items []T
}

type Comparator[T any] func(value T) bool

// Add a new item to DataStore
func (ds *DataStore[T]) Add(item T) {
	ds.items = append(ds.items, item)
}

func (ds *DataStore[T]) AddRef(item *T) {
	ds.items = append(ds.items, *item)
}

// List copies all the items in ds.items into a new slice and returns it
func (ds *DataStore[T]) List() []T {
	result := make([]T, len(ds.items))
	copy(result, ds.items)
	return result
}

// ListAsRef returns a new slice of pointers for all items in ds.items
func (ds *DataStore[T]) ListAsRef() []*T {
	// Copy first to avoid passing out internal references
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
	var results []T
	for _, item := range ds.items {
		if cmp(item) {
			results = append(results, item)
		}
	}
	return results
}

// Clear the datastore
func (ds *DataStore[T]) Clear() {
	ds.items = []T{}
}