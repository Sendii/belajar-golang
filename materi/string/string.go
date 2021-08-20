package string

import "fmt"

func CallString(){
	var firstName string = "John"
	lastName := "Doe" //kalo set variable pertama bisa pake :=
	lastName = "Doe baru"

	a, b, c := "A", "B", 90
	fmt.Printf("Halooo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")
	fmt.Println("haloo", a, b, c)
}