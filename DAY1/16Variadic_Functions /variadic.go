package main

func add(num ...int)int{
	 sum1:=0;
	 for _,val:=range num{
		sum1=sum1+val
	 }
	 return sum1
}



func main(){
	Ans:=add(1,2,3,4,5,6,7,8,9,10)
	println("sum:",Ans)
}