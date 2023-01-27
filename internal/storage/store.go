// Naive generic storage for basic functionality
// TODO(mlee): Replace this with real, go-routine safe storage. Preferably something with an external datastore that can be persisted to disk
package storage

// DataStore is a generic in-memory store for protobuf messages
type DataStore[T any] struct {
	items []T
	//m     sync.Mutex
}

type Comparator[T any] func(value T) bool

// Add a new item to DataStore
func (ds *DataStore[T]) Add(item T) {
	//ds.m.Lock()
	ds.items = append(ds.items, item)
	//ds.m.Unlock()
}

func (ds *DataStore[T]) AddRef(item *T) {
	//ds.m.Lock()
	ds.items = append(ds.items, *item)
	//ds.m.Unlock()
}

// List copies all the items in ds.items into a new slice and returns it
func (ds *DataStore[T]) List() []T {
	//ds.m.Lock()
	result := make([]T, len(ds.items))
	copy(result, ds.items)
	//ds.m.Unlock()
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
	var result T
	//ds.m.Lock()
	for _, item := range ds.items {
		if cmp(item) {
			result = item
		}
	}
	//ds.m.Unlock()
	return result
}

// FindAll values that match a comparator
func (ds *DataStore[T]) FindAll(cmp Comparator[T]) []T {
	var results []T
	//ds.m.Lock()
	for _, item := range ds.items {
		if cmp(item) {
			results = append(results, item)
		}
	}
	//ds.m.Unlock()
	return results
}

// Clear the datastore
func (ds *DataStore[T]) Clear() {
	//ds.m.Lock()
	ds.items = []T{}
	//ds.m.Unlock()
}
