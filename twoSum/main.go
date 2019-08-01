package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {

				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
func twoSum2(nums []int, target int) []int {
	t := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		t[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		two, ok := t[target-nums[i]]
		if ok && i != two {
			return []int{two, i}
		}
	}
	return []int{0, 0}
}

func main() {

}
