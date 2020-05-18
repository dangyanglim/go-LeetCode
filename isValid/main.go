package main

import (
	"fmt"
	"strings"
)
func main() {
	isValid("([)]")
}
func isValid(s string) bool {
    s_new:=s
    
    s_new=strings.Replace(s_new,"()","",-1)
    fmt.Printf(s_new)
	s_new=strings.Replace(s_new,"{}","",-1)
	fmt.Printf(s_new)
	s_new=strings.Replace(s_new,"[]","",-1)
	fmt.Printf(s_new)

    return s==s_new
}