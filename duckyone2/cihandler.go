package duckyone2

import (
	"github.com/yi-zhang/ducky-one-2-extreme-feedback/utils"
)

func (c Controller) executeCi(mode string) {
	switch mode {
	case Passing:
		c.handlePassing()
	case Broken:
		c.handleBroken()
	case Building:
		c.handleBuilding()
	case Built:
		c.handleBuilt()
	case BuildFailed:
		c.handleBuildFailed()
	}
}

func (c Controller) handlePassing() {
	if currentState == Passing {
		return
	}
	data := map[string]interface{}{
		"BackRgb":   "1,28,73",
		"ActiveRgb": "255,255,255",
		"Steps":     60,
	}
	utils.PostJSON(reactiveModeAPI, data)
	currentState = Passing
}

func (c Controller) handleBroken() {
	if currentState == Broken {
		return
	}
	data := map[string]interface{}{
		"BackRgb":  "255,0,0",
		"Interval": 550,
	}
	utils.PostJSON(blinkModeAPI, data)
	currentState = Broken
}

func (c Controller) handleBuilding() {
	if currentState == Building {
		return
	}
	data := map[string]interface{}{
		"BackRgb":    "255,0,255",
		"InnerRgb":   "0,255,225",
		"OuterRgb":   "255,99,71",
		"InnerSpeed": 25,
		"OuterSpeed": 55,
	}
	utils.PostJSON(progressModeAPI, data)
	currentState = Building
}

func (c Controller) handleBuilt() {

}

func (c Controller) handleBuildFailed() {

}
