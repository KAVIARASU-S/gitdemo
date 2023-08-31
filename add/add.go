package add

import "fmt"

func Add(a,b int)(sum int){
	sum=a+b
	sub:=a-b
	fmt.Println("The diff is ",sub)
	return sum
}
