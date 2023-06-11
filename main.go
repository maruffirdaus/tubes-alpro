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
	/* IS -
	   FS terinisiasi setiap jadwal */
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

	// Agenda manajer pertama pada hari selasa
	firstManagerAgenda[1].activity[1] = "Servis_mobil"
	firstManagerAgenda[1].identifier[1] = 0
	firstManagerAgenda[1].activity[2] = "Servis_mobil"
	firstManagerAgenda[1].identifier[2] = 0
	firstManagerAgenda[1].activity[3] = "Beli_Kue"
	firstManagerAgenda[1].identifier[3] = 0
	firstManagerAgenda[1].activity[5] = "pembahasan_finansial (rapat)"
	firstManagerAgenda[1].identifier[5] = 1
	firstManagerAgenda[1].activity[6] = "pembahasan_finansial (rapat)"
	firstManagerAgenda[1].identifier[6] = 1
	firstManagerAgenda[1].activity[7] = "presentasi (rapat)"
	firstManagerAgenda[1].identifier[7] = 1
	firstManagerAgenda[1].activity[8] = "Ke_Balai_Kota"
	firstManagerAgenda[1].identifier[8] = 0

	// Agenda manajer kedua pada hari Selasa
	secondManagerAgenda[1].activity[0] = "belanja"
	secondManagerAgenda[1].identifier[0] = 0
	secondManagerAgenda[1].activity[1] = "belanja"
	secondManagerAgenda[1].identifier[1] = 0
	secondManagerAgenda[1].activity[2] = "servis_mobil"
	secondManagerAgenda[1].identifier[2] = 0
	secondManagerAgenda[1].activity[5] = "pembahasan_finansial (rapat)"
	secondManagerAgenda[1].identifier[5] = 1
	secondManagerAgenda[1].activity[6] = "pembahasan_finansial (rapat)"
	secondManagerAgenda[1].identifier[6] = 1
	secondManagerAgenda[1].activity[7] = "presentasi (rapat)"
	secondManagerAgenda[1].identifier[7] = 1
	secondManagerAgenda[1].activity[8] = "Makan_bersama_keluarga"
	secondManagerAgenda[1].identifier[8] = 0

	// Agenda manajer pertama pada hari rabu
	firstManagerAgenda[2].activity[2] = "berkuda"
	firstManagerAgenda[2].identifier[2] = 0
	firstManagerAgenda[2].activity[3] = "berkuda"
	firstManagerAgenda[2].identifier[3] = 0
	firstManagerAgenda[2].activity[5] = "pembahasan_proposal (rapat)"
	firstManagerAgenda[2].identifier[5] = 0
	firstManagerAgenda[2].activity[7] = "presentasi (rapat)"
	firstManagerAgenda[2].identifier[7] = 0
	firstManagerAgenda[2].activity[8] = "mancing_bersama"
	firstManagerAgenda[2].identifier[8] = 0

	// Agenda manajer kedua pada hari rabu
	secondManagerAgenda[2].activity[0] = "berselancar"
	secondManagerAgenda[2].identifier[0] = 0
	secondManagerAgenda[2].activity[2] = "les_piano"
	secondManagerAgenda[2].identifier[2] = 0
	secondManagerAgenda[2].activity[5] = "istirahat"
	secondManagerAgenda[2].identifier[5] = 0
	secondManagerAgenda[2].activity[6] = "istirahat"
	secondManagerAgenda[2].identifier[6] = 0
	secondManagerAgenda[2].activity[7] = "tinju"
	secondManagerAgenda[2].identifier[7] = 0
	secondManagerAgenda[2].activity[8] = "tinju"
	secondManagerAgenda[2].identifier[8] = 0

	// Agenda manajer pertama pada hari Kamis
	firstManagerAgenda[3].activity[1] = "cuci_mobil"
	firstManagerAgenda[3].identifier[1] = 0
	firstManagerAgenda[3].activity[3] = "sarapan"
	firstManagerAgenda[3].identifier[3] = 0
	firstManagerAgenda[3].activity[5] = "ke_kantor"
	firstManagerAgenda[3].identifier[5] = 0
	firstManagerAgenda[3].activity[7] = "presentasi (rapat)"
	firstManagerAgenda[3].identifier[7] = 1
	firstManagerAgenda[3].activity[8] = "mancing_bersama"
	firstManagerAgenda[3].identifier[8] = 0

	// Agenda manajer kedua pada hari Kamis
	secondManagerAgenda[3].activity[0] = "yoga"
	secondManagerAgenda[3].identifier[0] = 0
	secondManagerAgenda[3].activity[2] = "sarapan"
	secondManagerAgenda[3].identifier[2] = 0
	secondManagerAgenda[3].activity[4] = "ke_mall"
	secondManagerAgenda[3].identifier[4] = 0
	secondManagerAgenda[3].activity[6] = "makan_bersama"
	secondManagerAgenda[3].identifier[6] = 0
	secondManagerAgenda[3].activity[7] = "presentasi (rapat)"
	secondManagerAgenda[3].identifier[7] = 1

	// Agenda manajer pertama pada hari Jumat
	firstManagerAgenda[4].activity[2] = "latihan_bela_diri"
	firstManagerAgenda[4].identifier[2] = 0
	firstManagerAgenda[4].activity[3] = "latihan_bela_diri"
	firstManagerAgenda[4].identifier[3] = 0
	firstManagerAgenda[4].activity[5] = "pembahasan_marketing (rapat)"
	firstManagerAgenda[4].identifier[5] = 1
	firstManagerAgenda[4].activity[6] = "makan_bersama"
	firstManagerAgenda[4].identifier[6] = 0
	firstManagerAgenda[4].activity[7] = "presentasi (rapat)"
	firstManagerAgenda[4].identifier[7] = 1
	firstManagerAgenda[4].activity[8] = "mancing_bersama"
	firstManagerAgenda[4].identifier[8] = 0

	// Agenda manajer kedua pada hari Jumat
	secondManagerAgenda[4].activity[0] = "sarapan"
	secondManagerAgenda[4].identifier[0] = 0
	secondManagerAgenda[4].activity[1] = "jogging"
	secondManagerAgenda[4].identifier[1] = 0
	secondManagerAgenda[4].activity[2] = "jogging"
	secondManagerAgenda[4].identifier[2] = 0
	secondManagerAgenda[4].activity[5] = "pembahasan_marketing (rapat)"
	secondManagerAgenda[4].identifier[5] = 1
	secondManagerAgenda[4].activity[6] = "makan_bersama"
	secondManagerAgenda[4].identifier[6] = 0
	secondManagerAgenda[4].activity[7] = "presentasi (rapat)"
	secondManagerAgenda[4].identifier[7] = 1

}
