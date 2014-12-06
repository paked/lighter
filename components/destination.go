package components

import (
	"github.com/paked/engi"
)

type DestinationComponent struct {
	engi.Point
}

func (dc DestinationComponent) Name() string {
	return "DesitinationComponent"
}
