package components

type KeyComponent struct {
	HasKey   bool
	Cooldown int
}

func (kc KeyComponent) Name() string {
	return "KeyComponent"
}
