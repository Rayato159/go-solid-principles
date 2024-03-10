package main

import "fmt"

type (
	Player  struct{}
	Display struct{}
)

func (p *Player) Move() {
	fmt.Println("Move")
}

func (p *Player) Attack() {
	fmt.Println("Attack")
}

func (d *Display) DisplayScoreBoard() {
	fmt.Println("DisplayScoreBoard")
}

func main() {
	player := &Player{}
	display := &Display{}

	player.Move()
	player.Attack()

	display.DisplayScoreBoard()
}
