// Package knap addresses the knapsack packing problem.
package knap

import (
	"fmt"

	"github.com/siuyin/dflt"
)

type Item struct {
	Name  string
	Cost  int
	Value int
}

// Pack takes a slice of items and a cost budget
// and returns total value
// and a slice of references to items selected.
func Pack(items []Item, budget int, m Memo) (int, []Item) {
	if len(items) == 0 {
		return 0, []Item{}
	}
	useMemo := dflt.EnvString("USE_MEMO", "1")
	if useMemo == "1" {
		v, i, ok := m.get(items, budget)
		if ok {
			//fmt.Printf("DEBUG: getting %v, %d\n", items, budget)
			return v, i
		}
	}
	if items[0].Cost > budget { // item 0 is too costly, pack remaining items.
		v, i := Pack(items[1:], budget, m)
		m.put(items, budget, v, i)
		return v, i
	}

	cand := items[0] // candidate item, to select or not?

	candIn, candInItems := Pack(items[1:], budget-cand.Cost, m) // candIn: scenario where we selected the candidate.
	candIn += cand.Value

	candOut, candOutItems := Pack(items[1:], budget, m)

	if candIn > candOut {
		candInItems = append(candInItems, cand)
		m.put(items, budget, candIn, candInItems)
		return candIn, candInItems
	}
	m.put(items, budget, candOut, candOutItems)
	return candOut, candOutItems
}

type ans struct {
	v int
	i []Item
}

type Memo map[string]ans

func NewMemo() Memo {
	m := Memo(make(map[string]ans))
	return m
}

func (m Memo) put(items []Item, budget int, valAns int, itemsAns []Item) {
	s := fmt.Sprintf("%d,%d", len(items), budget)
	m[s] = ans{valAns, itemsAns}
}

func (m Memo) get(items []Item, budget int) (int, []Item, bool) {
	s := fmt.Sprintf("%d,%d", len(items), budget)
	a, ok := m[s]
	if !ok {
		return 0, []Item{}, false
	}
	return a.v, a.i, true
}
