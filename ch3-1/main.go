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
	//  fmt.Printf("%d, %#[1]o, %#[1]x",x)
	 fmt.Printf("print with extra signs: %d, %#[1]o, %#[1]x\n",x)


	 ascii:='a'
	 unicode:='你'
	 newline:='\n'

	 fmt.Printf("%d, %[1]c, %[1]q\n",ascii)
	 fmt.Printf("%d, %[1]c, %[1]q\n",unicode)
	 fmt.Printf("%d, %[1]c, %[1]q\n",newline)

	 var float_x float32 = 1<<24
	 xx:=float_x+1
	 fmt.Println("x: ",float_x," ++: ",xx)
	 fmt.Println(xx==float_x)


	 if xxx:=float_x+1;xxx==float_x{
		 fmt.Println("equals")
	 }


 }