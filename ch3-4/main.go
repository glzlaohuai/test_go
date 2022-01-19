package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	fmt.Println(math.NaN())

	s := "hello world"
	fmt.Println(s)
	fmt.Println(len(s))

	s = "你好"
	fmt.Println(s)
	fmt.Println(len(s))

	fmt.Println(reflect.TypeOf(s[0:3]).Kind())

}

func itob(i int) bool {
	return i != 0
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
