// Rapat bersama
// identifier 0 = kegiatan pribadi
// identifier 1 = rapat bersama
// day 0,1,2,3,4 = hari senin, selasa, rabu, kamis, jumat
// time 0 = 08.00-09.00
// time 1 = 09.00-10.00
// time 2 = 10.00-11.00
// dst.

package main

import "fmt"

func insertMeeting(day, time, timeLength int, activityName string) {
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
		selectTimeMeetingMenu(day)
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
		if time+8+timeLength < 10 {
			fmt.Printf("    dengan waktu 0%v.00 - 0%v.00.\n", time+8, time+8+timeLength)
		} else if time+8+timeLength == 10 {
			fmt.Printf("    dengan waktu 0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else if time+8 < 10 {
			fmt.Printf("    dengan waktu 0%v.00 - %v.00.\n", time+8, time+8+timeLength)
		} else {
			fmt.Printf("    dengan waktu %v.00 - %v.00.\n", time+8, time+8+timeLength)
		}
		fmt.Print("                                      \n")
		fmt.Print("*------------------------------------*\n")
		selectTimeMeetingMenu(day)
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
		selectTimeMeetingMenu(day)
	}
}
