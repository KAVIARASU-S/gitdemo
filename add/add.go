package add

func Add(a,b int)(sum int){
	sum=a+b
	sub:=a-b
	fmt.Println("The diff is %d",sub)
	return sum
}
