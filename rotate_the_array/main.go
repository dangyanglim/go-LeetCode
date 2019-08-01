package main

func rotate(nums []int, k int)  {
	numsLenght:=len(nums)
    if(k<1){return}
	if(numsLenght<2){return}
	k%=numsLenght
    for i:=0;i<k;i++ {
		temp:=nums[numsLenght-1]
        for j:=0;j<numsLenght-1;j++ {			
            nums[numsLenght-1-j]=nums[numsLenght-2-j]
		}
		nums[0]=temp;
    }
    return;
}
func main(){

}