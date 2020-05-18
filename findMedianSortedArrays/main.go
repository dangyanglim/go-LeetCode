package main

import (
	"fmt"
)
func main(){

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1)==0{
        if len(nums2)%2==0{
            return float64(nums2[len(nums2)/2-1]+nums2[len(nums2)/2])/2
        }else{
            return float64(nums2[len(nums2)/2])
        }
    }
    if len(nums2)==0{
        if len(nums1)%2==0{
            return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2])/2
        }else{
            return float64(nums1[len(nums1)/2])
        }
    }
    var nums []int
    count:=0
    i,j:=0,0
    for count!=(len(nums1)+len(nums2)){
        if i==len(nums1) {
            for j!=len(nums2){
                nums=append(nums,nums2[j])
                count++
                j++                
            }
            break
        }
        if j==len(nums2) {
            for i!=len(nums1){
                nums=append(nums,nums1[i])
                count++
                i++                
            }
            break
        }
        if nums1[i]<nums2[j] {
            nums=append(nums,nums1[i])
            count++
            i++
        }else {
            nums=append(nums,nums2[j])
            count++
            j++
        }
    }
    fmt.Println(nums)
       if count%2==0{
            return float64(nums[count/2-1]+nums[count/2])/2
        }else{
            return float64(nums[count/2])
        }

}