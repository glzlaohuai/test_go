package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)


func main(){
	// scan_1()
	// scan_2()
	scan_3()
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
			fmt.Fprintf(os.Stderr,"open file failed: %v\n",err)
			continue
		}else{
			scanLines(file,xmap)
			file.Close()
		}
	}

	fmt.Println("map is: ",xmap)
}

func scan_3(){
	countmap:=make(map[string]int)
	files:=os.Args[1:]

	for _,fileName:=range files{
		data,err:=ioutil.ReadFile(fileName)

		if err==nil {
			fmt.Println("file readed, result is: ",string(data))
			var stringContent string =string(data)

			for _,line:=range strings.Split(stringContent,"\n"){
				countmap[line]++
			}
			
		}else{
			fmt.Fprintf(os.Stderr,"read file failed, %v\n",err)
			continue
		}

	}


	fmt.Println("result map is:", countmap)


}







func scanLines(f *os.File,counts map[string]int){
	fmt.Println("scan file: ",f.Name())
	reader:=bufio.NewScanner(f)
	for reader.Scan(){
		fmt.Println("readed: ",reader.Text())
		counts[reader.Text()]++
	}
}



