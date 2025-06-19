package main

import "fmt"

func main() {

	qwe()

	zx, xc := 54, 4
	sum := Sum(zx, xc)
	fmt.Println(sum)

	// возвращает удвоенное значение `x`
	x := 2
	qaz := double(x)
	fmt.Println(qaz)

	//возвращает `true`, если `age >= 18`, иначе `false`.
	age := 5

	if isAdult(age) {
		fmt.Println("true")
	} else {
		fmt.Println("false")

	}

	// большее из двух чисел
	fmt.Println(max(5, 10))

	// "отлично"
	fmt.Println(grade(95))

	//Считает индекс массы тела
	fmt.Println(bmi(50, 1.75))

	fmt.Println(isLeap(2001))

	personName := "oleg"
	fmt.Println(
		greet(personName, false),
	)

	fmt.Println(calc(10, 5, "*"))

}

func qwe() {
	fmt.Println("Hello World")
}

func Sum(zx, xc int) int {
	sum := zx + xc
	return sum
}

// возвращает удвоенное значение `x`
func double(x int) int {
	return x * 2
}

// возвращает `true`, если `age >= 18`, иначе `false`.
func isAdult(age int) bool {
	if age >= 18 {
		return true

	} else {
		return false
	}

}

// большее из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

// "отлично"
func grade(score int) string {
	if score >= 90 && score <= 100 {
		return "отлично"
	} else if score >= 75 {
		return "хорошо"
	} else if score >= 60 {
		return "удовлетворительно"
	} else if score >= 0 {
		return "плохо"
	}
	return "некорректный балл"
}

// Считает индекс массы тела
func bmi(weight float64, height float64) string {
	bim := weight / (height * height)
	if bim < 18.5 {
		return "недовес"
	} else if bim > 18.5 && bim < 24.9 {
		return "норма"
	} else if bim > 25 && bim < 29.9 {
		return "изб. вес"
	} else if bim > 30 {
		return "ожирение"
	} else {
		return "некорректный"
	}

}

func isLeap(year int) bool {
	return year%4 == 0
}

func greet(name string, isMorning bool) string {
	if isMorning {
		return "доброе утро " + name
	}
	return "привет " + name
}

func calc(a, b int, op string) int {
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "/" {
		if b == 0 {
			return 0
		}
		return a / b
	}
	return 0
}
