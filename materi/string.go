package materi

import (
	"strings"
	"fmt"
)

func ReplaceText(text string){
	new1 := strings.Replace(text, "g", "h", -1)
	fmt.Println(new1)
}