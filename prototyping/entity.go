package prototyping

import (
	"github.com/relvox/iridescence_go/sets"
)

type Entity struct {
	Id         int               `json:"Id"`
	Name       string            `json:"Name"`
	Properties map[string]string `json:"Properties"`
	Stats      map[string]int    `json:"Stats"`
	Tags       sets.Set[string]  `json:"Tags"`
}

func NewEntity(id int, name string) *Entity {
	return &Entity{
		Id:         id,
		Name:       name,
		Properties: make(map[string]string),
		Stats:      make(map[string]int),
		Tags:       sets.NewSet[string](),
	}
}

func (e *Entity) Clone() *Entity {
	result := NewEntity(e.Id, e.Name)
	for key, value := range e.Properties {
		result.Properties[key] = value
	}
	for key, value := range e.Stats {
		result.Stats[key] = value
	}
	for key, value := range e.Tags {
		result.Tags[key] = value
	}
	return result
}

func (e *Entity) WithProperty(key string, value string) *Entity {
	e.Properties[key] = value
	return e
}

func (e *Entity) WithStat(key string, value int) *Entity {
	e.Stats[key] = value
	return e
}

func (e *Entity) WithTag(key string) *Entity {
	e.Tags[key] = sets.U
	return e
}
