package structure_controls

import "fmt"

func main(){
	for_1()
	for_2()
	for_3()
	exp_switch()
}

func for_1() {
	i := 1
	for i <= 10 {
		fmt.Println("for_1 :", i)
		i++
	}
}

func for_2() {
	for i := 1; i <= 10; i ++ {
		fmt.Println("for_2", i)
	}
}

func for_3() {
	for i := 1; i <= 10; i ++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("for_2", i)
		}
	}
}

func exp_switch() {
	qualify := "e"
	switch qualify {
	case "m":
		fmt.Println("Oh noo!")
	case "b":
		fmt.Println("Jumm, not")
	case "e":
		fmt.Println("Ok, thats good")
	default:
		fmt.Println("Nothing")
	}
}