package helper
import "time"

func HariIndo()string{
	now := int(time.Now().Weekday())
	var day string
	switch now {
	case 1:
		day = "Senin"
	case 2:
		day = "Selasa"
	case 3:
		day = "Rabu"
	case 4:
		day = "Kamis"
	case 5:
		day = "Jum'at"
	case 6:
		day = "Sabtu"
	case 7:
		day = "Minggu"
	}
	return day
}

func BulanIndo()string{
	now := int(time.Now().Month())
	var month string
	switch now {
	case 1:
		month = "Januari"
	case 2:
		month = "Februari"
	case 3:
		month = "Maret"
	case 4:
		month = "April"
	case 5:
		month = "Mei"
	case 6:
		month = "Juni"
	case 7:
		month = "Juli"
	case 8:
		month = "Agustus"
	case 9:
		month = "September"
	case 10:
		month = "Oktober"
	case 11:
		month = "November"
	case 12:
		month = "Desember"	
	}
	return month
}

func TglIndo()(string, string, int){
	now := int(time.Now().Year())
	return HariIndo(), BulanIndo(), now
}