package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
)

var CONTROL_SCHEME_WASD = "WASD"
var CONTROL_SCHEME_ARROWS = "Arrows"

type ControlSystem struct {
	*engi.System
}

func (controls *ControlSystem) New() {
	controls.System = &engi.System{}
}

func (controls ControlSystem) Name() string {
	return "ControlSystem"
}

func (c *ControlSystem) Update(entity *engi.Entity, dt float32) {
	var controls *components.ControlComponent
	var speed *components.SpeedComponent
	var space *engi.SpaceComponent
	if !entity.GetComponent(&controls) || !entity.GetComponent(&speed) || !entity.GetComponent(&space) {
		return
	}

	var (
		up, down, left, right bool
	)

	switch controls.Scheme {
	case CONTROL_SCHEME_WASD:
		up = engi.Keys.KEY_W.Down()
		down = engi.Keys.KEY_S.Down()
		right = engi.Keys.KEY_A.Down()
		left = engi.Keys.KEY_D.Down()
	case CONTROL_SCHEME_ARROWS:
		up = engi.Keys.KEY_UP.Down()
		down = engi.Keys.KEY_DOWN.Down()
		right = engi.Keys.KEY_RIGHT.Down()
		left = engi.Keys.KEY_LEFT.Down()
	}
	accel := 100 * dt
	drag := float32(.7)

	speed.Acceleration = engi.Point{}
	if up {
		speed.Acceleration.Y -= accel
	}

	if down {
		speed.Acceleration.Y += accel
	}

	if left {
		speed.Acceleration.X += accel
	}

	if right {
		speed.Acceleration.X -= accel
	}

	speed.X += speed.Acceleration.X
	speed.Y += speed.Acceleration.Y

	space.Position.X += speed.X
	space.Position.Y += speed.Y

	speed.X *= drag
	speed.Y *= drag

	if speed.X <= 1 && speed.X >= -1 {
		speed.X = 0
	}

	if speed.Y <= 1 && speed.Y >= -1 {
		speed.Y = 0
	}
}
