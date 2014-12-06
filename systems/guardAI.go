package systems

import (
	"github.com/paked/engi"
	// "log"
)

type GuardAISystem struct {
	*engi.System
}

func (gai GuardAISystem) Name() string {
	return "GuardAISystem"
}

func (gai *GuardAISystem) New() {
	gai.System = &engi.System{}
}

func (gai *GuardAISystem) Receive(message engi.Message) {

}

func (gai *GuardAISystem) Update(e *engi.Entity, dt float32) {
	var (
		link        *engi.LinkComponent
		space       *engi.SpaceComponent
		targetSpace *engi.SpaceComponent
	)

	if !e.GetComponent(&link) || !e.GetComponent(&space) {
		return
	}

	if !link.Entity.GetComponent(&targetSpace) {
		return
	}

	vel := 100 * dt

	if space.Position.X < targetSpace.Position.X {
		space.Position.X += vel
	}

	if space.Position.X > targetSpace.Position.X {
		space.Position.X -= vel
	}

	if space.Position.Y < targetSpace.Position.Y {
		space.Position.Y += vel
	}

	if space.Position.Y > targetSpace.Position.Y {
		space.Position.Y -= vel
	}
}
