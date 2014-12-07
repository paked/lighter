package systems

import (
	"github.com/paked/engi"
	"math/rand"
)

type GemSystem struct {
	*engi.System
}

func (gs GemSystem) Name() string {
	return "GemSystem"
}

func (gs *GemSystem) New() {
	gs.System = &engi.System{}
}

func (gs *GemSystem) Pre() {
	if rand.Float32() > .995 {
		println("SPAWN")
		engi.Wo.AddEntity(NewGem())
	}
}

func (gs *GemSystem) Update(e *engi.Entity, dt float32) {

}

func NewGem() *engi.Entity {
	gem := engi.NewEntity([]string{"RenderSystem"})
	r := engi.NewRenderComponent(engi.Files.Image("gem"), engi.Point{1, 1}, "gem")
	x, y := rand.Float32()*engi.Width(), rand.Float32()*engi.Height()
	s := engi.SpaceComponent{Position: engi.Point{x, y}, Width: 16 * r.Scale.X, Height: 16 * r.Scale.Y}
	gem.AddComponent(&r)
	gem.AddComponent(&s)
	return gem
}
