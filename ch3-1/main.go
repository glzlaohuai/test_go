package main


 import(
	 "fmt"
 )


 func main(){
	 fmt.Println("main invoked")
	 var x int8 =7
	 var y int8 =5
	//  fmt.Println(^x+1)
	//  fmt.Println("invoked is: ",^x+1)
	fmt.Printf("x: %08b\n",x)
	fmt.Printf("y: %08b\n",y)

	 fmt.Printf("binary form: %08b\n",x^y) // the difference
 }