package main

import(
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	      fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    http.ListenAndServe(":8080", nil)
}
