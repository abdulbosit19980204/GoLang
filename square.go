package main

import (
	"fmt"
	"sort"
	"square/yuza"
	"time"
)

func hajm(a, b, c int) int {
	return a * b * c

}

func Student(Name, Surname string, sage int, Smestrniypodi bool) (string, string, int, bool) {
	return Name, Surname, sage, Smestrniypodi
}

//else if

func ifshart() {
	ism := "Robiya"

	if ism == "Abdulbosit" {
		fmt.Println("Tuychiyev")
	} else if ism == "Robiya" {
		fmt.Println("Tuychiyeva")
	} else {
		fmt.Println("Bunday shaxs mavjud emas")
	}
}

func funcSwitch() {
	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshaba")
	default:
		fmt.Println("Bunday xafta kuni yo'q")
	}

}

func farray() {
	var userarray [4]string
	userarray[0] = "Familya"
	userarray[1] = "Ism"
	userarray[2] = "time.Now()"
	userarray[3] = "Ish joyi"

	fmt.Println("|", userarray[0], "|")
	fmt.Println("2. ", userarray[1])
	fmt.Println("3. ", userarray[2])
	fmt.Println("Tug'ilgan sana:  ", userarray[3])

}
func pluralarray() {
	parray := [4][4]string{{"1", "2", "3", "5"},
		{"1", "2", "3", "5"},
		{"1", "8", "3", "4"},
		{"7", "2", "3", "0"}}
	fmt.Println(parray[1][1])

	for j := 0; j < 4; j++ {
		fmt.Printf("%v\n", parray[j])
	}

	//diagnalini chiqarish
	for i := 0; i < 4; i++ {

		for t := 0; t < 4; t++ {

			if i == t {
				fmt.Printf("%v ", parray[i][t])
			} else {
				// fmt.Printf("%v ", parray[i][t])
				fmt.Printf("%v ", "*")
			}
		}
		fmt.Println()
	}

	//teskari tartibda
	fmt.Println()

	for i := 0; i < 4; i++ {

		for t := 0; t < 4; t++ {

			if i == t {
				fmt.Printf("%v ", "*")
			} else {
				// fmt.Printf("%v ", parray[i][t])

				fmt.Printf("%v ", parray[i][t])
			}
		}
		fmt.Println()
	}

	//diagnalni teskari qilish

	//teskari tartibda
	fmt.Println()

	for i := 3; i >= 0; i-- {

		for t := 3; t >= 0; t-- {

			if i == t {
				fmt.Printf("%v ", parray[i][t])
			} else {
				// fmt.Printf("%v ", parray[i][t])
				fmt.Printf("%v ", "*")

			}
		}
		fmt.Println()
	}
	fmt.Println()
	// diagnal [1][3] xolda

	for i := 0; i < 4; i++ {

		for t := 0; t < 4; t++ {

			if i+t == 3 {
				fmt.Printf("%v ", parray[i][t])
				//fmt.Printf("%v ", "*")
			} else {
				//fmt.Printf("%v ", parray[i][t])
				fmt.Printf("%v ", "*")
			}
		}
		fmt.Println()
	}

}

// array elemntini o'zgartirish nusxalash bilan o'zgarmaydi

func charray() {
	chan1arry := [3]int{2, 8, 11}

	chanarray := chan1arry
	fmt.Println(chan1arry)
	fmt.Println(chanarray)
	fmt.Println()
	chan1arry[0] = 7
	fmt.Println(chan1arry)
	fmt.Println(chanarray)

}

//xotiraga nusxa olish bilan reference bilan bunda qator lar bitta yacheykaga qarab turadi

func xarray() {
	xan1arry := [3]int{2, 8, 11}

	xanarray := &xan1arry
	fmt.Println(xan1arry)
	fmt.Println(*xanarray)
	fmt.Println()
	xan1arry[2] = 7
	fmt.Println(xan1arry)
	fmt.Println(*xanarray)

}

// faqat go uchun slice cheksiz massiv xosil qilish uzinliksiz

func myslice() {
	slicearry := []int{5, 6, 7, 8, 9, 10}
	slicearry = append(slicearry, 25)

	fmt.Println("array uzunligi l = ", len(slicearry))
}

// slice dan elemntlarni ajratib olish
func catslice() {
	catslarray := []string{"olma", "anor", "nok", "bodring"}
	mycatslarry := catslarray[1:3]
	fmt.Println(mycatslarry)
	sort.Strings(catslarray)
	fmt.Println()
	fmt.Println(catslarray)
}

//map turi map[KeyType]ValueType shu strukturada ishlaydi

func mapfunc() {
	statuses := make(map[string]int)
	statuses["online"] = 1
	statuses["offline"] = 2
	statuses["joined"] = 3
	statuses["left"] = 4

	//mapdan ma'lumot o'qish
	var onlineStatus = statuses["online"]
	fmt.Println()

	delete(statuses, "online")
	//var onlineStatus = statuses["offline"]
	fmt.Println(onlineStatus)
}
func forloop() {
	for i := 0; i < 8; i++ {
		fmt.Println("Bu for loopdan odatiy xolda ishlatish i = ", i)
	}
	// for {
	// 	fmt.Println("Alhamdulillah " + strconv.Itoa(time.Now().Second()))
	// 	time.Sleep(1 * time.Second)
	// }

	K := 1
	for K < 10 {
		fmt.Println("K = ", K)
		K++
	}
}

func arryfor() {
	fmt.Println()
	forarray := []string{"nol", "bir", "ikki", "uch", "to'rt"}
	for index, value := range forarray {

		fmt.Println(index, value)
	}
	fmt.Println()
	formap := map[int]string{
		1: "Kompyuter",
		2: "Notebook",
		3: "Netbook",
		4: "Phone"}

	for key, value := range formap {
		fmt.Println(key, value)
	}
}

func main() {

	fmt.Println(yuza.Yuzax(10, 5))
	fmt.Println("Kubning hajmi V = ", hajm(10, 3, 2))
	fmt.Println(Student("Abdulbosit", "tuychiyev", 24, false))
	i := 8
	if i == 8 {
		fmt.Println("ishladi")
	}
	ifshart()
	funcSwitch()
	fmt.Println(time.Now().Weekday())
	farray()
	pluralarray()
	charray()
	xarray()
	myslice()
	catslice()
	mapfunc()
	forloop()
	arryfor()
}
