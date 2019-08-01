package main
import (
    "fmt"
    "testing"
)

func Test0001(t *testing.T) {
	nums :=[]int{1,5,3,2}
	fmt.Println(nums)
	rotate(nums,1)
	fmt.Println(nums)
}