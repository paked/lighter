package core

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/systems"
	"log"
	"math/rand"
	"time"
)

type Lighter struct {
	engi.World
}

func (l Lighter) Preload() {
	engi.Files.Add("player", "assets/player.png")
	engi.Files.Add("lightsource", "assets/lightsource.png")
}

func (l *Lighter) Setup() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to lighter, a game made by @paked_ for Ludum Dare 31")
	l.AddSystem(&engi.RenderSystem{})
	l.AddSystem(&systems.ControlSystem{})

	l.AddEntity(NewPlayer())

	for x := float32(0); x < 1; x += .25 {
		for y := float32(0); y < 1; y += .25 {
			l.AddEntity(NewLight(x*engi.Width(), y*engi.Height()))
		}

	}
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

func NewLight(x, y float32) *engi.Entity {
	light := engi.NewEntity([]string{"RenderSystem"})
	render := engi.NewRenderComponent(engi.Files.Image("lightsource"), engi.Point{2, 2}, "light")

	offset := engi.Point{rand.Float32() * (engi.Width() / 15), rand.Float32() * (engi.Height() / 15)}
	if rand.Float32() > .5 {
		offset.X *= -1
	}

	if rand.Float32() > .5 {
		offset.Y *= -1
	}
	space := engi.SpaceComponent{Position: engi.Point{x + offset.X, y + offset.Y}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	light.AddComponent(&render)
	light.AddComponent(&space)

	return light
}
