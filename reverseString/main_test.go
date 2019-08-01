package main

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	nums := []byte{'h','e','l','l','o'}
	fmt.Println(nums)
	reverseString(nums)
	fmt.Println(nums)

}
