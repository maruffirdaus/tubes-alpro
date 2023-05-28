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
	fmt.Print("|      2. Ubah Kegiatan              |\n")
	fmt.Print("|      3. Hapus Kegiatan             |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Kegiatan Rapat Bersama:          |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      4. Isi Rapat Baru             |\n")
	fmt.Print("|      5. Ubah Rapat                 |\n")
	fmt.Print("|      6. Hapus Rapat                |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      7. Optimalisasi Agenda        |\n")
	fmt.Print("|      8. Tampilkan Agenda           |\n")
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
		changeActivityMenu()
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		insertMeetingMenu()
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		changeMeetingMenu()
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		deleteMeetingMenu()
	} else if selectedMenu == "7" {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	} else if selectedMenu == "8" {
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
		selectTimeMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		selectTimeMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		selectTimeMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		selectTimeMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		selectTimeMenu(4)
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		insertActivityMenu()
	}
}

func selectTimeMenu(day int) {
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
		fmt.Print("|                                    |\n")
		fmt.Print("|   Isikan detail kegiatan!          |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Catatan: tidak dapat menerima    |\n")
		fmt.Print("|   input ( ) spasi, gunakan (_)     |\n")
		fmt.Print("|   underscore sebagai pengganti.    |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n\n")
		fmt.Print("Nama kegiatan: ")
		fmt.Scan(&activityName)
		fmt.Print("Berapa lama kegiatan (jam): ")
		fmt.Scan(&timeLength)
	}

	if selectedMenu == "a" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 0, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 0, timeLength, activityName)
		}
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 1, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 1, timeLength, activityName)
		}
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 2, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 2, timeLength, activityName)
		}
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 3, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 3, timeLength, activityName)
		}
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 4, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 4, timeLength, activityName)
		}
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 5, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 5, timeLength, activityName)
		}
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 6, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 6, timeLength, activityName)
		}
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 7, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 7, timeLength, activityName)
		}
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			insertActivity(&firstManagerAgenda, day, 8, timeLength, activityName)
		} else {
			insertActivity(&secondManagerAgenda, day, 8, timeLength, activityName)
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		insertActivityMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		selectTimeMenu(day)
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
	} else {
		fmt.Printf("\x1bc")
		changeActivityMenu()
	}
}

func deleteActivityMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu hapus kegiatan
	*/

	fmt.Print("\n")
	fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
	printHeader()
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
	} else {
		fmt.Printf("\x1bc")
		deleteActivityMenu()
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
	} else {
		fmt.Printf("\x1bc")
		insertMeetingMenu()
	}
}

func changeMeetingMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu ubah rapat
	*/

	fmt.Print("\n")
	fmt.Print("*------------ UBAH RAPAT ------------*\n")
	printHeader()
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
	} else {
		fmt.Printf("\x1bc")
		changeMeetingMenu()
	}
}

func deleteMeetingMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu hapus rapat
	*/

	fmt.Print("\n")
	fmt.Print("*----------- HAPUS RAPAT ------------*\n")
	printHeader()
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
	} else {
		fmt.Printf("\x1bc")
		deleteMeetingMenu()
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
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih nomor menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(0)
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(1)
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(2)
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(3)
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(4)
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	}
}

func selectTimeAgendaMenu(day int) {
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
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 0)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 0)
		}
	} else if selectedMenu == "b" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 1)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 1)
		}
	} else if selectedMenu == "c" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 2)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 2)
		}
	} else if selectedMenu == "d" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 3)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 3)
		}
	} else if selectedMenu == "e" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 4)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 4)
		}
	} else if selectedMenu == "f" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 5)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 5)
		}
	} else if selectedMenu == "g" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 6)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 6)
		}
	} else if selectedMenu == "h" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 7)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 7)
		}
	} else if selectedMenu == "i" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showTimeAgenda(firstManagerAgenda, day, 8)
		} else {
			showTimeAgenda(secondManagerAgenda, day, 8)
		}
	} else if selectedMenu == "j" {
		fmt.Printf("\x1bc")
		if selectedManager == 1 {
			showDailyAgenda(firstManagerAgenda, day)
		} else {
			showDailyAgenda(secondManagerAgenda, day)
		}
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(day)
	}
}
