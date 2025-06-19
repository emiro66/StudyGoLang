package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Принимает день недели и возвращает
	fmt.Println(dayType("понедельник"))

	// Принимает число от 1 до 5. Возвращает
	fmt.Println(rateGrade(6))

	// Использовать цикл `for`. Выводить числа в одну строку через пробел
	for i := 1; i < 11; i++ {
		fmt.Print(i, "  ")
	}
	// Найти и вывести сумму всех чётных
	plus := 1
	for plus < 100 {
		plus += plus
		fmt.Println(plus)
	}

	// Вывести числа от `1` до `30`, пропуская все кратные 3
	skipMultiples()

	//Сумма чётных чисел от 1 до 100:
	ggg := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			ggg += i
		}
	}
	fmt.Println("", ggg)

	fmt.Println(findFirstDivisible(100))

	fmt.Print(oneToTenSpaced())

}

// Принимает день недели и возвращает
func dayType(day string) string {
	switch {
	case day == "вторник" || day == "понедельник" || day == "чертверг" || day == "среда" || day == "пятница":
		return "будни"
	case day == "cуббота" || day == "воскресенье":
		return "выходные"
	default:
		return "неизвестно"

	}
}

// Принимает число от 1 до 5. Возвращает:
func rateGrade(score int) string {
	switch {
	case score == 5:
		return "Отлично"
	case score == 4:
		return "Хорошо"
	case score == 3:
		return "Удовлетворительно"
	case score >= 1 && score <= 2:
		return "Плохо"
	default:
		return "Ошибка"
	}

}
func skipMultiples() {
	for j := 1; j <= 30; j++ {
		if j%3 == 0 {
			fmt.Println(j)
		}

	}
}

func oneToTenSpaced() string {
	const max = 10
	buf := ""
	for i := 1; i <= max; i++ {
		con := strconv.Itoa(i)
		buf += con
		if i < max {
			buf += " "
		}
	}
	return buf
}

func findFirstDivisible(limit int) int {
	for i := 1; i <= limit; i++ {
		if i%5 == 0 && i%7 == 0 {
			return i
		}
	}
	return 0
}
