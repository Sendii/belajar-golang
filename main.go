package main

import (
	"fmt"
	// mat "belajar/materi"
	// "belajar/db"
	patt "belajar/pattern"
)


func main(){
	// ============ STRING,IF,ARRAY,MAP,FOR,POINTER ==================
	// mat.CallString()
	// mat.CallIf(80)
	// mat.CallFor(10)
	// mat.CallArray()
	// mat.CallArrayMultidimensional()
	// mat.CallMapMonth()
	// mat.PrintMessage([]string{"John", "Doe"}, "RPL")
	// mat.CallPointer()
	// mat.CallString()
	// =========================== END ==============================

	// =========================== STRUCT ===========================
	// mat.CallStudentStruct1()
	// mat.CallStudentStruct2()
	// mat.CallPersonStruct3()
	// mat.CallPersonStruct4()
	// mat.CallPersonStruct5()	
	// mat.CallStruct()
	// =========================== END ==============================


	// ==================== CALL FUNC INIT ==========================
	// fmt.Printf("Nama  : %s\n", mat.Student.Nama)
	// fmt.Printf("Kelas : %s\n", mat.Student.Kelas)
	// =========================== END ==============================
	// call with go run main.go partial.go
	// CallFuncRoot("Sendi")

	// ========================= INTERFACE ==========================
	// var bangunDatar mat.Hitung
	// bangunDatar = mat.Persegi{10.0}
	// fmt.Println("=== Persegi :")
	// fmt.Println("Luas :", bangunDatar.Luas())
	// fmt.Println("Keliling :", bangunDatar.Keliling())
	// fmt.Println("Total :", bangunDatar.Total())
	// fmt.Println("=================== :")

	// bangunDatar = mat.Lingkaran{14.0}
	// fmt.Println("=== Lingkaran :")
	// fmt.Println("Jari-jari :", bangunDatar.(mat.Lingkaran).JariJari())
	// fmt.Println("Luas :", bangunDatar.Luas())
	// fmt.Println("Keliling :", bangunDatar.Keliling())	
	// fmt.Println("Total :", bangunDatar.Total())


	// INTERFACE KOSONG
	// mat.GetInterfacePerson()
	// =========================== END ==============================

	// ========================= REFLECT ============================
	// mat.GetReflectKeyValue()
	// =========================== END ==============================

	// ====================== GO ROUTINE ============================
	// mat.GoRoutPrint()
	// mat.GoRoutChannel()
	// mat.GoRoutChannelParam()
	// =========================== END ==============================

	// ERROR
	// mat.CallErrorAndRecover()

	//random
	// mat.RandomNumber()
	// mat.CustRandom("string", 5)

	//TANGGAL
	// mat.TglSekarang()

	// STRING
	// mat.ReplaceText("Mgngga")
	// mat.EncodeText("John Doe")

	// FILE
	// mat.CreateFileTxt()
	// mat.DeleteFile()

	// web
	// mat.ActivWebsite()
	// mat.UrlParse()

	// JSON
	// mat.CallJson()
	// mat.CallJsonArray()
	// mat.CallApiSiswa()

	// DATABASE
	// db.SqlSelect()
	// db.SqlAction("delete")
	// db.SqlAction("insert")
	// db.SqlDetail("Z001")
	// db.SqlAction("update")
	// db.SqlDetail("Z001")
	// db.SqlPrepare()

	// PATTERN
	var query patt.PQuery

	query = patt.PSiswa{"A001", "Rizki", 19, "Inggris"}
	fmt.Println("show : ", query.Show())

	if query.Show() {
		query = patt.PSiswa{"A001", "Rizki", 19, "Inggris"}
		fmt.Println("delete : ", query.Delete())
	}	

	query = patt.PSiswa{"A001", "Rizki", 19, "Inggris"}
	fmt.Println("create : ", query.Create())

	if query.Show() {
		query = patt.PSiswa{"A001", "Rizki Update", 20, "Inggris Update"}
		fmt.Println("update : ", query.Update())
	}	

}