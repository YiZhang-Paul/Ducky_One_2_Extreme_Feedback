package duckyone2

import (
	"fmt"

	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"
)

func (c Controller) executeCi(mode string) {
	switch mode {
	case "passing":
		c.handlePassing()
	case "broken":
		c.handleBroken()
	case "building":
		c.handleBuilding()
	case "built":
		c.handleBuilt()
	case "build-failed":
		c.handleBuildFailed()
	}
}

func (c Controller) handlePassing() {
	data := map[string]interface{}{
		"BackRgb":   "1,28,73",
		"ActiveRgb": "255,255,255",
		"Steps":     60,
	}
	utils.PostJSON(fmt.Sprintf("%s/reactive", colorModeAPI), data)
}

func (c Controller) handleBroken() {

}

func (c Controller) handleBuilding() {

}

func (c Controller) handleBuilt() {

}

func (c Controller) handleBuildFailed() {

}
