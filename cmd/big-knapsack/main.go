package main

import (
	"fmt"
	"math/rand"
	"time"

	knap "github.com/siuyin/knapsack"
)

const budget = 15

func main() {
	fmt.Println("Big knapsack with lots of items to choose from.")
	fmt.Printf("budget: %d\n", budget)
	for i := 5; i < 100; i += 5 {
		items := genItems(i, 10, 20)
		m := knap.NewMemo()
		start := time.Now()
		c, _ := knap.Pack(items, budget, m)
		end := time.Now()
		fmt.Printf("n: %d, v: %d, time: %v seconds\n", i, c, end.Sub(start).Seconds())
	}
}

func genItems(n, maxCost, maxValue int) []knap.Item {
	r := []knap.Item{}
	for i := 0; i < n; i++ {
		c := rand.Intn(maxCost) + 1 // 1 .. maxCost
		v := rand.Intn(maxValue) + 1
		name := fmt.Sprintf("Item%d", i)
		r = append(r, knap.Item{Name: name, Cost: c, Value: v})
	}
	return r
}
