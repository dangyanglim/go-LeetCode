package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0,}
	nums1 := []int{4,5,6}
	fmt.Println(nums)

	merge(nums, 3, nums1, 3)
	fmt.Println(nums)

}
