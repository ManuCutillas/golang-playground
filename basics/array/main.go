package main

import (
	"fmt"
)

func main() {
	arr_1()
	multi_arr()
	useRange()
	sliceArr()
	utils()
}

func arr_1(){
	var my_array [5]int
	fmt.Println("my_array ==> ", my_array)

	my_array2 := [5]int{1,2,3,4,5}
	fmt.Println("my_array2 ==> ", my_array2)

	fruits := [3]string{"bananas", "Strawberry", "pear"}
	fmt.Println("fruits ===>", fruits)

	fmt.Println("Bananas:", fruits[0])
}

func multi_arr() {
	multi := [2][3]string{{"this","is"},{"awesome","cool","bad"}}
	fmt.Println(multi)
}

func sliceArr() {

	var slice []int
	fmt.Println("Print slice var ====>", slice)
	if slice == nil {
		fmt.Println("nil")
	}

	slice2 := make([]int, 0)
	fmt.Println(slice2)
	if slice2 == nil {
		fmt.Println("nil")
	}else{
		fmt.Println("slice 2 is not nil")
	}
}

func useRange(){
	n := []int{1,2,3,4,5}

	for nIndex, nValue :=range n {
		fmt.Println("range =====>",nIndex, nValue)
	}

	for _, nValue := range n {
		fmt.Println("range blank index =====>",nValue)
	}

}

func utils(){
	my_array3 := [4]string{"a","b","c","d"}
	fmt.Println(len(my_array3))
	fmt.Println(cap(my_array3))
	//Only make array with void length.
	my_array4 := make([]string, 5)
	fmt.Println("my_array4 ====>", my_array4)
	// Make array with void length and max capacity.
	my_array5 := make([]string, 5, 10)
	fmt.Println("my_array5 =====>", my_array5)

	onlySelected := my_array3[1:3]
	fmt.Println("onlySelected ====>", onlySelected)

	// Shared reference
	onlySelected[0] = "oh my god!"
	fmt.Println(my_array3);

	// Append elements
	onlySelected = append(onlySelected, "more more more")
	fmt.Println(onlySelected);
}