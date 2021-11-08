package main

import(
	"fmt"
	"time"
)
func main() {
	i := 2
	fmt.Println("write", i, "as ")
	switch	i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//basic switch statement

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")

	}
	//comma seperated expressions switch case

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("It's  afternoon")
	}
	//switch without expresssion
	whatAmI := func(i interface{}) {
		switch t := i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("i'm an int")
		default:
			fmt.Println("Don't know type %T/n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
