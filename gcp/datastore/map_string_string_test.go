package datastore

import (
	"reflect"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestMapStringString_AsEntity(t *testing.T) {
	input := MapStringString{
		"key1": "value1",
		"key2": "value2",
	}

	expectedProps := []datastore.Property{
		{Name: "key1", Value: "value1", NoIndex: true},
		{Name: "key2", Value: "value2", NoIndex: true},
	}

	entity := input.AsEntity()

	if !reflect.DeepEqual(entity.Properties, expectedProps) {
		t.Errorf("AsEntity() properties = %v, want %v", entity.Properties, expectedProps)
	}
}

func TestMapStringString_FromProps(t *testing.T) {
	props := []datastore.Property{
		{Name: "key1", Value: "value1"},
		{Name: "key2", Value: "value2"},
	}

	expected := MapStringString{
		"key1": "value1",
		"key2": "value2",
	}

	var result MapStringString = make(map[string]string)

	result.FromProps(props)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromProps() result = %v, want %v", result, expected)
	}
}
