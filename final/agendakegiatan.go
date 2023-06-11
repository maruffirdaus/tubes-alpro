package main

import "fmt"

type agenda struct {
	activity   [9]string
	identifier [9]int
}

type weekAgenda [5]agenda

var firstManagerAgenda weekAgenda  // Menyimpan data agenda manajer pertama
var secondManagerAgenda weekAgenda // Menyimpan data agenda manajer kedua

// Main

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

// Menu

var selectedMenu string // Menyimpan nilai menu terpilih
var selectedManager int // Menyimpan nilai manajer terpilih, 1 = manajer pertama, 2 = manajer kedua

func selectManager() {
	/* IS -
	FS menampilkan menu memilih manajer dan menyimpan nilai manajer terpilih, 1 untuk manajer pertama dan 2 untuk manajer kedua*/

	fmt.Printf("\x1bc")
	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Opsi Manajer:              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Manajer Pertama            |\n")
	fmt.Print("|      2. Manajer Kedua              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      0. Keluar                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		selectedManager = 1
		mainMenu()
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		selectedManager = 2
		mainMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
	} else {
		fmt.Printf("\x1bc")
		selectManager()
	}
}

func printHeader() {
	/* IS -
	FS menampilkan header */

	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|        [ MANAJER PERTAMA  ]        |\n")
	} else {
		fmt.Print("|         [ MANAJER KEDUA  ]         |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
}

func mainMenu() {
	/* IS -
	FS menampilkan main menu */

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
	printHeader()
	fmt.Print("|   Kegiatan Pribadi:                |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Isi Kegiatan Baru          |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Kegiatan Rapat Bersama:          |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      2. Isi Rapat Baru             |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Ubah/Hapus Kegiatan:             |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      3. Ubah Kegiatan              |\n")
	fmt.Print("|      4. Hapus Kegiatan             |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      5. Optimalisasi Agenda        |\n")
	fmt.Print("|      6. Tampilkan Agenda           |\n")
	fmt.Print("|      9. Ganti Opsi Manajer         |\n")
	fmt.Print("|      0. Keluar                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		insertActivityMenu()
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		insertMeetingMenu()
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		changeActivityMenu()
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		selectManager()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
	} else {
		fmt.Printf("\x1bc")
		mainMenu()
	}
}

func insertActivityMenu() {
	/* IS -
	FS menampilkan menu isi kegiatan baru */

	fmt.Print("\n")
	fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(4)
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		insertActivityMenu()
	}
}

func activitySelectTimeMenu(day int) {
	/* IS terdefinisi hari yang dipilih
	FS menampilkan menu memilih waktu mengisi kegiatan baru */

	var activityName string
	var timeLength int

	fmt.Print("\n")
	fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
	printHeader()
	fmt.Print("|   Pilih Waktu Mulai Kegiatan:      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      a. 08.00                      |\n")
	fmt.Print("|      b. 09.00                      |\n")
	fmt.Print("|      c. 10.00                      |\n")
	fmt.Print("|      d. 11.00                      |\n")
	fmt.Print("|      e. 12.00                      |\n")
	fmt.Print("|      f. 13.00                      |\n")
	fmt.Print("|      g. 14.00                      |\n")
	fmt.Print("|      h. 15.00                      |\n")
	fmt.Print("|      i. 16.00                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      x. Tampilkan Cuplikan         |\n")
	fmt.Print("|         Agenda                     |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "a" || selectedMenu == "b" || selectedMenu == "c" || selectedMenu == "d" || selectedMenu == "e" || selectedMenu == "f" || selectedMenu == "g" || selectedMenu == "h" || selectedMenu == "i" || selectedMenu == "j" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan detail kegiatan!           \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan: ")
		fmt.Scan(&activityName)
		fmt.Print("Berapa lama waktu kegiatan (jam): ")
		fmt.Scan(&timeLength)
	}

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		insertActivity(day, 0, timeLength, activityName)
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		insertActivity(day, 1, timeLength, activityName)
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		insertActivity(day, 2, timeLength, activityName)
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		insertActivity(day, 3, timeLength, activityName)
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		insertActivity(day, 4, timeLength, activityName)
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		insertActivity(day, 5, timeLength, activityName)
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		insertActivity(day, 6, timeLength, activityName)
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		insertActivity(day, 7, timeLength, activityName)
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		insertActivity(day, 8, timeLength, activityName)
	} else if selectedMenu == "x" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
		if selectedMenu == "9" {
			fmt.Printf("\x1bc")
			activitySelectTimeMenu(day)
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		insertActivityMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		activitySelectTimeMenu(day)
	}
}

