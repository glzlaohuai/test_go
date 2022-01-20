package main

import (
	"fmt"
	"crypto/sha256"
	"flag"
)
type Currency int

const(
	US Currency =iota
	EU
	RMB
)

var use256 *bool = flag.Bool("-b",false,"set if use sha256")
var inputString *string = flag.String("-s","no input specific","string to hash")

func main(){
	flag.Parse()

	fmt.Println("main invoked")

	currencies:=[...]string{US:"$",EU:"&",RMB:"￥"}
	fmt.Println(currencies)

	testArray:=[...]int{1,2,3}


	fmt.Println(testArray)
	testArray2:=[]int{99:-1}
	fmt.Println(testArray2)

	array1:=[2]int{1,2}
	array2:=[2]int{2,1}
	array3:=[2]int{1,2}
	array4:=[...]int{1,2}
	array5:=[...]int{1,2,3}
	fmt.Println(array5)

	fmt.Println("equals: ",array1 == array2)

	fmt.Println(array1==array3)
	fmt.Println(array1 == array4)
	// fmt.Println(array1 == array5) type mismatch



	bytes:=[]byte{10:100}
	fmt.Println("bytes is: ",bytes)
	fmt.Printf("add is :%p\n",&bytes)
	zeroit(bytes)
	// zeroit2(&bytes)
	fmt.Println("after zeroit: ",bytes)


	bytes2:=[32]byte{1,2,3}
	zeroit4(&bytes2)

	fmt.Println("after zero it: ",bytes2)



	c1:=sha256.Sum256([]byte{'a'})
	c2:=sha256.Sum256([]byte{'A'})

	fmt.Printf("c1: %x\n",c1)
	fmt.Printf("c2: %x\n",c2)

	fmt.Println("equals: ",c1==c2)

	fmt.Println("use 256? ",*use256)
	fmt.Println("string to hash: ",*inputString)

	if *use256 {
		fmt.Println(doSum256(*inputString))
	}else{
		fmt.Println(doSum224(*inputString))
	}
}




func doSum256(input string)[32]byte{
	return sha256.Sum256([]byte(input))
}


func doSum224(input string)[28]byte{
	return sha256.Sum224([]byte(input))
}







//send as ref ?
func zeroit(bytes []byte){
	fmt.Printf("zeroit addr is: %p\n",&bytes)
	for i,_:= range bytes{
		bytes[i]=0
	}
}

func zeroit2(bytes *[]byte){
	for i,_:=range *bytes{
		(*bytes)[i] = 0
	}
}

func zeroit3(bytes [32]byte){
	for i,_:=range bytes{
		bytes[i]=0
	}
}

func zeroit4(bytes *[32]byte){
	for i,_:=range(*bytes){
		(*bytes)[i] = 0
	}
}




