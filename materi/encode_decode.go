package materi
import (
	"fmt"
	"encoding/base64"
)

func EncodeText(text string){
	data := text
	dataEncoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded nya :", dataEncoded)

	dataDecodedByte, _ := base64.StdEncoding.DecodeString(dataEncoded)
	dataDecoded 	   := string(dataDecodedByte)
	fmt.Println("decoded nya :", dataDecoded)
}