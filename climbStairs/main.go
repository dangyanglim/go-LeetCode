package main


func climbStairs(n int) int {
    if n==0||n==1||n==2{
		return n
	}else{
		return climbStairs(n-1)+climbStairs(n-2)
	}
}
func climbStairs2(n int) int {
    if n==0||n==1||n==2{
		return n
	}else{
		a:=1
		b:=2
		temp:=0
		for i:=3;i<n+1;i++{
			temp=a+b
			a=b
			b=temp

		}
		return temp
	}
}
func main() {

}
