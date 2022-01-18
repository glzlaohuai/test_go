//variables declared in function scope is still available outof this scope
package main

import (
	"fmt"
	"flag"
)


var n = flag.Bool("n",false,"omit last line")
var s= flag.String("s","default","separator")
var m = flag.Int("m",0,"max value")



func main(){
	// fmt.Printf("test result: %t",f() == f())

	// flag.Parse()

	// fmt.Println("non flagged arguments: ",flag.Args())
	// fmt.Println("n value is: ", *n, ", s value is: ", *s, " max value is: ",*m)

	// defineTest()

	xmap:=map[string]int{
		"a":1,
		"b":2,
	}

	fmt.Println("map is:",xmap)

	fmt.Println("empty : ",xmap["d"])
	v,r:=xmap["d"]
	fmt.Println("value is: ",v ,"r is :",r )


	v,ok:=xmap["a"]
	fmt.Println("value is: ",v ,"r is :",r )



}



func f()*int{
	x:=1
	return &x
}


func defineTest(){
	x:=new(int)
	fmt.Println(*x)
	*x=100
	fmt.Println(*x)
}




