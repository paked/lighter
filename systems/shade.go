package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
)

type ShadeSystem struct {
	*engi.System
}

func (ss *ShadeSystem) New() {
	ss.System = &engi.System{}
}

func (ss ShadeSystem) Name() string {
	return "ShadeSystem"
}

func (ss ShadeSystem) Update(e *engi.Entity, dt float32) {
	var (
		playerLink  *engi.LinkComponent
		eSpace      *engi.SpaceComponent
		playerSpace *engi.SpaceComponent
	)
	if !e.GetComponent(&playerLink) || !e.GetComponent(&eSpace) {
		return
	}

	if !playerLink.Entity.GetComponent(&playerSpace) {
		return
	}
	var vulnerable *components.VulnerableComponent
	if !playerLink.Entity.GetComponent(&vulnerable) {
		return
	}

	// vulnerable.Is = false || vulnerable.Is

	if isPointInCircle(engi.Point{playerSpace.Position.X + playerSpace.Width/2, playerSpace.Position.Y + playerSpace.Height/2}, engi.Point{eSpace.Position.X + eSpace.Width/2, eSpace.Position.Y + eSpace.Height/2}, 128) {
		vulnerable.Is = true
	}
}
