package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "30"
	var str2 = "20"
	var num1, err1 = strconv.Atoi(str1)
	var num2, err2 = strconv.Atoi(str2)

	if err1 == nil {
		fmt.Println("x =", num1)
	}

	if err2 == nil {
		fmt.Println("y =", num2)
	}

	fmt.Println("x+y =", num1+num2) //penjumlahan
	fmt.Println("x-y=", num1-num2)  //pengurangan
	fmt.Println("x*y=", num1*num2)  //perkalian
	fmt.Println("x/y=", num1/num2)  //pembagian

}
