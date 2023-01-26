package storage

import (
	"fmt"
	"testing"

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

func benchmarkAdd(ds DataStore[TestItem], size int) {
	for i := 0; i < size; i++ {
		ds.Add(TestItem{Id: i, Value: "hello world"})
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()

	for _, v := range sizes {
		var ds DataStore[TestItem]
		b.Run(fmt.Sprintf("count_%d", v.count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkAdd(ds, v.count)

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

func benchmarkAddRef(ds DataStore[TestItem], size int) {
	for i := 0; i < size; i++ {
		ds.AddRef(&TestItem{Id: i, Value: "hello world"})
	}
}

func BenchmarkAddRef(b *testing.B) {
	b.ReportAllocs()

	for _, v := range sizes {
		var ds DataStore[TestItem]
		b.Run(fmt.Sprintf("count_%d", v.count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmarkAddRef(ds, v.count)
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

// NOTE(mlee): This is expected to fail
func TestGoroutineSafeUpdates(t *testing.T) {

}
