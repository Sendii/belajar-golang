package materi
import (
	"fmt"
	"net/http"
)

func index(a http.ResponseWriter, r *http.Request){
	fmt.Println(a, "apakabar")
}

func ActivWebsite(){
	http.HandleFunc("/", func(a http.ResponseWriter, r *http.Request){
		fmt.Fprintln(a, "halo")
	})
	http.HandleFunc("/index", index)
	fmt.Println("starting web server at http://localhost:1020")
	http.ListenAndServe(":1020")
}