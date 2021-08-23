package materi

import (
	"fmt"
	"strings"
)

type func_student struct {
	nama string
	kelas string
}

var Student = struct {
	Nama string
	Kelas string
}{}

// function yg pertama kali dipanggil ketika diload
func init(){
	Student.Nama = "John Wick"
    Student.Kelas = "RPL"

	fmt.Println("COmpleted exported")
}

func (s func_student) methodSayHello(){
	fmt.Println("Haloo :", s.nama)
}

func (s func_student) methodetGetNameAt(i int) string{
	return strings.Split(s.nama, " ")[i-1]
}

func (s func_student) methodChangeName1(nama string){
	fmt.Println("---> on Changename1, name changed to", nama)
	s.nama = nama
}

func (s *func_student) methodChangeName2(nama string){
	fmt.Println("---> on Changename1, name changed to", nama)
	s.nama = nama
}

func CallStruct(){
	// start struct
	var s1 = func_student{"Sendi Awal", "RPL"}
	fmt.Println("s1 before", s1.nama)
	fmt.Println("=================================================")

	s1.methodChangeName1("Sendi 1")
	fmt.Println("s1 after changename1", s1.nama)
	fmt.Println("=================================================")

	s1.methodChangeName2("Sendi 2")
	fmt.Println("s1 after changename2", s1.nama)
	// end struct
}