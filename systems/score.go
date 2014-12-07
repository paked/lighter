package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/messages"
	"strconv"
)

type ScoreSystem struct {
	*engi.System
	Score        int
	accountedFor map[string]bool
}

func (ss ScoreSystem) Name() string {
	return "ScoreSystem"
}

func (ss *ScoreSystem) New() {
	ss.System = &engi.System{}
	engi.Mailbox.Listen("ScoreMessage", ss)
}

func (ss *ScoreSystem) Receive(message engi.Message) {
	switch message.(type) {
	case messages.ScoreMessage:
		ss.Score += 1
	}
}

func (ss *ScoreSystem) Update(e *engi.Entity, dt float32) {
	var (
		r *engi.RenderComponent
		s *engi.SpaceComponent
	)

	if !e.GetComponent(&r) || !e.GetComponent(&s) {
		return
	}

	if t, ok := r.Display.(*engi.Text); ok {
		t.Content = strconv.Itoa(ss.Score) + " Points"
	}

	s.Position.X = (engi.Width() - s.Width) / 2
	s.Position.Y = engi.Height() - (s.Height * 2)

}
