package utils

import (
	"testing"
)

type Foo struct {
	Id   int64
	Text string
}

var (
	collectionFoo = []Foo{
		Foo{1, "Foo"},
		Foo{2, "Bar"},
	}
)

func TestCollectWithExistingField(t *testing.T) {
	newColl := Collect(collectionFoo, "Text")
	for i, foo := range collectionFoo {
		if newColl[i] != foo.Text {
			t.Errorf("Collect Array position %d should be %s, but was %s", i, foo.Text, newColl[i])
		}
	}
}

func TestCollectWithNoExistingField(t *testing.T) {
	newColl := Collect(collectionFoo, "NoField")
	if len(newColl) > 0 {
		t.Errorf("Collect Array len should be 0, but was %d", len(newColl))
	}
}

func TestToMap(t *testing.T) {
	newMap := ToMap(collectionFoo, "Id", false)
	for _, foo := range collectionFoo {
		if newMap[foo.Id] != foo {
			t.Errorf("%+v should not be different from %+v", foo, newMap[foo.Id])
		}
	}
}

func TestToMapMultiple(t *testing.T) {
	multipleCollectionFoo := append(collectionFoo, Foo{1, "Foo Bar"})
	newMap := ToMap(multipleCollectionFoo, "Id", true)
	multipleValues := newMap[1].([]interface{})

	if len(multipleValues) < 1 {
		t.Error("Len should be more than 0")
	}

	for _, foo := range multipleValues {
		if foo.(Foo).Id != 1 {
			t.Errorf("Id should be 1, but was %d", foo.(Foo).Id)
		}
	}
}
