// Menu

package main

import "fmt"

var selectedMenu string
var selectedManager int

func selectManager() {
	/*
		I.S. -
		F.S. Menampilkan menu untuk memilih manajer pertama atau kedua yang akan membuat agenda
	*/

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
	/*
		I.S. -
		F.S. Menampilkan header
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu utama pencatat agenda kegiatan
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu isi kegiatan baru
	*/

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
	/*
		I.S. Terdapat nilai dari hari yang dipilih dan lamanya waktu kegiatan
		F.S. Menampilkan menu isi kegiatan baru
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu isi rapat baru
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu ubah kegiatan
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu hapus kegiatan
	*/

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
	/*
		I.S. -
		F.S. Menampilkan menu optimalisasi agenda
	*/

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
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	}
}

func showAgendaMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu tampilkan agenda
	*/

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
	/*
		I.S. Terdapat nilai dari hari yang dipilih
		F.S. Menampilkan menu tampilkan agenda
	*/

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
