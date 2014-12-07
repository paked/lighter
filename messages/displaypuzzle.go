package messages

type DisplayPuzzle struct {
}

func (dp DisplayPuzzle) Type() string {
	return "DisplayPuzzleMessage"
}
