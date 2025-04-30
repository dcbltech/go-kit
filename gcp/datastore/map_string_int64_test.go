package datastore

import (
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/stretchr/testify/assert"
)

func TestMapStringInt64_AsEntity(t *testing.T) {
	input := MapStringInt64{"key1": 123, "key2": 456}
	expectedProps := []datastore.Property{
		{Name: "key1", Value: int64(123), NoIndex: true},
		{Name: "key2", Value: int64(456), NoIndex: true},
	}

	entity := input.AsEntity()

	assert.Nil(t, entity.Key)
	assert.ElementsMatch(t, expectedProps, entity.Properties)
}

func TestMapStringInt64_FromProps(t *testing.T) {
	props := []datastore.Property{
		{Name: "key1", Value: int64(123)},
		{Name: "key2", Value: int64(456)},
	}

	var output MapStringInt64 = make(map[string]int64)

	output.FromProps(props)

	expected := MapStringInt64{"key1": 123, "key2": 456}
	assert.Equal(t, expected, output)
}
