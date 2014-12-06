package systems

import (
	"github.com/paked/engi"
)

type StickySystem struct {
	*engi.System
}

func (ss StickySystem) Name() string {
	return "StickySystem"
}

func (ss *StickySystem) New() {
	ss.System = &engi.System{}
}

func (ss *StickySystem) Update(e *engi.Entity, dt float32) {
	var (
		space  *engi.SpaceComponent
		oSpace *engi.SpaceComponent
		link   *engi.LinkComponent
	)

	if !e.GetComponent(&space) || !e.GetComponent(&link) {
		return
	}

	if !link.Entity.GetComponent(&oSpace) {
		return
	}

	space.Position = engi.Point{(oSpace.Position.X + oSpace.Width/2) - (space.Width / 2), (oSpace.Position.Y + oSpace.Height/2) - (space.Height / 2)}
}
