package core

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/systems"
	"log"
)

type Lighter struct {
	engi.World
}

func (l Lighter) Preload() {
	engi.Files.Add("player", "assets/player.png")
}

func (l *Lighter) Setup() {
	log.Println("Welcome to lighter a game made by @paked_ for Ludum Dare 31")
	l.AddSystem(&engi.RenderSystem{})
	l.AddSystem(&systems.ControlSystem{})

	l.AddEntity(NewPlayer())
}

func NewPlayer() *engi.Entity {
	player := engi.NewEntity([]string{"RenderSystem", "ControlSystem"})
	render := engi.NewRenderComponent(engi.Files.Image("player"), engi.Point{2, 2}, "player")
	space := engi.SpaceComponent{Position: engi.Point{400, 400}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	control := components.ControlComponent{Scheme: systems.CONTROL_SCHEME_WASD}
	speed := components.SpeedComponent{}

	player.AddComponent(&render)
	player.AddComponent(&space)
	player.AddComponent(&control)
	player.AddComponent(&speed)

	return player
}
