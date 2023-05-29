// Lain
// identifier 0 = kegiatan pribadi
// identifier 1 = rapat bersama
// day 0,1,2,3,4 = hari senin, selasa, rabu, kamis, jumat
// time 0 = 08.00-09.00
// time 1 = 09.00-10.00
// time 2 = 10.00-11.00
// dst.

package main

import "fmt"

func dayString(day int) string {
	if day == 0 {
		return "Senin"
	} else if day == 1 {
		return "Selasa"
	} else if day == 2 {
		return "Rabu"
	} else if day == 3 {
		return "Kamis"
	} else {
		return "Jumat"
	}
}

func showAllAgenda() {
	var A weekAgenda
	var i, j, k int

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")

	for i = 0; i < 5; i++ {
		k = 8

		fmt.Printf("       Agenda pada hari %v:\n\n", dayString(i))
		fmt.Print("        Waktu           Kegiatan\n")

		for j = 0; j < 9; j++ {
			if k+1 < 10 {
				fmt.Printf("    0%v.00 - 0%v.00 | %v\n", k, k+1, A[i].activity[j])
			} else if k+1 == 10 {
				fmt.Printf("    0%v.00 - %v.00 | %v\n", k, k+1, A[i].activity[j])
			} else {
				fmt.Printf("    %v.00 - %v.00 | %v\n", k, k+1, A[i].activity[j])
			}
			k++
		}

		fmt.Print("\n")
	}

	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
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

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showAllAgenda()
	}
}

func showDailyAgenda(day int) {
	var A weekAgenda
	var i int
	var j int = 8

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")
	fmt.Printf("       Agenda pada hari %v:\n\n", dayString(day))
	fmt.Print("        Waktu           Kegiatan\n")

	for i = 0; i < 9; i++ {
		if j+1 < 10 {
			fmt.Printf("    0%v.00 - 0%v.00 | %v\n", j, j+1, A[day].activity[i])
		} else if j+1 == 10 {
			fmt.Printf("    0%v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		} else {
			fmt.Printf("    %v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		}
		j++
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
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

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(day)
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
	}
}

func peekDailyAgenda(day int) {
	var A weekAgenda
	var i int
	var j int = 8

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("\n")
	fmt.Printf("       Agenda pada hari %v:\n\n", dayString(day))
	fmt.Print("        Waktu           Kegiatan\n")

	for i = 0; i < 9; i++ {
		if j+1 < 10 {
			fmt.Printf("    0%v.00 - 0%v.00 | %v\n", j, j+1, A[day].activity[i])
		} else if j+1 == 10 {
			fmt.Printf("    0%v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		} else {
			fmt.Printf("    %v.00 - %v.00 | %v\n", j, j+1, A[day].activity[i])
		}
		j++
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
}

func showTimeAgenda(day, time int) {
	var A weekAgenda

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")
	fmt.Print("\n")
	if selectedManager == 1 {
		fmt.Print("         [ MANAJER PERTAMA  ]         \n")
	} else {
		fmt.Print("          [ MANAJER KEDUA  ]          \n")
	}
	fmt.Print("\n")
	fmt.Printf("       Agenda pada hari %v:\n\n", dayString(day))
	fmt.Print("        Waktu           Kegiatan\n")

	if time+9 < 10 {
		fmt.Printf("    0%v.00 - 0%v.00 | %v\n", time+8, time+9, A[day].activity[time])
	} else if time+9 == 10 {
		fmt.Printf("    0%v.00 - %v.00 | %v\n", time+8, time+9, A[day].activity[time])
	} else {
		fmt.Printf("    %v.00 - %v.00 | %v\n", time+8, time+9, A[day].activity[time])
	}

	fmt.Print("\n")
	fmt.Print("*------------------------------------*\n")

	fmt.Print("\n")
	fmt.Print("*----- PENCATAT AGENDA KEGIATAN -----*\n")
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

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		selectTimeAgendaMenu(day)
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, time)
	}
}
