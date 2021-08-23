package materi

import (
	"fmt"
	"runtime"
)

func print(number int, message string){
	for i := 0; i <= number; i++{
		fmt.Println(i, message)
	}
}

func GoRoutPrint(){
	runtime.GOMAXPROCS(1)
	go print(5, "Haii")
	print(5, "Apa kabarr")

	var input string
	fmt.Scanln(&input) //proses berhenti ketika menekan enter
}