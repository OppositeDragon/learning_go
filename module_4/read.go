// Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"fmt"
	"os"
)

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {
	var filename string
	fmt.Print("Filename: ")
	fmt.Scan(&filename)
	file, _ := os.Open(filename)
	b := make([]byte, 1)
	arrPerson := make([]Person, 0)
	var n int = -1
	//var err error
	var arrFName []byte
	var arrLName []byte
	var isName bool = true
	for n != 0 {
		n, _ = file.Read(b)
		if len(b) > 0 {

			if b[0] == 32 {
				isName = false
			} else if b[0] == 10 {
				isName = true
				arrPerson = append(arrPerson, Person{Fname: string(arrFName), Lname: string(arrLName)})
				arrFName = make([]byte, 0)
				arrLName = make([]byte, 0)
			} else {
				if isName == true {
					arrFName = append(arrFName, b[0])
				} else {
					arrLName = append(arrLName, b[0])
				}
			}
		}
	}
	arrPerson = append(arrPerson, Person{Fname: string(arrFName), Lname: string(arrLName)})
	for _, v := range arrPerson {
		fmt.Println(v.Fname, v.Lname)
	}
	// fmt.Print(barr)
	// fmt.Print(string(barr))
}
