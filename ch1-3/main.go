package main

import (
	"fmt"
	"os"
	"bufio"
)


func main(){
	scan_1()
}



func scan_1(){
	infomap:=make(map[string]int)

	scanner:=bufio.NewScanner(os.Stdin)

	for scanner.Scan() && len(infomap)<=10{
		infomap[scanner.Text()]++
	}

	for k,v :=range(infomap){
		fmt.Printf("%d\t%s\n",v,k)
	}

}

