package main

func myAtoi(str string) int {
	INT_Max := int(^uint32(0) >> 1)
    INT_Min := ^INT_Max
	x:=0
	base:=0
	sign := 1
	if str==""{
		return 0
	}
	for x<len(str) && str[x]==' ' {
		x++
	}

	if x<len(str) {
		if str[x]=='-'{
			x++
			sign=-1
		}else if str[x]=='+'{
			x++
			sign=1
		}
	}
	for x<len(str)&&str[x]>='0'&&str[x]<='9' {
        if base > INT_Max/10 || (base == INT_Max/10 && str[x]-'0' > 7) {
			if sign == 1 {
				return INT_Max
			} else {
				return INT_Min
			}
		}
		base=10*base+int(str[x]-'0')
		x++
	}
	return base*sign

}

func main() {

}
