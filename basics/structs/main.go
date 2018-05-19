package main

import (
	"fmt"
)

func main() {
	personalTyping()
	person()
	anonymousStructs()
	embeddingStructs()
}

type human struct {
	name string
	age int
	email string
}

type superHero struct {
	human
	superPowers bool
}

func (s superHero) getSuperHero() string {
	fmt.Println(s.name)
	return s.name
}

func person(){
 	iam := human{
 		"Manu Cutillas",
 		35,
 		"manucutillas@outlook.com",
 	}
	fmt.Println(iam)
 	fmt.Println(iam.name)
 	fmt.Println(iam.age)
 	fmt.Println(iam.email)

 	imDeadpool := superHero{
 		human : human{"DeadPool", 35, "nothing"},
 		superPowers:true,
	}

	superHeroName := imDeadpool.getSuperHero()

	fmt.Println("The superhero is:", superHeroName)

	imDeadpool.wonderWoman("Wonder Woman")
	fmt.Println("Memory Reference of imDeadpool hero -> ", &imDeadpool.name)
	fmt.Println("The new super hero is: ", imDeadpool)

	fmt.Println(imDeadpool)
 	fmt.Println(imDeadpool.name)
 	fmt.Println(imDeadpool.superPowers)
	fmt.Println(imDeadpool.human.email)


	batman := imDeadpool.changeSuperHero("Batman")
	fmt.Println("Chananana Batman, Chananana ", batman.name)
	fmt.Println("Memory Reference of batman hero -> ", &batman.name)
}


func personalTyping() {
	type myType int

	var a int
	var b int
	var c myType

	a = 10
	b = 10

	c = 30

	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(a + int(c))
}

func (d *superHero) wonderWoman(super string) {
	d.name = super
}

func (d superHero) changeSuperHero(super string) superHero {
	d.name = super
	return d
}

func anonymousStructs() {
	actor := struct{
		name string
		weapons bool
	}{"Chuck Norris", true}

	fmt.Println("actor: ", actor)

}

type MyString struct {
	value string
}

func NewMyString(s string) MyString {
	return MyString{value:s}
}

func (m MyString) Output() {
	fmt.Println(m.value)
}

func embeddingStructs() {
	hi := NewMyString("Hello Manu Cutillas!")
	hi.Output()
}
