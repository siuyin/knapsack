// Package knap addresses the knapsack packing problem.
package knap

type Item struct {
	Name  string
	Cost  int
	Value int
}

// Pack takes a slice of items and a cost budget
// and returns total value
// and a slice of references to items selected.
func Pack(items []Item, budget int) (int, []Item) {
	if len(items) == 0 {
		return 0, []Item{}
	}
	if items[0].Cost > budget { // item 0 is too costly, pack remaining items.
		return Pack(items[1:], budget)
	}

	cand := items[0] // candidate item, to select or not?

	candIn, candInItems := Pack(items[1:], budget-cand.Cost) // candIn: scenario where we selected the candidate.
	candIn += cand.Value

	candOut, candOutItems := Pack(items[1:], budget)

	if candIn > candOut {
		candInItems = append(candInItems, cand)
		return candIn, candInItems
	}
	return candOut, candOutItems
}
