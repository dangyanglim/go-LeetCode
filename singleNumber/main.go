package main

func singleNumber(nums []int) int {

	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}


func main() {

}
