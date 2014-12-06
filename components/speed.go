package components

import (
	"github.com/paked/engi"
)

type SpeedComponent struct {
	engi.Point              //Velocity
	Acceleration engi.Point //Acceleration
	// Drag         float32
}

func (s SpeedComponent) Name() string {
	return "SpeedComponent"
}
