package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(games ...string) {
	if len(games) < 1 {
		gamer.Games = append(gamer.Games, games[0])
	}
	for _, g := range games {
		gamer.Games = append(gamer.Games, g)
	}
}

func main() {
	julia := Gamer{Name: "Julia Ravenclaw"}
	julia.AddGame("Power")
	julia.AddGame("Mario cart", "Albion", "Guild wars 2")
	fmt.Println(julia.Games)
}
