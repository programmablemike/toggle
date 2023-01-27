package storage

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

type TestItem struct {
	Id    int
	Value string
}

var sizes = []struct {
	count int
}{
	{count: 1000},    // One thousand
	{count: 10000},   // Ten thousand
	{count: 100000},  // One hundred thousand
	{count: 1000000}, // One million
}

func TestAdd(t *testing.T) {
	var ds DataStore[TestItem]
	ds.Add(TestItem{Id: 1, Value: "hello world"})
	ds.Add(TestItem{Id: 2, Value: "goodbye moon"})
	ds.Add(TestItem{Id: 3, Value: "salutations sun"})
}

func benchmarkAdd(ds *DataStore[TestItem], size int) {
	for i := 0; i < size; i++ {
		ds.Add(TestItem{Id: i, Value: "hello world"})
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()

	for _, v := range sizes {
		ds := DataStore[TestItem]{}
		b.Run(fmt.Sprintf("count_%d", v.count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkAdd(&ds, v.count)

			}
		})
		ds.Clear()
	}
}

func TestAddRef(t *testing.T) {
	var ds DataStore[TestItem]
	ds.AddRef(&TestItem{Id: 1, Value: "hello world"})
	ds.AddRef(&TestItem{Id: 2, Value: "goodbye moon"})
	ds.AddRef(&TestItem{Id: 3, Value: "salutations sun"})
}

func benchmarkAddRef(ds *DataStore[TestItem], size int) {
	for i := 0; i < size; i++ {
		ds.AddRef(&TestItem{Id: i, Value: "hello world"})
	}
}

func BenchmarkAddRef(b *testing.B) {
	b.ReportAllocs()

	for _, v := range sizes {
		ds := DataStore[TestItem]{}
		b.Run(fmt.Sprintf("count_%d", v.count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkAddRef(&ds, v.count)
			}
		})
		ds.Clear()
	}
}

func TestFindOne(t *testing.T) {
	cmp := func(toCheck TestItem) bool {
		if toCheck.Value == "goodbye moon" {
			return true
		}
		return false
	}

	ds := DataStore[TestItem]{
		items: []TestItem{
			{Id: 1, Value: "hello world"},
			{Id: 2, Value: "goodbye moon"},
			{Id: 3, Value: "salutations sun"},
		},
	}
	res := ds.FindOne(cmp)
	assert.Equal(t, res.Id, 2)
}

func TestFindAll(t *testing.T) {
	cmp := func(toCheck TestItem) bool {
		if toCheck.Value == "goodbye moon" {
			return true
		}
		return false
	}

	ds := DataStore[TestItem]{
		items: []TestItem{
			{Id: 1, Value: "hello world"},
			{Id: 2, Value: "goodbye moon"},
			{Id: 3, Value: "goodbye moon"},
			{Id: 4, Value: "salutations sun"},
			{Id: 5, Value: "goodbye moon"},
		},
	}
	res := ds.FindAll(cmp)
	assert.Equal(t, len(res), 3)
	assert.Equal(t, res[0].Id, 2)
	assert.Equal(t, res[1].Id, 3)
	assert.Equal(t, res[2].Id, 5)
}

func TestList(t *testing.T) {
	ds := DataStore[TestItem]{
		items: []TestItem{
			{Id: 1, Value: "hello world"},
			{Id: 2, Value: "goodbye moon"},
			{Id: 3, Value: "salutations sun"},
		},
	}
	res := ds.List()
	assert.Equal(t, res[0].Id, 1)
	assert.Equal(t, res[0].Value, "hello world")
	assert.Equal(t, res[1].Id, 2)
	assert.Equal(t, res[1].Value, "goodbye moon")
	assert.Equal(t, res[2].Id, 3)
	assert.Equal(t, res[2].Value, "salutations sun")
	// Modify the returned list to verify the underlying items don't change
	res[0].Value = "not hello world"
	res2 := ds.List()
	assert.Equal(t, res[0].Value, "not hello world")
	assert.Equal(t, res2[0].Value, "hello world")
}

func TestListAsRef(t *testing.T) {
	ds := DataStore[TestItem]{
		items: []TestItem{
			{Id: 1, Value: "hello world"},
			{Id: 2, Value: "goodbye moon"},
			{Id: 3, Value: "salutations sun"},
		},
	}
	res := ds.ListAsRef()
	fmt.Printf("%v", *res[0])
	assert.Equal(t, (*res[0]).Id, 1)
	assert.Equal(t, (*res[0]).Value, "hello world")
	assert.Equal(t, (*res[1]).Id, 2)
	assert.Equal(t, (*res[1]).Value, "goodbye moon")
	assert.Equal(t, (*res[2]).Id, 3)
	assert.Equal(t, (*res[2]).Value, "salutations sun")
	// Modify the returned list to verify the underlying items don't change
	(*res[0]).Value = "not hello world"
	res2 := ds.List()
	assert.Equal(t, res[0].Value, "not hello world")
	assert.Equal(t, res2[0].Value, "hello world")
}

// TestGoroutineSafeUpdates runs a search over a DataStore's items while simultaneously clearing it
// Ideally this will not fail if proper mutex locking is used during update operations
func TestGoroutineSafeUpdates(t *testing.T) {
	var ds DataStore[TestItem]
	for i := 0; i < 1000; i++ {
		// Set the item with Id=500 to be different so we have something to search for
		if i == 500 {
			ds.Add(TestItem{Id: i, Value: "it"})
		} else {
			ds.Add(TestItem{Id: i, Value: "not it"})
		}
	}

	cmp := func(toCheck TestItem) bool {
		if toCheck.Id == 1 {
			// Pause on the first item to give the Clear() call time to run
			time.Sleep(2 * time.Second)
		}
		if toCheck.Value == "it" {
			return true
		}
		return false
	}

	// Read and check the result set
	// Concurrently clear the array
	var wg sync.WaitGroup
	output := make(chan []TestItem, 1)
	go func(out chan []TestItem) {
		wg.Add(1)
		defer wg.Done()
		res := ds.FindAll(cmp)
		output <- res
	}(output)
	go func() {
		wg.Add(1)
		defer wg.Done()
		// Wait a second to give FindAll() time to start
		time.Sleep(1 * time.Second)
		ds.Clear()
	}()
	res := <-output
	wg.Wait()
	assert.Equal(t, len(res), 1)
	assert.Equal(t, res[0].Id, 500)
	assert.Equal(t, res[0].Value, "it")
}
