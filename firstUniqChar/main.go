package main

func firstUniqChar(s string) int {
	var ret bool=true
    for i:=0;i<len(s);i++{
		ret=true
		for j:=0;j<len(s);j++{		
			if s[i]==s[j] && i!=j{
				ret=false
				break;
			}
		}
		if ret {
			return i
		}
	}
	return -1
}

func main() {

}
