package main

func removeDuplicates(nums []int) int {
	numsLenght := len(nums)
	if numsLenght < 1 {
		return 0
	}	
	if numsLenght < 2 {
		return 1
	}
	j := 0
	for i := 1; i < numsLenght; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
func main() {

}
