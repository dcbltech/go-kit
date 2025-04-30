package datastore

import (
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/stretchr/testify/assert"
)

func TestSliceString_AsAny(t *testing.T) {
	ss := SliceString{"a", "b", "c"}
	result := ss.AsAny()

	assert.Equal(t, []any{"a", "b", "c"}, result)
}

func TestSliceString_FromProp(t *testing.T) {
	ss := SliceString{}
	prop := datastore.Property{Value: []any{"x", "y", "z"}}

	ss.FromProp(prop)
	assert.Equal(t, SliceString{"x", "y", "z"}, ss)
}

func TestSliceString_AddUnique(t *testing.T) {
	ss := SliceString{"a", "b"}

	ss.AddUnique("c")
	assert.Equal(t, SliceString{"a", "b", "c"}, ss)

	ss.AddUnique("b")
	assert.Equal(t, SliceString{"a", "b", "c"}, ss)
}

func TestSliceString_AddUniques(t *testing.T) {
	ss := SliceString{"a"}

	ss.AddUniques([]string{"b", "c", "a"})
	assert.Equal(t, SliceString{"a", "b", "c"}, ss)
}

func TestSliceString_RemoveAll(t *testing.T) {
	ss := SliceString{"a", "b", "c", "d"}

	ss.RemoveAll([]string{"b", "d"})
	assert.Equal(t, SliceString{"a", "c"}, ss)
}
