package main

import "fmt"

func main() {
	// var myInt int
	// myInt = 5
	// //var myIntPointer *int
	// myIntPointer := &myInt

	// fmt.Println(&myInt)
	// fmt.Println(myIntPointer)

	// fmt.Println(myInt)
	// fmt.Println(*myIntPointer)
	// fmt.Println(reflect.TypeOf(myInt))
	// fmt.Println(reflect.TypeOf(myIntPointer))

	// var numPtr *int
	// num := 5
	// numPtr = &num
	// fmt.Println(num)
	// fmt.Println(numPtr)
	// pointertest(numPtr)
	// newNum := pointertest(&num)
	// fmt.Println(num)
	// fmt.Println(reflect.TypeOf(newNum))
	// fmt.Println(newNum)

	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)

}

// func pointertest(ptr *int) *int {
// 	*ptr = 2
// 	newNum := *ptr * 2
// 	return &newNum

// }

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
