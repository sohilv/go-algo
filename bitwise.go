package main

import (
	"fmt"
)

func main() {
	p:=1234 
	fmt.Println("p>>1", p>>1, "p/2",p/2)
	fmt.Println("p<<1", p<<1, "p*2",p*2)
	fmt.Println("p&1", p&1, "(p>>1)& 1", (p>>1)& 1)
	fmt.Printf("%b - p|1 - %b - value - %v \n",p, p|1, p|1)
	fmt.Printf("%b - p|4 - %b - value - %v \n",p, p|4, p|4)
	fmt.Printf("%b - p|1<<2 - %b - value - %v \n",p, p|1<<2 ,p|1<<2)
	fmt.Printf("%b - p^1 - %b - value - %v \n",p, p^1, p^1)
	
	fmt.Printf("%b - p^2 - %b - value - %v \n",p, p^2, p^2)
	fmt.Printf("\n---------------------------\n%b\n%b\n%b\n --- \n%v \n",p,3, p&3, p&3)
	fmt.Printf("\n---------------------------\n%b\n%b\n%b\n --- \n%v \n",p,3, p|3, p|3)
	fmt.Printf("\n---------------------------\n%b\n%b\n%b\n --- \n%v \n",p,3, p^3, p^3)
}
