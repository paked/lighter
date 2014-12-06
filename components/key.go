package components

type KeyComponent struct {
	HasKey bool
}

func (kc KeyComponent) Name() string {
	return "KeyComponent"
}
