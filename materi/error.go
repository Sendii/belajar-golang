package materi

import (
	"errors"
	"strings"
	"fmt"
)

func catch(){
	if r:= recover(); r != nil {
		fmt.Println("Ada error,", r)
	}
}

func validate(nama_input string, input string)(bool, error){
	if strings.TrimSpace(input) == ""{
		return false, errors.New(nama_input + " Tidak boleh kosong")
	}
	return true, nil
}

func checkValidate(nama_input string, nama string){
	if valid, err := validate(nama_input, nama); valid{
		fmt.Println("Haloo", nama)
	}else{
		panic(err.Error())
	}
}
func CallErrorAndRecover(){
	defer catch()
	var nama string 
	var kelas string 
	fmt.Print("Masukkan nama anda : ")
	fmt.Scanln(&nama)
	checkValidate("nama", nama)

	fmt.Print("Masukkan kelas anda : ")
	fmt.Scanln(&kelas)
	checkValidate("kelas", kelas)
}