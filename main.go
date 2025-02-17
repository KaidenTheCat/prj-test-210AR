package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       = 100
	pointsPerQuestion = 20
)

func main() {
	fmt.Println("Вітаємо у грі ---!")

	for i := 5.0; i > 0; i -= 0.1 {
		fmt.Printf("\rДо початку: %v", i)
		time.Sleep(100)
		if i == 1 {
			fmt.Print("\rГра почалась!")
		}
	}

	myPoints := 0

	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("\n%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)

		if err != nil {
			fmt.Println("Невдале знаяення, давай по новой")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У вас %v очок!", myPoints)
			} else {
				fmt.Println("Не праивльно!")
			}
		}
	}

}
