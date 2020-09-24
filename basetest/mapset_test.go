package mapset_test

import (
	"testing"
)

type Mapset map[interface{}]bool
type person struct {
	id   int
	name string
}

func (set Mapset) add(i interface{}) {
	if !set[i] {
		set[i] = true
	}
}

func (set Mapset) delete(i int) {
	delete(set, i)
}

func TestMapSet(t *testing.T) {
	mapset := make(Mapset)
	mapset.add(1)
	mapset.add(new(person))
	mapset.add(&person{name: "John"})
	mapset.add(new(person))
	mapset.add(&person{id: 1})
	mapset.add(3)
	mapset.add(3)
	t.Log(mapset)
}
