package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/messages"
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
	engi.Mailbox.Listen("CollisionMessage", gs)
}

func (gs *GemSystem) Pre() {
	if rand.Float32() > .995 {
		println("SPorrrawwwN")
		engi.Wo.AddEntity(NewGem())
	}
}

func (gs *GemSystem) Receive(message engi.Message) {
	switch message.(type) {
	case engi.CollisionMessage:
		cm := message.(engi.CollisionMessage)
		if cm.Entity.Pattern == "player" && cm.To.Pattern == "gem" {
			cm.To.Exists = false
			engi.Mailbox.Dispatch(messages.ScoreMessage{})
		}
	}
}

func (gs *GemSystem) Update(e *engi.Entity, dt float32) {

}

func NewGem() *engi.Entity {
	gem := engi.NewEntity([]string{"RenderSystem", "CollisionSystem"})
	gem.Pattern = "gem"
	r := engi.NewRenderComponent(engi.Files.Image("gem"), engi.Point{1, 1}, "gem")
	x, y := rand.Float32()*engi.Width(), rand.Float32()*engi.Height()
	s := engi.SpaceComponent{Position: engi.Point{x, y}, Width: 16 * r.Scale.X, Height: 16 * r.Scale.Y}
	c := engi.CollisionComponent{}
	gem.AddComponent(&r)
	gem.AddComponent(&s)
	gem.AddComponent(&c)
	return gem
}
