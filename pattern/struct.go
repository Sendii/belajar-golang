package pattern

type PSiswa struct{
	Id string
	Nama string
	Umur int
	Jurusan string
}

func(s PSiswa) Create() bool {
	return SqlInsert(s.Id, s.Nama, s.Umur , s.Jurusan)
}