package datastore

import (
	"time"

	"cloud.google.com/go/datastore"
)

const (
	entityFieldKey     = "__key__"
	entityFieldCreated = "_created"
)

type Entity struct {
	Key     *datastore.Key `json:"-"`
	Created time.Time      `json:"-"`
}

func (e *Entity) LoadKey(key *datastore.Key) error {
	e.Key = key

	return nil
}

type propertyLoader func(p datastore.Property) error

func (e *Entity) LoadProps(ps []datastore.Property, load propertyLoader) error {
	for i := 0; i < len(ps); i++ {
		switch ps[i].Name {
		case entityFieldCreated:
			e.Created = ps[i].Value.(time.Time)
		default:
			if err := load(ps[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Entity) SaveProps(ps []datastore.Property) ([]datastore.Property, error) {
	var props []datastore.Property

	if e.Created.IsZero() {
		e.Created = time.Now().UTC()
	}

	if e.Key != nil {
		props = append(
			props,
			datastore.Property{
				Name:    entityFieldKey,
				Value:   e.Key,
				NoIndex: false,
			},
		)
	}

	props = append(props,
		datastore.Property{
			Name:    entityFieldCreated,
			Value:   e.Created,
			NoIndex: true,
		},
	)

	props = append(props, ps...)

	return props, nil
}

func (e *Entity) LoadNullTime(p datastore.Property) *time.Time {
	var t *time.Time

	if p.Value != nil {
		if v, ok := p.Value.(time.Time); ok {
			t = &v
		} else {
			t = nil
		}
	}

	return t
}
