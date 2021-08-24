package materi
import (
	"fmt"
	"net/http"
	"net/url"
)

func index(a http.ResponseWriter, r *http.Request){
	fmt.Fprintln(a, "apakabar")
}

func ActivWebsite(){
	http.HandleFunc("/ea", func(a http.ResponseWriter, r *http.Request){
		fmt.Fprintln(a, "halo")
	})
	http.HandleFunc("/index", index)
	fmt.Println("starting web server at http://localhost:1020")
	http.ListenAndServe(":1020", nil)
}

func UrlParse(){
	var urlString = "http://google.com:9090/eaea/cari?nama=sendi dian&umur=2727272727272727"
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url : %s\n", urlString)
	fmt.Printf("protocol : %s\n", u.Host)
	fmt.Printf("Path : %s\n", u.Path)

	fmt.Printf("nama query parameter : %s\n", u.Query()["nama"][0])
	fmt.Printf("umur query parameter : %s", u.Query()["umur"][0])
}