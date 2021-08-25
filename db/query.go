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

func SqlPrepare() { //reuse sql
	db, err := ConenctDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select nama, umur, jurusan from siswa where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var r1 = SiswaSql{}
	stmt.QueryRow("E001").Scan(&r1.Nama, &r1.Umur, &r1.Jurusan)
	fmt.Printf("Nama : %s\nUmur: %d\nJurusan : %s\n", r1.Nama, r1.Umur, r1.Jurusan)

	var r2 = SiswaSql{}
	stmt.QueryRow("B001").Scan(&r2.Nama, &r2.Umur, &r2.Jurusan)
	fmt.Printf("Nama : %s\nUmur: %d\nJurusan : %s\n", r2.Nama, r2.Umur, r2.Jurusan)

	var r3 = SiswaSql{}
	stmt.QueryRow("B002").Scan(&r3.Nama, &r3.Umur, &r3.Jurusan)
	fmt.Printf("Nama : %s\nUmur: %d\nJurusan : %s\n", r3.Nama, r3.Umur, r3.Jurusan)
}