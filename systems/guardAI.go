package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/messages"
	"math/rand"
)

const (
	PROGRESS_MOVING = 1
	PROGRESS_NONE   = 0
)

type GuardAISystem struct {
	*engi.System
	progress map[string]bool
}

func (gai GuardAISystem) Name() string {
	return "GuardAISystem"
}

func (gai *GuardAISystem) New() {
	gai.System = &engi.System{}
	gai.progress = make(map[string]bool)
	engi.Mailbox.Listen("AttentionMessage", gai)
}

func (gai *GuardAISystem) Receive(message engi.Message) {
	switch message.(type) {
	case messages.AttentionMessage:
		attention := message.(messages.AttentionMessage)
		if gai.progress[attention.Entity.ID()] {
			return
		}

		for _, e := range gai.Entities() {
			var link *engi.LinkComponent
			if !e.GetComponent(&link) {
				break
			}

			if link.Entity == nil {
				gai.progress[attention.Entity.ID()] = true
				link.Entity = attention.Entity
				return
			}
		}
	}
}

func (gai *GuardAISystem) Update(e *engi.Entity, dt float32) {
	var (
		link        *engi.LinkComponent
		space       *engi.SpaceComponent
		targetSpace *engi.SpaceComponent
	)

	if !e.GetComponent(&link) || !e.GetComponent(&space) {
		return
	}

	point := engi.Point{}

	if link.Entity == nil {
		var dc *components.DestinationComponent
		if !e.GetComponent(&dc) {
			return
		}

		point = dc.Point
	} else {
		if link.Entity.GetComponent(&targetSpace) {
			point = targetSpace.Position
		}
	}

	vel := 300 * dt
	done := true
	if space.Position.X < (point.X - 5) {
		space.Position.X += vel
		done = false
	}

	if space.Position.X > (point.X + 5) {
		space.Position.X -= vel
		done = false
	}

	if space.Position.Y < (point.Y - 5) {
		space.Position.Y += vel
		done = false
	}

	if space.Position.Y > (point.Y + 5) {
		space.Position.Y -= vel
		done = false
	}

	if done {
		var dc *components.DestinationComponent
		if !e.GetComponent(&dc) {
			return
		}

		if link.Entity != nil {
			var shadeLink *engi.LinkComponent
			if !link.Entity.GetComponent(&shadeLink) {
				return
			}

			if !shadeLink.Entity.Exists {
				shadeLink.Entity.Exists = true
				gai.progress[link.Entity.ID()] = false
				link.Entity = nil
			}

			dc.Point = GenerateGuardPosition(point)
		} else {
			dc.Point = GenerateGuardPosition(point)

		}

	}

}

func GenerateGuardPosition(old engi.Point) engi.Point {
	point := engi.Point{engi.Width() * rand.Float32(), engi.Height() * rand.Float32()}

	if rand.Float32() > .5 {
		point.X = old.X
	} else {
		point.Y = old.Y
	}
	return point
}
