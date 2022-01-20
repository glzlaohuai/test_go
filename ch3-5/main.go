package main

import(
	"fmt"
	"unicode/utf8"
	"strings"
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

	a:="/abc/x/x/x/x.text"
	b:="empty"
	c:="abc/x.go.test"
	d:="/a/b/c/d/"

	fmt.Println(basename(a))
	fmt.Println(basename(b))
	fmt.Println(basename(c))
	fmt.Println(basename(d))


	fmt.Println(basename2(a))
	fmt.Println(commaIt("123234546"))
}



const char_dot rune = '.'
const char_slash rune = '/'

func basename(s string)string{

	result:=s
	runeArray:=[]rune(result)


	//remove first encountered slash char from rever order
	for i:=len(result)-1;i>=0;i--{
		if runeArray[i] == char_slash {
			result = s[i+1:]
			break
		}
	}

	// fmt.Println("result after remove slash char: ",result)

	runeArray = []rune(result)

	//remove first encountered dot from reverse order
	for i:=len(runeArray)-1;i>0;i--{
		if runeArray[i] == char_dot{
			//slice
			result = result[:i]
			break
		}

	}

	return result
}



func basename2(s string) string{

	result:=s
	// runeArray:=[]rune(result)


	lastIndexOfSlash:=strings.LastIndex(result,"/")
	fmt.Println("last index of slash: ",lastIndexOfSlash)

	result = s[lastIndexOfSlash+1:]

	lastIndexOfDot:=strings.LastIndex(result,".")

	fmt.Println("last index of dot: ",lastIndexOfDot)

	result = result[:lastIndexOfDot]



	return result

}

func commaIt(s string) string{
	if len(s)<=3 {
		return s
	}

	return s[:3]+","+commaIt(s[3:])
}


