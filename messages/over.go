package messages

type GameOverMessage struct {
}

func (gom GameOverMessage) Type() string {
	return "GameOverMessage"
}
