package main
import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	fmt.Println(m)
	fmt.Println(nums1)
	fmt.Println(n)
	if n == 0 {
		for m > 0 {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	if m == 0 {
		for n > 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

}

func main() {

}
