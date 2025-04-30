package datastore

import "cloud.google.com/go/datastore"

type SetString map[string]any

func (t *SetString) Add(s string) {
	(*t)[s] = nil
}

func (t *SetString) AddAll(s []string) {
	for _, v := range s {
		(*t)[v] = nil
	}
}

func (t *SetString) Remove(s string) {
	delete(*t, s)
}

func (t *SetString) AsSlice() []string {
	if len(*t) == 0 {
		return []string{}
	}

	var data []string

	for k := range *t {
		data = append(data, k)
	}

	return data
}

func (t *SetString) AsProp(name string, noIndex bool) datastore.Property {
	var data []any

	for k := range *t {
		data = append(data, k)
	}

	return datastore.Property{Name: name, Value: data, NoIndex: noIndex}
}

func (t *SetString) FromProp(p datastore.Property) {
	for _, v := range p.Value.([]any) {
		(*t)[v.(string)] = nil
	}
}
