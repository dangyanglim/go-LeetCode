package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	fmt.Println(nums)
	removeDuplicates(nums)
	fmt.Println(nums)
}
