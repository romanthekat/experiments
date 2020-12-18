package core

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	basePopulation = 420
	ticksPerSecond = 30
	ticksPerDay    = ticksPerSecond * 1
)

type Game struct {
	Cities        []*City
	playerCompany *Company
}

func NewGame(citiesCount int) *Game {
	rand.Seed(time.Now().UTC().UnixNano())

	game := &Game{}

	for i := 0; i < citiesCount; i++ {
		game.GenerateCity()
	}

	return game
}

func (g *Game) GenerateCity() {
	city := NewCity(fmt.Sprintf("Town %d", len(g.Cities)+1), basePopulation+rand.Intn(basePopulation/2))
	g.Cities = append(g.Cities, city)
}

func (g *Game) Tick() {
	for _, city := range g.Cities {
		city.Tick()
	}
}

func LaunchGameCycle(game *Game) {
	ticker := time.NewTicker(time.Second / ticksPerSecond)

	for {
		select {
		case <-ticker.C:
			game.Tick()
			visualize(game)
		}
	}
}

func visualize(game *Game) {
	fmt.Println()
	for _, city := range game.Cities {
		fmt.Println(city)
	}
}
