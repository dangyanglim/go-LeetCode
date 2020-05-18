package main
import (
    "fmt"
    "testing"
)

func Test0001(t *testing.T) {
	nums1 :=[]int{1,2,3}
	nums2 :=[]int{4,5}
	fmt.Println(nums1)
	findMedianSortedArrays(nums1,nums2)
}