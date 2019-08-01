package main

func reverse(x int) int {
	const INT_MIN int=-2147483648
	const INT_MAX int=2147483648
	res:=0
	temp:=0
	
	for {
		if x==0 {
			break
		}
		temp=res*10+x%10
		if temp/10!=res {
			return 0
		}
		res=temp
		if res < INT_MIN || res > INT_MAX {
			return 0;
		}
		x/=10
	}
	return res
}

func main() {

}
