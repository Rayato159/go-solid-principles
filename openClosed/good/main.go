package main

import (
	"fmt"
)

type (
	Player struct{}
	Item   struct {
		Name      string
		Abilities []Ability
	}
	Heal        struct{}
	DamageBluff struct{}

	Ability interface {
		Execute()
	}
)

func (p *Player) Use(item Item) {
	fmt.Println("Use:", item.Name)

	for _, item := range item.Abilities {
		item.Execute()
	}
}

func (h Heal) Execute() {
	fmt.Println("Heal")
}

func (d DamageBluff) Execute() {
	fmt.Println("Increase Damage 100%")
}

func main() {
	p := &Player{}
	steak := Item{
		Name: "Steak",
		Abilities: []Ability{
			Heal{},
			DamageBluff{},
		},
	}

	p.Use(steak)
}
