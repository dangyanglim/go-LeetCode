package main

import (

	"strconv"
)

func countAndSay(n int) string {
	var res string = "1"
	var cur string = ""
	var cnt int = 1
	if n < 0 {
		return ""
	}
	for i := 0; i < n-1; i++ {
		cur = ""

		for j := 0; j < len(res); j++ {
			cnt = 1
			for {
				if j+1 < len(res) && res[j] == res[j+1] {
					cnt++
					j++
				} else {
					break
				}
			}

			cur += strconv.Itoa(cnt) + string(res[j])
		}
		res = cur
	}
	return res
}
func countAndSay2(n int) string {
	var res string = "1"
	var cur string = ""
	var cnt int = 1
	if n < 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	res=countAndSay2(n-1)
	for i := 0; i < len(res); i++ {
		cnt = 1
		for {
			if i+1 < len(res) && res[i] == res[i+1] {
				cnt++
				i++
			} else {
				break
			}
		}

		cur += strconv.Itoa(cnt) + string(res[i])		
	}
	return cur
}
func main() {

}
