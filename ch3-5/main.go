package main

import(
	"fmt"
	"unicode/utf8"
)

func main(){
	fmt.Println("main invoked")

	s:="hello, 你好"

	// for i,c:= range s{
	// 	fmt.Printf("%d \t %c\n",i,c)
	// }

	for i:=0;i<len(s);{
		r,size:=utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d \t %c\n",size,r)

		i+=size
	}


	for i,r:=range s{
		fmt.Printf("char at %d, is %d\n",i,r)
	}


	count:=0
	for _,_=range s{
		count++
	}
	fmt.Println(" coumnt is: ",count)

	fmt.Println("count is: ",utf8.RuneCountInString(s))

	count=0
	for range s{
		count++
	}
	fmt.Println("count is: ",count)

	runearray:=[]rune(s)
	fmt.Println("rune array is:",runearray)
	for i,r:= range runearray{
		fmt.Printf("i is %d, r is: %#x\n",i,r)
	}

}