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

func insertActivity(day, time, timeLength int, activityName string) {
	/*
		I.S. Terdapat nilai dari hari dan waktu yang dipilih, lamanya waktu, dan nama kegiatan pribadi
		F.S. Mengisi kegiatan pribadi ke dalam array
	*/

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
		selectTimeActivityMenu(day)
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
		fmt.Print("    Kegiatan berhasil diisikan        \n")
		fmt.Print("    pada agenda dengan waktu yang     \n")
		fmt.Print("    dipilih.                          \n")
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")
		mainMenu()
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
		selectTimeActivityMenu(day)
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
	fmt.Print("                                      \n")
	fmt.Print("    Kegiatan/rapat berhasil           \n")
	fmt.Print("    dihapus pada agenda dengan        \n")
	fmt.Print("    waktu yang dipilih.               \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteActivityByName(day int, activityName string) {
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
	fmt.Printf("    pada hari %v berhasil         \n", dayString(day))
	fmt.Print("    dihapus.                          \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteAllActivity() {
	var A weekAgenda
	var i, j int

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	for i = 0; i < 5; i++ {
		for j = 0; j < 9; j++ {
			A[i].activity[j] = ""
			A[i].identifier[j] = 0
		}
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
	} else {
		secondManagerAgenda = A
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Semua kegiatan/rapat berhasil     \n")
	fmt.Print("    dihapus.                          \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteAllActivityByDay(day int) {
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
	fmt.Printf("    %v berhasil dihapus.              \n", dayString(day))
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	mainMenu()
}

func deleteAllActivityByName(activityName string) {
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
	mainMenu()
}
