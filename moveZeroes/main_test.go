package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{14, 0, 23, 22, 0, 25, 24}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)

}
