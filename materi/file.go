package materi
import (
	"fmt"
	"os"
)

func isError(err error) bool{
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}

var path string = "/Users/asus/Documents/Project/Golang/src/belajar/satu.txt"

func CreateFileTxt(){
	// cek file, ada apa ngga
	var _, err = os.Stat(path)

	if os.IsNotExist(err){
		var file, err = os.Create(path)
		fmt.Println("Berhasil dibuat", path)		
		if isError(err) {return}
		defer file.Close()
	}else{
		fmt.Println("File sudah ada !")
	}
	writeFile()
}

func writeFile(){
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {return}
	defer file.Close()

	_,err = file.WriteString("Haloo\n")
	if isError(err) {return}
	_,err = file.WriteString("belajar golang yuu\n")
	if isError(err) {return}

	err   = file.Sync()
	if isError(err) {return}
	fmt.Println("Berhasil diisi ")
}

func DeleteFile(){
	var err = os.Remove(path)
	if isError(err) {return}
	fmt.Println("Berhasil dihapus ")	
}