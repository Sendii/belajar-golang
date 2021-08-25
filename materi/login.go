package materi

type Login struct {
	Username string
	Password string
}

var credential = Login{"sendi", "8888"}

func AttemptLogin(u string, p string) (string, bool){
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
