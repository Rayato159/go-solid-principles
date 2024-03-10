package main

import "fmt"

type (
	Aircraft interface {
		Fly()
		Fire()
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

func (s *Spitfire) Fly() {
	fmt.Println("Fly")
}

func (s *Spitfire) Fire() {
	fmt.Println("Fire")
}

func main() {
	spitfire := &Spitfire{}
	spitfire.Fly()
	spitfire.Fire()

	boeing747 := &Boeing747{}
	boeing747.Fly()
	boeing747.Fire()
}
