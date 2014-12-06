package messages

import (
	"github.com/paked/engi"
)

type TargetMessage struct {
	Entity *engi.Entity
	Guard  *engi.Entity
}

func (tm TargetMessage) Type() string {
	return "TargetMessage"
}
