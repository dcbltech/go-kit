package datastore

import (
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestSetString_Add(t *testing.T) {
	set := SetString{}
	set.Add("test")

	if _, exists := set["test"]; !exists {
		t.Errorf("expected 'test' to be added to the set")
	}
}

func TestSetString_AddAll(t *testing.T) {
	set := SetString{}
	set.AddAll([]string{"a", "b", "c"})

	expected := SetString{"a": nil, "b": nil, "c": nil}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v, got %v", expected, set)
	}
}

func TestSetString_Remove(t *testing.T) {
	set := SetString{"a": nil, "b": nil}
	set.Remove("a")

	if _, exists := set["a"]; exists {
		t.Errorf("expected 'a' to be removed from the set")
	}
}

func TestSetString_AsSlice(t *testing.T) {
	set := SetString{"a": nil, "b": nil}
	slice := set.AsSlice()

	expected := []string{"a", "b"}

	if !reflect.DeepEqual(slice, expected) && !reflect.DeepEqual(slice, []string{"b", "a"}) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}

func TestSetString_AsProp(t *testing.T) {
	set := SetString{"a": nil, "b": nil}
	prop := set.AsProp("test", true)

	expected := datastore.Property{
		Name:    "test",
		Value:   []any{"a", "b"},
		NoIndex: true,
	}

	if !reflect.DeepEqual(prop, expected) && !reflect.DeepEqual(prop.Value, []any{"b", "a"}) {
		t.Errorf("expected %v, got %v", expected, prop)
	}
}

func TestSetString_FromProp(t *testing.T) {
	prop := datastore.Property{
		Value: []any{"a", "b"},
	}

	set := SetString{}
	set.FromProp(prop)

	expected := SetString{"a": nil, "b": nil}

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("expected %v, got %v", expected, set)
	}
}
