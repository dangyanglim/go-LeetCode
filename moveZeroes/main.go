package main



func moveZeroes(nums []int) {

	for i := len(nums) - 1; i >= 0; i-- {

		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
				nums[j+1] = 0
			}
		}
	}

}

func main() {

}
