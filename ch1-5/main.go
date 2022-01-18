//fetch url using http lib
package main


import(

	"fmt"
	"os"
	"net/http"
	"io"
	"strings"
)


func main(){
	urls:=os.Args[1:]
	fmt.Println("fetching url: ",urls)

	fetch(urls)
}


func fetch(urls []string){
	for _,url:=range urls{

		if  !strings.HasPrefix(url,"http"){
			url="http://"+url
			
		}
		resp,err:=http.Get(url)
		if err!=nil {
			fmt.Fprintf(os.Stderr,"get url failed with error %v",err)
			os.Exit(1)
		}else{
			// fmt.Println("get url succeeded, ",url)
			// fmt.Println("status: ",resp.Status)
		// _,err:=ioutil.ReadAll(resp.Body)

		written,err:=	io.Copy(os.Stdout,resp.Body)
		if err==nil {
			fmt.Println("written: ",written)
		}else{
			fmt.Fprintf(os.Stderr,"copy failed, %v\n",err)
		}


			resp.Body.Close()
			if err==nil {
				fmt.Println("content of url :",url)
				fmt.Println()
			}else{
				fmt.Fprintf(os.Stderr,"read content from response failed, error: %v",err)
			}
		}

	}

}
