package main

import (
	"fmt"
)

func main() {
 sum:= sum(2,2)
	 multiply:= multiplication(2,2)
		division:= division(200,10, 2)
		subtract:=	subtract(2,2)
	fmt.Println(sum, multiply,division,subtract)
}

func sum(i ... int) int{
		total:=0

		for _, v := range i {
			total+= v
		}
	return total
}

func multiplication(i ... int) int{
	total:= 1
 	for _, v := range i {
		fmt.Println(v)

		total = total * v
	}
return total
}

func division(i ... float32) float32{
	result := i[0]

	for _, num := range i[1:] {
					result /= num
	}
	return result
}

func subtract(i ...int) int { 
	result := i[0]

	for _, num := range i[1:] {
					result -= num
	}
	return result
}