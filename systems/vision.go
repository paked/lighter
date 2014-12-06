package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"log"
	"math"
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

			if isPointInCircle(oSpace.Position, engi.Point{space.Position.X + space.Width/2, space.Position.Y + space.Height/2}, 64) {
				log.Println("Hey")
			}

		}
	}
}

func isPointInCircle(point engi.Point, center engi.Point, radius float32) bool {
	return math.Pow(float64(point.X-center.X), 2)+math.Pow(float64(point.Y-center.Y), 2) < math.Pow(float64(radius), 2)
}
