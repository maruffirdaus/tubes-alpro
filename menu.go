package main

import "fmt"

func mainMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu utama pencatat agenda kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
	fmt.Print("|                                    |\n")
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
	fmt.Print("|      0. Keluar                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == 1 {
		firstMenu()
	} else if selectedMenu == 2 {
		secondMenu()
	} else if selectedMenu == 3 {
		thirdMenu()
	} else if selectedMenu == 4 {

	} else if selectedMenu == 5 {

	} else if selectedMenu == 6 {

	} else if selectedMenu == 7 {

	} else if selectedMenu == 8 {

	}
}

func firstMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu isi kegiatan baru
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
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
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == 1 {

	} else if selectedMenu == 2 {

	} else if selectedMenu == 9 {
		mainMenu()
	}
}

func secondMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu ubah kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*---------- UBAH KEGIATAN -----------*\n")
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
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == 1 {

	} else if selectedMenu == 2 {

	} else if selectedMenu == 9 {
		mainMenu()
	}
}

func thirdMenu() {
	/*
		I.S. -
		F.S. Menampilkan menu hapus kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
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
	fmt.Print("|      9. Kembali                    |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == 1 {

	} else if selectedMenu == 2 {

	} else if selectedMenu == 9 {
		mainMenu()
	}
}
