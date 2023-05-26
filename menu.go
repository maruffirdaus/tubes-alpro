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

	fmt.Print("Pilih Nomor Menu: ")
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "1" {
		fmt.Printf("\x1bc")
		firstMenu()
	} else if selectedMenu == "2" {
		fmt.Printf("\x1bc")
		secondMenu()
	} else if selectedMenu == "3" {
		fmt.Printf("\x1bc")
		thirdMenu()
	} else if selectedMenu == "4" {
		fmt.Printf("\x1bc")
		fourthMenu()
	} else if selectedMenu == "5" {
		fmt.Printf("\x1bc")
		fifthMenu()
	} else if selectedMenu == "6" {
		fmt.Printf("\x1bc")
		sixthMenu()
	} else if selectedMenu == "7" {
		fmt.Printf("\x1bc")
		seventhMenu()
	} else if selectedMenu == "8" {
		fmt.Printf("\x1bc")
		eighthMenu()
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

func firstMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu isi kegiatan baru
	*/

	fmt.Print("\n")
	fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
	printHeader()
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		firstMenu()
	}
}

func secondMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		secondMenu()
	}
}

func thirdMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		thirdMenu()
	}
}

func fourthMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		fourthMenu()
	}
}

func fifthMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		fifthMenu()
	}
}

func sixthMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		sixthMenu()
	}
}

func seventhMenu() {
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

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		seventhMenu()
	}
}

func eighthMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu tampilkan agenda
	*/

	fmt.Print("\n")
	fmt.Print("*--------- TAMPILKAN AGENDA ---------*\n")
	printHeader()
	fmt.Print("|   Menu Lain:                       |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		eighthMenu()
	}
}
