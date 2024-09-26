//KamalDeep Kaur
// Subtraction 
package main

import "fmt"

func main() {
	fmt.Println("Welcome to GroupE's Week 4 Project")
	fmt.Println(substract(4, 3))
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
