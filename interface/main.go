package main

import "fmt"

type People struct {
}
type Pig struct {
}
type Cow struct {
}

type Noisers interface {
	makeNoise() string
}

func main() {
	pig := Pig{}
	cow := Cow{}
	people := People{}

	fmt.Println(pig, cow, people)
	Greet(cow)

}

func Greet(n Noisers) {

	fmt.Println(n.makeNoise())
}

func (p Pig) makeNoise() string {
	return "хрю хрю хрю"
}

func (pe People) makeNoise() string {
	return "Э БЛЯТЬ"
}

func (c Cow) makeNoise() string {
	return "МУ МУ МУ"
}
