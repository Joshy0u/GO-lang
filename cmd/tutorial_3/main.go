package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hellowworld"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	}else{
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
		//reminder that THIS ROUNDS DOWN
	}

	//interesting swtich statemnets
	switch remainder{
	case 0:
			fmt.Println("The division was exact")
	case 1,2:
			fmt.Println("the division was close")
	default:
			fmt.Println("The division was not close")
	}

}

func printMe(printVal string){
	fmt.Println(printVal)
}

func intDivision(numerator int, denominator int) (int, int, error){//define return value here 
	var err error //defaults to nil 
	if denominator==0{
		err = errors.New("cannot divide by zero")
		return 0 , 0, err
	}

	var result int = numerator/denominator
	var remanider int = numerator%denominator
	return result , remanider, err 

}
