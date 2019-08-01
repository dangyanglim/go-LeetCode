package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{2, 14, 18, 22, 23}
	fmt.Println(nums)
	fmt.Println(containsDuplicate(nums))

}
func Test0002(t *testing.T) {
	nums := []int{2, 14, 18, 22, 22}
	fmt.Println(nums)
	fmt.Println(containsDuplicate2(nums))

}
