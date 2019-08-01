package main

func reverseString(s []byte)  {
	length:=len(s)
	var	temp byte
	for i := 0; i < length; i++ {

		if i<length/2 {		
			temp=s[length-i-1]
			s[length-i-1]=s[i]
			s[i]=temp
		}
	}
}

func main() {

}
