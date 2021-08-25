package pattern

type PSiswa struct{
	Id string
	Nama string
	Umur int
	Jurusan string
}

func(s PSiswa) Show() bool {
	return sqlShow(s.Id)
}

func(s PSiswa) Create() bool {
	return SqlInsert(s.Id, s.Nama, s.Umur , s.Jurusan)
}

func(s PSiswa) Update() bool {
	return sqlUpdate(s.Id, s.Nama, s.Umur , s.Jurusan)
}

func(s PSiswa) Delete() bool {
	return sqlDelete(s.Id)
}