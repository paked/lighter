package globals

import (
	"github.com/paked/lighter/core"
)

var (
	W *core.Lighter
)

func init() {
	W = &core.Lighter{}
}
