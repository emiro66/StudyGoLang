package main

import (
	"fmt"
)

type People struct {
	name string
	age  int
}

func main() {
	//crw()
	//slices()
	//variadicFunctions()
	//convertPointer()
	getSlice()
}

// массив
func crw() {
	var intArr [3]int //длина (len), вместимость (cap)

	intArr[0] = 6 //c нуля счет
	intArr[1] = 10
	fmt.Println(intArr)

	qwe := [...]People{ // не обязательно указывать число, можно указать ... и добавлять больше
		{
			name: "катя",
			age:  14,
		}, {
			name: "оля",
			age:  17,
		}, {
			name: "оля",
			age:  17,
		}, {
			name: "оля",
			age:  17,
		},
	}
	fmt.Println(qwe)

	//длина (len), вместимость (cap)

	stringArr := [...]string{
		"оля", "катя", "ева",
	}
	fmt.Println(stringArr) //[оля катя ева]

	fmt.Println(len(stringArr), cap(stringArr)) //3 3

	for i := 0; i < len(stringArr); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, stringArr[i])
	}

	for inx, value := range stringArr {
		fmt.Println(inx, value)
	}
	for _, value := range stringArr {
		fmt.Println(value)
	}

	//интеграция по массиву

	newIntArr := changeArray(&intArr) // по указателю
	fmt.Println(intArr)
	fmt.Println(newIntArr)

}

// интеграция по массиву
func changeArray(arr *[3]int) [3]int {
	arr[2] = 5
	return *arr
}

// слайсы как массивы, но без элемента
func slices() { //слайс (динамический массив) ссылается на массив, а не создает
	var defaultSlice []int //дефолтное значение слайсов - nil
	fmt.Println(len(defaultSlice), cap(defaultSlice))

	stringSlice := []string{"оля", "катя", "маша"}
	fmt.Println(stringSlice)

	sliceByMake := make([]int, 0, 5) //0 - длина, 5 - вместимость
	fmt.Println(sliceByMake)
	fmt.Printf("len: %d, Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5) //append - функ, можно добавить элементы с помощью нее в слайс
	fmt.Println(sliceByMake)
	fmt.Printf("len: %d, Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 6) // если добавить элемент, вместимость увеличится в 2 раза
	fmt.Println(sliceByMake)
	fmt.Printf("len: %d, Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	for ind, value := range sliceByMake {
		fmt.Println(ind, value)
	}
}

// слайсы2

func variadicFunctions() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6)

	firstSlices := []int{1, 2, 3, 4, 5}
	//secondSlices := []int{4, 3, 2, 1}
	showAllElements(firstSlices...)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println("Value: ", val)
	}
	fmt.Println()
}

//	type _slices struct {
//		elements unsafe.Pointer - указательь на массив с определенным типом данных
//		len int
//
// cap int
// }
func convertPointer() {
	initialSlice := []int{1, 2}
	fmt.Println(initialSlice)

	intArray := (*[2]int)(initialSlice) //указатель на массив, (*[2]int) - массив с 2 элем, обязательно что бы длина массива совпадала с длиной слайса
	fmt.Println(intArray)

	changeValue(initialSlice)
	fmt.Println(initialSlice)

	newSlice := append(initialSlice, 3) // добавили элемент, capaсity увеличивается на 2
	fmt.Println(newSlice)

	newSlice2 := appendValue(newSlice)
	fmt.Println(newSlice)
	fmt.Printf("len: %d, Capacity: %d\n", len(newSlice2), cap(newSlice2))
}

func changeValue(slice []int) {
	slice[1] = 15
}

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5, 7, 5, 8, 4)
	fmt.Println(slice)
	return slice
}

// получение слайса из массива и из слайса
func getSlice() {
	intArr := []int{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	intSlice := intArr[1:3] //выводит знач в заданном диапазоне, но последний не включается
	fmt.Println(intSlice)

	fullSlice := intArr[:] // выводит все значения
	fmt.Println(fullSlice)
	//
}
