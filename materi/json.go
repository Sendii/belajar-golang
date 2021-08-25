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

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user, password = r.FormValue("username"), r.FormValue("password")

	if r.Method == "POST"{
		var res, err = json.Marshal(dataSiswa)
		var msg, errmsg = AttemptLogin(user, password)
		if !errmsg {
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST"{
		var id, user, password = r.FormValue("id"), r.FormValue("username"), r.FormValue("password") //get id from form-data postman
		var res []byte
		var err error

		var msg, errmsg = AttemptLogin(user, password)
		if !errmsg {
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}
		for _, siswa := range dataSiswa {
			if siswa.ID == id {
				res, err = json.Marshal(siswa)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(res)
				return
			}else if id == ""{
				http.Error(w, "Masukkan id siswa nya.", http.StatusBadRequest)
				return
			}
		}
		http.Error(w, "Siswa tidak ditemukan", http.StatusBadRequest)
		return
	}
}

func CallApiSiswa(){	
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)

	fmt.Println("starting web server at http://localhost:1020")
		http.ListenAndServe(":1020", nil)
	}