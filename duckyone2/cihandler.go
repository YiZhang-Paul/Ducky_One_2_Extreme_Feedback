package duckyone2

import (
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
	utils.PostJSON(reactiveModeAPI, data)
}

func (c Controller) handleBroken() {
	data := map[string]interface{}{
		"BackRgb":  "255,0,0",
		"Interval": 550,
	}
	utils.PostJSON(blinkModeAPI, data)
}

func (c Controller) handleBuilding() {

}

func (c Controller) handleBuilt() {

}

func (c Controller) handleBuildFailed() {

}
