package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(games ...string) {
	if len(games) > 1 {
		for _, g := range games {
			gamer.Games = append(gamer.Games, g)
		}
	}
	gamer.Games = append(gamer.Games, games[0])
}

func main() {
	julia := Gamer{Name: "Julia Ravenclaw"}
	// julia.AddGame("Guild Wars 2")
	julia.AddGame("Mario cart", "Albion", "Guild wars 2")
	fmt.Println(julia.Games)
}
