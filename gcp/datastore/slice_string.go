package datastore

import (
	"slices"

	"cloud.google.com/go/datastore"
)

type SliceString []string

func (t *SliceString) AsAny() []any {
	v := make([]any, len(*t))

	for i := range *t {
		v[i] = (*t)[i]
	}

	return v
}

func (t *SliceString) FromProp(p datastore.Property) {
	if p.Value == nil {
		return
	}

	for _, v := range p.Value.([]any) {
		*t = append(*t, v.(string))
	}
}

func (t *SliceString) AddUnique(s string) {
	if !slices.Contains(*t, s) {
		*t = append(*t, s)
	}
}

func (t *SliceString) AddUniques(s []string) {
	for _, v := range s {
		if !slices.Contains(*t, v) {
			*t = append(*t, v)
		}
	}
}

func (t *SliceString) RemoveAll(r []string) {
	*t = slices.DeleteFunc(*t, func(s string) bool {
		return slices.Contains(r, s)
	})
}
