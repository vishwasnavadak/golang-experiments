package main

import "fmt"

func main() {
	flag := 1
	var dd, mm, yyyy int
	fmt.Println("Enter a date(dd mm yyyy): ")
	fmt.Scanln(&dd)
	fmt.Scanln(&mm)
	fmt.Scanln(&yyyy)
	if mm <= 0 || mm > 12 || yyyy <= 0 {
		flag = 0
	}
	switch mm {
	case 1, 3, 5, 7, 8, 10, 12:
		if dd > 31 {
			flag = 0
		}
	case 4, 6, 9, 11:
		if dd > 30 {
			flag = 0
		}
	case 2:
		if yyyy%4 == 0 {
			if dd > 29 {
				flag = 0
			}

		} else if dd > 28 {
			flag = 0
		}

	}
	if flag == 1 {
		fmt.Println("Date is valid")
	} else {
		fmt.Println("Date is invalid")
	}
	fmt.Printf("Entered values: %d %d %d", dd, mm, yyyy)

}
