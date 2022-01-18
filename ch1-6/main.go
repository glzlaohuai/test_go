// goroutine
package main

import(
	"fmt"
	"time"
	"io/ioutil"
	"io"
	"net/http"
	"os"
)



func main(){

	urls:=os.Args[1:]
	fmt.Println("fetch urls: ",urls)

	start:=time.Now()
	fetchAll(urls)
	fmt.Printf("total: %.2f\n",time.Since(start).Seconds())
}



func fetchAll(urls []string){
	ch:=make(chan string)

	for _,url:=range urls{
		go fetch(url,ch)
	}


	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

}


func fetch(url string, ch chan<-string){
	start:=time.Now()
	fmt.Println("fetch url: ",url)

	resp,err:=http.Get(url)
	if err==nil {
		written,err:=io.Copy(ioutil.Discard,resp.Body)
		if err!=nil {
			ch<-fmt.Sprint(err)
		}else{
			ch<-fmt.Sprintf("%.2fs\t%d\t%s",time.Since(start).Seconds(),written,url)
		}
		
	}else{
		ch<-fmt.Sprint(err)
	}
}