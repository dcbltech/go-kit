package datastore

import "cloud.google.com/go/datastore"

type MapStringString map[string]string

func (t *MapStringString) AsEntity() *datastore.Entity {
	var props []datastore.Property

	for k, v := range *t {
		props = append(props, datastore.Property{Name: k, Value: v, NoIndex: true})
	}

	return &datastore.Entity{Key: nil, Properties: props}
}

func (t *MapStringString) FromProps(ps []datastore.Property) {
	for _, p := range ps {
		(*t)[p.Name] = p.Value.(string)
	}
}
