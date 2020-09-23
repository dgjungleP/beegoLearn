package mapset_test

import (
	"testing"
)

type Mapset map[int]bool

func (set Mapset) add(i int) {
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
	mapset.add(2)
	mapset.add(3)
	mapset.add(3)
	t.Log(mapset)
}
