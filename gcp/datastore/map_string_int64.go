package datastore

import "cloud.google.com/go/datastore"

type MapStringInt64 map[string]int64

func (t *MapStringInt64) AsEntity() *datastore.Entity {
	var props []datastore.Property

	for k, v := range *t {
		props = append(props, datastore.Property{Name: k, Value: v, NoIndex: true})
	}

	return &datastore.Entity{Key: nil, Properties: props}
}

func (t *MapStringInt64) FromProps(ps []datastore.Property) {
	for _, p := range ps {
		(*t)[p.Name] = p.Value.(int64)
	}
}
