package materi
import (
	"fmt"
	"time"
	"math/rand"
)

func RandomNumber(){
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1", rand.Int())
	fmt.Println("random ke-2", rand.Int())
	fmt.Println("random ke-3", rand.Int())
}