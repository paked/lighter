package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/messages"
)

type LightSystem struct {
	*engi.System
}

func (ls LightSystem) Name() string {
	return "LightSystem"
}

func (ls *LightSystem) New() {
	ls.System = &engi.System{}
	engi.Mailbox.Listen("CollisionMessage", ls)
}

func (ls *LightSystem) Receive(message engi.Message) {
	switch message.(type) {
	case engi.CollisionMessage:
		cm := message.(engi.CollisionMessage)
		var (
			link     *engi.LinkComponent
			controls *components.ControlComponent
			key      *components.KeyComponent
		)

		if !cm.Entity.GetComponent(&key) || !cm.Entity.GetComponent(&controls) || !cm.To.GetComponent(&link) {
			return
		}
		// log.Println(cm.Entity.Pattern, cm.To.Pattern)
		if cm.Entity.Pattern == "player" && cm.To.Pattern == "light" {
			if key.HasKey {
				if link.Entity != nil && link.Entity.Pattern == "shade" {
					link.Entity.Exists = false
				}
			}
		}
	}
}

func (ls *LightSystem) Update(e *engi.Entity, dt float32) {
	var (
		link *engi.LinkComponent
	)

	if !e.GetComponent(&link) {
		return
	}

	if !link.Entity.Exists {
		engi.Mailbox.Dispatch(messages.AttentionMessage{e})
	}
}
