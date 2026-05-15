package main

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum+1
	fmt.Println(intNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2 
	var result float32 = floatNum32 +float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3 
	var intNum2 int = 2 
	fmt.Println(intNum1/intNum2) //rounds down
	fmt.Println(intNum1%intNum2)

	var myString string = "hello" + " " + "World"
	fmt.Println(myString)

	var myBoolean bool = true

	myVar := "infered"//infered ?
	fmt.Println("this text is",myVar,"=",myBoolean)// + only used to concatentate strings?
	var1, var2 := 1,2 
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst) // shows const value (must be declared) 
}
