package main

func plusOne(digits []int) []int {
	carry := 1
	tmp := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 0 {
			return digits
		}
		tmp = digits[i] + carry
		carry = tmp / 10
		digits[i] = tmp % 10
	}
	if carry != 0 {
		re := make([]int, len(digits)+1)
		re[0] = 1
		return re
	}
	return digits
}

func main() {

}
