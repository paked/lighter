package components

type VulnerableComponent struct {
	Is bool
}

func (vc VulnerableComponent) Name() string {
	return "VulnerableComponent"
}
