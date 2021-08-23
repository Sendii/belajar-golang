package materi

import (
	"math"
	"fmt"
)

type rumus interface{
	Luas() float64
	Keliling() float64
}

type Hitung interface{
	rumus //embeded interface rumus
	Total() float64
}

type Lingkaran struct {
	Diameter float64
}

type Persegi struct {
	Sisi float64
}

func (l Lingkaran) JariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.JariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

func (l Lingkaran) Total() float64 {
	return l.Keliling() + l.Luas()
}

func (p Persegi) Luas() float64{
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) Keliling() float64{
	return p.Sisi * 4
}

func (p Persegi) Total() float64 {
	return p.Keliling() + p.Luas()
}

var intPerson = []map[string]interface{}{
	{
		"nama": "Sendi Dian",
		"umur": 20,
	},
	{
		"nama": "Rizky",
		"umur": 21,
	},
	{
		"nama": "Amar",
		"umur": 22,
	},
}

func GetInterfacePerson(){
	for _, p := range intPerson{
		fmt.Println("Nama anda adalah :", p["nama"], "dan umur anda :", p["umur"])
	}
}