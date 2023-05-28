// Kegiatan pribadi
// identifier 0 = kegiatan pribadi
// identifier 1 = rapat bersama
// day 0,1,2,3,4 = hari senin, selasa, rabu, kamis, jumat
// time 0 = 08.00-09.00
// time 1 = 09.00-10.00
// time 2 = 10.00-11.00
// dst.

package main

import "fmt"

func checkActivity(A weekAgenda, day, time, timeLength int) bool {
	// Mengecek apakah agenda pada hari dan waktu yang dipilih kosong ataukah tidak
	// true = kosong
	// false = tidak kosong

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

func insertActivity(A *weekAgenda, day, time, timeLength int, activityName string) {
	/*
		I.S. Terdapat nilai dari hari dan waktu yang dipilih, lamanya waktu, dan nama kegiatan pribadi
		F.S. Mengisi kegiatan pribadi ke dalam array
	*/

	var i int

	if 9 < time+timeLength {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Lama kegiatan melebihi batas     |\n")
		fmt.Print("|   waktu yang bisa digunakan,       |\n")
		fmt.Print("|   silahkan pilih waktu yang lain   |\n")
		fmt.Print("|   atau kurangi lama waktu          |\n")
		fmt.Print("|   kegiatan!                        |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n")
		selectTimeMenu(day)
	} else if checkActivity(*A, day, time, timeLength) {
		for i = time; i < time+timeLength; i++ {
			A[day].activity[i] = activityName
			A[day].identifier[i] = 0
		}
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Kegiatan telah berhasil          |\n")
		fmt.Print("|   diisikan pada agenda dengan      |\n")
		fmt.Print("|   waktu yang dipilih.              |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n")
		mainMenu()
	} else {
		fmt.Print("\n")
		fmt.Print("*------------ PERINGATAN ------------*\n")
		fmt.Print("|                                    |\n")
		fmt.Print("|   Agenda pada waktu tersebut       |\n")
		fmt.Print("|   telah terisi, silahkan pilih     |\n")
		fmt.Print("|   waktu yang lain atau kurangi     |\n")
		fmt.Print("|   lama waktu kegiatan!             |\n")
		fmt.Print("|                                    |\n")
		fmt.Print("*------------------------------------*\n")
		selectTimeMenu(day)
	}
}

func searchActivityName(A weekAgenda, day, i int, activityName string) int {

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

func changeActivity(A *weekAgenda, day, time int) {

}

func deleteActivity(day, time int) {
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
	fmt.Print("|                                    |\n")
	fmt.Print("|   Kegiatan telah berhasil          |\n")
	fmt.Print("|   dihapus pada agenda dengan       |\n")
	fmt.Print("|   waktu yang dipilih.              |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteAllActivity(day int) {
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
	fmt.Print("|                                    |\n")
	fmt.Print("|   Semua kegiatan berhasil          |\n")
	fmt.Print("|   dihapus.                         |\n")
	fmt.Print("|                                    |\n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteActivityByName(day int, activityName string) {
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
		if searchActivityName(A, day, i, activityName) != -1 {
			if A[day].identifier[searchActivityName(A, day, i, activityName)] == 0 {
				A[day].activity[searchActivityName(A, day, i, activityName)] = ""
			} else {
				A[day].activity[searchActivityName(A, day, i, activityName)] = ""
				B[day].activity[searchActivityName(A, day, i, activityName)] = ""
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
	fmt.Printf("    Kegiatan %v\n", activityName)
	fmt.Print("    berhasil dihapus.                 \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}
