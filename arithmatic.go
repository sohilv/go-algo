package main

import (
	"fmt"
)

const(
	 MAXINT int64 = 1<<63-1  
)


func main() {

	 fmt.Println(pow(3, 234))
	 
	fmt.Println(Multiply(34342, 27))
	
	
}
func Multiply(a, b int) int {
	fmt.Println(">",a, b )
	if a == 0 || b == 0 {
		return 0
	}
	
	if b == 1{
	
		return a
	}
	
	temp:= Multiply(a, b >> 1)
	
	if b &1 == 0 {
		return temp+temp
	}
	
	temp= temp+ temp+a
	
	return temp
}

func pow(a, b int) int{

	
	
	fmt.Println(">",a, b )
	if a==0 && b!=0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	
	res:=pow(a, b >>1)
	
	if (b & 1) == 0 {
	 	return res*res
	}
	return res* res*  a

}