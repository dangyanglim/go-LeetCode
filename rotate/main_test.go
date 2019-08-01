package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(nums)

	rotate(nums)
	fmt.Println(nums)

}
