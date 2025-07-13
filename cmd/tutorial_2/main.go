package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")

	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum2 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("test")) //is the number of bytes, not the number of characters

	fmt.Println(utf8.RuneCountInString("y")) //1

	var myRune rune = 'a'
	fmt.Println(myRune) //97

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3) //0 (default value)

	myVar := "text" //var myVar = "text" (without string)
	fmt.Println(myVar)

	var1, var2 := 1, 2 //var var1, var2 int = 1, 2 (defining multiple variables)
	fmt.Println(var1, var2)

	//!!! Once a variable is created, its value cannot be changed.

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415

}
