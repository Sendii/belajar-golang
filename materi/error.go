package materi

import (
	"errors"
	"strings"
)

func Validate(input string)(bool, error){
	if strings.TrimSpace(input) == ""{
		return false, errors.New("Tidak boleh kosong")
	}
	return true, nil
}