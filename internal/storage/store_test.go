// Naive generic storage for basic functionality
// TODO(mlee): Replace this with real, go-routine safe storage. Preferably something with an external datastore that can be persisted to disk
package storage

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestItem struct {
	Id    int
	Value string
}

func TestAdd(t *testing.T) {
	var ds DataStore[TestItem]
	i1 := TestItem{Id: 1, Value: "hello world"}
	i2 := TestItem{Id: 2, Value: "goodbye moon"}
	i3 := TestItem{Id: 3, Value: "salutations sun"}
	ds.Add(i1)
	ds.Add(i2)
	ds.Add(i3)
}

func TestFindOne(t *testing.T) {
	cmp := func(toCheck TestItem) bool {
		if toCheck.Value == "goodbye moon" {
			return true
		}
		return false
	}

	ds := DataStore[TestItem]{
		Items: []TestItem{
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
		Items: []TestItem{
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
