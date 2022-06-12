package main

import (
	"aperez24/heaps/heap"
	"fmt"
)

func main() {
	fmt.Println("running")
	comparator := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}
	heap := heap.NewHeap(comparator)
	heap.Put(12)
	heap.Put(3)
	heap.Put(4)
	heap.Put(1)
	heap.Put(5)
	for !heap.IsEmpty() {
		fmt.Printf("%v\n", heap.Peak())
		heap.Remove()
	}
}
