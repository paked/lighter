package messages

type ScoreMessage struct {
}

func (sm ScoreMessage) Type() string {
	return "ScoreMessage"
}