func insertMeetingMenu() {
	/* IS -
	FS menampilkan menu mengisi rapat baru*/

	fmt.Print("\n")
	fmt.Print("*---------- ISI RAPAT BARU ----------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(4)
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		insertMeetingMenu()
	}
}

func meetingSelectTimeMenu(day int) {
	/* IS terdefinisi hari yang dipilih
	FS menampilkan menu memilih waktu saat akan mengisi rapat baru */

	var activityName string
	var timeLength int

	fmt.Print("\n")
	fmt.Print("*---------- ISI RAPAT BARU ----------*\n")
	printHeader()
	fmt.Print("|   Pilih Waktu Mulai Kegiatan:      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      a. 08.00                      |\n")
	fmt.Print("|      b. 09.00                      |\n")
	fmt.Print("|      c. 10.00                      |\n")
	fmt.Print("|      d. 11.00                      |\n")
	fmt.Print("|      e. 12.00                      |\n")
	fmt.Print("|      f. 13.00                      |\n")
	fmt.Print("|      g. 14.00                      |\n")
	fmt.Print("|      h. 15.00                      |\n")
	fmt.Print("|      i. 16.00                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      x. Tampilkan Cuplikan         |\n")
	fmt.Print("|         Agenda                     |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "a" || selectedMenu == "b" || selectedMenu == "c" || selectedMenu == "d" || selectedMenu == "e" || selectedMenu == "f" || selectedMenu == "g" || selectedMenu == "h" || selectedMenu == "i" || selectedMenu == "j" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*---------- ISI RAPAT BARU ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan detail rapat!           \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama rapat: ")
		fmt.Scan(&activityName)
		fmt.Print("Berapa lama waktu rapat (jam): ")
		fmt.Scan(&timeLength)
	}

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 0, timeLength, activityName)
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 1, timeLength, activityName)
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 2, timeLength, activityName)
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 3, timeLength, activityName)
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 4, timeLength, activityName)
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 5, timeLength, activityName)
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 6, timeLength, activityName)
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 7, timeLength, activityName)
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		insertMeeting(day, 8, timeLength, activityName)
	} else if selectedMenu == "x" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
		if selectedMenu == "9" {
			fmt.Printf("\x1bc")
			meetingSelectTimeMenu(day)
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		insertMeetingMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		meetingSelectTimeMenu(day)
	}
}

func changeActivityMenu() {
	/* IS -
	FS menampilkan menu ubah kegiatan/rapat */

	fmt.Print("\n")
	fmt.Print("*---------- UBAH KEGIATAN -----------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(4)
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		changeActivityMenu()
	}
}

func changeSelectTimeMenu(day int) {
	/* IS terdefinisi hari yang dipilih
	FS menampilkan menu memilih waktu saat akan mengubah kegiatan/rapat */

	var activityName, newActivityName string

	fmt.Print("\n")
	fmt.Print("*----------- UBAH KEGIATAN ----------*\n")
	printHeader()
	fmt.Print("|   Pilih Waktu:                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      a. 08.00 - 09.00              |\n")
	fmt.Print("|      b. 09.00 - 10.00              |\n")
	fmt.Print("|      c. 10.00 - 11.00              |\n")
	fmt.Print("|      d. 11.00 - 12.00              |\n")
	fmt.Print("|      e. 12.00 - 13.00              |\n")
	fmt.Print("|      f. 13.00 - 14.00              |\n")
	fmt.Print("|      g. 14.00 - 15.00              |\n")
	fmt.Print("|      h. 15.00 - 16.00              |\n")
	fmt.Print("|      i. 16.00 - 17.00              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      j. Ubah Berdasarkan Nama      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      x. Tampilkan Cuplikan         |\n")
	fmt.Print("|         Agenda                     |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "a" || selectedMenu == "b" || selectedMenu == "c" || selectedMenu == "d" || selectedMenu == "e" || selectedMenu == "f" || selectedMenu == "g" || selectedMenu == "h" || selectedMenu == "i" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*----------- UBAH KEGIATAN ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan detail kegiatan/rapat      \n")
		fmt.Print("    baru!                             \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan/rapat baru: ")
		fmt.Scan(&activityName)
	}

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		changeActivity(day, 0, activityName)
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		changeActivity(day, 1, activityName)
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		changeActivity(day, 2, activityName)
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		changeActivity(day, 3, activityName)
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		changeActivity(day, 4, activityName)
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		changeActivity(day, 5, activityName)
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		changeActivity(day, 6, activityName)
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		changeActivity(day, 7, activityName)
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		changeActivity(day, 8, activityName)
	} else if selectedMenu == "j" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*----------- UBAH KEGIATAN ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan detail kegiatan/rapat      \n")
		fmt.Print("    yang akan diubah!                 \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan/rapat: ")
		fmt.Scan(&activityName)
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*----------- UBAH KEGIATAN ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan detail kegiatan/rapat      \n")
		fmt.Print("    baru!                             \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan/rapat baru: ")
		fmt.Scan(&newActivityName)
		fmt.Printf("\x1bc")
		changeActivityByName(day, activityName, newActivityName)
	} else if selectedMenu == "x" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
		if selectedMenu == "9" {
			fmt.Printf("\x1bc")
			changeSelectTimeMenu(day)
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		changeActivityMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		changeSelectTimeMenu(day)
	}
}

