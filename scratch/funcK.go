package main

import "fmt"

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