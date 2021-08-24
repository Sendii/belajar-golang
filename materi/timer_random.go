package materi
import (
	"fmt"
	"time"
	"math/rand"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomNumber(){
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1", rand.Int())
	fmt.Println("random ke-2", rand.Int())
	fmt.Println("random ke-3", rand.Int())
}

func CustRandom(tipe string, n int) {	
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if tipe == "number" {
		letterRunes = []rune("1234567890")
	}

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    fmt.Println("Hasil random", tipe, "nya adalah", string(b))
}