func deleteActivityMenu() {
	/* IS -
	FS menampilkan menu hapus kegiatan/rapat */

	var activityName string

	fmt.Print("\n")
	fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      6. Hapus Semua Berdasarkan    |\n")
	fmt.Print("|         Nama                       |\n")
	fmt.Print("|      7. Hapus Semua                |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)
	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(4)
	}

	if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan nama kegiatan/rapat        \n")
		fmt.Print("    yang akan dihapus!                \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan/rapat: ")
		fmt.Scan(&activityName)
	}

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(4)
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		deleteAllActivityByName(activityName)
	} else if selectedMenu == "7" {
		fmt.Printf("\x1bc")
		deleteAllActivity()
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	}
}

func deleteSelectTimeMenu(day int) {
	/* IS terdefinisi hari yang dipilih
	FS menampilkan menu memilih waktu saat akan menghapus kegiatan/rapat */

	var activityName string

	fmt.Print("\n")
	fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
	printHeader()
	fmt.Print("|   Pilih Waktu:                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      a. 08.00 - 09.00              |\n")
	fmt.Print("|      b. 09.00 - 10.00              |\n")
	fmt.Print("|      c. 10.00 - 11.00              |\n")
	fmt.Print("|      d. 11.00 - 12.00              |\n")
	fmt.Print("|      e. 12.00 - 13.00              |\n")
	fmt.Print("|      f. 13.00 - 14.00              |\n")
	fmt.Print("|      g. 14.00 - 15.00              |\n")
	fmt.Print("|      h. 15.00 - 16.00              |\n")
	fmt.Print("|      i. 16.00 - 17.00              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      j. Hapus Berdasarkan Nama     |\n")
	fmt.Print("|      k. Hapus Semua                |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      x. Tampilkan Cuplikan         |\n")
	fmt.Print("|         Agenda                     |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "j" {
		fmt.Printf("\x1bc")
		fmt.Print("\n")
		fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Isikan nama kegiatan/rapat        \n")
		fmt.Print("    yang akan dihapus!                \n")
		fmt.Print("                                      \n")
		fmt.Print("    Catatan: tidak dapat menerima     \n")
		fmt.Print("    input ( ) spasi, gunakan (_)      \n")
		fmt.Print("    underscore sebagai pengganti.     \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan/rapat: ")
		fmt.Scan(&activityName)
	}

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 0)
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 1)
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 2)
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 3)
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 4)
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 5)
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 6)
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 7)
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		deleteActivity(day, 8)
	} else if selectedMenu == "j" {
		fmt.Printf("\x1bc")
		deleteActivityByName(day, activityName)
	} else if selectedMenu == "k" {
		fmt.Printf("\x1bc")
		deleteAllActivityByDay(day)
	} else if selectedMenu == "x" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
		if selectedMenu == "9" {
			fmt.Printf("\x1bc")
			deleteSelectTimeMenu(day)
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(day)
	}
}

func optimizeAgendaMenu() {
	/* IS -
	FS menampilkan menu optimalisasi agenda */

	fmt.Print("\n")
	fmt.Print("*------- OPTIMALISASI AGENDA --------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      6. Optimalkan Semua           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		optimizeAgenda(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		optimizeAgenda(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		optimizeAgenda(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		optimizeAgenda(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		optimizeAgenda(4)
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		optimizeAllAgenda()
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	}
}

