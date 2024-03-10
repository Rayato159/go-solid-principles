package main

import "fmt"

type (
	Player struct{}
	Steak  struct{}
)

func (p *Player) EatSteak(steak *Steak) {
	steak.Heal()
	steak.DamageBuff()
}

func (s *Steak) Heal() {
	fmt.Println("Heal")
}

func (s *Steak) DamageBuff() {
	fmt.Println("Increase Damage 100%")
}

func main() {
	p := &Player{}
	s := Steak{}

	p.EatSteak(&s)
}
