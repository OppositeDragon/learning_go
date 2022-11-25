// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	Name    string `json:"name"`
// 	Address string `json:"address"`
// }

func main() {
	var name string
	var address string
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Address: ")
	fmt.Scan(&address)
	//var p1 Person = Person{Name: name1, Address: address1}
	// jsonString, err := json.Marshal(p1)
	var map1 = map[string]string{"name": name, "address": address}
	jsonString, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%s\n", string(jsonString))
	}

}
