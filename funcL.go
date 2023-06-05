package main

import "fmt"

func dayString(day int) string {
	//mengembalikan string hari dari indeks

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
	/* IS -
	FS menampilkan seluruh agenda */

	var A weekAgenda
	var i, j, k int

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
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
	/* IS terpilih hari yang ditentukan
	FS menampilkan data agenda pada hari yang dipilih */

	var A weekAgenda
	var i int
	var j int = 8

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
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

	if selectedMenu != "9" && selectedMenu != "0" {
		fmt.Printf("\x1bc")
		showDailyAgenda(day)
	}
}

func showTimeAgenda(day, time int) {
	/* IS terpilih hari dan waktu yang ditentukan
	FS menampilkan data agenda pada hari dan waktu yang dipilih */

	var A weekAgenda

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	fmt.Print("\n")
	fmt.Print("*-------------- AGENDA --------------*\n")
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

	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		showSelectTimeMenu(day)
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		showTimeAgenda(day, time)
	}
}

func searchActivityLength(A weekAgenda, day, i int) int {
	//mengembalikan durasi kegiatan dari agenda pada hari yang ditentukan

	var status bool
	var tempName string
	var j int
	tempName = A[day].activity[i]
	for i < 9 && !status {
		i++
		if i < 9 {
			status = tempName != A[day].activity[i]
		} else {
			status = true
		}
		j++
	}
	return j
}

func optimizeAgenda(day int) {
	/* IS terpilih hari yang ditentukan
	FS mengoptimalkan data agenda pada hari yang dipilih */

	var i, j, k, tempA, tempB int
	var A weekAgenda
	var status bool

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	for i < 9 {
		if A[day].activity[i] == "" {
			tempA = searchActivityLength(A, day, i)
			if i+tempA < 9 {
				j = i + tempA
				if A[day].activity[j] != "" && A[day].identifier[j] == 0 {
					A[day].activity[i] = A[day].activity[j]
					A[day].activity[j] = ""
				}
			}
		}
		i++
	}

	i = 0

	for i < 9 {
		if A[day].activity[i] == "" {
			tempA = searchActivityLength(A, day, i)
			j = i + tempA
			status = false
			for j < 9 && !status {
				tempB = searchActivityLength(A, day, j)
				if A[day].activity[j] != "" && A[day].identifier[j] == 0 {
					if tempB <= tempA {
						for k = i; k < i+tempB; k++ {
							A[day].activity[k] = A[day].activity[j]
						}
						for k = j; k < j+tempB; k++ {
							A[day].activity[k] = ""
						}
						status = true
					}
				}
				j = j + tempB
			}
		}
		i++
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
	} else {
		secondManagerAgenda = A
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Printf("    Agenda pada hari %v telah\n", dayString(day))
	fmt.Print("    dioptimalkan.                     \n")
	fmt.Print("                                      \n")
	fmt.Print("*------------------------------------*\n")
	showDailyAgenda(day)
	if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	}
}

func optimizeAllAgenda() {
	/* IS -
	FS mengoptimalkan seluruh data agenda */

	var A weekAgenda
	var i, j, k, l, m, tempA, tempB int
	var status bool

	if selectedManager == 1 {
		A = firstManagerAgenda
	} else {
		A = secondManagerAgenda
	}

	for i < 5 {
		j = 0
		for j < 9 {
			if A[i].activity[j] == "" {
				tempA = searchActivityLength(A, i, j)
				if j+tempA < 9 {
					k = j + tempA
					if A[i].activity[k] != "" && A[i].identifier[k] == 0 {
						A[i].activity[j] = A[i].activity[k]
						A[i].activity[k] = ""
					}
				}
			}
			j++
		}
		i++
	}

	i = 0
	j = 0
	k = 0

	for i < 5 {
		j = 0
		for j < 9 {
			if A[i].activity[j] == "" {
				tempA = searchActivityLength(A, i, j)
				k = i
				l = j + tempA
				status = false
				for k < 5 && !status {
					if i < k {
						l = 0
					}
					for l < 9 {
						tempB = searchActivityLength(A, k, l)
						if A[k].activity[l] != "" && A[k].identifier[l] == 0 {
							if tempB <= tempA {
								for m = j; m < j+tempB; m++ {
									A[i].activity[m] = A[k].activity[l]
								}
								for m = l; m < l+tempB; m++ {
									A[k].activity[m] = ""
								}
								status = true
							}
						}
						l = l + tempB
					}
					k++
				}
			}
			j++
		}
		i++
	}

	if selectedManager == 1 {
		firstManagerAgenda = A
	} else {
		secondManagerAgenda = A
	}

	fmt.Print("\n")
	fmt.Print("*------------ PERINGATAN ------------*\n")
	fmt.Print("                                      \n")
	fmt.Print("    Semua agenda telah                \n")
	fmt.Print("    dioptimalkan.                     \n")
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
	} else if selectedMenu == "9" {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	} else if selectedMenu == "0" {
		fmt.Printf("\x1bc")
		mainMenu()
	} else {
		fmt.Printf("\x1bc")
		optimizeAgendaMenu()
	}
}
