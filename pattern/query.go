package pattern
import (
	"belajar/db"
	"fmt"
)

func SqlInsert(a string, b string, c int, d string)bool{
	db, err := db.ConenctDB()
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	_, err = db.Exec("insert into siswa values(?, ?, ?, ?)", a, b, c, d)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}