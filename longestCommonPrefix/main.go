package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var minLen int = len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	fmt.Printf("%d", minLen)
	var temp byte = 0
	for j := 0; j < minLen; j++ {
		temp = strs[0][j]
		for k := 0; k < len(strs); k++ {
			if strs[k][j] != temp {
				return strs[0][0:j]
			}
		}
		// if j == 0 && j == minLen-1 {
		// 	return strs[0][0:1]
		// }
		if j == minLen-1 {
			return strs[0][0:minLen]
		}
	}
	return ""
}

func main() {

}