func showAgendaMenu() {
	/* IS -
	FS menampilkan menu tampilkan agenda */

	fmt.Print("\n")
	fmt.Print("*--------- TAMPILKAN AGENDA ---------*\n")
	printHeader()
	fmt.Print("|   Pilih Hari:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Senin                      |\n")
	fmt.Print("|      2. Selasa                     |\n")
	fmt.Print("|      3. Rabu                       |\n")
	fmt.Print("|      4. Kamis                      |\n")
	fmt.Print("|      5. Jumat                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      6. Tampilkan Semua            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(4)
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		showAllAgenda()
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	}
}

func showSelectTimeMenu(day int) {
	/* IS terdefinisi hari yang dipilih
	FS menampilkan menu memilih waktu saat akan menampilkan agenda */

	fmt.Print("\n")
	fmt.Print("*--------- TAMPILKAN AGENDA ---------*\n")
	printHeader()
	fmt.Print("|   Pilih Waktu:                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      a. 08.00 - 09.00              |\n")
	fmt.Print("|      b. 09.00 - 10.00              |\n")
	fmt.Print("|      c. 10.00 - 11.00              |\n")
	fmt.Print("|      d. 11.00 - 12.00              |\n")
	fmt.Print("|      e. 12.00 - 13.00              |\n")
	fmt.Print("|      f. 13.00 - 14.00              |\n")
	fmt.Print("|      g. 14.00 - 15.00              |\n")
	fmt.Print("|      h. 15.00 - 16.00              |\n")
	fmt.Print("|      i. 16.00 - 17.00              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilihan Lain:                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      j. Tampilkan Semua            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 0)
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 1)
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 2)
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 3)
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 4)
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 5)
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 6)
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 7)
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, 8)
	} else if selectedMenu == "j" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
		if selectedMenu == "9" {
			fmt.Printf("\x1bc")
			showSelectTimeMenu(day)
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(day)
	}
}

// Search Function

func checkActivity(A weekAgenda, day, time, timeLength int) bool {
	// mengembalikan true jika slot kosong dan false jika sudah terisi

	var status bool = true
	var i int = time

	for i < time+timeLength && status {
		if A[day].activity[i] != "" {
			status = false
		}
		i++
	}

	return status
}

func searchActivityName(A weekAgenda, day, i int, activityName string) int {
	// mengembalikan indeks dari suatu kegiatan/rapat jika ditemukan dan -1 jika tidak ditemukan

	var found bool = false
	for i < 9 && !found {
		found = A[day].activity[i] == activityName
		i = i + 1
	}
	if found {
		return i - 1
	} else {
		return -1
	}
}

func searchActivityLength(A weekAgenda, day, i int) int {
	//mengembalikan durasi kegiatan dari agenda pada hari yang ditentukan

	var status bool
	var tempName string
	var j int
	tempName = A[day].activity[i]
	for i < 9 && !status {
		i++
		if i < 9 {
			status = tempName != A[day].activity[i]
		} else {
			status = true
		}
		j++
	}
	return j
}

// Extra Function

func dayString(day int) string {
	//mengembalikan string hari dari indeks

	if day == 0 {
		return "Senin"
	} else if day == 1 {
		return "Selasa"
	} else if day == 2 {
		return "Rabu"
	} else if day == 3 {
		return "Kamis"
	} else {
		return "Jumat"
	}
}

// Insert Function

