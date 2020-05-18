package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{5, 3, 4, 4, 1, 5, 2, 2, 9, 10, 23, 13, 7, 6, 8, 9, 11}
	//nums :=[]int {2,2,1,4,5,4,3}
	//heapify_sort(nums)
	///quickSort(nums)
	bubbleSort(nums)
	fmt.Println(nums)

}

func Benchmark_heap_sort(b *testing.B) {
	nums :=[]int {1,2,9,5,8,7,4,6,8,5,9,5, 3, 4, 4, 1, 5, 2, 2, 9, 10, 23, 13, 7, 6, 8, 9, 11}
	for i := 0; i < b.N; i++ {
		heapify_sort(nums[:])
	}
}
func Benchmark_qsort(b *testing.B) {
	nums :=[]int {1,2,9,5,8,7,4,6,8,5,9,5, 3, 4, 4, 1, 5, 2, 2, 9, 10, 23, 13, 7, 6, 8, 9, 11}
	for i := 0; i < b.N; i++ {
		quickSort(nums[:])
	}
}
func Benchmark_bsort(b *testing.B) {
	nums :=[]int {1,2,9,5,8,7,4,6,8,5,9,5, 3, 4, 4, 1, 5, 2, 2, 9, 10, 23, 13, 7, 6, 8, 9, 11}
	for i := 0; i < b.N; i++ {
		bubbleSort(nums)
	}
}
