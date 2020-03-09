package knap

import (
	"fmt"
	"testing"
)

var c1 = []Item{
	Item{"pen", 1, 10},      // 0
	Item{"paper", 1, 9},     // 1
	Item{"water", 5, 20},    // 2
	Item{"peanuts", 2, 10},  // 3
	Item{"raincoat", 6, 12}, // 4
}

func TestPack(t *testing.T) {
	type d struct {
		b int    // budget
		v int    // value
		o []Item // output -- items selected
	}
	ds := []d{
		d{1, 10, []Item{c1[0]}},
		d{2, 19, []Item{c1[1], c1[0]}},
		d{10, 0, []Item{}},
		d{11, 0, []Item{}},
		d{12, 0, []Item{}},
		d{13, 0, []Item{}},
	}
	m := newMemo()
	for i := range ds {
		v, o := Pack(c1, ds[i].b, m) // output
		if fmt.Sprintf("%v", o) != fmt.Sprintf("%v", ds[i].o) {
			t.Errorf("case %d: budget: %d expected %v got %v", i, ds[i].b, ds[i].o, o)
		}
		if v != ds[i].v {
			t.Errorf("case %d: budget: %d expected %v got %v", i, ds[i].b, ds[i].v, v)
		}
	}
}
