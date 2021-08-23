package materi

import (
	"fmt"
	"runtime"
)

func init(){
	runtime.GOMAXPROCS(2)
}

func print(number int, message string){
	for i := 0; i <= number; i++{
		fmt.Println(i, message)
	}
}

func GoRoutPrint(){	
	go print(5, "Haii")
	print(5, "Apa kabarr")

	var input string
	fmt.Scanln(&input) //proses berhenti ketika menekan enter
}

var glob_messages = make(chan string) //membuat channgel dengan menuliskan keyword make dengan isi keyword chan, dan diikuti tipe data channel nya.
var sayHelloTo    = func (nama string){
	var data 	  = fmt.Sprintf("Hallo %s", nama)
	glob_messages <- data
}

func GoRoutChannel(){
	go sayHelloTo("Sendi")
	go sayHelloTo("Rizky")
	go sayHelloTo("Amar")

	var message1 = <- glob_messages
	fmt.Println(message1)

	var message2 = <- glob_messages
	fmt.Println(message2)

	var message3 = <- glob_messages
	fmt.Println(message3)
}

var sayHelloToParam = func(nama string){
	var data = fmt.Sprintf("Halooo %s", nama)
	glob_messages <- data
}

func printMessageChannel(what chan string){
	fmt.Println(<-what)
}

func GoRoutChannelParam(){
	var total = 0
	for _, nama := range []string{"sendi", "dian", "hadiwijaya"}{
		total += 1
		go sayHelloToParam(nama)
	}

	for i := 0; i < total; i++ {
		printMessageChannel(glob_messages)
	}
}