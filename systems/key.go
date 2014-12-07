package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"log"
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
	if cm.Entity.Pattern == "player" && cm.To.Pattern == "key" && !key.HasKey {
		if key.Cooldown != 0 {
			return
		}
		link.Entity = cm.Entity
		key.HasKey = true
		log.Println("Now haz key")
	}
}

func (ks *KeySystem) Update(e *engi.Entity, dt float32) {
	var (
		link       *engi.LinkComponent
		space      *engi.SpaceComponent
		otherSpace *engi.SpaceComponent
		key        *components.KeyComponent
		mKey       *components.KeyComponent
	)

	if !e.GetComponent(&mKey) {
		return
	}

	// if mKey.Cooldown != 0 {
	// 	mKey.Cooldown -= 1
	// 	return
	// }

	if !e.GetComponent(&link) || !e.GetComponent(&space) {
		return
	}

	if link.Entity == nil {
		return
	}

	if !link.Entity.GetComponent(&otherSpace) {
		return
	}
	// log.Println("GETTING KEY")
	if !link.Entity.GetComponent(&key) {
		return
	}
	// log.Println("GOT KEY")
	// log.Println(link.Entity.Pattern, e.Pattern)
	if key.HasKey {
		// log.Println("MOVIUNG KEY")
		space.Position = otherSpace.Position
		space.Position.X += space.Width - 8
		space.Position.Y += space.Height / 2
	}

}
