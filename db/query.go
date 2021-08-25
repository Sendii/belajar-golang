package db

import (
	"fmt"
)

func SqlSelect(){ //select menggunakan .query
	db, err := ConenctDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var umur = 21
	rows, err := db.Query("select id, nama, umur, jurusan from siswa where umur = ?", umur)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var res []SiswaSql
	for rows.Next(){
		var siswa = SiswaSql{}
		var err = rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Umur, &siswa.Jurusan)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		res = append(res, siswa)
	}
	if err = rows.Err();err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, siswa := range res {
		fmt.Println(siswa.Nama)
	}
}

func SqlDetail(){ //select menggunakan .queryRow
	db, err := ConenctDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var res = SiswaSql{}
	var id  = "E001"
	err = db.
			QueryRow("select nama, umur, jurusan from siswa where id = ?", id).
			Scan(&res.Nama, &res.Umur, &res.Jurusan) // yg di select, kudu dipanggil semua biar ga error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Nama : %s\nUmur : %d\n", res.Nama, res.Umur)
}