package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // рамдомне число в кожен період часу
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 // рандомне число (target - ціль. вданому випадку - число))
	fmt.Println("Я обираю число від 1 до 100")
	fmt.Println("Число вибрано")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) // виклик ко ристувача
	success := false                    // success- успіх

	//Цикл

	for guesses := 0; guesses < 10; guesses++ { // цикл ( початкове знаення/кінцеве/ крок)
		//guesses -припущення
		fmt.Println("У вас залишилось: ", 10-guesses, "спроб")

		fmt.Print("Напишіть ваше чило:  ")

		input, err := reader.ReadString('\n') // читаємо рядок введення користувачем
		// до переведення рядка на наступний
		if err != nil {
			log.Fatal(err) // виводимо помилку і завершуємо програму
		}
		input = strings.TrimSpace(input) //перезаписуємо інпут, бо вище він має
		// ще додатковий символ ('\n') -переводу рядка, а нам того не треба
		//TrimSpase - прибирає всі лишні символи, табуляції.

		guess, err := strconv.Atoi(input) // щоб перетворити рядок в ціле число int
		if err != nil {
			log.Fatal(err) // якщо помилка - то показати її
		}
		if guess > target {
			fmt.Println("Ваше число більше ніж загадане")
		} else if guess < target {
			fmt.Println("Ваше число менше ніж загадане")
		} else {
			success = true
			fmt.Println("Вітаю!!!")
			break //щоб зупинити цикл
		}
	}
	if !success {
		fmt.Println("Ваші спроби закінчились!Це було число - ", target)
	}
}
