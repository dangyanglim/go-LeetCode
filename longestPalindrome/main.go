package main

import (
	"fmt"
)

func main(){
	ret:=longestPalindrome("eabcb")
	fmt.Printf(ret)
}
func isPalindrome(s string) bool {
	var st int=0
	var ss string=s
	var end int=len(ss)-1
	for {
		if st>end{
			return true
		}
				if ss[st]!=ss[end] {
					return false
				}else {
					st++
					end--
				}
	}
}
func longestPalindrome(s string) string {
    m:=len(s)
    ret:=""
	for i:=0;i<m;i++{
		for j:=i+1;j<=m;j++{
			//fmt.Println(s[i:j])
			if isPalindrome(s[i:j])&&len(ret)<j-i{
				ret=s[i:j]
			}
		}
	}
    return ret
}