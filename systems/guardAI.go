package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/messages"
	// "log"
	"math"
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
		animation   *engi.AnimationComponent
		// vision      *components.VisionComponent
	)

	if !e.GetComponent(&link) || !e.GetComponent(&space) || !e.GetComponent(&animation) {
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
	var speed *components.SpeedComponent

	if !e.GetComponent(&speed) {
		return
	}
	maxVel := 75 * dt
	accel := 10 * dt
	if link.Entity != nil {
		if link.Entity.Pattern == "player" {
			maxVel *= 2
		}
	}

	drag := float32(.4)
	done := true
	if space.Position.X < (point.X - 5) {
		speed.Acceleration.X += accel
		done = false
		animation.SelectAnimation("right")

	}

	if space.Position.X > (point.X + 5) {
		speed.Acceleration.X -= accel
		animation.SelectAnimation("left")
		done = false
	}

	if !done {
		speed.Acceleration.Y = 0
	}

	if done {
		speed.Acceleration.X = 0
		if space.Position.Y < (point.Y - 5) {
			speed.Acceleration.Y += accel
			done = false
			animation.SelectAnimation("down")
		}

		if space.Position.Y > (point.Y + 5) {
			speed.Acceleration.Y -= accel
			animation.SelectAnimation("up")
			done = false
		}
	}

	speed.X += speed.Acceleration.X
	speed.Y += speed.Acceleration.Y

	if speed.X > maxVel {
		speed.X = maxVel
	}

	if speed.X < -maxVel {
		speed.X = -maxVel
	}

	if speed.Y > maxVel {
		speed.Y = maxVel
	}

	if speed.Y < -maxVel {
		speed.Y = -maxVel
	}

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

		speed.X = 0
		speed.Y = 0
		speed.Acceleration.X = 0
		speed.Acceleration.Y = 0
	}

}

func distanceBetween(one, two engi.Point) float64 {
	return math.Sqrt(math.Pow(float64(one.X-two.X), 2) + math.Pow(float64(one.Y-two.Y), 2))
}

func GenerateGuardPosition(old engi.Point) engi.Point {
	point := engi.Point{engi.Width() * rand.Float32(), engi.Height() * rand.Float32()}
	// log.Println("NEW KEY")
	// if rand.Float32() > .5 {
	// 	point.X = old.X
	// } else {
	// 	point.Y = old.Y
	// }
	return point
}
