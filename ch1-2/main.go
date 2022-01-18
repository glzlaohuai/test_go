package main

import(
	"fmt"
	"os"
	"strings"
)


func main(){
	fmt.Println("the running program is: ",os.Args[0])

	// echo_2()
	// echo3()
	echo4()
}


func echo_1(){
	var s,seg string ="",""


	for i:=1;i<len(os.Args);i++{
		s+=seg+os.Args[i]
		seg =" "
	}

	fmt.Println("echo:",s)
}

func echo_2(){

	s,seg:="",""

	for _,v:= range os.Args[1:]{
		s+=seg+v
		seg = " "
	}


	fmt.Println("what you input is: ",s)

}


func echo3(){
	fmt.Println("what you inpout is:",strings.Join(os.Args[1:]," "))
}

func echo4(){
	for i,v:=range os.Args[1:]{
		fmt.Println("index is: ",i," value is: ",v)
	}
}
	