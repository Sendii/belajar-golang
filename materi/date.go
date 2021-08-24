package materi
import (
	"fmt"
	"time"
	help "belajar/helper"
)

func TglSekarang(){
	now := time.Now()
	fmt.Println(help.TglIndo(now))
}