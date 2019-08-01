package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums1 := []int{14, 14, 22, 22, 23, 24, 24}
	nums2 := []int{14, 15, 14, 21, 22}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(intersect(nums1, nums2))

}
