package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"log"
	// "math"
)

const (
	LOOKING_UP    = 0
	LOOKING_DOWN  = 1
	LOOKING_LEFT  = 2
	LOOKING_RIGHT = 3
)

type VisionSystem struct {
	*engi.System
}

func (vs VisionSystem) Name() string {
	return "VisionSystem"
}

func (vs *VisionSystem) New() {
	vs.System = &engi.System{}
}

func (vs *VisionSystem) Update(e *engi.Entity, dt float32) {
	var v *components.VisionComponent
	if !e.GetComponent(&v) {
		return
	}

	for _, entity := range vs.Entities() {
		if entity.ID() != e.ID() {
			var (
				space  *engi.SpaceComponent
				oSpace *engi.SpaceComponent
			)

			if !e.GetComponent(&space) || !entity.GetComponent(&oSpace) {
				return
			}
			var v1, v2, v3 engi.Point
			switch v.Direction {
			case LOOKING_DOWN:
				v1 = engi.Point{space.Position.X - 25, space.Position.Y}
				v2 = engi.Point{space.Position.X, space.Position.Y + 50}
				v3 = engi.Point{space.Position.X + 25, space.Position.Y}
			case LOOKING_UP:
				v1 = engi.Point{space.Position.X - 50, space.Position.Y}
				v2 = engi.Point{space.Position.X, space.Position.Y - 50}
				v3 = engi.Point{space.Position.X + 50, space.Position.Y}
			case LOOKING_LEFT:
				// v1 = engi.Point{space.Position.X, space.Position.Y - 50}
				// v2 = engi.Point{space.Position.X - 50, space.Position.Y}
				// v1 = engi.Point{space.Position.X, space.Position.Y + 50}
			case LOOKING_RIGHT:
				// v1 = engi.Point{space.Position.X, space.Position.Y - 50}
				// v2 = engi.Point{space.Position.X + 50, space.Position.Y}
				// v3 = engi.Point{space.Position.X, space.Position.Y + 50}
			}

			log.Println(IsPointInTriangle(oSpace.Position, v1, v2, v3), entity.Pattern)
			// log.Println(canSee(space, oSpace))
		}
	}
}

// func canSee(base *engi.SpaceComponent, other *engi.SpaceComponent) bool {

// }

func IsPointInTriangle(point, v1, v2, v3 engi.Point) bool {
	var b1, b2, b3 bool
	b1 = sign(point, v1, v2) < 0
	b1 = sign(point, v2, v3) < 0
	b1 = sign(point, v3, v1) < 0

	return (b1 == b2) && b2 == b3
}

func sign(v1, v2, v3 engi.Point) float32 {
	return (v1.X-v3.X)*(v2.Y-v3.Y) - (v2.X-v3.X)*(v1.Y-v3.Y)
}
