package main

import (
	"Practice_HomeWork/people"
	"fmt"
)

func main() {
	people.PeopleList.Sort()
	people.PeopleList.Let()
	fmt.Println(people.PeopleList.Let())
}
