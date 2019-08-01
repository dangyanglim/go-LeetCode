package main

func maxSubArray(nums []int) int {
    if len(nums)==0{
        return 0
    }
    sum:=nums[0]
    temp:=nums[0]
    for i:=1;i<len(nums);i++{
        if temp<0{
            temp=nums[i]
        }else{
            temp+=nums[i]
		}
        if sum<temp{
            sum=temp
        }
    }
    return sum
}

func main() {

}
