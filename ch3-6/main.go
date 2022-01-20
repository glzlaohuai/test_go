package main

import(
	"fmt"
)


const(
	deadbeef = 0xdeadbeef
	etst = float32(deadbeef)
)


func main(){
	// fmt.Printf("original: %d\n",deadbeef)
	// fmt.Printf("float is: %f",etst)



	uint_test:=uint32(deadbeef)
	float32_test:=float32(uint_test)


	var y int64 = int64(etst)
	fmt.Println("result is: ",y)
	fmt.Println(y-deadbeef)


	fmt.Println(uint_test,float32_test)
	fmt.Printf("%f",float32_test)
	
}