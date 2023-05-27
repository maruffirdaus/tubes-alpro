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

func changeActivity(A *weekAgenda, day, time int) {

}

func deleteActivity(A *weekAgenda, day, time int) {

}
