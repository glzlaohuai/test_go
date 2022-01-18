package main


import(
	"fmt"
	"net/http"
	"log"
	"sync"
)

var count int
var lock sync.Mutex



func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	http.HandleFunc("/info",func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w," %s %s %s ", r.Method, r.URL, r.Proto)
		for k,v:=range r.Header{
			fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
		}
		fmt.Fprintf(w,"HOST = %q", r.Host)
		fmt.Fprintf(w,"Remote Addr = %q",r.RemoteAddr)

		if err:=r.ParseForm();err!=nil{
			log.Print(err)
		}

		for k,v := range r.Form{
			fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8989",nil))
}


func handler(writer http.ResponseWriter, r *http.Request){
	fmt.Fprintf(writer,"retrieved path is: %s",r.URL.Path)
	lock.Lock()
	count++
	lock.Unlock()
}

func counter(writer http.ResponseWriter, r *http.Request){
	lock.Lock()
	fmt.Fprintf(writer,"path is: %s, visited count: %d",r.URL.Path, count)
	lock.Unlock()
}





