package components

import (
// "github.com/paked/engi"
)

type HealthComponent struct {
	Points    float32
	MaxPoints float32
}

func (hc HealthComponent) Name() string {
	return "HealthComponent"
}
