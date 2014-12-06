package components

type ControlComponent struct {
	Scheme string
}

func (c ControlComponent) Name() string {
	return "ControlComponent"
}
