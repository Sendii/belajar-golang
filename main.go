package main

import ("fmt"
	"strings"
	lib "belajar/library"
	mat "belajar/materi"
)

type func_student struct {
	nama string
	kelas string
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

func main(){	
	// mat.CallString()
	// mat.CallIf(80)
	// mat.CallFor(10)
	// mat.CallArray()
	// mat.CallArrayMultidimensional()
	// mat.CallMapMonth()
    // mat.PrintMessage([]string{"John", "Doe"}, "RPL")
	// mat.CallPointer()
	// mat.CallStudentStruct1()
	// mat.CallStudentStruct2()
	// mat.CallPersonStruct3()
	// mat.CallPersonStruct4()
	// mat.CallPersonStruct5()

	// start struct
	// var s1 = func_student{"Sendi Awal", "RPL"}
	// fmt.Println("s1 before", s1.nama)
	// fmt.Println("=================================================")

	// s1.methodChangeName1("Sendi 1")
	// fmt.Println("s1 after changename1", s1.nama)
	// fmt.Println("=================================================")

	// s1.methodChangeName2("Sendi 2")
	// fmt.Println("s1 after changename2", s1.nama)
	// end struct

	lib.SayHello("Sendi Hadi")

	mat.CallString()
}