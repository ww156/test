package main

import(
    	"fmt"
    	"net/http"
    	"html"
	"os"
)

func main(){
    	f, err := os.OpenFile("/log/test/log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		f.WriteString(r.URL.Path)
		fmt.Println("Hello", r.URL.Path)
	    	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    	})
   	http.ListenAndServe(":8080", nil)
}
