package main

import ("fmt"
	"strings"
)

func callString(){
	var firstName string = "John"
	lastName := "Doe" //kalo set variable pertama bisa pake :=
	lastName = "Doe baru"

	a, b, c := "A", "B", 90
	fmt.Printf("Halooo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")
	fmt.Println("haloo", a, b, c)
}

func callIf(numb int8){
	point := numb
	if point == 10 {
		fmt.Println("Nilai anda adalah 10")
	}else if point == 50 {
		fmt.Println("Nilai anda adalah 50")
	}else if point == 80 {
		fmt.Println("Nilai anda adalah 80")
	}else if point == 100 {
		fmt.Println("Nilai anda adalah 100")
	}else{
		fmt.Println("Nilai anda adalah ?")
	}
}

func callFor(numb_max int){
	for i := 1; i <= numb_max; i++ {
		fmt.Println("Angka", i);
	}
}

func callArray(){
	//set array dengan gaya horizontal
	var names [4]string
	names[0] = "Sendi"
	names[1] = "Dian"
	names[2] = "Hadi"
	names[3] = "wijaya"

	// set array dengan gaya vertikal tanpa jumlah 
	var fruits = [...]string{"apple", "grape", "banana", "watermelon", "ea"}

	fmt.Println(names[0], names[1], names[2], names[3], len(names))

	for i := 0; i < len(fruits); i++ {
		fmt.Print(fruits[i], " ")
	}
}

func callArrayMultidimensional(){
	var numbers1 = [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}

func callMapMonth(){
	var month = map[int]string{1: "januari", 2: "februari", 3: "maret", 4: "april"}
	delete(month, 1) //delete map by key

	for key, val := range month {
		fmt.Println("bulan ke :", key, val)
	}
}

func callPointer(){
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	nama := "Sendi Dian Hadiwijaya"
	pointerNama := &nama
	fmt.Println("Pointer nya :", pointerNama)
	fmt.Println("Nilai Pointer nya :", *pointerNama)
}

func printMessage(arr []string, kelas string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(nameString, kelas)
}

// struct
type student struct {
	nama string
	kelas int
}

func callStudentStruct1(){
	//penerapan struct
	var s1 student
	s1.nama  = "Sendi Hadi"
	s1.kelas = 10

	fmt.Println("Nama anda adalah", s1.nama)
	fmt.Println("Kelas anda adalah", s1.kelas)
}

func callStudentStruct2(){
	//cara-cara inisialisasi struct
	var s1   = student{}
	s1.nama  = "Sendi"
	s1.kelas = 10

	var s2 = student{"Sendi 2", 10}
	var s3 = student{nama: "Sendi 3"}

	fmt.Println("Student 1", s1.nama)
	fmt.Println("Student 2", s2.nama)
	fmt.Println("Student 3", s3.nama)
}

func main(){	
	// callString()
	// callIf(80)
	// callFor(10)
	// callArray()
	// callArrayMultidimensional()
	// callMapMonth()
    // printMessage([]string{"John", "Doe"}, "RPL")
	// callPointer()
	// callStudentStruct1()
	callStudentStruct2()
}