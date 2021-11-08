package main

import(
	"fmt"
	"math"
)
const s string = "constant"
//declare comst s and fill it with string content "constant"

func main() {
	fmt.Println(s)
	//print out const s

	const n = 500000000
	//declare const n and fill it with the int 500000000

	const  d = 3e20 / n
	//declare d and use n

	fmt.Println(int64(n))
	//print out n as a int64

	fmt.Println(math.Sinh(n))
	//take the sin of n and print it out

}
