package library

import "fmt"

// kalo public function huruf depan nama function nya harus kapital
func SayHello(){
	fmt.Println("Hello ini fungsi publik")
}

func introduce(nama string){
	fmt.Println("nama saya adalah :", nama)
}