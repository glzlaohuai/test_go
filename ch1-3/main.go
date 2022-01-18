package main

import (
	"fmt"
	"os"
	"bufio"
)


func main(){
	// scan_1()
	scan_2()
}



func scan_1(){
	infomap:=make(map[string]int)

	scanner:=bufio.NewScanner(os.Stdin)

	for scanner.Scan() && len(infomap)<=3{
		infomap[scanner.Text()]++
	}

	for k,v :=range(infomap){
		fmt.Printf("%d\t%s\n",v,k)

		fmt.Printf("type is: %T\n",v)
		fmt.Printf("literal value is: %v\n",v)
	}

}


func scan_2(){
	files:=os.Args[1:]
	fmt.Println("incoming files: ",files)
	xmap:=make(map[string]int)

	for _,f:=range files{
		file,err:=os.Open(f)
		if err!=nil {
			fmt.Println("open file failed, error is: ",err)
		}else{
			scanLines(file,xmap)
		}
	}

	fmt.Println("map is: ",xmap)
}



func scanLines(f *os.File,counts map[string]int){
	fmt.Println("scan file: ",f.Name())
	reader:=bufio.NewScanner(f)
	for reader.Scan(){
		fmt.Println("readed: ",reader.Text())
		counts[reader.Text()]++
	}
}



