package main

import "fmt"

func main() {
	map1()
	iterate()
	myDelete()
}

func map1() {
	fmt.Println("hoo!")
	dictionary := make(map[string]string)
	fmt.Println("dictionary", dictionary)
	dictionary2 := map[string]string{}
	dictionary2["first"] = "My first value"
	fmt.Println("dictionary2", dictionary2)
	dictionary3 := map[string]string{
		"first": "My first value",
		"second": "My second value",
	}
	fmt.Println("dictionary3", dictionary3)

	myValue, exist := dictionary3["first"]
	if exist {
		fmt.Println(myValue)
	}else{
		fmt.Println("oh noo!")
	}

	mySuperHeroValue, exist2 := dictionary3["superman"]
	if exist2 {
		fmt.Println(mySuperHeroValue)
	}else{
		fmt.Println("ooh noo!")
	}
}

func iterate(){
	dictionary := map[string]string{
		"A": "It´s A",
		"B": "It´s B",
		"C": "It´s C",
	}
	for myKey, myValue := range dictionary {
		fmt.Println(myKey, myValue)
	}
}

func myDelete(){
	dictionary := map[string]string{
		"A": "It´s A",
		"B": "It´s B",
		"C": "It´s C",
	}
	delete(dictionary, "C")
	fmt.Println("DELETED C ===>", dictionary)
}