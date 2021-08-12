package main

import ("fmt"
	"strings"
)

func main(){
	// ==1
	// var firstName string = "John"
	// var lastName string
	// lastName = "Doe"
	// lastName := "Doe" //kalo set variable pertama bisa pake :=
	// lastName = "Doee"

	// a, b, c := "A", "B", 90
	// fmt.Printf("Halooo %s %s!\n", firstName, lastName)
	// fmt.Printf("halo %s %s!\n", firstName, lastName)
	// fmt.Println("halo", firstName, lastName + "!")

	// =========================================================
	// ==2
	// name := new(string)
	// fmt.Println(*name)

	// var positiveNumber uint8 = 89
	// negativeNumber := -1243423644

	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// value := 10 + (20 / 5) + (10 * 5)
	// fmt.Println("Nilai nya adalah", value)


	// point := 85
	// if point == 10 {
	// 	fmt.Println("Nilai anda adalah 10")
	// }else if point == 50 {
	// 	fmt.Println("Nilai anda adalah 50")
	// }else if point == 80 {
	// 	fmt.Println("Nilai anda adalah 80")
	// }else if point == 100 {
	// 	fmt.Println("Nilai anda adalah 100")
	// }else{
	// 	fmt.Println("Nilai anda adalah ?")
	// }


	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Angka", i);
	// }

	//set array dengan gaya horizontal
	// var names [4]string
	// names[0] = "Sendi"
	// names[1] = "Dian"
	// names[2] = "Hadi"
	// names[3] = "wijaya"

	//set array dengan gaya vertikal tanpa jumlah 
	// var fruits = [...]string{"apple", "grape", "banana", "watermelon", "ea"}

	// fmt.Println(names[0], names[1], names[2], names[3], len(names))

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Print(fruits[i], " ")
	// }


	//ARRAY MULTIDIMENSI
	// var numbers1 = [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)


	// MAP
	// var month = map[int]string{1: "januari", 2: "februari", 3: "maret", 4: "april"}
	// delete(month, 1) //delete map by key

	// for key, val := range month {
	// 	fmt.Println("bulan ke :", key, val)
	// }	


	// CALL FUNCTION
	// var names = []string{"John", "Doe"}
    // printMessage(names, "RPL")
    // var numb = []int{30, 2}
    // calculate("*", numb)


    // POINTER
    // var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)
	// fmt.Println("numberA (address) :", &numberA)
	// fmt.Println("numberB (value)   :", *numberB)
	// fmt.Println("numberB (address) :", numberB)

	// nama := "Sendi Dian Hadiwijaya"
	// pointerNama := &nama
	// fmt.Println("Pointer nya :", pointerNama)
	// fmt.Println("Nilai Pointer nya :", *pointerNama)
}

func printMessage(arr []string, kelas string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(nameString, kelas)
}

func calculate(operate string, number []int){
	if operate == "*"{
		total := number[0] * number[1]
		var totalPointer *int = &total
		fmt.Println(total)
		fmt.Println(*totalPointer)
	}else if operate == "/"{
		fmt.Println(number[0] / number[1])
	}else if operate == "+"{
		fmt.Println(number[0] + number[1])
	}else if operate == "-"{
		fmt.Println(number[0] - number[1])
	}else{
		fmt.Println("Tidak ditemukan")
	}
}