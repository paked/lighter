package systems

import (
	"github.com/paked/engi"
	// "github.com/paked/lighter/core"
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

}

func (ls *LightSystem) Update(e *engi.Entity, dt float32) {
	// log.Printf("%T", 0xffffff)
}
