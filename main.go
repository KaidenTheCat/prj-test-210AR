package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"prj_test/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 100
	pointsPerQuestion = 20
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі!")

	users := getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}

	sortAndSave(users)
	for {
		menu()

		choise := ""
		fmt.Scan(&choise)

		switch choise {
		case "1":
			user := play()
			users = getUsers()
			users = append(users, user)
			sortAndSave(users)
		case "2":
			users = getUsers()
			if len(users) == 0 {
				terCl()
				fmt.Println("Рейтинг порожній")
			} else {
				terCl()
				for _, u := range users {
					fmt.Printf(
						"Id: %v, Name: %s, Time: %v\n",
						u.Id, u.Name, u.TimeSpent,
					)
				}
			}
		case "3":
			terCl()
			fmt.Println("Ви впевнені?")
			fmt.Println("1. Так\n2. Ні")
		label:
			for {
				confirm := ""
				fmt.Scan(&confirm)

				switch confirm {
				case "1":
					terCl()
					os.Truncate("users.json", 0)
					fmt.Println("Рейтинг очищено")
					break label
				case "2":
					break label
				default:
				}
			}
		case "4":
			return
		default:
		}
	}

}

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Очистити рейтинг")
	fmt.Println("4. Вийти")
}

func play() domain.User {
	TimeStart := time.Now()
	myPoints := 0

	for myPoints < totalPoints {
		example := rand.Intn(5) + 1
		switch example {
		case 1:
			x, y := rand.Intn(9), rand.Intn(9)
			fmt.Printf("\n%v + %v = ", x, y)

			ans := ""
			fmt.Scan(&ans)

			ansInt, err := strconv.Atoi(ans)

			if err != nil {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x+y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 2:
			x, y := rand.Intn(9), rand.Intn(9)
			fmt.Printf("\n%v - %v = ", x, y)

			ans := ""
			fmt.Scan(&ans)

			ansInt, err := strconv.Atoi(ans)

			if err != nil {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x-y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 3:
			x, y := rand.Intn(9), rand.Intn(9)
			fmt.Printf("\n%v * %v = ", x, y)

			ans := ""
			fmt.Scan(&ans)

			ansInt, err := strconv.Atoi(ans)

			if err != nil {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x*y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 4:
			x, y := rand.Intn(9), rand.Intn(9)
			fmt.Printf("\n%v / %v = ", x, y)

			ans := ""
			fmt.Scan(&ans)

			ansInt, err := strconv.Atoi(ans)

			if err != nil {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x/y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 5:
			x, y, z := rand.Intn(9), rand.Intn(9), rand.Intn(6)
			fmt.Printf("\n(%v + %v) * %v = ", x, y, z)

			ans := ""
			fmt.Scan(&ans)

			ansInt, err := strconv.Atoi(ans)

			if err != nil {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == (x+y)*z {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		}
	}

	TimeFinish := time.Now()
	timeSpent := TimeFinish.Sub(TimeStart)

	terCl()
	fmt.Printf("Ваш час: %v", timeSpent)
	fmt.Print("Введіть ваше ім'я: ")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:        id,
		Name:      name,
		TimeSpent: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i int, j int) bool {
		return users[i].TimeSpent < users[j].TimeSpent
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Printf("sortAndSave(os.OpenFile): %s", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("sortAndSave(file.Close()): %s", err)
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		log.Printf("sortAndSave(encoder.Encode)): %s", err)
		return
	}
}

func getUsers() []domain.User {
	var users []domain.User
	file, err := os.Open("users.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = os.Create("users.json")
			if err != nil {
				log.Printf("getUsers(os.Create): %s", err)
			}
			return nil
		}
		log.Printf("getUsers(os.Open): %s", err)
		return nil
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		log.Printf("getUsers(decoder.Decode): %s", err)
	}

	return users
}

func terCl() {
	fmt.Print("\033[H\033[2J")
}
