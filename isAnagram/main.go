package main

func isAnagram(s string, t string) bool {
	var temp byte=0;
	var des []byte;
	var tar []byte;
	for l:=0;l<len(s);l++{
		des=append(des,s[l])
	}
	for l:=0;l<len(t);l++{
		tar=append(tar,t[l])
	}
	for i:=0;i<len(des);i++ {
		for j:=i+1;j<len(des);j++ {
			if des[j]>des[i] {
				temp=des[i]
				des[i]=des[j]
				des[j]=temp
			}
		}
	}
	for m:=0;m<len(tar);m++ {
		for n:=m+1;n<len(tar);n++ {
			if tar[n]>tar[m] {
				temp=tar[m]
				tar[m]=tar[n]
				tar[n]=temp
			}
		}
	}
	if len(des)==len(tar) {
		for m:=0;m<len(tar);m++{
			if tar[m]!=des[m]{
				return false
			}
		}
		return true
	}
	return false
    
}

func main() {

}
