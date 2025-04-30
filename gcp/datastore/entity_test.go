package datastore

import (
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/stretchr/testify/assert"
)

func TestEntity_LoadKey(t *testing.T) {
	key := &datastore.Key{Kind: "TestKind", Name: "TestName"}
	entity := &Entity{}

	err := entity.LoadKey(key)

	assert.NoError(t, err)
	assert.Equal(t, key, entity.Key)
}

func TestEntity_LoadProps(t *testing.T) {
	entity := &Entity{}
	created := time.Now().UTC()
	props := []datastore.Property{
		{Name: entityFieldCreated, Value: created},
		{Name: "OtherField", Value: "OtherValue"},
	}

	load := func(p datastore.Property) error {
		if p.Name == "OtherField" {
			assert.Equal(t, "OtherValue", p.Value)
		}

		return nil
	}

	err := entity.LoadProps(props, load)

	assert.NoError(t, err)
	assert.True(t, entity.Created.Equal(created))
}

func TestEntity_SaveProps(t *testing.T) {
	entity := &Entity{
		Key:     &datastore.Key{Kind: "TestKind", Name: "TestName"},
		Created: time.Now().UTC(),
	}

	props, err := entity.SaveProps([]datastore.Property{
		{Name: "OtherField", Value: "OtherValue"},
	})
	assert.NoError(t, err)

	foundKey := false
	foundCreated := false

	for _, p := range props {
		if p.Name == entityFieldKey {
			foundKey = true
			assert.Equal(t, entity.Key, p.Value)
		}

		if p.Name == entityFieldCreated {
			foundCreated = true
			assert.Equal(t, entity.Created, p.Value)
		}
	}

	assert.True(t, foundKey, "key property not found")
	assert.True(t, foundCreated, "created property not found")
}

func TestEntity_LoadNullTime(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name     string
		property datastore.Property
		expected *time.Time
	}{
		{
			name:     "Valid time",
			property: datastore.Property{Value: now},
			expected: &now,
		},
		{
			name:     "Nil value",
			property: datastore.Property{Value: nil},
			expected: nil,
		},
		{
			name:     "Invalid type",
			property: datastore.Property{Value: "invalid"},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := &Entity{}
			result := entity.LoadNullTime(tt.property)
			assert.Equal(t, tt.expected, result)
		})
	}
}
