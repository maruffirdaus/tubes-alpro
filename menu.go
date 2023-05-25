package main

import "fmt"

func selectManager() {
	/*
		I.S. -
		F.S. Menampilkan menu untuk memilih manajer pertama atau kedua yang akan membuat agenda
	*/

	var selectedMenu int

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

	if selectedMenu == 1 {
		mainMenu(1)
	} else if selectedMenu == 2 {
		mainMenu(2)
	} else if selectedMenu == 0 {

	} else {
		selectManager()
	}
}

func mainMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu utama pencatat agenda kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
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
	fmt.Print("|      9. Ganti Opsi Manajer         |\n")
	fmt.Print("|      0. Keluar                     |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n\n")

	fmt.Print("Pilih Nomor Menu: ")
	fmt.Scan(&selectedMenu)

	if selectedMenu == 1 {
		firstMenu(selectedManager)
	} else if selectedMenu == 2 {
		secondMenu(selectedManager)
	} else if selectedMenu == 3 {
		thirdMenu(selectedManager)
	} else if selectedMenu == 4 {
		fourthMenu(selectedManager)
	} else if selectedMenu == 5 {
		fifthMenu(selectedManager)
	} else if selectedMenu == 6 {
		sixthMenu(selectedManager)
	} else if selectedMenu == 7 {
		seventhMenu(selectedManager)
	} else if selectedMenu == 8 {
		eighthMenu(selectedManager)
	} else if selectedMenu == 9 {
		selectManager()
	} else if selectedMenu == 0 {

	} else {
		mainMenu(selectedManager)
	}
}

func firstMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu isi kegiatan baru
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*-------- ISI KEGIATAN BARU ---------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		firstMenu(selectedManager)
	}
}

func secondMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu ubah kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*---------- UBAH KEGIATAN -----------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		secondMenu(selectedManager)
	}
}

func thirdMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu hapus kegiatan
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*---------- HAPUS KEGIATAN ----------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		thirdMenu(selectedManager)
	}
}

func fourthMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu isi rapat baru
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*---------- ISI RAPAT BARU ----------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		fourthMenu(selectedManager)
	}
}

func fifthMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu ubah rapat
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*------------ UBAH RAPAT ------------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		fifthMenu(selectedManager)
	}
}

func sixthMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu hapus rapat
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*----------- HAPUS RAPAT ------------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		sixthMenu(selectedManager)
	}
}

func seventhMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu optimalisasi agenda
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*------- OPTIMALISASI AGENDA --------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		seventhMenu(selectedManager)
	}
}

func eighthMenu(selectedManager int) {
	/*
		I.S. -
		F.S. Menampilkan menu tampilkan agenda
	*/

	var selectedMenu int

	fmt.Print("\n")
	fmt.Print("*--------- TAMPILKAN AGENDA ---------*\n")
	fmt.Print("|                                    |\n")
	if selectedManager == 1 {
		fmt.Print("|          MANAJER PERTAMA           |\n")
	} else {
		fmt.Print("|           MANAJER KEDUA            |\n")
	}
	fmt.Print("|                                    |\n")
	fmt.Print("|            PILIHAN MENU            |\n")
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
		mainMenu(selectedManager)
	} else {
		eighthMenu(selectedManager)
	}
}
