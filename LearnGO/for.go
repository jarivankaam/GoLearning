package main

import "fmt"

func main() {
	i := 1
	for 1 <+ 3 {
		fmt.Println(i)
		i = i +1
	}
	//single condition for loop

	for j := 7; j <+ 9; j++{
		fmt.Println(j)
	}
	// initial/condition/after for loop

	for {
		fmt.Println("loop")
		break
	}
	// not conditional for loop with break

	for n := 0; n <= 5; n++{
		if n%2 == 0{
		continue}
		fmt.Println(n)
	}
	// continues after next iteration
}

