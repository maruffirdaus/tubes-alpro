package main

type agenda struct {
	activity   [9]string
	identifier [9]int
}

type weekAgenda [5]agenda

var firstManagerAgenda weekAgenda
var secondManagerAgenda weekAgenda

func main() {
	initialization()
	selectManager()
}

func initialization() {
	// Inisialisasi data untuk keperluan demo program

	// Agenda manajer pertama pada hari Senin
	firstManagerAgenda[0].activity[2] = "bermain_golf"
	firstManagerAgenda[0].identifier[2] = 0
	firstManagerAgenda[0].activity[3] = "bermain_golf"
	firstManagerAgenda[0].identifier[3] = 0
	firstManagerAgenda[0].activity[5] = "pembahasan_proposal (rapat)"
	firstManagerAgenda[0].identifier[5] = 1
	firstManagerAgenda[0].activity[7] = "presentasi (rapat)"
	firstManagerAgenda[0].identifier[7] = 1
	firstManagerAgenda[0].activity[8] = "mancing_bersama"
	firstManagerAgenda[0].identifier[8] = 0

	// Agenda manajer kedua pada hari Senin
	secondManagerAgenda[0].activity[0] = "belanja"
	secondManagerAgenda[0].identifier[0] = 0
	secondManagerAgenda[0].activity[2] = "servis_mobil"
	secondManagerAgenda[0].identifier[2] = 0
	secondManagerAgenda[0].activity[5] = "pembahasan_proposal (rapat)"
	secondManagerAgenda[0].identifier[5] = 1
	secondManagerAgenda[0].activity[6] = "makan_bersama"
	secondManagerAgenda[0].identifier[6] = 0
	secondManagerAgenda[0].activity[7] = "presentasi (rapat)"
	secondManagerAgenda[0].identifier[7] = 1
}
