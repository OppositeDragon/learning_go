/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/
package main

import (
	"fmt"

	"strconv"
)

func main() {
	var mainArray = make([]int, 0)
	var n string
	fmt.Println("Enter numbers to fill an array, enter a letter to stop")
	for n != "x" {
		fmt.Scan(&n)
		number, err := strconv.Atoi(n)
		if err == nil {
			mainArray = append(mainArray, number)
		} else {
			break
		}
	}

	lenght := len(mainArray)
	ipart := float64(lenght) / 4
	slice1 := mainArray[0:int(ipart)]
	slice2 := mainArray[int(ipart):int(ipart*2)]
	slice3 := mainArray[int(ipart*2):int(ipart*3)]
	slice4 := mainArray[int(ipart*3):lenght]

	channel := make(chan []int, 4)
	go bubbleSort(slice1, channel)
	go bubbleSort(slice2, channel)
	go bubbleSort(slice3, channel)
	go bubbleSort(slice4, channel)
	a, b, c, d := <-channel, <-channel, <-channel, <-channel // receive from c
	slice := append(a, b...)
	slice = append(slice, c...)
	slice = append(slice, d...)
	
	bubbleSort(slice, channel)
	fmt.Println(slice)
}

func bubbleSort(slice []int, c chan []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
	c <- slice
}

func swap(slice []int, i int) {
	var temp int
	temp = slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = temp
}
