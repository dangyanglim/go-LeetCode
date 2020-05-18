package main

import (
    "strings"
)
func main(){

}

func intToRoman(num int) string {
    var ret int
    var build strings.Builder
    for ret=num/1000;ret>0; {
        num=num-1000
        build.WriteString("M")
        ret=num/1000
    }
    for ret=num/100;ret>0;{
        if ret==9{
            build.WriteString("CM")
            num=num-900
        }else if ret>=5{
            build.WriteString("D")
            num=num-500
        }else if ret==4{
            build.WriteString("CD")
            num=num-400
        }else{
            build.WriteString("C")
            num=num-100
        }
        ret=num/100
    }
    for ret=num/10;ret>0;{
        if ret==9{
            build.WriteString("XC")
            num=num-90
        }else if ret>=5{
            build.WriteString("L")
            num=num-50
        }else if ret==4{
            build.WriteString("XL")
            num=num-40
        }else{
            build.WriteString("X")
            num=num-10
        }
        ret=num/10       
    }
    for ;num>0;{
        if num==9{
            build.WriteString("IX")
            num=num-9
        }else if num>=5{
            build.WriteString("V")
            num=num-5
        }else if num==4{
            build.WriteString("IV")
            num=num-4
        }else{
            build.WriteString("I")
            num=num-1
        }
      
    }    
    return build.String()

}