package main

import "fmt"
type Currency int

const(
	US Currency =iota
	EU
	RMB
)


func main(){

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





}
