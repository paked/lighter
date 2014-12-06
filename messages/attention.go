package messages

import (
	"github.com/paked/engi"
)

type AttentionMessage struct {
	Entity *engi.Entity
}

func (am AttentionMessage) Type() string {
	return "AttentionMessage"
}
