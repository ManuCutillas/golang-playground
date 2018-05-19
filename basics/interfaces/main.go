package main

import "fmt"

type superHero interface{
	getName() string
	getAge() int
	getSuperPowers() bool
}

type sName struct {
	name string
	age int
	superPowers bool
}

func (str sName)getName() string {
	return str.name
}

func (a sName)getAge() int {
	return a.age
}

func (b sName)getSuperPowers() bool{
	return b.superPowers
}

func main() {
	var s = sName{"Batman", 35, true,}
	superHeroName := getSuperHero(s)
	fmt.Println("Super Hero Name : ",superHeroName)
}

func getSuperHero(s superHero) string {
	name := s.getName()
	fmt.Println(name)
	return name
}