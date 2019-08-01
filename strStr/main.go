package main

func strStr(haystack string, needle string) int {
    if len(needle)==0 {
		return 0
	}
	st:=0
	for i:=0;i<len(haystack)-len(needle)+1;i++{
		st=i
		for j:=0;j<len(needle);j++{
			if needle[j]==haystack[st]{
				st++
			}else{
				break;
			}
			if j==len(needle)-1{
				return i
			}
		}

	}
	return -1
}

func main() {

}
