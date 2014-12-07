package messages

type PuzzleDoneMessage struct {
}

func (pdz PuzzleDoneMessage) Type() string {
	return "PuzzleDoneMessage"
}
