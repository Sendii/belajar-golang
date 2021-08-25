package materi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Siswa struct{
	NamaLengkap string `json:"Nama"` //menyesuaikan dengan isi json string nya
	Umur 		int
}

type ApiSiswa struct{
	ID	   string
	Nama   string
	Umur   int
}

func CallJson(){
	var jsonString = `{
		"Nama"	: "Sendi Dian Hadiwijaya",
		"Umur"	: 20
	}`
	var jsonData   = []byte(jsonString)

	var data Siswa
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nama lengkap anda adalah :", data.NamaLengkap)
	fmt.Println("Umur anda adalah :", data.Umur)
}

func CallJsonArray(){
	var jsonString = `[
	{
		"Nama"	: "Sendi Dian Hadiwijaya",
		"Umur"	: 20
	},
	{
		"Nama"	: "Muhammad Rizky",
		"Umur"	: 22
	}
	]`

	var data []Siswa
	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("Siswa", i + 1 ,"namanya :", data[i].NamaLengkap)
		fmt.Println("Umurnya :", data[i].Umur)
		fmt.Println("===========================================")
	}
}

var dataSiswa = []ApiSiswa{
	ApiSiswa{"A001", "sendi", 20},
	ApiSiswa{"A002", "rizky", 21},
	ApiSiswa{"A003", "dian", 22},
	ApiSiswa{"A004", "wijaya", 23},
}

type Login struct {
	Username string
	Password string
}

var credential = Login{"sendi", "8888"}
func attemptLogin(u string, p string) (string, bool){
	if u == "" && p == "" {
		return "Username dan Password tidak boleh kosong", false
	}else if u == "" {
		return "Username tidak boleh kosong", false
	}else if p == "" {
		return "Password tidak boleh kosong", false
	}else if u == credential.Username && p == credential.Password {
		return "", true
	}else{
		return "Ada kesalahan", false
	}
}

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user, password = r.FormValue("username"), r.FormValue("password")

	if r.Method == "POST"{
		var res, err = json.Marshal(dataSiswa)
		var msg, errmsg = attemptLogin(user, password)
		if !errmsg {
			http.Error(w, msg, http.StatusUnauthorized)
		}else{
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(res)
			return
		}
	}
}

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST"{
		var id, user, password = r.FormValue("id"), r.FormValue("username"), r.FormValue("password") //get id from form-data postman
		var res []byte
		var err error

		var msg, errmsg = attemptLogin(user, password)
		if !errmsg {
			http.Error(w, msg, http.StatusUnauthorized)
		}else{
			for _, siswa := range dataSiswa {
				if siswa.ID == id {
					res, err = json.Marshal(siswa)

					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(res)
					return
				}else{
					http.Error(w, "Siswa tidak ditemukan", http.StatusBadRequest)
					return
				}
			}
		}		
	}
}

func CallApiSiswa(){	
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)

	fmt.Println("starting web server at http://localhost:1020")
		http.ListenAndServe(":1020", nil)
	}