package main

import "fmt"

type (
	Aircraft interface {
		Fly()
		Fire()
		Aircondition()
	}

	Boeing747 struct{}
	Spitfire  struct{}
)

func (b *Boeing747) Fly() {
	fmt.Println("Fly")
}

func (b *Boeing747) Fire() {
	panic("Error: Boeing747 cannot fire")
}

func (b *Boeing747) Aircondition() {
	fmt.Println("Aircondition")
}

func (s *Spitfire) Fly() {
	fmt.Println("Fly")
}

func (s *Spitfire) Fire() {
	fmt.Println("Fire")
}

func (s *Spitfire) Aircondition() {
	panic("Error: Spitfire doesn't have aircondition")
}

func main() {
	spitfire := &Spitfire{}
	spitfire.Fly()
	spitfire.Fire()
	spitfire.Aircondition()

	boeing747 := &Boeing747{}
	boeing747.Fly()
	boeing747.Fire()
	boeing747.Aircondition()
}
