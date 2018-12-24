package main

import(
    	"fmt"
    	"net/http"
    	"html"
	"os"
	"os/exec"
)

func main(){
	cmd := exec.Command("sh", "-c", "mkdir -p /log/test")
	cmd.Run()
    	f, err := os.OpenFile("/log/test/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		n, err := f.WriteString(r.URL.Path + "\n")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Hello", r.URL.Path)
	    	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    	})
   	http.ListenAndServe(":8080", nil)
}
