//KamalDeep Kaur
// Subtraction 
package main

import "fmt"

func main() {
	fmt.Println("Welcome to GroupE's Week 4 Project")
	fmt.Println(substract(4, 3))
	fmt.Println(("Manpreet Kaur"))
	fmt.Println(multiply(6, 3))


}
}
func substract(a int, b int) int {
	return a - b
}
// Armaandeep Singh
// Bubble sort
func bubbleSort(arr []int) 
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

//Manpreet Kaur
//Addition of strings

func string (a string ,b string) string{
	return a+b
}

//priyanka rani
//multiply of integers

func multiply(a int, b int) int {
	return a * b
}
//Prabhjot Kaur
// Variadic Functions
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
