package library

import "fmt"

// kalo public function huruf depan nama function nya harus kapital
func SayHello(nama string){
	fmt.Println("Hello ini fungsi publik")
	introduce(nama)
}

// kalo public function huruf depan nama function nya harus kecil
func introduce(nama string){
	fmt.Println("nama saya adalah :", nama)
}