package components

type VisionComponent struct {
	Looking   bool
	Direction int
}

func (vs VisionComponent) Name() string {
	return "VisionComponent"
}
