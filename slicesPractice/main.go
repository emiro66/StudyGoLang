package main

import (
	"errors"
	"fmt"
)

func main() {
	//allNumbers()
	//allStrings()
	//floats()
	allMaps()
}

//СЛАЙСЫ

// числа
func allNumbers() {
	slices()

	numbers := []int{-1, 5, 6, 8, 4, 16, 9, 10, 2, 3} // создаёт слайс из 10 целых чисел
	fmt.Println(numbers)
	fmt.Println(findMinMax(numbers))

}

func findMinMax(numbers []int) (int, int) { //находит максимум и минимум
	max := numbers[0]
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min, max
}

//if max < numbers[2] {
//	max = numbers[2]
//}
//if max > numbers[3] {
//	max = numbers[3]
//}

// создаёт новый слайс только из чётных чисел
func slices() {
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numSlice); i++ {
		if i%2 != 0 {
			fmt.Println(numSlice[i])
		}
	}
}

// строки

func allStrings() {
	fmt.Println(strings([]string{"Apple", "banana", "Cat", "dog", "Elephant", " "}))
}

// строки длиной больше 3 символов, начинаются с заглавной буквы
func strings(input []string) (result []string) {
	for _, s := range input {
		if len(s) > 3 && s[0] >= 'A' && s[0] <= 'Z' { //если длина s больше 3 символов и s [0] (s первая буква) начин с большой )
			result = append(result, s) // то добавляется
		}
	}
	return
}

//МАССИВЫ

func floats() {

	// создаёт массив из 5 чисел с плавающей точкой
	float32s()

	fmt.Println(isAllTrue([]bool{true, true}))

}

// создаёт массив из 5 чисел с плавающей точкой
func float32s() {
	num := [5]float32{1.2, 3.4, 5.6, 7.8, 9.1} //создали массив

	fmt.Println(num)

	var sum float32 = 0.0       //sum равен 0.0
	for _, value := range num { // создаем перем. value и через range перебираю каждое значение массива
		value *= 1.1 //умножаем каждое знач на 1.1 (10%)
		fmt.Println(value)
		sum += value //массив = массив + value
	}
	average := sum / float32(len(num)) //считает среднее значение
	fmt.Println(average)

}

func isAllTrue(slice []bool) string {
	for _, v := range slice {
		if !v { // ! - true на false и наоборот (инвертация)
			return "что-то не так"
		}

	}
	return "все ок"
}

// МАПЫ
func allMaps() {
	ballMap()

}

func ballMap() {
	scores := map[string]int{ // создание мапы (перем присваиваетя мар[ключ]условие )
		"Александр": 57,
		"Маша":      63,
		"Алиса":     75,
		"Инокентий": 91,
		"Алла":      88,
	}

	//вывод
	for name, score := range scores {
		fmt.Printf("Имя: %s, балы: %d\n", name, score)
	}

	//вывести имя студента с макс. балом
	fmt.Println(maxScore(scores))

	// добавление студента
	fmt.Println(appendStudent(scores, "Александр", 77))

	// удаление студента
	fmt.Println(deleteStudent(scores, "Александр", 88))

}

// вывести имя студента с макс. балом
func maxScore(scoreMap map[string]int) string {
	max := 0                     //создаю переменную (копилку)
	maxScoreName := ""           //создаю переменную, куда записывается имя студента с макс. балом
	for k, v := range scoreMap { //с помощью range перебираю каждое значение мапы
		if max < v { //если максимальное меньше v (v - бал)
			max = v          //то максимальное становится v
			maxScoreName = k //соответственно сюда попадает имя студента с макс. балом
		}
	}
	return maxScoreName //после завершения цикла возвр. имя студента с макс. балом
}

// добавление студента
func appendStudent(studentMap map[string]int, name string, score int) (map[string]int, error) {
	_, exists := studentMap[name] //создаю переменную exists и присваиваю ей мапу с арг name
	if exists {
		return nil, errors.New("студент уже существует") //если else, тогда возвр. ошибка (nil для _)

	}
	studentMap[name] = score // иначе добавляется имя и бал
	return studentMap, nil
}

// удаление студента
func deleteStudent(studentMap map[string]int, name string, score int) (map[string]int, error) {
	_, exists := studentMap[name]
	if !exists {

		return nil, errors.New("студент не найден ")
	}
	delete(studentMap, name)
	return studentMap, nil
}
