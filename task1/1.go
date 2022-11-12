// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import (
	"fmt"
)

type Human struct { // parent structure
	Name   string
	Height uint8
	Weight uint16
}

type Action struct { // child structure
	Human
	Job   string
	Hobby string
}

func (T Human) WhatIsYourName() { // Human Method
	fmt.Printf("My name is %s.\n", T.Name)
}

func (T Human) IntroduceYourself() { // Human method
	fmt.Printf("My name is %s. My height is %v, my weight is %v.\n", T.Name, T.Height, T.Weight)
}

func (T Action) IntroduceYourself() { // override Human method
	fmt.Printf("My name is %s. My height is %v, my weight is %v. I'm a %s\n", T.Name, T.Height, T.Weight, T.Job)
}

func (T Action) WhatIsYourHobby() { // Action method
	fmt.Printf("My name is %s. My hobby is %s\n", T.Name, T.Hobby)
}

func main() {
	var Jon = Human{Name: "Jon", Height: 173, Weight: 70}
	var CoolerJon = Action{Human: Human{Name: "Cooler Jon", Height: 175, Weight: 73}, Job: "programmer", Hobby: "Snorkeling"}

	Jon.WhatIsYourName()          // Human method
	Jon.IntroduceYourself()       // Human method
	CoolerJon.WhatIsYourName()    // inherit Human method
	CoolerJon.IntroduceYourself() // override Human Method
	CoolerJon.WhatIsYourHobby()   // Action method
}
