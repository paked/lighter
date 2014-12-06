package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/messages"
)

type GuardAISystem struct {
	*engi.System
}

func (gai GuardAISystem) Name() string {
	return "GuardAISystem"
}

func (gai *GuardAISystem) New() {
	gai.System = &engi.System{}
	engi.Mailbox.Listen("AttentionMessage", gai)
}

func (gai *GuardAISystem) Receive(message engi.Message) {
	switch message.(type) {
	case messages.AttentionMessage:
		attention := message.(messages.AttentionMessage)
		for _, e := range gai.Entities() {
			var link *engi.LinkComponent
			if !e.GetComponent(&link) {
				break
			}

			if link.Entity == nil {
				link.Entity = attention.Entity
				return
			}
		}
	}
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

	if link.Entity == nil {
		return
	}

	if !link.Entity.GetComponent(&targetSpace) {
		return
	}

	vel := 100 * dt
	done := true
	if space.Position.X < (targetSpace.Position.X - 5) {
		space.Position.X += vel
		done = false
	}

	if space.Position.X > (targetSpace.Position.X + 5) {
		space.Position.X -= vel
		done = false
	}

	if space.Position.Y < (targetSpace.Position.Y - 5) {
		space.Position.Y += vel
		done = false
	}

	if space.Position.Y > (targetSpace.Position.Y + 5) {
		space.Position.Y -= vel
		done = false
	}

	if done {
		var shadeLink *engi.LinkComponent
		if !link.Entity.GetComponent(&shadeLink) {
			return
		}

		if !shadeLink.Entity.Exists {
			shadeLink.Entity.Exists = true
			link.Entity = nil
		}
	}

}
