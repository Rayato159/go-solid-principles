package main

import "fmt"

type (
	Aircraft interface {
		Fly()
	}

	Boeing747 struct{ Aircraft }

	Spitfire struct{ Aircraft }
)

func (b *Boeing747) Fly() {
	fmt.Println("Fly")
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
}
