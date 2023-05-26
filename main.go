// Program utama

package main

type agenda struct {
	activity [9]string
}

type weekAgenda [5]agenda

var firstManagerAgenda weekAgenda
var secondManagerAgenda weekAgenda

func main() {
	selectManager()
}
