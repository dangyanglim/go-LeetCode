package main
import (
	"strings"
)
func isPalindrome(s string) bool {
	var st int=0
	var ss string=strings.ToLower(s)
	var end int=len(ss)-1
	for {
		if st>end{
			return true
		}
		if (ss[st]>='0'&&ss[st]<='9')||(ss[st]>='a'&&ss[st]<='z'){
			if (ss[end]>='0'&&ss[end]<='9')||(ss[end]>='a'&&ss[end]<='z'){
				if ss[st]!=ss[end] {
					return false
				}else {
					st++
					end--
				}
			}else{
				end--
			}
		}else{
			st++
		}
	}

}

func main() {

}
