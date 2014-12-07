package messages

type ScoreMessage struct {
	ID string
}

func (sm ScoreMessage) Type() string {
	return "ScoreMessage"
}
