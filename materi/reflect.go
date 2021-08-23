package materi

import (
	"fmt"
	"reflect"
)

func GetReflectKeyValue(){
	var number int
	number = 23
	var character = "Sendi"

	var reflectValue  = reflect.ValueOf(&number) //number to pointer
	var reflectValue2 = reflect.ValueOf(character)
	fmt.Println("Tipe variabel nya :", reflectValue.Type())
	fmt.Println("Tipe variabel nya2 :", reflectValue2.Type())

	fmt.Println("====== :", reflectValue.Kind())
	fmt.Println("====== :", reflectValue2.Kind())
	fmt.Println("Nilai variabel nya :", reflectValue.Pointer())
	fmt.Println("Nilai variabel2 nya :", reflectValue2.String())
}