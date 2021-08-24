package materi
import (
	"fmt"
	// "time"
	help "belajar/helper"
)

func TglSekarang(){
	// time1 := time.Now()
	// fmt.Println("hari:", time1.Weekday().String(), "bulan:", time1.Month(), "tahun", time1.Year())
	hari, bulan, tahun := help.TglIndo()
	fmt.Println("jadii..", hari, bulan, tahun)
}