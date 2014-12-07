package systems

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/messages"
	"log"
)

const (
	PUZZLE_UP    = "UP"
	PUZZLE_DOWN  = "DOWN"
	PUZZLE_LEFT  = "LEFT"
	PUZZLE_RIGHT = "RIGHT"
	PUZZLE_NONE  = "NONE"
)

type PuzzleSystem struct {
	*engi.System
	s       *engi.Spritesheet
	puzzles [][]string
	Busy    bool
}

func (pz PuzzleSystem) Name() string {
	return "PuzzleSystem"
}

func (pz *PuzzleSystem) New() {
	pz.System = &engi.System{}
	engi.Mailbox.Listen("DisplayPuzzleMessage", pz)
	// engi.Mailbox.Listen("MovePuzzle", pz)

	pz.s = engi.NewSpritesheet("arrows", 32)

	ps := [][]string{{PUZZLE_UP, PUZZLE_DOWN, PUZZLE_DOWN, PUZZLE_RIGHT}, {PUZZLE_DOWN, PUZZLE_RIGHT, PUZZLE_UP}, {PUZZLE_RIGHT, PUZZLE_LEFT, PUZZLE_UP}}
	pz.puzzles = ps
}
func (pz *PuzzleSystem) Receive(message engi.Message) {
	switch message.(type) {
	case messages.DisplayPuzzle:
		var (
			pc *components.PuzzleComponent
		)

		if !pz.Entities()[0].GetComponent(&pc) {
			return
		}

		pc.Current = pz.puzzles[0]
		pc.Progress = 0
		pz.Entities()[0].Exists = true

		log.Println("SHOULD BE DISPLAYING")
	}
}

func (pz *PuzzleSystem) Update(e *engi.Entity, dt float32) {
	var (
		pc *components.PuzzleComponent
		r  *engi.RenderComponent
	)

	if !e.GetComponent(&pc) || !e.GetComponent(&r) {
		return
	}

	if len(pc.Current) == 0 {
		return
	}

	current := PUZZLE_NONE

	if engi.Keys.KEY_DOWN.JustPressed() || engi.Keys.KEY_S.JustPressed() {
		current = PUZZLE_DOWN
	}

	if engi.Keys.KEY_UP.JustPressed() || engi.Keys.KEY_W.JustPressed() {
		current = PUZZLE_UP
	}

	if engi.Keys.KEY_LEFT.JustPressed() || engi.Keys.KEY_A.JustPressed() {
		current = PUZZLE_LEFT
	}

	if engi.Keys.KEY_RIGHT.JustPressed() || engi.Keys.KEY_D.JustPressed() {
		current = PUZZLE_RIGHT
	}

	// log.Println(len(pc.Current), pc.Progress)
	if pc.Progress == len(pc.Current) {
		engi.Mailbox.Dispatch(messages.PuzzleDoneMessage{})
		e.Exists = false
		return
	}

	if current == pc.Current[pc.Progress] {
		pc.Progress += 1
	} else if current != pc.Current[pc.Progress] && current != PUZZLE_NONE {
		e.Exists = false
	}

	if pc.Progress != len(pc.Current) {
		next := pc.Current[pc.Progress]
		switch next {
		case PUZZLE_UP:
			r.Display = pz.s.Cell(0)
		case PUZZLE_DOWN:
			r.Display = pz.s.Cell(1)
		case PUZZLE_LEFT:
			r.Display = pz.s.Cell(2)
		case PUZZLE_RIGHT:
			r.Display = pz.s.Cell(3)
		}
	}
}
