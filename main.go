package main

import "os"

type Student struct {
	Id        int64
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func cariBanyakMurid(students []Student, inputs []string) {
	for _, student := range students {
		for _, input := range inputs {
			if input == student.Nama {
				println("ID : ", student.Id)
				println("Nama : ", student.Nama)
				println("Alamat : ", student.Alamat)
				println("Pekerjaan : ", student.Pekerjaan)
				println("Alasan : ", student.Alasan)
			}
		}
	}
}

func main() {
	var args = os.Args
	students := []Student{
		{
			Id:        0,
			Nama:      "Ariq",
			Alamat:    "Jakarta",
			Pekerjaan: "IT",
			Alasan:    "Kerja",
		},
		{
			Id:        1,
			Nama:      "Jagabaya",
			Alamat:    "Jakarta",
			Pekerjaan: "IT",
			Alasan:    "Kerja",
		},
	}
	cariBanyakMurid(students, args)
}
