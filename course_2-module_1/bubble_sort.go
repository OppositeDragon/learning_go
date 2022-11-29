/**
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/

package main

import "fmt"

func main() {
	var quantity int
	fmt.Print("How many numbers does the array contain? (min:2 ,max:10) ")
	fmt.Scan(&quantity)
	if quantity > 10 {
		quantity = 10
	} else if quantity < 2 {
		quantity = 2
	}
	var slice = make([]int, quantity)
	askUserForNumbers(slice, quantity)
	bubbleSort(slice)
	fmt.Println("Ordered slice: ", slice)
}

func askUserForNumbers(slice []int, quantity int) {
	var number int
	fmt.Printf("Please enter %d numbers\n", quantity)
	for i := 0; i < quantity; i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Scan(&number)
		slice[i] = number
	}
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if(slice[j] > slice[j+1]) {
				swap(slice, j)
			}
		}
	}
}

func swap(slice []int, i int) {
	var temp int
	temp = slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = temp
}
