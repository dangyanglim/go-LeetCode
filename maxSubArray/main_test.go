package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(nums)

	
	fmt.Println(maxSubArray(nums))

}
