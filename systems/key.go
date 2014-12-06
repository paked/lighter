package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
)

type KeySystem struct {
	*engi.System
}

func (ks KeySystem) Name() string {
	return "KeySystem"
}

func (ks *KeySystem) New() {
	ks.System = &engi.System{}
	engi.Mailbox.Listen("CollisionMessage", ks)
}

func (ks KeySystem) Receive(message engi.Message) {
	cm, ok := message.(engi.CollisionMessage)
	if !ok {
		return
	}
	var (
		key *components.KeyComponent
		// oKey    *components.KeyComponent
		control *components.ControlComponent
		link    *engi.LinkComponent
	)

	if !cm.Entity.GetComponent(&control) || !cm.To.GetComponent(&link) || !cm.Entity.GetComponent(&key) {
		return
	}

	link.Entity = cm.Entity
	key.HasKey = true
	// oKey.HasKey = true
}

func (ks *KeySystem) Update(e *engi.Entity, dt float32) {
	var (
		link       *engi.LinkComponent
		space      *engi.SpaceComponent
		otherSpace *engi.SpaceComponent
	)

	if !e.GetComponent(&link) || !e.GetComponent(&space) {
		return
	}

	if link.Entity == nil {
		return
	}

	if !link.Entity.GetComponent(&otherSpace) {
		return
	}

	space.Position = otherSpace.Position
	space.Position.X += space.Width - 8
	space.Position.Y += space.Height / 2

}
