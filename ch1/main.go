package main

import(
	"fmt"
	"os"
)


func main(){

	var s,seg string ="",""

	fmt.Println("the running program is: ",os.Args[0])

	for i:=1;i<len(os.Args);i++{
		s+=seg+os.Args[i]
		seg =" "
	}

	fmt.Println("echo:",s)
}
