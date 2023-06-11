package main

import "fmt"

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