func insertActivity(day, time, timeLength int, activityName string) {
	/* IS terdefinis hari, waktu, durasi, & nama kegiatan
	FS durasi, & nama kegiatan masuk ke dalam agenda hari dan waktu yang dipilih */

	var A weekAgenda
	var i int

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	if 9 < time+timeLength {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Lama waktu kegiatan melebihi      \n")
		fmt.Print("    batas waktu yang bisa             \n")
		fmt.Print("    digunakan, silahkan pilih         \n")
		fmt.Print("    waktu yang lain atau kurangi      \n")
		fmt.Print("    lama waktu kegiatan!              \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			activitySelectTimeMenu(day)
		}
	} else if checkActivity(A, day, time, timeLength) {
		for i = time; i < time+timeLength; i++ {
			A[day].activity[i] = activityName
			A[day].identifier[i] = 0
		}

		if selectedManager == 1 {
			firstManagerAgenda = A
		} else {
			secondManagerAgenda = A
		}

		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Printf("    Kegiatan %v\n", activityName)
		fmt.Print("    berhasil diisikan pada agenda     \n")
		fmt.Printf("    hari %v dengan waktu\n", dayString(day))
		if time+8+timeLength < 10 {
			fmt.Printf("    0%v.00 - 0%v.00.\n", time+8, time+8+timeLength)
		} else if time+8+timeLength == 10 {
			fmt.Printf("    0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else if time+8 < 10 {
			fmt.Printf("    0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else {
			fmt.Printf("    %v.00 - %v.00.\n", time+8, time+8+timeLength)
		}
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Pilih Aksi:                      |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      1. Tampilkan Agenda           |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "1" {
			fmt.Printf("\x1bc")
			showAgendaMenu()
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			activitySelectTimeMenu(day)
		}
	} else {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Agenda pada waktu tersebut        \n")
		fmt.Print("    telah terisi, silahkan pilih      \n")
		fmt.Print("    waktu yang lain atau kurangi      \n")
		fmt.Print("    lama waktu kegiatan!              \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			activitySelectTimeMenu(day)
		}
	}
}

func insertMeeting(day, time, timeLength int, activityName string) {
	/* IS terdefinis hari, waktu, durasi, & nama rapat
	FS durasi, & nama rapat masuk ke dalam agenda hari dan waktu yang dipilih */

	var A, B weekAgenda
	var i int

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	if 9 < time+timeLength {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Lama waktu rapat melebihi         \n")
		fmt.Print("    batas waktu yang bisa             \n")
		fmt.Print("    digunakan, silahkan pilih         \n")
		fmt.Print("    waktu yang lain atau kurangi      \n")
		fmt.Print("    lama waktu rapat!                 \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			meetingSelectTimeMenu(day)
		}
	} else if checkActivity(A, day, time, timeLength) && checkActivity(B, day, time, timeLength) {
		for i = time; i < time+timeLength; i++ {
			A[day].activity[i] = activityName + " (rapat)"
			A[day].identifier[i] = 1
			B[day].activity[i] = activityName + " (rapat)"
			B[day].identifier[i] = 1
		}

		if selectedManager == 1 {
			firstManagerAgenda = A
			secondManagerAgenda = B
		} else {
			secondManagerAgenda = A
			firstManagerAgenda = B
		}

		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Printf("    Rapat %v\n", activityName)
		fmt.Print("    berhasil diisikan pada agenda     \n")
		fmt.Printf("    hari %v dengan waktu\n", dayString(day))
		if time+8+timeLength < 10 {
			fmt.Printf("    0%v.00 - 0%v.00.\n", time+8, time+8+timeLength)
		} else if time+8+timeLength == 10 {
			fmt.Printf("    0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else if time+8 < 10 {
			fmt.Printf("    0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else {
			fmt.Printf("    %v.00 - %v.00.\n", time+8, time+8+timeLength)
		}
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Pilih Aksi:                      |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      1. Tampilkan Agenda           |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "1" {
			fmt.Printf("\x1bc")
			showAgendaMenu()
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			meetingSelectTimeMenu(day)
		}
	} else {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Agenda salah satu atau kedua      \n")
		fmt.Print("    manajer pada waktu tersebut       \n")
		fmt.Print("    telah terisi, silahkan pilih      \n")
		fmt.Print("    waktu yang lain atau kurangi      \n")
		fmt.Print("    lama waktu rapat!                 \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			meetingSelectTimeMenu(day)
		}
	}
}

// Modify Function

func changeActivity(day, time int, activityName string) {
	/* IS terdefinis hari, waktu, dan nama kegiatan/rapat baru
	FS kegiatan/rapat pada hari dan waktu yang dipilih sudah terubah */

	var A, B weekAgenda

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	if A[day].activity[time] != "" {
		if A[day].identifier[time] == 0 {
			A[day].activity[time] = activityName
		} else {
			A[day].activity[time] = activityName + " (rapat)"
			B[day].activity[time] = activityName + " (rapat)"
		}

		if selectedManager == 1 {
			firstManagerAgenda = A
			secondManagerAgenda = B
		} else {
			secondManagerAgenda = A
			firstManagerAgenda = B
		}

		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Kegiatan/rapat pada agenda        \n")
		fmt.Printf("    hari %v dengan waktu\n", dayString(day))
		if time+9 < 10 {
			fmt.Printf("    0%v.00 - 0%v.00 berhasil\n", time+8, time+9)
		} else if time+9 == 10 {
			fmt.Printf("    0%v.00 - %v.00 berhasil\n", time+8, time+9)
		} else {
			fmt.Printf("    %v.00 - %v.00 berhasil\n", time+8, time+9)
		}
		fmt.Printf("    diubah menjadi %v.\n", activityName)
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Pilih Aksi:                      |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      1. Tampilkan Agenda           |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "1" {
			fmt.Printf("\x1bc")
			showAgendaMenu()
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			changeSelectTimeMenu(day)
		}
	} else {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Tidak terdapat kegiatan/rapat     \n")
		fmt.Print("    pada agenda dengan waktu          \n")
		if time+9 < 10 {
			fmt.Printf("    0%v.00 - 0%v.00, silahkan\n", time+8, time+9)
		} else if time+9 == 10 {
			fmt.Printf("    0%v.00 - %v.00, silahkan\n", time+8, time+9)
		} else {
			fmt.Printf("    %v.00 - %v.00, silahkan\n", time+8, time+9)
		}
		fmt.Print("    pilih waktu yang lain!            \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			changeSelectTimeMenu(day)
		}
	}
}

func changeActivityByName(day int, activityName, newActivityName string) {
	/* IS terdefinis hari, nama kegiatan/rapat yang akan diubah, nama kegiatan/rapat baru
	FS kegiatan/rapat pada hari dan nama yang dipilih sudah terubah */

	var A, B weekAgenda
	var i, j int
	var status bool

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	for i = 0; i < 9; i++ {
		if searchActivityName(A, day, i, activityName) != -1 {
			j = searchActivityName(A, day, i, activityName)
			A[day].activity[j] = newActivityName
			if !status {
				status = true
			}
		} else if searchActivityName(A, day, i, activityName+" (rapat)") != -1 {
			j = searchActivityName(A, day, i, activityName+" (rapat)")
			A[day].activity[j] = newActivityName + " (rapat)"
			B[day].activity[j] = newActivityName + " (rapat)"
			if !status {
				status = true
			}
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	if status {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Printf("    Kegiatan/rapat %v\n", activityName)
		fmt.Printf("    pada hari %v berhasil\n", dayString(day))
		fmt.Printf("    diubah menjadi %v.\n", newActivityName)
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Pilih Aksi:                      |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      1. Tampilkan Agenda           |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "1" {
			fmt.Printf("\x1bc")
			showAgendaMenu()
		} else if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			changeSelectTimeMenu(day)
		}
	} else {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("                                      \n")
		fmt.Print("    Tidak terdapat kegiatan/rapat     \n")
		fmt.Printf("    dengan nama %v\n", activityName)
		fmt.Printf("    pada hari %v, silahkan\n", dayString(day))
		fmt.Print("    isikan nama kegiatan/rapat        \n")
		fmt.Print("    yang lain!                        \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")

		fmt.Print("\n")
		fmt.Print("*------------------------------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|            PILIHAN MENU            |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Menu Lain:                       |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|      9. Kembali                    |\n")
		fmt.Print("|      0. Menu Utama                 |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")

		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&selectedMenu)

		if selectedMenu == "0" {
			fmt.Printf("\x1bc")
			mainMenu()
		} else {
			fmt.Printf("\x1bc")
			changeSelectTimeMenu(day)
		}
	}
}

// Delete Function

func deleteActivity(day, time int) {
	/* IS terdefinis hari, dan waktu kegiatan yang akan dihapus
	FS kegiatan/rapat pada hari dan waktu yang dipilih terhapus */

	var A, B weekAgenda

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	if A[day].identifier[time] == 0 {
		A[day].activity[time] = ""
	} else {
		A[day].activity[time] = ""
		B[day].activity[time] = ""
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Kegiatan/rapat pada agenda        \n")
	fmt.Printf("    hari %v dengan waktu\n", dayString(day))
	if time+9 < 10 {
		fmt.Printf("    0%v.00 - 0%v.00 berhasil\n", time+8, time+9)
	} else if time+9 == 10 {
		fmt.Printf("    0%v.00 - %v.00 berhasil\n", time+8, time+9)
	} else {
		fmt.Printf("    %v.00 - %v.00 berhasil\n", time+8, time+9)
	}
	fmt.Print("    dihapus.                          \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(day)
	}
}

func deleteActivityByName(day int, activityName string) {
	/* IS terdefinis hari dan nama kegiatan/rapat yang akan dihapus
	FS kegiatan/rapat pada hari dan nama yang dipilih terhapus */

	var A, B weekAgenda
	var i, j int

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	for i = 0; i < 9; i++ {
		if searchActivityName(A, day, i, activityName) != -1 {
			j = searchActivityName(A, day, i, activityName)
			A[day].activity[j] = ""
		} else if searchActivityName(A, day, i, activityName+" (rapat)") != -1 {
			j = searchActivityName(A, day, i, activityName+" (rapat)")
			A[day].activity[j] = ""
			B[day].activity[j] = ""
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Printf("    Kegiatan/rapat %v\n", activityName)
	fmt.Printf("    pada hari %v berhasil\n", dayString(day))
	fmt.Print("    dihapus.                          \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(day)
	}
}

func deleteAllActivity() {
	/* IS -
	FS semua kegiatan terhapus */

	var A, B weekAgenda
	var i, j int

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	for i = 0; i < 5; i++ {
		for j = 0; j < 9; j++ {
			if A[i].identifier[j] == 0 {
				A[i].activity[j] = ""
				A[i].identifier[j] = 0
			} else {
				A[i].activity[j] = ""
				A[i].identifier[j] = 0
				B[i].activity[j] = ""
				B[i].identifier[j] = 0
			}
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Semua kegiatan/rapat berhasil     \n")
	fmt.Print("    dihapus.                          \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	}
}

func deleteAllActivityByDay(day int) {
	/* IS terdefinisi hari pada kegiatan/rapat yang akan dihapus
	FS kegiatan/rapat pada hari yang dipilih terhapus */

	var A, B weekAgenda
	var i int

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	for i = 0; i < 9; i++ {
		if A[day].identifier[i] == 0 {
			A[day].activity[i] = ""
		} else {
			A[day].activity[i] = ""
			B[day].activity[i] = ""
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Semua kegiatan/rapat pada hari    \n")
	fmt.Printf("    %v berhasil dihapus.\n", dayString(day))
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteSelectTimeMenu(day)
	}
}

func deleteAllActivityByName(activityName string) {
	/* IS terdefinisi nama kegiatan/rapat yang akan dihapus
	FS semua kegiatan/rapat pada nama yang dipilih terhapus */

	var A, B weekAgenda
	var i, j, k int

	if selectedManager == 1 {
		A = firstManagerAgenda
		B = secondManagerAgenda
	} else {
		A = secondManagerAgenda
		B = firstManagerAgenda
	}

	for i = 0; i < 5; i++ {
		for j = 0; j < 9; j++ {
			if searchActivityName(A, i, j, activityName) != -1 {
				k = searchActivityName(A, i, j, activityName)
				A[i].activity[k] = ""
			} else if searchActivityName(A, i, j, activityName+" (rapat)") != -1 {
				k = searchActivityName(A, i, j, activityName+" (rapat)")
				A[i].activity[k] = ""
				B[i].activity[k] = ""
			}
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
		secondManagerAgenda = B
	} else {
		secondManagerAgenda = A
		firstManagerAgenda = B
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Printf("    Semua kegiatan/rapat %v\n", activityName)
	fmt.Print("    berhasil dihapus.                 \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	}
}

// Sorting Function

func optimizeAgenda(day int) {
	/* IS terpilih hari yang ditentukan
	FS mengoptimalkan data agenda pada hari yang dipilih */

	var i, j, k, tempA, tempB int
	var A weekAgenda
	var status bool

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	for i < 9 {
		if A[day].activity[i] == "" {
			tempA = searchActivityLength(A, day, i)
			if i+tempA < 9 {
				j = i + tempA
				if A[day].activity[j] != "" && A[day].identifier[j] == 0 {
					A[day].activity[i] = A[day].activity[j]
					A[day].activity[j] = ""
				}
			}
		}
		i++
	}

	i = 0

	for i < 9 {
		if A[day].activity[i] == "" {
			tempA = searchActivityLength(A, day, i)
			j = i + tempA
			status = false
			for j < 9 && !status {
				tempB = searchActivityLength(A, day, j)
				if A[day].activity[j] != "" && A[day].identifier[j] == 0 {
					if tempB <= tempA {
						for k = i; k < i+tempB; k++ {
							A[day].activity[k] = A[day].activity[j]
						}
						for k = j; k < j+tempB; k++ {
							A[day].activity[k] = ""
						}
						status = true
					}
				}
				j = j + tempB
			}
		}
		i++
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
	} else {
		secondManagerAgenda = A
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Printf("    Agenda pada hari %v telah\n", dayString(day))
	fmt.Print("    dioptimalkan.                     \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	}
}

func optimizeAllAgenda() {
	/* IS -
	FS mengoptimalkan seluruh data agenda */

	var A weekAgenda
	var i, j, k, l, m, tempA, tempB int
	var status bool

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	for i < 5 {
		j = 0
		for j < 9 {
			if A[i].activity[j] == "" {
				tempA = searchActivityLength(A, i, j)
				if j+tempA < 9 {
					k = j + tempA
					if A[i].activity[k] != "" && A[i].identifier[k] == 0 {
						A[i].activity[j] = A[i].activity[k]
						A[i].activity[k] = ""
					}
				}
			}
			j++
		}
		i++
	}

	i = 0
	j = 0
	k = 0

	for i < 5 {
		j = 0
		for j < 9 {
			if A[i].activity[j] == "" {
				tempA = searchActivityLength(A, i, j)
				k = i
				l = j + tempA
				status = false
				for k < 5 && !status {
					if i < k {
						l = 0
					}
					for l < 9 {
						tempB = searchActivityLength(A, k, l)
						if A[k].activity[l] != "" && A[k].identifier[l] == 0 {
							if tempB <= tempA {
								for m = j; m < j+tempB; m++ {
									A[i].activity[m] = A[k].activity[l]
								}
								for m = l; m < l+tempB; m++ {
									A[k].activity[m] = ""
								}
								status = true
							}
						}
						l = l + tempB
					}
					k++
				}
			}
			j++
		}
		i++
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
	} else {
		secondManagerAgenda = A
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Semua agenda telah                \n")
	fmt.Print("    dioptimalkan.                     \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Pilih Aksi:                      |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      1. Tampilkan Agenda           |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	}
}

// Print Function

func showTimeAgenda(day, time int) {
	/* IS terpilih hari dan waktu yang ditentukan
	FS menampilkan data agenda pada hari dan waktu yang dipilih */

	var A weekAgenda

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")
	fmt.Printf("       Agenda pada hari %v:\n\n", dayString(day))
	fmt.Print("        Waktu           Kegiatan\n")

	if time+9 < 10 {
		fmt.Printf("    0%v.00 - 0%v.00 | %v\n", time+8, time+9, A[day].activity[time])
	} else if time+9 == 10 {
		fmt.Printf("    0%v.00 - %v.00 | %v\n", time+8, time+9, A[day].activity[time])
	} else {
		fmt.Printf("    %v.00 - %v.00 | %v\n", time+8, time+9, A[day].activity[time])
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(day)
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, time)
	}
}

func showDailyAgenda(day int) {
	/* IS terpilih hari yang ditentukan
	FS menampilkan data agenda pada hari yang dipilih */

	var A weekAgenda
	var i int
	var j int = 8

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")
	fmt.Printf("       Agenda pada hari %v:\n\n", dayString(day))
	fmt.Print("        Waktu           Kegiatan\n")

	for i = 0; i < 9; i++ {
		if j+1 < 10 {
			fmt.Printf("    0%v.00 - 0%v.00 | %v\n", j, j+1, A[day].activity[i])
		} else if j+1 == 10 {
			fmt.Printf("    0%v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		} else {
			fmt.Printf("    %v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		}
		j++
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu != "9" && selectedMenu != "0" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
	}
}

func showAllAgenda() {
	/* IS -
	FS menampilkan seluruh agenda */

	var A weekAgenda
	var i, j, k int

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")

	for i = 0; i < 5; i++ {
		k = 8

		fmt.Printf("       Agenda pada hari %v:\n\n", dayString(i))
		fmt.Print("        Waktu           Kegiatan\n")

		for j = 0; j < 9; j++ {
			if k+1 < 10 {
				fmt.Printf("    0%v.00 - 0%v.00 | %v\n", k, k+1, A[i].activity[j])
			} else if k+1 == 10 {
				fmt.Printf("    0%v.00 - %v.00 | %v\n", k, k+1, A[i].activity[j])
			} else {
				fmt.Printf("    %v.00 - %v.00 | %v\n", k, k+1, A[i].activity[j])
			}
			k++
		}

		fmt.Print("\n")
	}

	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|      0. Menu Utama                 |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showAllAgenda()
	}
}
