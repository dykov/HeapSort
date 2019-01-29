package main

import "fmt"

type MaxHeap struct {
	slice    []int
	heapSize int
}

func BuildMaxHeap(slice []int) MaxHeap {

	heap := MaxHeap{
		slice: slice,
		heapSize: len(slice),
	}

	for i := len(slice)/2; i>=0; i-- {
		heap.MaxHeapify(i)
	}

	return heap
}

func (h MaxHeap) MaxHeapify(i int) {

	l := 2*i+1
	r := 2*i+2
	max := i

	if l < h.heapSize && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.heapSize && h.slice[r] > h.slice[max] {
		max = r
	}

	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}

}

func heapSort(slice []int) []int {

	heap := BuildMaxHeap(slice)
	for i := len(heap.slice) - 1; i >= 1; i-- {
		heap.slice[0], heap.slice[i] = heap.slice[i], heap.slice[0]
		heap.heapSize--
		heap.MaxHeapify(0)
	}
	return heap.slice

}

func main() {

	var array = []int{20,30,1,2,10,0,7,-1}
	fmt.Println(heapSort(array))

}
