package components

type PuzzleComponent struct {
	Current  []string
	Progress int
}

func (pc PuzzleComponent) Name() string {
	return "PuzzleComponent"
}
