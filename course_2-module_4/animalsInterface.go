/**
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

| Animal | Food eaten | Locomotion method | Spoken sound |
| ------ | ---------- | ----------------- | ------------ |
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name, food, locomotion, noise string
}

func (animal Cow) Eat() {
	fmt.Println(animal.food)
}
func (animal Cow) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Cow) Speak() {
	fmt.Println(animal.noise)
}

type Bird struct {
	name, food, locomotion, noise string
}

func (animal Bird) Eat() {
	fmt.Println(animal.food)
}
func (animal Bird) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Bird) Speak() {
	fmt.Println(animal.noise)
}

type Snake struct {
	name, food, locomotion, noise string
}

func (animal Snake) Eat() {
	fmt.Println(animal.food)
}
func (animal Snake) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Snake) Speak() {
	fmt.Println(animal.noise)
}

type AnimalsContainer struct {
	cows   []Cow
	birds  []Bird
	snakes []Snake
}

var animalsContainer AnimalsContainer
var animals []Animal

func main() {
	animals = make([]Animal, 0)
	var request string
	for true {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request = scanner.Text()
		request = strings.ToLower(request)
		command, name, animalOrAction, err := separateCommands(request)
		if err != "" {
			fmt.Println(err)
		} else {
			newOrQuery(command, name, animalOrAction)
		}
	}
}

func separateCommands(request string) (command, name, animalOrAction string, err string) {
	slice := strings.Split(request, " ")
	if len(slice) != 3 {
		err = "Invalid request"
		return
	}
	return slice[0], slice[1], slice[2], err
}
func newOrQuery(command, name, animalOrAction string) {
	if command == "newanimal" {
		newAnimal(name, animalOrAction)
	} else if command == "query" {
		if len(animals) == 0 {
			fmt.Println("No animals registered")
			return
		}
		query(name, animalOrAction)
	} else {
		fmt.Println("Command not supported")
	}
}

func newAnimal(name, animal string) {

	switch animal {
	case "cow":
		animals = append(animals, Cow{name, "grass", "walk", "moo"})
		fmt.Println("Created it!")
		break
	case "bird":
		animals = append(animals, Bird{name, "worms", "fly", "peep"})
		fmt.Println("Created it!")
		break
	case "snake":
		animals = append(animals, Snake{name, "mice", "slither", "hsss"})
		fmt.Println("Created it!")
		break
	default:
		fmt.Println("Animal not supported")
	}
}

func query(name, action string) {
	if action == "eat" || action == "move" || action == "speak" {
		for _, value := range animals {
			switch animal := value.(type) {
			case Cow:
				if animal.name == name {
					printAnimalAction(animal, action)
					return
				}
			case Bird:
				if animal.name == name {
					printAnimalAction(animal, action)
					return
				}
			case Snake:
				if animal.name == name {
					printAnimalAction(animal, action)
					return
				}
			default:
				fmt.Println("Animal not supported")
			}
		}
	} else {
		fmt.Println("Action not supported")
		return
	}
}

func printAnimalAction(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
		break
	case "move":
		animal.Move()
		break
	case "speak":
		animal.Speak()
		break
	default:
		fmt.Println("Action not suported")
	}
}
