package storage

// DataStore is a generic in-memory store for protobuf messages
type DataStore[T any] struct {
	Items []T
}

type Comparator[T any] func(value T) bool

func (ds *DataStore[T]) Add(item T) {
	ds.Items = append(ds.Items, item)
}

func (ds *DataStore[T]) List() []T {
	return ds.Items
}

// FindOne value that matches a comparator
func (ds *DataStore[T]) FindOne(cmp Comparator[T]) T {
	var result T
	for _, item := range ds.Items {
		if cmp(item) {
			result = item
		}
	}
	return result
}

// FindAll values that match a comparator
func (ds *DataStore[T]) FindAll(cmp Comparator[T]) []T {
	var results []T
	for _, item := range ds.Items {
		if cmp(item) {
			results = append(results, item)
		}
	}
	return results
}
